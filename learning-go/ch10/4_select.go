package main

import (
	"fmt"
)

/*
동일한 두개의 채널을 접근하는 두 고루틴이 있다면, 두 고루틴 내에서 같은 순서로 반드시 접근해야 교착 상태에 빠지지 않는다.
*/
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		// 고루틴은 ch1을 읽을 때까지 진행할 수 없다.
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()

	v := 2
	// main은 go 런타임에서 처음 시작할 때 동작하는 고루틴에서 실행된다.
	// main은 ch2를 읽을 때까지 진행할 수 없다.
	/*
		ch2 <- v
		v2 := <-ch1
		fmt.Println(v, v2)
	*/
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println(v, v2)
	fmt.Println("main done")
}
