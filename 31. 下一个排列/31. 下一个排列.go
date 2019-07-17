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
