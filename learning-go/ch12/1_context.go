package main

import (
	"context"
	"fmt"
	"net/http"
)

func main1() {
	ctx := context.Background()
	result, err := logic(ctx, "a string")
	fmt.Println(result, err)
}
func logic(ctx context.Context, info string) (string, error) {
	return "", nil
}

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 요청에서 기존 컨텍스트 추출
		ctx := r.Context()
		// 이전 요청과 현재 생성된 context를 기반으로 새 요청을 만든다.
		r = r.WithContext(ctx)
		// 새로운 요청 + 기존 ResponseWriter를 전달하여 호출한다.
		handler.ServeHTTP(w, r)
	})
}

func handler(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	err := req.ParseForm()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	data := req.FormValue("data")
	result, err := logic(ctx, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte(result))
}

// 응용 프로그램에서 다른 http 서비스로 호출하는 경우 WithContext를 사용하여 새 컨텍스트를 만들어야 한다.
type ServiceCaller struct {
	client *http.Client
}

func (sc ServiceCaller) callAnotherService(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)
	res, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	// 응답 처리를 위한 나머지 부분
	return "", nil
}
