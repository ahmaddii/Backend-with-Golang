package main

import (
	"fmt"

	"time" // use time delay to use a delay between lights
)

func main() {

	for i := 1; i < 5; i++ {

		fmt.Println("---Traffic Light Simulater---")

		showLight("Red", "- Stop 🚫")
		time.Sleep(4 * time.Second)
		showLight("Yellow", "- Ready  ⚠️")
		time.Sleep(4 * time.Second)
		showLight("Green", "-Starts ✅")

	}

}

func showLight(color string, signal string) {

	fmt.Println("Light", color, signal)

}
