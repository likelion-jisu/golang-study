package main

import (
	"context"
	"fmt"
	"net/http"
)

type guidKey int

const key2 guidKey = 1

func contextWithGUID(ctx context.Context, guid string) context.Context {
	return context.WithValue(ctx, key2, guid)
}

func guidFromContext(ctx context.Context) (string, bool) {
	g, ok := ctx.Value(key).(string)
	return g, ok
}

// 요청에서 GUID를 추출 혹은 새로운 GUID를 생성한다.
func Middleware3(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		if guid := req.Header.Get("X-GUID"); guid != "" {
			ctx = contextWithGUID(ctx, guid)
		} else {
			ctx = contextWithGUID(ctx, "00000000-0000-0000-0000-000000000000")
		}
		// 업데이트된 컨텍스트로 새 요청 만든 다음 진행한다.
		// 이렇게 하면 메모리 너무 많이 쓰지 않을까?
		req = req.WithContext(ctx)
		h.ServeHTTP(rw, req)
	})
}

type Logger struct{}

func (Logger) Log(ctx context.Context, message string) {
	if guid, ok := guidFromContext(ctx); ok {
		message = guid + ": " + message
	}
	// do logging
	fmt.Println(message)
}

func Request(req *http.Request) *http.Request {
	ctx := req.Context()
	if guid, ok := guidFromContext(ctx); ok {
		req.Header.Add("X-GUID", guid)
	}
	return req
}
