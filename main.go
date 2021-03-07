package main

import (
	"fmt"
	"github.com/shourshort/test-go-module/lesson1"
	"os"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	var a uintptr = 12

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(time)
	}

	os.Exit(0)
}
