package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := TimeNow()
	fmt.Println(currentTime)
}

func TimeNow() int64 {
	return time.Now().Unix()
}
