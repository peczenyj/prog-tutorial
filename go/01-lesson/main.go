package main

import (
	"fmt"
)

func sum(a, b int) int {
	return 0 // TODO: implement
}

func test() {
	fmt.Printf("%d + %d = %d\n", 1, 1, sum(1, 1))
}


func main() {
  test()
}
// this must print the expected answer: 1 + 1 = 2
// go run main.go
