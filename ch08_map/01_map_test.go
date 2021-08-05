package ch08_map

import "testing"

func TestMapWithFunc(t *testing.T) {
	map1 := map[int]func(op int) int{}
	map1[0] = func(op int) int {
		return op + 1
	}

	map1[1] = func(op int) int {
		return op * 2
	}

	map1[2] = func(op int) int {
		return op * op
	}

	t.Log(map1[0](2), map1[1](2), map1[2](4))

}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d 已存在", n)
	}else {
		t.Logf("%d 不已存在", n)
	}
	
}