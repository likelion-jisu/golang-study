package main

import (
	"fmt"
	"time"
)

func main1() {
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	output := p.String()
	fmt.Println(output)

	var c Counter
	fmt.Println(c.String())
	// c가 값 타입임에도 포인터 리시버 메서드 호출 가능하다.
	// Go는 자동으로 지역 변수를 포인터 타입으로 변환한다.
	// c.Increment()은 (&c).Increment()로 변환된다.
	c.Increment()
	fmt.Println(c.String())

	doUpdateWrong(c)
	fmt.Println("in main:", c.String())
	doUpdateRight(&c)
	fmt.Println("in main:", c.String())

	// nil 인스턴스를 위한 메서드 작성
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(7)
	it = it.Insert(2)
	fmt.Println(it.Contains(3))
	fmt.Println(it.Contains(10))

	// method is function
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))

	// method value
	// 메서드 값은 생성된 인스턴스 항목이 있는 값에 접근할 수 있다. -> 클로저와 유사하다.
	f1 := myAdder.AddTo
	fmt.Println(f1(10))

	// method expression
	// type 자체로 함수를 생성할 수 있다.
	// 함수 시그니처 : func(Adder, int) int -> 첫 번째 파라미터는 메서드를 위한 리시버이다.
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 15))

	// 타입 선언은 상속되지 않는다.
	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	hs = HighScore(s) // hs = s -> 타입이 다르기 때문에 할당 불가능
	s = Score(i)      // s = i -> error
	fmt.Println(hs, s)
}

// 패키지 블록 ~ 모든 레벨에서 타입 선언 가능
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Score int
type Converter func(string) Score
type TeamScores map[string]Score

// 타입을 위한 메서드는 패키지 블록 레벨에서 정의한다.
// 함수 선언 + 리시버를 명시해야 한다.
// 오버로딩 불가 (같은 타입의 다른 메서드를 동일한 이름으로 사용할 수 없다.)
func (p Person) String() string {
	return fmt.Sprintf("%s %s: %d", p.FirstName, p.LastName, p.Age)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

type HighScore Score

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)
