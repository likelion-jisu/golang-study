package main

import (
	"fmt"
	"sort"
)

// 클로저 = 함수 내부에 선언된 함수가 외부 함수에서 선언한 변수 접근, 수정할 수 있는 것
func main3() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"John", "Doe", 30},
		{"Jane", "Smith", 50},
		{"Mike", "Smith", 40},
	}
	fmt.Println(people)

	// people은 closure에 의해 capture 되었다.
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

	// 함수에서 함수 반환
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
	fmt.Println(makeMult(4)(2))
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
