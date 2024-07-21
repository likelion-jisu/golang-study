package main

import "fmt"

type MyInt int

func main4() {
	var i interface{}
	var mine MyInt = 20
	i = mine

	// i2 = MyInt type
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)

	// panic: interface conversion: interface {} is main.MyInt, not string
	// i3 := i.(string)
	// fmt.Println(i3)

	// panic: interface conversion: interface {} is main.MyInt, not int
	// i3 := i.(int)
	// fmt.Println(i3 + 1)

	i3, ok := i.(int)
	if !ok {
		fmt.Println("unexcpeted type for ", i3) // i3 is 0
		return
	}
	fmt.Println(i3 + 1)
}
