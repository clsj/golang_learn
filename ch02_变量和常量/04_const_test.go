package ch02_变量和常量

import "testing"

const (
	Monday = iota + 1
	Tuesday
)
func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday)
}

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConst2(t *testing.T) {
	t.Log(Readable, Writeable, Executable)
}
