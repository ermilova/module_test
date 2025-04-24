package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveTwo(t *testing.T) {
	A := 1
	B := 0
	C := -1
	expected := []float64 {1, -1}
	actual, err := solve(A, B, C)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
func TestSolveZero(t *testing.T) {
	A := 1
	B := 0
	C := 1
	expected := make([]float64, 0)
	actual, err := solve(A, B, C)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestSolveOne(t *testing.T) {
	A := 1
	B := 2
	C := 1
	expected := [1]float64 {1}
	actual, err := solve(A, B, C)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestSolveA(t *testing.T) {
	A := 0
	B := 1
	C := 1
	
	actual, err := solve(A, B, C)

	assert.NilError(t, err)
}
