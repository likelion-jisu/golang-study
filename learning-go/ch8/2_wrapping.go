package main

import (
	"errors"
	"fmt"
	"os"
)

func main2() {
	err := fileChecker("not-exist.txt")
	if err != nil {
		// in fileChecker: open not-exist.txt: no such file or directory
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			// open not-exist.txt: no such file or directory
			fmt.Println(wrappedErr)
		}
	}

	// 오류 생성하지만 래핑하고 싶지 않다면 fmt.Errorf에 %w 대신 %v를 사용하자.
	originalError := errors.New("original error")

	wrappedError := fmt.Errorf("wrapped error: %w", originalError)
	fmt.Println(wrappedError)
	fmt.Println(errors.Unwrap(wrappedError)) // original error

	notWrappedError := fmt.Errorf("wrapped error: %v", originalError)
	fmt.Println(notWrappedError)
	fmt.Println(errors.Unwrap(notWrappedError)) // nil
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

type StatusErr2 struct {
	Status  Status
	Message string
	err     error
}

func (se StatusErr2) Error() string {
	return se.Message
}

func (se StatusErr2) Unwrap() error {
	return se.err
}

func Login2() ([]byte, error) {
	return nil, StatusErr2{
		Status:  InvalidLogin,
		Message: "invalid login",
		err:     errors.New("invalid login"),
	}
}

func GetData2() ([]byte, error) {
	return nil, StatusErr2{
		Status:  NotFound,
		Message: "User not found",
		err:     errors.New("invalid login"),
	}
}
