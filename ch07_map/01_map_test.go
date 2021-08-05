package ch07_map

import "testing"

func TestMapInit(t *testing.T) {

	map1 := map[string]int {"one": 1, "two" : 2}

	t.Log(len(map1))
	for s, i := range map1 {
		t.Log(s, i)
	}
}

func TestMapNotExist(t *testing.T) {

	map1 := make(map[int]int, 10)

	if v,ok := map1[1]; ok{
		t.Log("key 存在", v)
	}else {
		t.Log("key 不存在")
	}
}
