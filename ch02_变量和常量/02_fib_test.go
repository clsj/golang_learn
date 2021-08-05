package ch02_变量和常量

import "testing"

func TestFib(t *testing.T) {
	a := 1;
	b := 1;
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		temp := a
		a = b;
		b = temp + a
	}
}

