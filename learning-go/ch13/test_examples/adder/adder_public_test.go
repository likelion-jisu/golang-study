// _test를 붙임으로써 패키지를 '블랙 박스'로 취급한다.
package adder_test

import (
	"testing"
	// 같은 디렉터리에 있음에도 패키지가 달라 import한다.
	"test_examples/adder"
)

func TestAddNumbers(t *testing.T) {
	// 다른 패키지에서 노출된 함수를 호출하므로 adder.Adders를 사용한다.
	result := adder.AddNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
