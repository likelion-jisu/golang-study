package main

import (
	"fmt"
	"strings"
)

func main1() {
	var s StackOld
	s.Push(1)
	s.Push(2)
	s.Push(3)
	v, ok := s.Pop()
	fmt.Println(v, ok)

	// panic: interface conversion: interface {} is main.OrderableString, not main.OrderableInt
	var it *TreeOld
	it = it.Insert(OrderableInt(3))
	it = it.Insert(OrderableString("nope"))
}

type OrderableOld interface {
	OrderOld(interface{}) int
}

type TreeOld struct {
	val         OrderableOld
	left, right *TreeOld
}

func (t *TreeOld) Insert(val OrderableOld) *TreeOld {
	if t == nil {
		return &TreeOld{val: val}
	}
	if t.val.OrderOld(val) < 0 {
		t.left = t.left.Insert(val)
	} else {
		t.right = t.right.Insert(val)
	}
	return t
}

type OrderableInt int

func (o OrderableInt) OrderOld(other interface{}) int {
	return int(o) - int(other.(OrderableInt))
}

type OrderableString string

func (o OrderableString) OrderOld(other interface{}) int {
	return strings.Compare(string(o), other.(string))
}

type StackOld struct {
	vals []interface{}
}

func (s *StackOld) Push(val interface{}) {
	s.vals = append(s.vals, val)
}

func (s *StackOld) Pop() (interface{}, bool) {
	if len(s.vals) == 0 {
		return nil, false
	}
	val := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return val, true
}
