package main

import (
	"fmt"
	"os"
)

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func MyFunc(opts MyFuncOpts) error {
	fmt.Println("FirstName:", opts.FirstName)
	fmt.Println("LastName:", opts.LastName)
	fmt.Println("Age:", opts.Age)
	return nil
}

func main1() {
	// 이름이 지정된 파라미터
	MyFunc(MyFuncOpts{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	})

	// 선택적 파라미터 사용
	MyFunc(MyFuncOpts{
		FirstName: "Jane",
	})

	fmt.Println(divide(10, 0))

	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	// reminder 반환값은 무시
	result, _, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)

	x, y, z := divAndRemainder2(5, 2)
	fmt.Println(x, y, z)

	x1, y1, z1 := divAndRemainder3(5, 2)
	fmt.Println(x1, y1, z1)
}

// 가변 인자
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// 다중 반환값 허용
func divAndRemainder(numerator int, demonimator int) (result int, remainder int, err error) {
	if demonimator == 0 {
		err = fmt.Errorf("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/demonimator, numerator%demonimator
	return result, remainder, err
}

func divAndRemainder2(numerator, demonimator int) (result int, remainder int, err error) {
	// 20, 30은 사용되지 않는다.
	result, remainder = 20, 30
	if demonimator == 0 {
		return 0, 0, fmt.Errorf("cannot divide by zero")
	}
	return numerator / demonimator, numerator % demonimator, nil
}

// 빈 반환을 사용할 수 있다. 어떤 값을 반환하는지 알아내기 어려워서 사용하지 말자.
func divAndRemainder3(numerator, demonimator int) (result int, remainder int, err error) {
	if demonimator == 0 {
		err = fmt.Errorf("cannot divide by zero")
		return
	}
	result, remainder = numerator/demonimator, numerator%demonimator
	return
}
