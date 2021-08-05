package ch11_面向对象

import (
	"fmt"
	"testing"
	"time"
)

type InnerType func(op int) int

func timeSpent(inner InnerType) InnerType {
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
