package main

import (
	"fmt"
	"strings"
)

func main2() {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)
	v, ok := s.Pop()
	fmt.Println(v, ok)
	// s.Push("nope") -> compile error

	var it *Tree[OrderableInt]
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(12))
	a := 10
	// uncomment to see a compile-time error
	// it = it.Insert(a)
	it = it.Insert(OrderableInt(a))
	// uncomment to see a compile-time error
	// it = it.Insert(OrderableString("nope"))
}

// type 파라미터는 대괄호 뒤에 위치한다.
type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		// var는 다른 값이 할당되지 않으면 항상 제로 값을 변수에 초기화한다.
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

type Orderable[T any] interface {
	Order(T) int
}

func (oi OrderableInt) Order(val OrderableInt) int {
	return int(oi - val)
}

func (os OrderableString) Order(val OrderableString) int {
	return strings.Compare(string(os), string(val))
}

type Tree[T Orderable[T]] struct {
	val         T
	left, right *Tree[T]
}

func (t *Tree[T]) Insert(val T) *Tree[T] {
	if t == nil {
		return &Tree[T]{val: val}
	}

	switch comp := val.Order(t.val); {
	case comp < 0:
		t.left = t.left.Insert(val)
	case comp > 0:
		t.right = t.right.Insert(val)
	}
	return t
}

func (t *Tree[T]) Contains(val T) bool {
	if t == nil {
		return false
	}
	switch comp := val.Order(t.val); {
	case comp < 0:
		return t.left.Contains(val)
	case comp > 0:
		return t.right.Contains(val)
	default:
		return true
	}
}
