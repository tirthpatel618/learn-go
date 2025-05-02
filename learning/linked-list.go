package main

import (
	"fmt"
	"strings"
)

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Head() T {
	return l.val
}

func (l *List[T]) Tail() *List[T] {
	return l.next
}

func (l *List[T]) Prepend(val T) *List[T] {
	return &List[T]{next: l, val: val}
}

func (l *List[T]) Len() int {
	count := 0
	for curr := l; curr != nil; curr = curr.next {
		count++
	}
	return count
}

// Stringer method
func (l *List[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for curr := l; curr != nil; curr = curr.next {
		sb.WriteString(fmt.Sprintf("%v", curr.val))
		if curr.next != nil {
			sb.WriteString(" -> ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func main() {
	var list *List[int]
	list = list.Prepend(3)
	list = list.Prepend(2)
	list = list.Prepend(1)

	fmt.Println("List:", list)
	fmt.Println("Head:", list.Head())
	fmt.Println("Length:", list.Len())
}
