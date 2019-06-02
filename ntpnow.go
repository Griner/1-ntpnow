package main

import (
	"fmt"
	"os"
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
		fmt.Printf("NTP error: %v\n", err)
		os.Exit(1)
	} else {
		PrintTime("NTP time", t)
	}
}
