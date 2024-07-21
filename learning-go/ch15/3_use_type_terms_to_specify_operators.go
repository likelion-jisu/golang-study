package main

import (
	"errors"
	"fmt"
)

func main3() {
	var a uint = 18_446_744_073_709_551_615
	var b uint = 9_223_372_036_854_775_808
	fmt.Println(divAndRemainder(a, b))

	type MyInt int
	var myA MyInt = 10
	var myB MyInt = 20
	fmt.Println(divAndRemainder(myA, myB))

	// compile error
	/*
		s := ImpossibleStruct[int]{10}
		fmt.Println(s)
		s2 := ImpossibleStruct[MyInt]{10}
		fmt.Println(s2)
	*/
}

// byte와 rune은 각각 uint8과 int32의 유형 별칭이므로 생략할 수 있습니다.

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func divAndRemainder[T Integer](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

// 기본 유형 int와 String() 문자열 메소드를 가져야 함을 지정할 수 있습니다.

type PrintableInt interface {
	~int
	String() string
}

// 실제로 인스턴스화할 수 없는 유형 매개변수 인터페이스를 선언할 수 있다는 점에 유의하세요.
// int에는 메서드가 없으므로 이를 충족하는 유효한 유형이 없습니다.
// 컴파일 오류가 발생하므로 실제로 사용할 수 없습니다.

type ImpossiblePrintableInt interface {
	int
	String() string
}

type ImpossibleStruct[T ImpossiblePrintableInt] struct {
	val T
}

type MyInt int

func (mi MyInt) String() string {
	return fmt.Sprint(mi)
}
