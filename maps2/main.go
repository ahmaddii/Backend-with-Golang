package main

import "fmt"

func main() {

	myMap := map[string]string{

		"user1": "Ahmad",
		"user2": "Ali",
	}

	fmt.Println(myMap)

	for key, value := range myMap {

		fmt.Println("Key: ", key, "Value :", value)
	}

	myMap2 := make(map[string]int)

	myMap2["Orange"] = 20 // key or value di

	myMap2["Grapes"] = 25

	fmt.Println(myMap2["Orange"])

	fmt.Println(myMap2["Grapes"])

	// Checking keys or values

	value, exist := myMap2["Grapes"]

	if exist {

		fmt.Println("Found Key", value)
	} else {

		fmt.Println("key Not Found")
	}

	// Real World Usage in Golang Backend

	headers := map[string]string{

		"content-type":  "application/json",
		"Authorization": "Access Token 123",
	}

	fmt.Println("Authorization with ", headers["Authorization"])

	// Delete

	delete(myMap2, "Grapes")

	fmt.Println("After Deleting Grapes Key", myMap2)

}
