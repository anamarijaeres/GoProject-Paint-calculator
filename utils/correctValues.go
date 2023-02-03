package utils

import (
	"fmt"
	"strconv"
)

func CheckIfPosInt(line string) int {
	if rooms, err := strconv.Atoi(line); err != nil {
		fmt.Printf("Looks like this is not a number :( \n")
	} else {
		if rooms <= 0 {
			fmt.Printf("I'm sorry but you can't enter a negative number or zero :( \n")
		} else {
			return rooms
		}
	}
	return -1
}
