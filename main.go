package main

import (
	"os"
	"time"

	"github.com/SamaraRuizSandoval/GolangAndAlgorithms/fundamentals/mocking"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}
