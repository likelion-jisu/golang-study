package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// bool의 zero 값은 false
	var flag bool
	var isAwesome = true
	fmt.Println("flag:", flag)
	fmt.Println("isAwesome:", isAwesome)

	// 복소수
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)

	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))
	fmt.Println(cmplx.NaN())

	// 명시적 타입 변환
	var x1 int = 10
	var y1 float64 = 30.2
	var z1 float64 = float64(x1) + y1
	var d int = x1 + int(y1)
	fmt.Println("x1:", x1)
	fmt.Println("y1:", y1)
	fmt.Println("z1:", z1)
	fmt.Println("d:", d)

	// 변수 선언
	var x2 int
	var x3, y3 int = 10, 20
	var x4, y4 int
	var x5, y5 = 10, "hello"
	//y5[0] = 'c'
	fmt.Println("x2:", x2)
	fmt.Println("x3, y3:", x3, y3)
	fmt.Println("x4, y4:", x4, y4)
	fmt.Println("x5, y5:", x5, y5)

	var (
		x6     int
		y6         = 20
		z6     int = 30
		d6, e6     = 40, "hello"
		f, g   string
	)
	fmt.Println("x6:", x6)
	fmt.Println("y6:", y6)
	fmt.Println("z6:", z6)
	fmt.Println("d6, e6:", d6, e6)
	fmt.Println("f, g:", f, g)

	x7, y7 := 10, "hello"
	x8, z8 := 20, "hello"
	fmt.Println("x7, y7:", x7, y7)
	fmt.Println("x8, z8:", x8, z8)

	fmt.Println("x9:", x9)
	fmt.Println("idKey:", idKey)
	fmt.Println("nameKey:", nameKey)
	fmt.Println("z9:", z9)
}

const x9 int64 = 10
const (
	idKey   = "id"
	nameKey = "name"
)
const z9 = 20 * 10
