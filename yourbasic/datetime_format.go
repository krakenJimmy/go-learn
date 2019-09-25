// Reference: https://programming.guide/go/format-parse-string-time-date-example.html
package main

import (
	"fmt"
	"time"
)

func main() {
	timeStr := "2019-10-01 10:00:00"
	t, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(timeStr)
	fmt.Println("now the time is ", time.Now().Format("2006-01-02 15:04:05"))
}
