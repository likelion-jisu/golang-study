package main

import (
	"fmt"
	"io"
)

func doThings(i interface{}) {
	switch j := i.(type) {
	case nil:
		// j의 타입은 interface{} 이다.
		fmt.Println("nil")
	case int:
		// j의 타입은 int다.
		fmt.Println("int")
	case io.Reader:
		// j의 타입은 io.Reader다.
		fmt.Println("io.Reader")
	case string:
		// j의 타입은 string이다.
		fmt.Println("string")
	case bool, rune:
		// j의 타입은 interface{} 이다.
		fmt.Println("bool or rune")
	default:
		// j의 타입은 interface{} 이다.
		fmt.Println("unknown")
		fmt.Println(j)
	}
}
