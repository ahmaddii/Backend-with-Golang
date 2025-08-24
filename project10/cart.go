package main

import (
	"fmt"
	"time"
)

func main() {

	var price, quantity, grandTotal, Total int
	var itemName string
	var choice string

	for {

		fmt.Println("Enter the Name of Item")
		fmt.Scanln(&itemName)

		fmt.Println("Enter the Price Of Item")
		fmt.Scanln(&price)

		fmt.Println("Enter the Quantity Of Item")
		fmt.Scanln(&quantity)

		Total = price * quantity

		if Total > 4000 {

			discount := Total * 10 / 100 // total ager 4000 se zayada ha to discount lagao 10 %
			Total = Total - discount     // ab actaul total mein discount jo nikala ha us ko minus krdo

		}

		tax := Total * 5 / 100 // apply 5 % tax
		Total = Total + tax    // then add the tax to the total

		grandTotal += Total // now add to actual grand total

		fmt.Println("Total Amount with Discount:", Total)
		fmt.Println("Grand Total ", grandTotal)

		fmt.Println("Do you want to add Another Item: yes/no")
		fmt.Scanln(&choice)

		if choice == "no" {

			break

		}

	}

	time.Sleep(2 * time.Second) // just 2 sec delay

	fmt.Println("============")
	fmt.Println("Grand Total:", grandTotal)
	fmt.Println("Thank You for Shopping")

}
