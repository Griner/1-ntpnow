package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func PrintTime(description string, t time.Time) {
	fmt.Printf("%-20s %s\n", description, t)
}

func main() {

	PrintTime("Current local time", time.Now())

	t, err := ntp.Time("0.debian.pool.ntp.org")
	if err != nil {
		fmt.Errorf("NTP error: %v\n", err)
	} else {
		PrintTime("NTP time", t)
	}
}
