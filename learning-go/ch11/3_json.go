package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main3() {
	data := `{
		"id": "1234",
		"items": [],
		"customer_id": "5678"
	}`
	var o Order
	// io.Reader와 마찬가지로 구조체를 재사용, 메모리 제어 가능하다.
	// go는 제네릭이 없어서 어떤 타입을 인스턴스화 해야하는지 지정할 방법이 없다.
	err := json.Unmarshal([]byte(data), &o)
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := json.Marshal(o)
	fmt.Println("hi")
	fmt.Println(out)

	toFile := Person{
		Name: "John",
		Age:  30,
	}

	tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(tmpFile.Name())
	err = json.NewEncoder(tmpFile).Encode(toFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tmpFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	var fromFile Person
	err = json.NewDecoder(tmpFile2).Decode(&fromFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tmpFile2.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n\n", fromFile)

	// json stream의 encoding, decoding
	const data2 = `
		{"name": "Fred", "age": 40}
		{"name": "Mary", "age": 21}
		{"name": "Pat", "age": 30}
	`
	var t struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	dec := json.NewDecoder(strings.NewReader(data2))
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t)
		err = enc.Encode(t)
		if err != nil {
			panic(err)
		}
	}
	out2 := b.String()
	fmt.Println(out2)

	const jsonStream = `
	[
		{"Name": "Ed", "Text": "Knock knock."},
		{"Name": "Sam", "Text": "Who's there?"},
		{"Name": "Ed", "Text": "Go fmt."},
		{"Name": "Sam", "Text": "Go fmt who?"},
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	]
`
	type Message struct {
		Name, Text string
	}
	dec2 := json.NewDecoder(strings.NewReader(jsonStream))

	// read open bracket
	t2, err := dec2.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t2, t2)

	// while the array contains values
	for dec2.More() {
		var m Message
		// decode an array value (Message)
		err := dec2.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v: %v\n", m.Name, m.Text)
	}

	// 사용자 지정 JSON parsing
	data3 := `
	{
		"id": "12345",
		"items": [
			{
				"id": "xyz123",
				"name": "Thing 1"
			},
			{
				"id": "abc789",
				"name": "Thing 2"
			}
		],
		"date_ordered": "01 May 20 13:01 +0000",
		"customer_id": "3"
	}`

	var o2 Order
	err4 := json.Unmarshal([]byte(data3), &o2)
	if err4 != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o2)
	fmt.Println(o2.DateOrdered.Month())
	out, err5 := json.Marshal(o2)
	if err5 != nil {
		panic(err5)
	}
	fmt.Println(string(out))
}

// go vet - 포맷 검증 가능
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Order struct {
	ID          string      `json:"id"`
	Items       []Item      `json:"items"`
	DateOrdered RFC822ZTime `json:"date_ordered"`
	CustomerID  string      `json:"customer_id"`
}

type RFC822ZTime struct {
	time.Time
}

func (rt RFC822ZTime) MarshalJSON() ([]byte, error) {
	out := rt.Time.Format(time.RFC822Z)
	return []byte(`"` + out + `"`), nil
}

func (rt *RFC822ZTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	t, err := time.Parse(`"`+time.RFC822Z+`"`, string(b))
	if err != nil {
		return err
	}
	*rt = RFC822ZTime{t}
	return nil
}
