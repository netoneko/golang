package main

import (
    "fmt"
    "math"
)

func iterate(z, x float64) float64 {
    return z - (z * z - x) / 2 / z   
}

func Sqrt(x float64)(z float64) {
  z = x / 2
    
  for ; math.Abs(z * z - x) > 0.00000001; z = iterate(z, x) {
  }
    
  return
}

func main() {
  fmt.Println(Sqrt(2))
}

