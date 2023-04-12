package main

import "fmt"

func main() {
	var firstVal int64
	var SecondVale int64
	fmt.Scan(&firstVal)
	fmt.Scan(&SecondVale)

	if firstVal == SecondVale {
		fmt.Println("Equal Value")
	} else {
		fmt.Println("Doest not Equal")
	}

}
