package order

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	m "tsi.co/mod/go-paintProject/menu"
	u "tsi.co/mod/go-paintProject/utils"
)

func InputRoomNum(scanner *bufio.Scanner) int {

	for scanner.Scan() {

		line := scanner.Text()
		line = strings.TrimSpace(line)

		rooms := u.CheckIfPosInt(line)
		if rooms > 0 {
			return rooms
		}
		fmt.Println("Please enter the number of rooms you want to paint:")
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
				fmt.Println("Please enter the number of rooms you want to paint:")
				rooms = InputRoomNum(scanner)
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

func inputDoorNum(scanner *bufio.Scanner) int64 {
	for scanner.Scan() {

		line := scanner.Text()
		line = strings.TrimSpace(line)

		doors := u.CheckIfPosInt(line)
		if doors > 0 {
			return int64(doors)
		}
		fmt.Println("Could you please give me the number of doors in the room?")
	}
	return 0

}
