package Solve

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveTwo(t *testing.T) {
	var A float64 = 1
	var B float64 = 0
	var C float64 = -1
	expected := []float64 {1, -1}
	actual, err := solve(A, B, C)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
func TestSolveZero(t *testing.T) {
	var A float64 = 1
	var B float64 = 0
	var C float64 = 1
	expected := make([]float64, 0)
	actual, err := solve(A, B, C)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestSolveOne(t *testing.T) {
	var A float64 = 1
	var B float64 = 2
	var C float64= 1
	expected := []float64 {-1}
	actual, err := solve(A, B, C)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestSolveA(t *testing.T) {
	var A float64 = 0
	var B float64 = 1
	var C float64 = 1
	
	_, err := solve(A, B, C)

	assert.NotNil(t, err)
}
