package helper

import (
	"fmt"
	"strconv"
)

func reverse(n int) int {
	intString := strconv.Itoa(n)
	newString := ""
	for x := len(intString); x > 0; x-- {
		newString += string(intString[x-1])
	}
	newInt, err := strconv.Atoi(newString)
	if err != nil {
		fmt.Println(err)
	}
	return newInt
}

func Palindrom(x int, y int) int {
	total := 0
	for i := x; i <= y; i++ {
		awal := i
		hasilReverse := reverse(i)
		if awal == hasilReverse {
			total += 1
		}
	}
	return total
}
