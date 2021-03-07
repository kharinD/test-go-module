package lesson1

import (
	"fmt"
	"github.com/beevik/ntp"
)

func GetTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(time)
	}
}
