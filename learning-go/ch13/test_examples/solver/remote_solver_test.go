package solver

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteSolver_Resolve(t *testing.T) {
	type info struct {
		expression string
		code       int
		body       string
	}
	// 함수에 전달된 데이터가 서버에 도착하는지 확인하고 싶다.
	// 입/출력 저장을 위해 io 변수를 사용한다.
	var io info
	// 가상의 원격 서버를 설정한다.
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			expression := req.URL.Query().Get("expression")
			if expression != io.expression {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte("invalid expression: " + io.expression))
				return
			}
			rw.WriteHeader(io.code)
			rw.Write([]byte(io.body))
		}))
	defer server.Close()
	rs := RemoteSolver{
		MathServerURL: server.URL,
		Client:        server.Client(),
	}
	data := []struct {
		name   string
		io     info
		result float64
		errMsg string
	}{
		{"case1", info{"2 + 2 * 10", http.StatusOK, "22"}, 22, ""},
		{"case2", info{"( 2 + 2 ) * 10", http.StatusOK, "40"}, 40, ""},
		{"case3", info{"( 2 + 2 * 10", http.StatusBadRequest,
			"invalid expression: ( 2 + 2 * 10"},
			0, "invalid expression: ( 2 + 2 * 10"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			io = d.io
			result, err := rs.Resolve(context.Background(), d.io.expression)
			if result != d.result {
				t.Errorf("io `%f`, got `%f`", d.result, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("io error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
