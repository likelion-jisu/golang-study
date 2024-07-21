package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

func main3() {
	err := fileChecker2("not_here.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
		}
	}

	if errors.Is(err, ResourceErr{Resource: "Database"}) {
		fmt.Println("The database is broken:", err)
		// do something
	}

}

func fileChecker2(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

type MyErr struct {
	Codes []int
}

func (e MyErr) Error() string {
	return fmt.Sprintf("codes: %v", e.Codes)
}

// 오류가 비교 불가능한 타입일 경우 Is 메서드 직접 구현 가능하다.
func (e MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return reflect.DeepEqual(e, me2)
	}
	return false
}

// 동일한 인스턴스가 아닌 오류에 대해 비교 가능하다.
type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("resource: %s, code: %d", re.Resource, re.Code)
}

func (re ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code
		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}
	return false
}
