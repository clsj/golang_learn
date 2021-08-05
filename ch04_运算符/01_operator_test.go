package ch04_运算符

import "testing"

func TestComparArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 2, 3, 5}

	d := [...]int{1, 2, 3}

	// invalid operation: a == b (mismatched types [3]int and [4]int)
	// t.Log(a == b)
	t.Log(b == c)
	t.Log(a == d)

}
