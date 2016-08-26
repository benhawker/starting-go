package main

import "os"
import "fmt"

func main() {
  fmt.Println("All args", os.Args)

  fmt.Println("Length of os args", len(os.Args))

  fmt.Println("Concating 2 args", os.Args[1] + os.Args[2])

  fmt.Println("The first arg provided", os.Args[1])

  fmt.Println("the agrs provided without the command", os.Args[1:])

  fmt.Println("the final arg provided", os.Args[len(os.Args) - 1])
}
