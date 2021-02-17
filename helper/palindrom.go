package helper

import (
	"fmt"
	"strconv"
	"strings"
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

func Palindrom(x string) int {
	strs := strings.Split(x, " ")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}
	total := 0
	for i := ary[0]; i <= ary[1]; i++ {
		awal := i
		hasilReverse := reverse(i)
		if awal == hasilReverse {
			total += 1
		}
	}
	return total
}
