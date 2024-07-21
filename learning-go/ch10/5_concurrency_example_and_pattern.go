package main

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main5() {
	// 고루틴에서 값 캡처하기
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		// 이렇게 하면 안된다. 모든 고루틴을 위한 클로저가 같은 변수를 캡처했기 때문이다.
		go func() {
			ch <- v * 2
		}()
	}

	for _, v := range a {
		// 방법 1 : 루프 내에서 값을 섀도잉하고 클로저에 전달한다.
		v := v
		go func() {
			ch <- v * 2
		}()
	}

	for _, v := range a {
		// 방법 2 : 파라미터로 값을 고루틴으로 넘긴다.
		v := v
		go func(val int) {
			ch <- val * 2
		}(v)
	}

	for i := 0; i < 15; i++ {
		fmt.Println(<-ch)
	}

	// 취소 함수 사용
	ch2, cancel := countTo(10)
	for i := range ch2 {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	cancel()

	// 여러 고루틴을 기다려야 할 때
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		doThingThatShouldBeLimited()
	}()
	go func() {
		defer wg.Done()
		doThingThatShouldBeLimited()
	}()
	go func() {
		defer wg.Done()
		doThingThatShouldBeLimited()
	}()
	wg.Wait()

	// 배압
	pg := New(10)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too Many Requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}

// Done 채널 패턴
// 여러 검색 함수를 동시에 실행하여 병렬 처리를 하고, 가장 빠르게 결과를 반환하는 방식으로 동작한다.
func searchData(s string, searchers []func(string) []string) []string {
	// 값이 중요하지 않기 때문에 타입으로 빈 구조체를 사용한다.
	// 이 채널은 아무 것도 쓰지 않을 것이고 단지 닫기만 할 것이다.
	done := make(chan struct{})
	result := make(chan []string)
	for _, searcher := range searchers {
		go func(searcher func(string) []string) {
			select {
			case result <- searcher(s):
			case <-done:
			}
		}(searcher)
	}
	r := <-result
	close(done)
	return r
}

// 취소 함수 사용
// done 채널을 직접 반환하는 것보다 done 채널을 닫는 클로저를 반환하는 것이 더 낫다.
func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			default:
				ch <- i
			}
		}
		close(ch)
	}()
	return ch, cancel
}

// 버퍼가 있는 채널의 사용
// 얼마나 많은 고루틴이 실행될지 알고 있을 때, 실행시킬 고루틴의 개수/대기중인 작업의 양을 제한하는 경우에 유용하다.
func processChannel(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func() {
			for v := range ch {
				results <- v * 2
			}
		}()
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

// 배압
type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{ch: ch}
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	case <-pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	// 타임아웃
	case <-time.After(2 * time.Second):
		return errors.New("work timed out")
	default:
		return errors.New("no more capacity")
	}
}

func doThingThatShouldBeLimited() string {
	time.Sleep(2 * time.Second)
	return "done"
}

// 정확히 한 번만 코드 실행하기
type SlowComplecatedParser interface {
	Parse(string) string
}

var parser SlowComplecatedParser
var once sync.Once

/*
func Parse(dataToParse string) string {
	// Parse가 한 번 이상 호출된다면, once.Do()는 더 이상 실행되지 않는다.
	once.Do(func() {
		parser = initParser()
	})
	return parser.Parse(dataToParse)
}

func initParser() SlowComplecatedParser {
	// 설정/로딩
	return &slowComplecatedParserImpl{}
}
*/
