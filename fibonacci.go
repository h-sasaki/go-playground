package main

import (
  "fmt"
  "math/big"
)

func main () {
  m := big.NewInt(1)
  n := big.NewInt(1)

  fmt.Println(m)
  fmt.Println(n)

  for i := 0; i < 100; i++ {
//  n = m + n
    n = new(big.Int).Add(m, n)
//  m = n - m
    m = new(big.Int).Sub(n, m)
    fmt.Println(n)
  }
}
