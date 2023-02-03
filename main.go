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
	scanner := bufio.NewScanner(os.Stdin)
	rooms := 0
	fmt.Println("Please enter the number of rooms you want to paint:")
	rooms = o.InputRoomNum(scanner)

	fmt.Println("Great!Let's look at the possible discounts for the quantity:")
	discount := m.DisplayDiscounts(rooms)
	fmt.Println(discount)
	fmt.Println("Would you like to change the number of rooms(y/n)?")

	rooms = o.ChangeRoomNum(scanner, rooms)

	fmt.Println(rooms)

	surfaceOfRooms := []float64{}

	for i := 1; i <= rooms; i++ {

		message := fmt.Sprintf("Could you please give me the width of room %d in meters?", i)
		fmt.Println(message)
		width := o.InputSize(scanner, message)
		message = fmt.Sprintf("Could you please give me the height of room %d in meters?", i)
		fmt.Println(message)
		height := o.InputSize(scanner, message)
		surfaceW := width * height
		fmt.Println("Great! The size of the first room is " + strconv.FormatInt(int64(surfaceW), 10) + " square meters.")

		fmt.Println("How many doors there are in the room?")
		doorNum := o.InputRoomNum(scanner)
		message = fmt.Sprintf("Could you please give me the width of doors in room %d in meters?", i)
		fmt.Println(message)
		doorWidth := o.InputSize(scanner, message)
		message = fmt.Sprintf("Could you please give me the height of doors in room %d in meters?", i)
		fmt.Println(message)
		doorHeight := o.InputSize(scanner, message)
		surfaceDoors := int64(doorNum) * int64(doorHeight) * int64(doorWidth)
		fmt.Println("Great! The total surface of the doors in the room is:" + strconv.FormatInt(int64(surfaceDoors), 10) + " square meters.")

		surfaceToPaint := surfaceW - float64(surfaceDoors)
		surfaceOfRooms = append(surfaceOfRooms, surfaceToPaint)
		fmt.Println("The surface that needs to be painted is " + strconv.FormatInt(int64(surfaceToPaint), 10) + " square meters.")
	}

	total := 0.0

	for _, roomSurface := range surfaceOfRooms {
		total += roomSurface
	}

	fmt.Println("The total size is:" + strconv.FormatInt(int64(total), 10))

	fmt.Println("The price by square meter is 3 euros.")

	totalPrice := total * (1 - discount)
	fmt.Printf("With the discount of %f the total price is %f.\n", discount, totalPrice)
}
