package main

import "fmt"

func main() {

	var item1, item2, item3, item4 float64

	fmt.Println("----Welcome to Lasani Pizza Time----")

	fmt.Println("Enter the Price of Item1")
	fmt.Scanln(&item1)

	fmt.Println("Enter the Price of Item2")
	fmt.Scanln(&item2)

	fmt.Println("Enter the Price of Item3")
	fmt.Scanln(&item3)

	fmt.Println("Enter the Price of Item4")
	fmt.Scanln(&item4)

	totalBill := item1 + item2 + item3 + item4

	if totalBill > 2000 {

		totalBill -= 200
	}

	tax := totalBill * 0.1 // 10 % Tax

	fmt.Println("Sub Total", totalBill)
	fmt.Println("With Tax", tax)
	fmt.Println("With Discount", totalBill)

}
