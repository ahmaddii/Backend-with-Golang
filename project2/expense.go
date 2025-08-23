package main

import "fmt"

func main() {

	var day1, day2, day3, day4 int
	var budget int

	fmt.Println("Enter the Budget to Set")
	fmt.Scanln(&budget)

	fmt.Println("Enter the Expense of Day 1")
	fmt.Scanln(&day1)

	fmt.Println("Enter the Expense of Day 2")
	fmt.Scanln(&day2)

	fmt.Println("Enter the Expense of Day 3")
	fmt.Scanln(&day3)

	fmt.Println("Enter the Expense of Day 4")
	fmt.Scanln(&day4)

	total := day1 + day2 + day3 + day4
	fmt.Println("Total Spending:", total)

	fmt.Println("Average Spending :", total/4)

	if total > budget {

		println("Alert ! Your Budget Exceeds")
	}

}
