package ch09_字符串

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s := "hello"
	// cannot assign to s[1] 不可以修改s[1]的值
	//s[1] = 'a'
	s2 := s[1:2]
	t.Log(s2)
}

func TestStringFunc(t *testing.T) {
	s := "he,ll,o"

	array := strings.Split(s, ",")

	for _, value := range array {
		t.Log(value)
	}

	s = strings.Join(array, "-")

	t.Log(s)
}

func TestStringConv(t *testing.T) {
	s  := strconv.Itoa(10)

	t.Log(s)

	if a,err := strconv.Atoi("20"); err == nil {
		t.Log(a)
	}else {
		t.Log(err)
	}

}