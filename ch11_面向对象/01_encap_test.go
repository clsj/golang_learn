package ch11_面向对象

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id string
	Name string
	Age int
}


func TestCreate(t *testing.T) {

	e := Employee{"11", "li", 2}

	e1 := Employee{Age: 20}

	// 返回的是指针类型
	e2 := new(Employee)

	e2.Id = "222"
	e2.Name = "xxx"
	e2.Age = 50

	t.Log(e, e1, e2)

	t.Logf("%T, %T", e, e2)
	t.Logf("%T, %T", &e, e2)
}

func (e Employee) String() string{
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// 没有数据复制
func (e *Employee) String2() string{
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestMethod(t *testing.T) {
	e := Employee{"11", "li", 2}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	//t.Log(e.String())
	t.Log(e.String2())
}