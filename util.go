package main

import (
	"time"
	"bufio"
"strings"
	"errors"
	"strconv"
	"fmt"
	"os"
)

const commands string = `
Enter "help" to see this list again:

	focus [mins] - set/reset the focus duration
	break [mins] - take a break
	reset - reset the duration of the current state: focus or break

`

func clientInput(focusC, breakC chan int, resetC, doneC chan bool) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input := readSantizedString(reader)
		command, value,err := splitUserInput(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if command == "focus" {
			focusC<-value
		} else if command == "break" {
			breakC<-value
		} else if command == "reset" {
			resetC<-true
		} else if command == "help" {
			fmt.Println(commands)
		}
	}
}

func minute(d int) time.Duration {
	return time.Duration(d) * time.Minute
}

func readSantizedString(reader *bufio.Reader) (input string){
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.Replace(input, "\n", "", -1)
	input = strings.ToLower(input)
	return
}

func splitUserInput(input string) (command string, value int, err error) {
	s := strings.Split(input, " ")
	if len(s) > 2 || len(s) < 1 || (len(s) == 2 && s[1] == ""){
		return command, value, errors.New("Input must be in the form COMMAND [duration - optional]")
	}

	command = s[0]
	if len(s) == 2 {
		value, err = strconv.Atoi(s[1])
	}
	return command, value, err
}
