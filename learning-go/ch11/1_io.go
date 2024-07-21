package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main1() {
	s := "The quick brown fox jumps over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(counts)

	r, closer, err := buildGZipReader("my_data.txt.gz")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closer()
	counts, err = countLetters(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(counts)
}

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		// io.Reader한테 버퍼를 전달, 메모리 할당을 개발자의 통제 하에 둘 수 있다.
		n, err := r.Read(buf)
		// 얼마나 많은 바이트가 쓰여졌는지 n을 통해 알 수 있다.
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		// io.EOF - 읽기 완료되었다는 뜻
		if err == io.EOF {
			return out, nil
		}
		// 데이터 스트림의 끝이나 예상치 못한 조건에서 발생한 오류 전에 읽은 데이터가 있을 수 있으므로,
		// 오류 처리 전 데이터를 읽는다.
		if err != nil {
			return nil, err
		}
	}
}

// *gzip.Reader는 io.Reader를 구현한 것이기 때문에 countLetters 함수에 전달할 수 있다.
func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}
