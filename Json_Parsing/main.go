// Parsing Json data from api in slices

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int    `json:id`
	Name string `json:name` // map the name with Name mostly we done in json

}

func main() {

	data := `[{"id":1,"name":"Ali"},{"id":2,"name":"Ayesha"},{"id":3,"name":"Bilal"}]` // API Response se data aya

	var users []User // slice of user objects

	err := json.Unmarshal([]byte(data), &users) // parses json data using unmarshal fro bytes to slice  and also pass memory addres of users to updated

	if err != nil {

		fmt.Println("Error", err) // if error occur give this 
		return

	}

	for _, user := range users { // run loop 

		fmt.Printf("Id: %d and Name: %s\n", user.Id, user.Name) // give placholders and get raw data from json formating
	}

}
