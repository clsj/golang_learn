package ch03_基本数据类型

import "testing"

// 类型转换
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64

	// cannot use a (type int32) as type int64 in assignment
	//b = a
	t.Log(a, b)
}

func TestPoint(t *testing.T) {
	a := 1
	apt := &a
	t.Log(a, apt)

	// 不支持指针运算
	//apt = apt + 1

	t.Logf("%T, %T", a, apt)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))

	if s == "" {
		t.Log("s")
	}

}