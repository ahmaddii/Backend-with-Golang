package main

import "fmt"

func main() {

	// We are learning Maps so

	//myMap := map[string]int // we just declare map till now it is un usable [] key ha or int = value ki type ha

	myMap := make(map[string]int) // by using make we initialze the map now we can use it

	// now we insert some values in mymap

	myMap["Apple"] = 14 // map key ka name ka likha or us ke ander value dali

	myMap["Banna"] = 25 // same there also

	myMap["Apple"] = 50 // Update the Value of Apple Key

	myMap["Orange"] = 60

	fmt.Println(myMap) // show full key value pair  maps

	fmt.Println(myMap["Apple"]) // show only value using key

	// Check weather the key is available or not

	value,exsist := myMap["Apple"]

	if exsist {

		fmt.Println("Found the Key",value)
	} else {

		fmt.Println("Not Found")
	}

	// Deleting Keys

	delete(myMap,"Orange")

	fmt.Println("After Deleting Orange Key")

	fmt.Println(myMap)


	// Loop through the Map

	for key,value := range myMap{

		fmt.Println("Key: ",key, "Value :",value)


	}

	// Example of Headers in using go maps

	headers := map[string]string{

		"content-type": "Application/json",
		"Authorization": "Bearer token 123",

	}

	fmt.Println("Authorization: ",headers["Authorization"])

}
