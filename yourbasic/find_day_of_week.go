// The Weekday function returns the day of the week of a time.Time
package main

import (
	"fmt"
	"time"
)

func main() {
	weekday := time.Now().Weekday()
	fmt.Println(weekday)
	fmt.Println(int(weekday))
}
