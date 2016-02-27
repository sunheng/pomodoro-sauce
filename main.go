package main

import "fmt"

func main() {

	focusC := make(chan int)
	breakC := make(chan int)
	resetC := make(chan bool)
	doneC := make(chan bool)

	fmt.Println(commands)

	go clientInput(focusC, breakC, resetC, doneC)
	go pomodoro(focusC, breakC, resetC, doneC)

	<-doneC
}