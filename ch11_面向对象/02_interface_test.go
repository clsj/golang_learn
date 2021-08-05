package ch11_面向对象

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {

}
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"hello world\")"
}

type JavaProgrammer struct {

}

func (j *JavaProgrammer) WriteHelloWorld() string {
	return "sout(\"hello world\")"
}

func TestInterface(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())

	p = new(JavaProgrammer)
	t.Log(p.WriteHelloWorld())
}
