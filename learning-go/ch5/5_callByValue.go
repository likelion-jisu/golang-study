package main

import "fmt"

type person struct {
	age  int
	name string
}

func main() {
	// go는 함수 파라미터로 넘겨지는 변수가 있다면, 항상 복사본을 만들어서 넘긴다.
	// 따라서 함수 내부에서 변수를 변경해도 원본에는 영향을 미치지 않는다. (구조체도 동일)
	p := person{}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p) // 2 Hello {0 }

	// 맵이나 슬라이스는 참조 타입이므로 함수 내부에서 변경하면 원본에 영향을 미친다.
	// 슬라이스에 길이를 늘리는 것은 안된다. (포인터로 구현되어 있기 때문이다.)

	m := map[int]string{1: "one", 2: "two"}
	modMap(m)
	fmt.Println(m) // map[2:hello 3:goodbye]
	s2 := []int{1, 2, 3}
	modSlice(s2)
	fmt.Println(s2) // [2 4 6]
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}
