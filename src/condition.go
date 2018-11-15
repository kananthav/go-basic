package main

import "fmt"

func main() {
  // If else
  /*
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
  */

  // Temporary if else
  /*
  point := 8840.0

  if percent := point / 100; percent >= 100 {
    fmt.Printf("%.1f%s perfect!\n", percent, "%")
  } else if percent >= 40 {
    fmt.Printf("%.1f%s good\n", percent, "%")
  } else {
    fmt.Printf("%.1f%s bad\n", percent, "%")
  }
  */

  // Nested if else
  /*
  point := 10

  if point > 7 {
    switch point {
      case 10:
        fmt.Println("Perfect")
      default:
        fmt.Println("Nice")
    }
  } else {
    if point == 5 {
      fmt.Println("not bad")
    } else if point == 3 {
      fmt.Println("keep trying")
    } else {
      fmt.Println("you can do it")
      if point == 0 {
        fmt.Println("try harder!")
      }
    }
  }
  */

  // For
  /*
  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }
  */

  // For with condition (argument)
  /*
  i := 0

  for i < 5 {
    fmt.Println(i)
    i++
  }
  */

  // For without argument
  /*
  i := 0

  for {
    fmt.Println(i)

    i++
    if i == 5 {
      break
    }
  }
  */

  // Break and continue
  /*
  for i:= 1; i <= 10; i++ {
    if i % 2 == 1 {
      continue
    }

    if i > 8 {
      break
    }

    fmt.Println(i)
  }
  */

  // Nested for
  /*
  for i := 0; i < 5; i++ {
    for j := i; j < 5; j ++ {
      fmt.Print(j, " ")
    }
    fmt.Println()
  }
  */
}
