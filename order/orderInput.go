package order

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	m "tsi.co/mod/go-paintProject/menu"
	u "tsi.co/mod/go-paintProject/utils"
)

func InputRoomNum(scanner *bufio.Scanner, message string) int {

	for scanner.Scan() {

		line := scanner.Text()
		line = strings.TrimSpace(line)

		rooms := u.CheckIfPosInt(line)
		if rooms > 0 {
			return rooms
		}
		fmt.Println(message)
	}
	return 0
}

func ChangeRoomNum(scanner *bufio.Scanner, roomsOld int) int {
	rooms := roomsOld
	for scanner.Scan() {
		l := scanner.Text()
		l = strings.TrimSpace(l)
		l = strings.ToLower(l)
		if l != "y" && l != "n" {
			fmt.Println("Didn't quite catch that, try again.")
		} else {
			if l == "y" {
				message := "Please enter the number of rooms you want to paint:"
				fmt.Println(message)
				rooms = InputRoomNum(scanner, message)
				m.DisplayDiscounts(rooms)
			} else {
				return rooms
			}
		}
		fmt.Println("Would you like to change the number of rooms(y/n)?")
	}
	return 0
}

func InputSize(scanner *bufio.Scanner, message string) float64 {
	for scanner.Scan() {

		line := scanner.Text()
		line = strings.TrimSpace(line)

		if width, err := strconv.ParseFloat(line, 64); err != nil {
			fmt.Printf("Looks like this is not a number :( \n")
		} else {
			if width <= 0 {
				fmt.Printf("I'm sorry but you can't enter a negative number for width/height or zero :( \n")
			} else {
				return width
			}
		}
		fmt.Println(message)
	}
	return 0
}
