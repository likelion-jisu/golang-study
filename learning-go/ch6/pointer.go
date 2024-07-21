package main

import "fmt"

func main() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX) // 역참조, dereferencing
	z := 5 + *pointerToX
	fmt.Println(z)

	var y *int
	fmt.Println(y == nil)
	// fmt.Println(*y) -> 패닉 (SIGSEGV: segmentation violation)

	// new 함수는 드물게 사용된다.
	// 기본 타입을 위한 포인터가 필요하다면 변수를 선언, 해당 변수를 가리키도록 하자.
	a := &Foo{}
	var b string
	c := &b
	fmt.Println(a, b, c)

	/*
		p := person{
			FirstName:  "John",
			MiddleName: "Quincy", // cannot use "Quincy" (untyped string constant) as *string value in struct literal
			LastName:   "Doe",
		}
	*/
	p := person{
		FirstName:  "John",
		MiddleName: stringp("Quincy"),
		LastName:   "Doe",
	}
	fmt.Println(p)

	var f *int
	failedUpdate(f)
	fmt.Println(f) // nil

	// 포인터를 변경하면 값이 바뀌지 않는다. 역참조해서 값을 설정해야 바뀐다.
	x2 := 20
	failedUpdate(&x2)
	fmt.Println(x2)
	update(&x2)
	fmt.Println(x2)
}

// 상수를 함수로 전달했을 때 파라미터로 복사되면서 주소를 가진다.
func stringp(s string) *string {
	return &s
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func update(px *int) {
	*px = 30
}

type Foo struct {
	bar string
}

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

// don't do this
func MakeFoo(f *Foo) error {
	f.bar = "hello"
	return nil
}

// do this
// 구조체 전달을 포인터로 해서 항목을 채우는 것 보다, 함수 내에서 구조체를 초기화하고 반환하는게 좋다.
func MakeFoo2() (Foo, error) {
	f := Foo{
		bar: "hello",
	}
	return f, nil
}
