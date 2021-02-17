package helper

import "testing"

func TestPalindrom(t *testing.T) {
	input1 := Palindrom(1, 10)
	if input1 != 9 {
		t.Error("Result must be 9")
	}
	input2 := Palindrom(99, 100)
	if input2 != 1 {
		t.Error("Result must be 1")
	}
	input3 := Palindrom(21, 31)
	if input3 != 1 {
		t.Error("Result must be 1")
	}
}
