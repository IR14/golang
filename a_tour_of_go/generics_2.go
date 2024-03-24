package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) isEnd() bool {
	if l.next == nil {
		return true
	}
	return false
}

func main() {
	l := List[int]{val: 4}
	
	l.next = &List[int]{val: 5}
	
	fmt.Printf("%v - %t\n", l.val, l.isEnd())
	
	curr := l.next
	
	fmt.Printf("%v - %t\n", curr.val, curr.isEnd())
}

// 4 - false
// 5 - true
