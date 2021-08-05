package ch12_继承

import (
	"fmt"
	"testing"
)

type Pet struct {

}

func (p *Pet) Speak() {
	fmt.Println(".....")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(".....")
}

type Dog struct {
	Pet
}

// 无法重载
func (d *Dog) Speak() {
	fmt.Println("wangwang")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	// .....
	dog.SpeakTo("Chao")
}
