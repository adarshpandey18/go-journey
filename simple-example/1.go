package main

import (
  "fmt"
  "os" 
)

func main() {
  if len(os.Args) > 1 {
    fmt.Println("Hello", os.Args[1])
  } else {
    fmt.Println("Hello, world!")
  }
}

/*
 
CMD 
 go run .             --->      output : Hello, world!
 go run . Adarsh      --->      output : Hello, Adarsh

*/
