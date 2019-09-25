package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("just a text")

	duration := time.Since(start)

	fmt.Println(duration)
}
