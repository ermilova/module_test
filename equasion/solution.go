package main
import "errors"
func solve(a float64,b float64,c float64) ([]float64, error){
  result := make([]float64, 0, 3) 
  if a == 0 { 
    return result, errors.New("this is an error")
  }
  d := b * b - 4 * a * c
  if d == 0 {
    result = append(result, ((-1)* b) / (2 * a))
  } else if d > 0 {
    result = append(result, ((-1) * b - d) / (2 * a))
    result = append(result, ((-1) * b + d) / (2 * a)) 
  }
  return result, nil
}