package main // package name, mandatory

import ( // import packages to current file
  "fmt"  // fmt is part of standard library, see: https://golang.org/pkg/fmt/
)

func main() { // main function, will be executed like in java or C
  fmt.Println("Hello, world") // note the first letter of this function is capital 'P' ? why ?
}

// you can run by this command line ( or you can compile and run the binary )
// go run main.go
