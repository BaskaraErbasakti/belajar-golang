package helper

import "testing"

func TestMissingNumber(t *testing.T) {
	input1 := MissingNumber("23242526272830")
	if input1 != 29 {
		t.Error("Result must be 29")
	}
	input2 := MissingNumber("101102103104105106107108109111112113")
	if input2 != 110 {
		t.Error("Result must be 110")
	}
	input3 := MissingNumber("12346789")
	if input3 != 5 {
		t.Error("Result must be 5")
	}
	input4 := MissingNumber("79101112")
	if input4 != 8 {
		t.Error("Result must be 8")
	}
	input5 := MissingNumber("7891012")
	if input5 != 11 {
		t.Error("Result must be 11")
	}
	input6 := MissingNumber("9799100101102")
	if input6 != 98 {
		t.Error("Result must be 98")
	}
	input7 := MissingNumber("100001100002100003100004100006")
	if input7 != 100005 {
		t.Error("Result must be 100005")
	}
}
