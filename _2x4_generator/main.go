/*
Script for generate 2x4 mode with Braille patterns
See: https://en.wikipedia.org/wiki/Braille_Patterns

Usage:

	go run _2x4_generator/main.go > box_chars_2x4.go
*/
package main

import (
	"fmt"
	"math/bits"
	"strings"
)

const tmpl = `package  tcg

var pixelChars2x4Braille = []rune{
	%s
}

`

func main() {
	rows := []string{"// symbol // code - graphics bits"}
	for i := 0; i < 256; i++ {
		code := transform(byte(i))
		rows = append(rows, fmt.Sprintf("'%c', // 0x%04X - %08b", code, code, i))
	}
	fmt.Printf(tmpl, strings.Join(rows, "\n"))
}

/*
|1|4|    |1|2|
|2|5| => |3|4|
|3|6|    |5|6|
|7|8|    |7|8|
*/
var tr = [8]byte{
	1, 4,
	2, 5,
	3, 6,
	7, 8,
}

func transform(in byte) rune {
	out := byte(0)

	for fromBit, toBit := range tr {
		bit := in & (0b1000_0000 >> (fromBit))
		if bit != 0 {
			out |= (0b1000_0000 >> (toBit - 1))
		}
	}

	out = bits.Reverse8(out)
	return 0x2800 + rune(out)
}
