package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	repetitions := flag.Int("r", 1, "The amount of pomodoro repetitions")
	flag.Parse()

	if *repetitions <= 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}	

	for i := 0; i < *repetitions; i++ {
		pomodoro()
	}
}

func pomodoro() {
	for i := 0; i < 3; i++ {
		countdownTimer(25, afterWork)
		countdownTimer(5, afterBreak)
	}
	countdownTimer(25, afterWork)
	countdownTimer(15, afterBreak)
}

func countdownTimer(minutes time.Duration, messageAfterTimeIsUp string) {
	timeAfter := time.Now().Add(time.Minute * minutes)

	for range time.Tick(1 * time.Second) {
		timeRemaining := getTimeRemaining(timeAfter)

		if timeRemaining.t <= 0 {
			fmt.Println(messageAfterTimeIsUp)
			break
		}

		fmt.Printf("\r%02d:%02d", timeRemaining.m, timeRemaining.s)
	}
}

type countdown struct {
	t int
	m int
	s int
}


func getTimeRemaining(t time.Time) countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		m: minutes,
		s: seconds,
	}
}


const afterWork string = "\nWork time is over!\n"
const afterBreak string = "\nBreak time is over!\n"