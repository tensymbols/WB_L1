package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	select {
	case <-time.After(duration):
		fmt.Printf("%d ms has passed\n", duration.Milliseconds())
	}
}

func main() {
	then := time.Now().UnixMilli()
	mySleep(1 * time.Second)
	fmt.Println(time.Now().UnixMilli() - then)
}
