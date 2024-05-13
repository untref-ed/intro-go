package main

import (
	"fmt"

	"github.com/untref-ed/intro-go/tda/stack"
)

func main() {
	s := stack.Stack{}

	s.Push("world!")
	s.Push("Hello, ")

	for s.Size() > 0 {
		if x, err := s.Pop(); err == nil {
			fmt.Print(x)
		}
	}

	fmt.Println()
}
