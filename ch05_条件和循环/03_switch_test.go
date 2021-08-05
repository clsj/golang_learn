package ch05_条件和循环

import (
	"runtime"
	"testing"
)

func TestSwitch(t *testing.T) {

	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("OS.X")
	case "linux":
		t.Log("linux")
	default:
		t.Log(os)
	}

}

func TestSwitch2(t *testing.T) {
	const Num = 4
	switch  {
	case 0 <= Num && Num <= 3:
		t.Log("0-3")
	case 4 <= Num && Num <= 10:
		t.Log("4-10")
	default:
		t.Log(">10")
	}

}