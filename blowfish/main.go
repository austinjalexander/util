package main

import (
	"fmt"
)

var (
	// P ...
	P [18]uint32
	// S ...
	S [4][256]uint32
)

func f(x uint32) uint32 {
	h := S[0][x>>24] + S[1][x>>16&0xff]
	return (h ^ S[2][x>>8&0xff]) + S[3][x&0xff]
}

func encrypt(L, R *uint32) {
	for i := 16; i > 0; i -= 2 {
		*L ^= P[i]
		*R ^= f(*L)
		*R ^= P[i+1]
		*L ^= f(*R)
	}
	*L ^= P[1]
	*R ^= P[0]
}

func main() {
	fmt.Printf("%+v, %+v", P, S)
}
