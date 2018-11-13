package main

import "fmt"

func main() {
  point := 8

  if point == 10 {
    fmt.Println("Perfect = 10")
  } else if point > 5 {
    fmt.Println("Congrats = 5")
  } else if point == 4 {
    fmt.Println("Almost = 4")
  } else {
    fmt.Printf("Not lucky %d", point)
  }
}
