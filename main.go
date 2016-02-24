package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

const commands string = `
Enter "help" to see this list again:

	focus [mins] - set/reset the focus duration
	break [min] - take a break
	reset - reset the duration of the current state: focus or break

`

const focusPrompt string = `Enter the desired FOCUS duration`
const breakPrompt string = `Enter the desired BREAK duration`

const minimumMinErr string = `Minutes must be greater than 0`


func main() {
	fmt.Println(commands)

	//focusMin, breakMin := initInteraction()


}

func initInteraction() (focusMin, breakMin int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(focusPrompt)
		input, err := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		if err != nil {
			continue
		}
		focusMin, err = strconv.Atoi(input)
		if err != nil || focusMin < 1{
			fmt.Println(minimumMinErr)
			continue
		}
		break
	}

	for {
		fmt.Println(breakPrompt)
		input, err := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		if err != nil {
			continue
		}
		breakMin, err = strconv.Atoi(input)
		if err != nil || breakMin < 1{
			fmt.Println(minimumMinErr)
			continue
		}
		break
	}

	return
}