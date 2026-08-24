package utils

import "fmt"

func Print() {
	fmt.Println("I'm a print util")
}

func Max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

type GInteger interface {
	~int | ~uint
}

func GIMax[T GInteger](a T, b T) T {
	if a >= b {
		return a
	} else {
		return b
	}
}

type Box[T any] struct {
	V T
}

func (b Box[T]) Map[X GInteger](f func(T) X) Box[X] {
	return Box[X]{f(b.V)}
}
