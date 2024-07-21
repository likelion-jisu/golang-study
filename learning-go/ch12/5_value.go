package main

import (
	"context"
	"net/http"
)

type userKey int

const key userKey = 1

func ContextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, key, user)
}

func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(key).(string)
	return user, ok
}

func extractUser(req *http.Request) (string, error) {
	userCookie, err := req.Cookie("user")
	if err != nil {
		return "", err
	}
	return userCookie.Value, nil
}

func Middleware2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// 사용자 값을 얻는다.
		user, err := extractUser(req)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		// 컨텍스트 추출
		ctx := req.Context()
		// 사용자 포함 새 컨텍스트 생성
		ctx = ContextWithUser(ctx, user)
		// 이전 요청과 새로운 컨텍스트로 새로운 요청을 만든다.
		req = req.WithContext(ctx)
		// 핸들러 체인의 다음 함수를 호출한다.
		h.ServeHTTP(rw, req)
	})
}
