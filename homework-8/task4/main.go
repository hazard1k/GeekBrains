package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type privateMessage struct {
	from string
	to   string
	msg  string
}

type cli struct {
	who string
	ch  chan<- string
}

type client chan<- string
type privatemsg chan<- privateMessage

var (
	entering = make(chan cli)
	leaving  = make(chan cli)
	messages = make(chan string)
	prvmsg   = make(chan privateMessage)
)

func main() {
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	log.Printf("Server running at %s", listener.Addr().String())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[string]client)
	for {
		select {
		case msg := <-messages:
			for _, cli := range clients {
				cli <- msg
			}
		case msg := <-prvmsg:
			if cli, ok := clients[msg.to]; ok {
				cli <- "from " + msg.from + " " + msg.msg
			}
		case cli := <-entering:
			clients[cli.who] = cli.ch

		case cli := <-leaving:
			delete(clients, cli.who)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	ch <- "Enter your nick name"

	input := bufio.NewScanner(conn)
	input.Scan()
	who := input.Text()

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli{who, ch}

	for input.Scan() {
		msg := input.Text()
		if strings.HasPrefix(msg, "to:") {
			privateInput := strings.Split(msg, " ")
			prvmsg <- privateMessage{
				from: who,
				msg:  strings.Join(privateInput[1:], " "),
				to:   strings.TrimPrefix(privateInput[0], "to:")}
			continue
		}
		messages <- who + ": " + msg
	}
	leaving <- cli{who, ch}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
