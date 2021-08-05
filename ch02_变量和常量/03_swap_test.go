package ch02_变量和常量

import "testing"

func TestSwap(t *testing.T) {
	a := 1
	b := 2
	a,b = b,a
	t.Log(a, b)
}