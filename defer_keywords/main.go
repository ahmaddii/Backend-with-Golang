// File Handling In Backend Using Defer keyword


package main 


import "fmt"
import "os"


func main(){


	// I want to create a file

    file,error := os.Create("example.txt") // file gives reference as a handler

	if (error != nil) {

		fmt.Println("Error Occur While Creating a File",error)
		return
	}

	defer file.Close() // It runs everytime at the end of every function no matter function executs or gives error
	// it help us to manage resouse allocation


	_,err := file.WriteString("Hello from Golang Backend") // _ ignores the bytes data

	  if(err != nil){

		fmt.Println("Error Occur during Writing a file ",err)
		return
	  }

	  fmt.Println("File Written Successfully")




}