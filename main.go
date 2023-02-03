package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	m "tsi.co/mod/go-paintProject/menu"
	o "tsi.co/mod/go-paintProject/order"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Oops try again")
			main()
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	rooms := 0

	//entering the number of rooms
	message := "Please enter the number of rooms you want to paint:"
	fmt.Println(message)
	rooms = o.InputRoomNum(scanner, message)

	//look at discount menu
	fmt.Println("Great!Let's look at the possible discounts for the quantity:")
	discount := m.DisplayDiscounts(rooms)

	//offer the change of the number of rooms after displaying the discount
	fmt.Println("Would you like to change the number of rooms(y/n)?")
	rooms = o.ChangeRoomNum(scanner, rooms)

	surfaceOfRooms := []float64{}
	//iterate over the number of rooms to get the room sizes
	for i := 1; i <= rooms; i++ {

		message = fmt.Sprintf("Could you please give me the width of all room %d walls together in meters?", i)
		fmt.Println(message)
		width := o.InputSize(scanner, message)

		message = fmt.Sprintf("Could you please give me the height of room %d in meters?", i)
		fmt.Println(message)
		height := o.InputSize(scanner, message)

		surfaceW := width * height
		fmt.Println("Great! The size of the first room is " + strconv.FormatInt(int64(surfaceW), 10) + " square meters.")

		//get the number of doors in the room
		message := "Could you please give me the number of doors in the room?"
		fmt.Println(message)
		doorNum := o.InputRoomNum(scanner, message)

		//get the width of all the doors in the room
		message = fmt.Sprintf("Could you please give me the width of doors in room %d in meters?", i)
		fmt.Println(message)
		doorWidth := o.InputSize(scanner, message)
		//make sure that users input is valid
		for doorWidth*float64(doorNum) > width {
			fmt.Println("Sorry, can't have doors wider than your wall :(")
			fmt.Println(message)
			doorWidth = o.InputSize(scanner, message)
		}

		//get the height of all the doors in the room
		message = fmt.Sprintf("Could you please give me the height of doors in room %d in meters?", i)
		fmt.Println(message)
		doorHeight := o.InputSize(scanner, message)
		//make sure that users input is valid
		for doorHeight > height {
			fmt.Println("Sorry, can't have doors higher than your wall :(")
			fmt.Println(message)
			doorHeight = o.InputSize(scanner, message)
		}

		surfaceDoors := int64(doorNum) * int64(doorHeight) * int64(doorWidth)
		fmt.Println("Great! The total surface of the doors in the room is:" + strconv.FormatInt(int64(surfaceDoors), 10) + " square meters.")

		surfaceToPaint := surfaceW - float64(surfaceDoors)
		//make sure that surface to paint is not negative
		if surfaceToPaint < 0 {
			panic("Oh no something went wrong :(((")
		}

		//put all the surfaces to paint in a slice
		surfaceOfRooms = append(surfaceOfRooms, surfaceToPaint)
		fmt.Println("The surface that needs to be painted is " + strconv.FormatInt(int64(surfaceToPaint), 10) + " square meters.")
	}

	total := 0.0

	for _, roomSurface := range surfaceOfRooms {
		total += roomSurface
	}

	fmt.Println("The total surface that you have requested to paint is:" + strconv.FormatInt(int64(total), 10))

	colorMenu := m.GetColorMenu()
	m.DisplayColor(colorMenu)

	fmt.Println("At the moment we only have white paint with the price od 3 euros/square meter.")

	totalPrice := total * (1 - discount)
	fmt.Printf("With the discount of %.2f the total price is %.2f.\n", discount, totalPrice)
}
