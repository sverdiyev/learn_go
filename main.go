package main

import (
	"fmt"
	"sort"
)

func main() {

	//Maps
	myMap := make(map[string]string)
	myMap["Sasha"] = "me"
	fmt.Println(myMap["Sasha"])

	//Slices (aka Arrays)
	mySlice := []int{10, 11}
	mySlice = append(mySlice, 9, 8, 7, 6, 5, 2)

	fmt.Println(mySlice)
	sort.Ints(mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[0:2])

}
