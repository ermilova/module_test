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

func TestSolveAZero(t *testing.T) {
	var A float64 = 0
	var B float64 = 1
	var C float64 = 1
	
	_, err := solve(A, B, C)

	assert.NotNil(t, err)
}
func TestSolveAEps(t *testing.T) {
	var A float64 = 1e-10
	var B float64 = 1
	var C float64 = 1
	
	_, err := solve(A, B, C)

	assert.NotNil(t, err)
}
func TestSolveDEps(t *testing.T) {
	var A float64 = 1e-5
	var B float64 = -2.5 * 1e-5
	var C float64= 1e-5
	expected := []float64 {1.9999999999999996, 0.5000000000000001}
	actual, err := solve(A, B, C)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestSolveNan(t *testing.T) {
	A := math.NaN()
	B := math.NaN()
	C := math.NaN()
	
	_, err := solve(A, B, C)
	
	assert.NotNil(t, err)
}
func TestSolveInf(t *testing.T) {
	A := math.Inf()
	B := math.Inf()
	C := math.Inf()
	
	_, err := solve(A, B, C)
	
	assert.NotNil(t, err)
}
