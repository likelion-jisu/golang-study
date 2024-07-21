package main

import (
	"context"
	"database/sql"
	"io"
	"log"
	"os"
)

// defer = 임시 자원 정리에 사용
// 함수 실행시 해당 부분을 함수 종료되기 직전에 실행시킨다.
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer f.Close()

	var result []byte
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		result = append(result, data[0:count]...)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
				return "", err // f.Close() 호출
			}
			break
		}
	}
	return string(result), nil // f.Close() 호출
}

func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		err = tx.Commit()
	}()
	_, err = tx.ExecContext(ctx, "INSERT INTO table1 (val) VALUES (?)", value1)
	if err != nil {
		return err
	}
	// 여기서 DB 삽입 연산이나 이후의 추가 처리를 위해 tx 변수 사용
	return nil
}

func example() {
	defer func() int {
		return 10 // 해당 값을 읽을 방법이 없다.
	}()
}
