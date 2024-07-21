package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"os"
)

// 관용적으로 마지막 반환값으로 error를 반환한다.
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main1() {
	numerator := 20
	denominator := 3
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)

	// Exception이 없다. if문 사용해서 오류 변수가 nil이 아닌지 확인 필요
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder, mod)

	// sentinel error
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err = zip.NewReader(notAZipFile, int64(len(data)))
	if errors.Is(err, zip.ErrFormat) {
		fmt.Println("This is not a zip file")
	}

	// 기본 타입이 nil이 아닌 한, 인터페이스도 nil이 아니다.
	err2 := GenerateError(false)
	fmt.Println(err2 != nil) // true
	err2 = GenerateError(true)
	fmt.Println(err2 != nil) // true
}

// error is value
type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func Login() ([]byte, error) {
	return nil, StatusErr{
		Status:  InvalidLogin,
		Message: "invalid credentials for user",
	}
}

func GetData() ([]byte, error) {
	return nil, StatusErr{
		Status:  NotFound,
		Message: "User not found",
	}
}

// 자신만의 오류 타입을 사용한다면, 초기화되지 않은 인스턴스를 반환하지 않도록 하자.
// 이것은 사용자 오류의 타입이 되게 하기 위해 변수를 선언하지 않고 변수를 반환한다는 의미이다.
func GenerateError(flat bool) error {
	var genErr StatusErr
	if flat {
		genErr = StatusErr{
			Status: InvalidLogin,
		}
	}
	return genErr
}

func GenerateError2(flat bool) error {
	if flat {
		return StatusErr{
			Status: InvalidLogin,
		}
	}
	return nil
}
