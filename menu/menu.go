package menu

import (
	"fmt"
	"strconv"
)

func DisplayDiscounts(rooms int) float64 {
	discountMenu := make(map[int64]string)

	discountMenu[3] = "5%"
	discountMenu[4] = "10%"
	discountMenu[5] = "20%"
	discountMenu[6] = "40%"

	fmt.Println("Discount for the quantity:")
	for no, percentage := range discountMenu {
		if no == 6 {
			fmt.Println("For " + strconv.FormatInt(no, 10) + " or more:" + percentage)
		} else {
			fmt.Println("For " + strconv.FormatInt(no, 10) + ":" + percentage)
		}
	}
	if rooms < 3 {
		fmt.Println("No discount for you :(")
		return 0
	} else if rooms > 6 {
		fmt.Println("What a discount!")
		return 0.4
	} else {
		fmt.Println("=======Your discount is " + discountMenu[int64(rooms)] + " ========")
		perc := discountMenu[int64(rooms)][0:2]
		percNum, _ := strconv.ParseFloat(perc, 32)
		return percNum / 100
	}

}

func GetColorMenu() map[string]float64 {
	colorMenu := make(map[string]float64)

	colorMenu["red"] = 3.5
	colorMenu["blue"] = 4.1
	colorMenu["white"] = 3
	colorMenu["green"] = 4.8

	return colorMenu
}

func DisplayColor(colorMenu map[string]float64) {

	fmt.Println("The color menu is:")
	for color, price := range colorMenu {

		fmt.Printf("For color %s, the price is %.1f for square meter. \n", color, price)

	}

}
