package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(w io.Writer, sleeper Sleeper) {

	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprint(w, "Go!")
}

func main() {

	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
