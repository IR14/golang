package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  prev := -1
	curr := 1
	return func() int {
		prev, curr = curr, prev + curr
		return curr
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// 0
// 1
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
