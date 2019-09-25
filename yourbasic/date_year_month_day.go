package main

import (
	"fmt"
	"time"
)

// The Date function returns the year, month and day of a time.TIme.
// func (t Time) Date() (year int, month Month, day int)
func main() {
	year, month, day := time.Now().Date()
	fmt.Println(year, month, day)
	fmt.Println(year, int(month), day)

	fmt.Println("----------------")

	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
}
