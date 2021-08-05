package ch10_函数

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(40)
}

func TestFuncMultiValues(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		res := inner(op)
		fmt.Println("time spent", time.Since(start).Seconds())
		return res
	}
}

func TestTimeSpent(t *testing.T) {
	timeSpent(func(op int) int {
		return op * op
	})(20)
}

func sum(op ...int) int {
	ret := 0
	for _, value := range op {
		ret += value
	}
	return ret
}
func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3))
}

func clear() {
	fmt.Println("clear...")
}

func TestDef(t *testing.T) {
	defer clear()
	fmt.Println("start")
	panic("err")
	fmt.Println("end")
}