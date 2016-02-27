package main

import (
	"os/exec"
	"fmt"
	"time"
)

func pomodoro(focusC, breakC chan int, resetC, doneC chan bool) {
	var lastTime int
	timer := time.NewTimer(time.Duration(0))
	firstTime := true
	overtime := false
	overTimeAmount := 0
	var state string
	for {
		select {
		case focusTime := <-focusC:
			state = "focus"
			lastTime = focusTime
			overtime = false
			overTimeAmount = 0
			timer = time.NewTimer(minute(focusTime))
		case breakTime := <-breakC:
			state = "break"
			lastTime = breakTime
			overtime = false
			overTimeAmount = 0
			timer = time.NewTimer(minute(breakTime))

		case <-resetC:
			overtime = false
			overTimeAmount = 0
			fmt.Printf("Reseting %s back to %d mins \n", state, lastTime)
			timer = time.NewTimer(minute(lastTime))
		case <-timer.C:
			if firstTime {
				firstTime = false
				continue
			} else if overtime {
				overTimeAmount += 5
				fmt.Printf("Total overtime: %d mins \n", overTimeAmount)
				timer = time.NewTimer(minute(5))
			} else {
				overtime = true
				fmt.Printf("%d mins of %s has expired \n", lastTime, state)
				exec.Command("say", fmt.Sprintf("Your %s has expired! Do you want to change or reset this state?", state)).Output()
				timer = time.NewTimer(minute(5))
			}
		}
	}
}
