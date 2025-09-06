package main

import "fmt"

func main() {

	// We are learning Maps so

	//myMap := map[string]int // we just declare map till now it is un usable [] key ha or int = value ki type ha

	myMap := make(map[string]int) // by using make we initialze the map now we can use it

	// now we insert some values in mymap

	myMap["Apple"] = 14 // map key ka name ka likha or us ke ander value dali

	myMap["Banna"] = 25 // same there also

	fmt.Println(myMap) // show full key value pair  maps

	fmt.Println(myMap["Apple"]) // show only value using key

}
