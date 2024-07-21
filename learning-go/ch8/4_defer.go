package main

import "fmt"

func DoSomeThings(val1 int, val2 string) (string, error) {
	_, err := doThing(val1)
	if err != nil {
		return "", fmt.Errorf("in DoSomethings, :w", err)
	}
	_, err = doThing2(val1)
	if err != nil {
		return "", fmt.Errorf("in DoSomethings, :w", err)
	}

	result, err := doThing3(val1)
	if err != nil {
		return "", fmt.Errorf("in DoSomethings, :w", err)
	}
	return result, nil
}

// defer를 사용하여 코드를 간결하게 만들 수 있다.
func DoSomeThings2(val1 int, val2 string) (_ string, err error) {
	// 오류가 반환되었다면 새로운 오류를 재할당해준다.
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomethings, :w", err)
		}
	}()

	_, err = doThing(val1)
	if err != nil {
		return "", err
	}
	_, err = doThing2(val1)
	if err != nil {
		return "", err
	}

	return doThing3(val1)
}

func doThing(val int) (string, error) {
	return "", nil
}
func doThing2(val int) (string, error) {
	return "", nil
}
func doThing3(val int) (string, error) {
	return "", nil
}
