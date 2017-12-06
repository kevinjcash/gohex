package main

import "fmt"
import "math"

type Hex struct {
	q int
	r int
	s int
}

func intabs(x int) int {
	return int(math.Abs(float64(x)))
}

func NewHexCube(q_ int, r_ int, s_ int) Hex {
	if q_+r_+s_ != 0 {
		fmt.Println("Invalid hex coordinates!")
	}
	h := Hex{q: q_, r: r_, s: s_}
	return h
}

func NewHexAxial(q_ int, r_ int) Hex {
	s_ := -q_ - r_
	h := Hex{q: q_, r: r_, s: s_}
	return h
}

func HexAdd(a Hex, b Hex) Hex {
	return NewHexCube(a.q+b.q, a.r+b.r, a.s+b.s)
}

func HexSubtract(a Hex, b Hex) Hex {
	return NewHexCube(a.q-b.q, a.r-b.r, a.s-b.s)
}

func HexMultiply(a Hex, k int) Hex {
	return NewHexCube(a.q*k, a.r*k, a.s*k)
}

func HexLength(hex Hex) int {
	return int((intabs(hex.q) + intabs(hex.r) + intabs(hex.s)) / 2)
}

func HexDistance(a Hex, b Hex) int {
	return HexLength(HexSubtract(a, b))
}

func main() {
}
