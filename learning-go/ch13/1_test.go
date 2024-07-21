package ch13

import (
	"fmt"
	"os"
	"testing"
)

func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func TestFileProcessing(t *testing.T) {
	fName, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fName)
	// 테스트를 수행, 정리를 걱정하지 않아도 된다.

}
