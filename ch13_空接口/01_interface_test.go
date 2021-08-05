package ch13_空接口

import (
	"fmt"
	"testing"
)

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

func WriteProgram(p Programmer)  {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)

	WriteProgram(goProg)
	WriteProgram(javaProg)

}