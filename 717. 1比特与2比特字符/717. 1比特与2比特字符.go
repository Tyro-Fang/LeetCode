package main

import (
	"fmt"
)

func main() {
	a := []int{1, 1, 1, 0}
	b := isOneBitCharacterTX(a)
	fmt.Println(b)
	//c := maximalRectangle(a)
	//fmt.Println(c)
}
func isOneBitCharacter(bits []int) bool {
	blen := len(bits)
	i := 0
	for i < blen {
		if i == blen-1 {
			return true
		}
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}

	}
	return false
}
func isOneBitCharacterTX(bits []int) bool { //贪心法，从倒数第二个字符回溯，到遇到第一个0结束，若是其中1为奇数，则最后一个字符为2比特，否则为1比特
	blen := len(bits)
	count := 0
	for i := blen - 2; i >= 0; i-- {
		if bits[i] == 1 {
			count++
		} else {
			break
		}
	}
	y := count % 2
	if y == 1 {
		return false
	}
	return true
}
