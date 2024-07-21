package main

import (
	"fmt"
	"time"
)

func main2() {
	// 기간 - time.Duration으로 표시
	// d is time.Duration type
	d := 2*time.Hour + 30*time.Minute
	fmt.Println(d)

	// 순간 - time.Time으로 표시
	t, err := time.Parse("2006-02-01 15:04:05 -0700", "2016-13-03 00:00:00 +0000")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
}
