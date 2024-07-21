package main

import "context"

func longRunningThingManager(ctx context.Context, data string) (string, error) {
	type wrapper struct {
		result string
		err    error
	}
	// 버퍼 크기가 1인 타입이 wrapper인 채널을 만든다.
	// 채널을 버퍼링하면 버퍼링된 값이 취소로 인해 읽히지 않더라도 고루틴이 종료될 수 있다.
	ch := make(chan wrapper, 1)
	go func() {
		// do the long running thing
		result, err := longRunningThing(ctx, data)
		ch <- wrapper{result, err}
	}()
	select {
	// 함수 실행/타임아웃으로 인해 컨텍스트가 취소
	case data := <-ch:
		return data.result, data.err
		// 컨텍스트가 취소
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func longRunningThing(ctx context.Context, data string) (string, error) {
	return "", nil
}
