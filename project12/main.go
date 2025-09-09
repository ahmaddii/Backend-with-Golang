package main

import "fmt"

// Student Managment Syestem

type Student struct {
	Id    int
	Name  string
	Grade string
}

func main() {

	Student1 := []Student{ // store student data in student1 var in slices

		{Id: 1, Name: "Ahmad", Grade: "A"},
		{Id: 2, Name: "Hassan", Grade: "B"},
	}

	fmt.Println(Student1)

	StudentMap := make(map[int]Student)

	Student1 = append(Student1, Student{Id: 3, Name: "Sara", Grade: "C"})

	// fill map from slice to search by id

	for _, s := range Student1 {

		StudentMap[s.Id] = s // by this you can search student by Id

	}

	fmt.Println(Student1)

	fmt.Println("Id: ", StudentMap[2])

	fmt.Println("Id: ", StudentMap[1])

}
