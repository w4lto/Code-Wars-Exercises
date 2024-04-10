package main

import (
	"strconv"
	"strings"
)

func main() {
	CountBits(20)
}

/*
Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number. You can guarantee that input is non-negative.
Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case
*/
func CountBits(input uint) int {
	bin := strconv.FormatInt(int64(input), 2) // Gets the binary version of input number
	return strings.Count(bin, "1")            // Counts the occurencys of bit 1
}
