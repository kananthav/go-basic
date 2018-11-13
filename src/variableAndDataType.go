package main

import "fmt"

func main() {
  // Declare variable with data type
  /*
  var firstName string = "Kevin"
  var lastName string = "Anantha"

  fmt.Printf("Hallo Kevin Anantha!\n")
  fmt.Printf("Hallo %s %s!\n", firstName, lastName)
  fmt.Println("Hallo", firstName + " " + lastName + "!")
  */

  // Declare variable without data type
  /*
  firstName := "Kevin"
  lastName := "Anantha"

  fmt.Printf("Hallo Kevin Anantha\n")
  fmt.Printf("Hallo %s %s\n", firstName, lastName)
  fmt.Println("Hallo", firstName + " " + lastName)
  */

  // Declare multi variable
  /*
  var firstName, lastName string
  firstName, lastName = "Kevin", "Anantha"

  var firstName, lastName string = "Kevin", "Anantha"

  firstName, lastName := "Kevin", "Anantha"
  */

  // Declare variable with underscore
  /*
  _ = "Learn Golang"
  _ = "Golang is easy"
  name, _ := "John", "Wick"

  fmt.Println(name)
  */

  // Declate variable with keyword "new"
  /*
  name := new(string)

  fmt.Println(name)
  fmt.Println(*name)
  */

  // Data type number non-decimal
  /*
  var positiveNumber uint8 = 89
  var negativeNumber = -1243423644

  fmt.Printf("Positive: %d\n", positiveNumber)
  fmt.Printf("Negative: %d\n", negativeNumber)
  */

  // Data type number decimal
  /*
  var decimal = 2.62

  fmt.Printf("Decimal: %f\n", decimal)
  fmt.Printf("Decimal: %.3f\n", decimal)
  fmt.Printf("Decimal: %.10f\n", decimal)
  */

  // Data type boolean
  /*
  exist := true
  fmt.Printf("exists? %t\n", exist) // use %t to format data boolean
  */

  // Data type string
  /*
  msg := "Hello"
  msg1 := `
    Hello, my name is "John Wick"
    nice to meet up
    let's learn Golang
  `
  fmt.Printf("message: %s\n", msg)
  fmt.Println(msg1)
  */

  // Const
  /*
  const firstName = "Kevin"
  fmt.Println(firstName)
  */
}
