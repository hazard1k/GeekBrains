package main

import (
	"reflect"
	"testing"
)

type expected struct {
	isRoot bool
	x1, x2 float32
}

var tests = []struct {
	a   float32
	b   float32
	c   float32
	exp expected
}{
	{2, 7, -4, expected{true, 0.5, -4}},
	{1, -4, 4, expected{true, 2, 2}},
	{3, -1, 7, expected{false, 0, 0}},
	{2, 5, 3, expected{true, -1, -1.5}},
}

func TestGetRoots(t *testing.T) {
	for _, e := range tests {
		isRoot, x1, x2 := getRoots(e.a, e.b, e.c)
		if !reflect.DeepEqual(e.exp, expected{isRoot, x1, x2}) {
			t.Errorf("get(%v, %v, %v) = %v, %v, %v , expected %v, x1=%v, x2=%v", e.a, e.b, e.c, isRoot, x1, x2, e.exp.isRoot, e.exp.x1, e.exp.x2)
		}
	}
}
