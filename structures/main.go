package main

import (
	"fmt"
)

// Structures in golang

// struct declare outside of the main function and
// in which e.g the data store in pieces in different variables but now we combine them in single container like struct

type User struct {
	name  string
	Age   int
	Email string
} // we create a user which is just like a complete container as it has its own data

func main() {

	// now create instance or object so that we store things or data in our container

	user1 := User{

		name:  "Ahmad",
		Age:   18,
		Email: "ahmad@gmail.com",
	}

	// we can create multiple objects or instance dekha kitna asan ha
	// bajae is ke ke ap bar bar variables declare kre

	user2 := User{

		name:  "Hassan",
		Age:   20,
		Email: "hassan@gmail.com",
	}

	fmt.Println("Name :", user1.name)
	fmt.Println("Age :", user1.Age)
	fmt.Println("Email :", user1.Email)

	fmt.Println("Name :", user2.name)
	fmt.Println("Age :", user2.Age)
	fmt.Println("Email :", user2.Email)

	// We can also used structs with slices

	users := []User{ // instance with slice

		// struct with slicess

		{name: "Ahmad", Age: 18},
		{name: "Ali", Age: 20},
	}

	// now just give the index number of the struct row which you want to access
	fmt.Println(users[1].name)

	// how to update the struct elements within slices

	users[1].name = "Sara"

	fmt.Println(users[1].name)

	fmt.Println(users[0].Age)

}
