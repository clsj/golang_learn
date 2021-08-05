package ch06_数组和切片

import "testing"

func TestSliceInit(t *testing.T) {
	//var arr []int
	//t.Log(cap(arr), len(arr))
	//arr = append(arr, 1)
	//t.Log(cap(arr), len(arr))
	//
	//arr2 := []int{1, 2, 3, 4}
	//t.Log(cap(arr2), len(arr2))
	//arr2 = append(arr2, 1)
	//t.Log(cap(arr2), len(arr2))

	// 长度为3 容量为5
	arr3 := make([]int, 3, 5)
	t.Log(cap(arr3), len(arr3))

}

func TestSliceCap(t *testing.T) {

	var arr []int

	for i := 0; i < 20; i++ {
		arr = append(arr, i)
		t.Log(len(arr), cap(arr))
	}
}

func TestSliceShare(t *testing.T) {

	// 共享内存
	arr := []string{"abc", "cde", "fgh", "ijk", "lmn", "opq"}
	arr2 := arr[0:2]
	t.Log(len(arr2), cap(arr2))

	arr3 := arr[1:3]
	t.Log(len(arr3), cap(arr3))
	arr3[0] = "xyz"

	t.Log(arr2, arr3)
}

func TestSliceCompare(t *testing.T) {
	// 共享内存
	arr1 := []int{1,2,3}
	arr2 := []int{1,2,3}
	// invalid operation: arr1 == arr2 (slice can only be compared to nil)
	//t.Log(arr1 == arr2)
	t.Log(arr1, arr2)
}