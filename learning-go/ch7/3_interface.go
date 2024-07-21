package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main3() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()

	var s *string
	fmt.Println(s == nil)

	var i interface{}
	fmt.Println(i == nil)
	i = s
	// 타입이 nil이 아닌 한 인터페이스는 nil이 아니다.
	fmt.Println(i == nil)

	// 빈 인터페이스는 어떤 것도 표현하지 않는다.
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName  string
	}{"John", "Doe"}

	data := map[string]interface{}{}
	contents, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer contents.Close()
	json.Unmarshal(contents, &data)

}

type Stringer interface {
	String() string
}

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// business logic
	return "processed: " + data
}

// 인터페이스는 호출자가 무엇을 원하는지 명시한다.
// 사용자 코드는 그것이 필요한 기능이 무엇인지 명시하기 위해 인터페이스를 정의한다.
type Logic3 interface {
	Process(data string) string
}

// Client가 interface를 알고있다.
type Client struct {
	L Logic3
}

func (c Client) Program() {
	// get data from somewhere
	result := c.L.Process("data")
	fmt.Println(result)
}

type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

// go에 제네릭이 없어서 interface{}를 사용한다.
func (ll *LinkedList) Inserts(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		}
	}
	ll.Next = ll.Next.Inserts(pos-1, val)
	return ll
}
