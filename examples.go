// Declare a list of variables.

var (
  name string
  age int
  location string
)

// multiple assignment
var (
  nam, location string
)


// Individually
var name string
// etc.


// Vars can also be assigned inside a function
// Inside a function, the := short assignment statement
// can be used in place of a var declaration with implicit type.
func main () {
  name, location := "Bob", "New York"
  age :=  78
  fmt.Printf(("%s (%d) of %s", name, age, location)
}


// COnstant's are declare with `const`
const Pi = 3.14
const (
        StatusOK                   = 200
        StatusCreated              = 201
)


//Printing vars (or const)
func main() {
  someNumber := 6
  fmt.Println(someNumber)
}


// fmt.Println prints the passed in variables’ values and appends a newline.
// fmt.Printf is used when you want to print one or multiple values using a defined format specifier.


// Programs start running in package main.


// A function can take zero or more typed arguments.
// The type comes after the variable name.
// Functions can be defined to return any number
// of values that are always typed.

package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(add(42, 13))
}

// No need to declare the type of each parameters in this example.

package main

import "fmt"

func add(x, y int) int {
  return x + y
}

func main() {
  fmt.Println(add(42, 13))
}

// Functions take parameters.
// In Go, functions can return multiple “result parameters”,
// not just a single value. They can be named and act just like variables.

func location(city string) (string, string) {
  var region string
  var continent string

  switch city {
  case "Los Angeles", "LA", "Santa Monica":
    region, continent = "California", "North America"
  case "New York", "NYC":
    region, continent = "New York", "North America"
  default:
    region, continent = "Unknown", "Unknown"
  }
  return region, continent
}

func main() {
  region, continent := location("Santa Monica")
  fmt.Printf("Matt lives in %s, %s", region, continent)
}



