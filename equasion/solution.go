package Solve
import (
  "errors"
  "math"
  )
  
func solve(a float64,b float64,c float64) ([]float64, error){
  result := make([]float64, 0, 3) 
  epsilon := math.Nextafter(1.0,2.0)-1.0
  if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
    return result, errors.New("One of parameters is NaN")
  }
  if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
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