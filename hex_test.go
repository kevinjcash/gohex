package main

import (
	"testing"
)

func TestHexAdd(t *testing.T) {
	h1 := NewHexCube(2, -1, -1)
	h2 := NewHexCube(-1, 0, 1)
	if HexAdd(h1, h2) != NewHexCube(1, -1, 0) {
		t.Error()
	}
}

func TestHexSubtract(t *testing.T) {
	h1 := NewHexCube(2, -1, -1)
	h2 := NewHexCube(-1, 0, 1)
	if HexSubtract(h1, h2) != NewHexCube(3, -1, -2) {
		t.Error()
	}
}
