package main

import "fmt" 

func main () {
  m := 1 
  n := 1 

  fmt.Println(m)
  fmt.Println(n)

  for i := 0; i < 100; i++ {
    n = m + n
    m = n - m
    fmt.Println(n)
  }
}
