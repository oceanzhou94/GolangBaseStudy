package fib

import (
	"fmt"
	"testing"
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

//斐波那契数的非递归实现
func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func TestConstantTry(t *testing.T) {
	a := 2
	t.Log(a&Readable == Readable)
	t.Log(a&Writable == Writable)
	t.Log(a&Executable == Executable)
}
