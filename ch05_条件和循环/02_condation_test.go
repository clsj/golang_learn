package ch05_条件和循环

import "testing"

func TestIfMultiSec(t *testing.T) {
	// a := 1 == 1 可以调用方法
	if a := 1 == 1; a {
		t.Log(a)
	}

	//if a, err := someFunc(); err != nil {
	//	t.Log(a)
	//}

}
