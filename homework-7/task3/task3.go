package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// Server ..
type Server struct {
	ns       net.Listener
	clients  []chan struct{} // Пустая структура не занимает памяти
	isRunned bool
}

// Start Запуск сервера.
func (srv *Server) Start() error {
	fmt.Println("Server starting at", srv.ns.Addr().String())
	go srv.handleInput()
	srv.isRunned = true
	for {
		conn, err := srv.ns.Accept()
		if err != nil {
			// Если сюда попали и флаг статуса сервера false, то это shutdown
			if !srv.isRunned {
				return nil
			}
			log.Print(err)
			continue
		}
		go srv.handleConn(conn)
	}
}

// Stop Остановка сервера. Закрывает все коннекты клиентов
func (srv *Server) Stop() error {
	fmt.Println("Server stopping...")
	srv.isRunned = false
	err := srv.ns.Close()
	for _, cancelCh := range srv.clients {
		cancelCh <- struct{}{}
		close(cancelCh)
	}
	return err
}

// handleInput метод слушает ввод комманд сервера
func (srv *Server) handleInput() {
	var command string
	for {
		_, err := fmt.Scan(&command)
		if err != nil {
			fmt.Println("Error", err.Error())
			continue
		}
		if command == "exit" {
			srv.Stop()
		}
	}
}

// handleConn метод делает полезную работу для клиента
func (srv *Server) handleConn(conn net.Conn) {
	fmt.Println("Connected from", conn.RemoteAddr().String())
	clientCh := make(chan struct{})
	srv.clients = append(srv.clients, clientCh)
	for {
		select {
		case <-clientCh:
			conn.Close()
			break
		default:
			_, err := io.WriteString(conn, time.Now().Format("15:04:05\n\r"))
			if err != nil {
				fmt.Println("err", err.Error())
				return
			}
			time.Sleep(1 * time.Second)
		}
	}
}

// NewTCP Создает TCP сервер и возвращает указатель на него
func NewTCP(address string) *Server {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	return &Server{
		ns:       listener,
		clients:  []chan struct{}{},
		isRunned: false,
	}
}

func main() {
	server := NewTCP("localhost:8000")
	server.Start()
}
