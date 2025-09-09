package main

import "fmt"

type Student struct {
	Id         int
	Name       string
	CGPA       float64
	Department string
	Semester   int
}

func main() {

	// store all the strucures values in slices list

	Student1 := []Student{

		{Id: 1, Name: "Ahmad", CGPA: 3.42, Department: "Software Engineering", Semester: 6},
		{Id: 2, Name: "Ali", CGPA: 3.2, Department: "Buisness", Semester: 4},
		{Id: 3, Name: "Hamza", CGPA: 3.43, Department: "CE", Semester: 3},
		{Id: 4, Name: "Hafsa", CGPA: 3.45, Department: "Civil", Semester: 1},
		{Id: 5, Name: "Aqsa", CGPA: 3.46, Department: "Telecom", Semester: 2},
		{Id: 6, Name: "Sara", CGPA: 3.48, Department: "Electrical Engineering", Semester: 3},
	}

	fmt.Println(Student1)

	StudentMap := make(map[int]Student)

	// now how can user search for specific student data with id

	for _, s := range Student1 {

		StudentMap[s.Id] = s

	}

	fmt.Println(StudentMap[1])
}
