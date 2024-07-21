package main

import "fmt"

func main1() {
	ch := make(chan int)
	b := 3
	// fatal error: all goroutines are asleep - deadlock!
	// 채널(ch)이 값을 읽는 부분과 값을 쓰는 부분이 메인 고루틴(main goroutine)에서 동시에 실행되지 않기 때문입니다.
	ch <- b   // b의 값을 ch에 쓴다.
	a := <-ch // ch에서 값을 읽어 a에 할당한다.

	bufferedCh := make(chan int, 10) // 버퍼가 있는 채널

	// 다른 for-range 루프와 다르게 단일 변수만 있다.
	// break, return 문에 도달할 때까지 지속된다.
	for v := range ch {
		fmt.Println(v)
	}

	for v := range bufferedCh {
		fmt.Println(v)
	}

	fmt.Println("a:", a)
}
