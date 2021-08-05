package ch05_条件和循环

import "testing"

func TestWhile(t *testing.T) {
	n := 0

	// 没有while 只有for
	for n < 5 {
		t.Log(n)
		n++
	}
}
