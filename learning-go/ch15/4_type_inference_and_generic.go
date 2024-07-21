package main

import (
	"fmt"
	"reflect"
)

func Convert[T1, T2 Integer](in T1) T2 {
	return T2(in)
}

func main4() {
	var a int = 10
	b := Convert[int, int64](a)
	fmt.Println(b, reflect.TypeOf(b))
}
