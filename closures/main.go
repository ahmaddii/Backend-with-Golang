package main

import "fmt"
 
	func counter() func () int {

		count := 0

		return func() int{

			count ++
			return count // body of funtion
		}
		
	}

func main() {

	mycounter := counter()

	fmt.Println(mycounter())
	fmt.Println(mycounter())
	fmt.Println(mycounter())
	fmt.Println(mycounter())


	
	// Closures

	//GFB := 0 // Declare global vaiable with = 0 value

	//counter := func() int {

	//	GFB += 1

	//	return GFB

	//}

	//fmt.Println(counter())
	//fmt.Println(counter())

}
