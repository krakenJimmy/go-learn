package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	date := "1999-12-31"
	t, _ := time.Parse(layoutISO, date)
	fmt.Println(t)
	fmt.Println(t.Format(layoutUS))
}
