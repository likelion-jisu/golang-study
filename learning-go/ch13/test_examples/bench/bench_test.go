package bench

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	makeData()
	exitVal := m.Run()
	os.Remove("testdata/data.txt")
	os.Exit(exitVal)
}

// makeData makes our data file for us. Rather than checking in a large file, we recreate it for the test.
// By setting the random seed to the same value every time, we ensure that we generate the same file every time.
// This random seed generates a file that's 65,204 bytes long.
func makeData() {
	file, err := os.Create("testdata/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rand.Seed(1)
	for i := 0; i < 10000; i++ {
		data := makeWord(rand.Intn(10) + 1)
		file.Write(data)
	}
}

func makeWord(l int) []byte {
	out := make([]byte, l+1)
	for i := 0; i < l; i++ {
		out[i] = 'a' + byte(rand.Intn(26))
	}
	out[l] = '\n'
	return out
}

func TestFileLen(t *testing.T) {
	result, err := FileLen("testdata/data.txt", 1)
	if err != nil {
		t.Fatal(err)
	}
	if result != 65204 {
		t.Error("Expected 65204, got", result)
	}
}

/*
컴파일러는 코드 최적화를 통해 불필요한 연산을 제거하려고 시도합니다.
예를 들어, 만약 FileLen 함수 호출 결과가 이후 코드에서 사용되지 않는다면, 컴파일러는 이를 최적화 과정에서 제거할 수 있습니다.
그러나 벤치마크 테스트의 목적은 실제 함수 호출 시간을 측정하는 것이므로, 이러한 최적화는 부적절합니다.

이를 방지하기 위해 result 값을 blackhole 변수에 할당합니다.
이 변수는 패키지 레벨에서 선언되었기 때문에, 컴파일러는 FileLen 함수 호출을 최적화하여 제거하지 못합니다.
따라서 실제로 함수가 호출되고 그 실행 시간이 정확하게 측정될 수 있습니다.
*/
var blackhole int

func BenchmarkFileLen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := FileLen("testdata/data.txt", 1)
		if err != nil {
			b.Fatal(err)
		}
		blackhole = result
	}
}

func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := FileLen("testdata/data.txt", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
