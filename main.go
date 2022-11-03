package main

import "fmt"

func main() {

	//loop over map
	myMap := map[string]string{"me": "Sasha", "you": "unknown", "him": "my Boi"}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	fmt.Println()
	fmt.Println()

	//loop over slice
	mySlice := []string{"Sasha", "Masha"}
	mySlice = append(mySlice, "Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta")

	for i, myString := range mySlice {
		fmt.Println(i, myString)
	}

}
