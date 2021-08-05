package ch06_数组和切片

import "testing"

func TestArrayInit(t *testing.T) {

	var arr1 [3]int

	arr2 := [4]int{1,2,3,4}

	arr3 := []int{1,2,3,4, 5}

	arr3[1] = 1

	t.Log(arr1, arr2, arr3)
}

func TestArrayTravel(t *testing.T) {

	arr := [4]int{1,2,3,4}

	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for index, value := range arr {
		t.Log(index, value)
	}
}

func TestArraySection(t *testing.T) {

	arr := [4]int{1,2,3,4}
	arr1 := arr[:3]

	for _, value := range arr1 {
		t.Log(value)
	}
}