package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := 1.0
  y := 1.0
  for i := 0; i < 100; i++ {
    z -= (z*z - x) / (2*z)
    fmt.Println(z)
    if math.Abs(z - y) < 1e-6 {
      fmt.Println("count: ", i)
      break
    }
    y = z
  }
  return z
}

func main() {
  fmt.Println("sqrt: ", Sqrt(2))
  fmt.Println("math: ", math.Sqrt(2))
}
