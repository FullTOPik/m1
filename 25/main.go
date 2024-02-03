package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	<-time.After(d * time.Millisecond)
}

func main() {
	fmt.Println(time.Now())

	Sleep(5000)

	fmt.Println(time.Now())
}
