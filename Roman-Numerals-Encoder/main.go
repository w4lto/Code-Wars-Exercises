package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(TransformToRomanNumber(182))
}

func TransformToRomanNumber(input int) string {
	if input <= 0 {
		return ""
	}
	var sb strings.Builder
	for i, v := range []int{1000, 900, 500, 400, 100, 90, 50, 40, 9, 5, 5, 1} {
		for input >= v {
			input -= v
			sb.WriteString([]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}[i])
		}
	}

	return sb.String()
}
