package ch15_package

import (
	"fmt"
	"golang_learn/ch15_package/series"
	"testing"
)

func TestPackage(t *testing.T) {
	v, _ := series.GetFib(3)

	fmt.Println(v)

}


