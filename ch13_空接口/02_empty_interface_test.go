package ch13_空接口

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{})  {
	//if i, ok := p.(int); ok {
	//	fmt.Println("int ", i)
	//	return
	//}
	//
	//if i, ok := p.(string); ok {
	//	fmt.Println("string ", i)
	//	return
	//}

	switch i := p.(type) {
	case int:
		fmt.Println("int ", i)
	case string:
		fmt.Println("string ", i)
	default:
		fmt.Println("unknow ", i)
	}
}

func TestEmpty(t *testing.T) {
	DoSomething(true)
}