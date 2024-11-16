package main

import "fmt"

func main() {
	var availablePencils int
	var availablePupils int
	//input values
	fmt.Print("Enter number of available pencils")
	fmt.Scan(&availablePencils)
	fmt.Print("Enter number of available pupils")
	fmt.Scan(&availablePupils)
	//calculate if the pencils are enough for the pupils
	if availablePencils >= availablePupils {
		fmt.Println("The pencils are enough")
	} else {
		fmt.Println("The pencils are not enough")
	}

}
