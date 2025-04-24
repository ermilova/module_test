package Solve
import (
  "errors"
  "math"
  )
  
func solve(a float64,b float64,c float64) ([]float64, error){
  result := make([]float64, 0, 3) 
  epsilon := 1e-9
  if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
    return result, errors.New("One of parameters is NaN")
  }
  if math.IsInf(a) || math.IsInf(b) || math.IsInf(c) {
    return result, errors.New("One of paramets is Inf")
  }
  if math.Abs(a) < epsilon { 
    return result, errors.New("a == 0")
  }
  d := b * b - 4 * a * c
  if math.Abs(d) < epsilon {
    result = append(result, ((-1)* b) / (2 * a))
  } else if !math.IsNaN(math.Sqrt(d)) {
    result = append(result, ((-1) * b + math.Sqrt(d)) / (2 * a))
    result = append(result, ((-1) * b - math.Sqrt(d)) / (2 * a))
  }
  return result, nil
}