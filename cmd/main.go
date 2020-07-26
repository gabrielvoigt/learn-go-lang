package main

import (
	"fmt"
)

const (
	author string = "Gabriel"
	version string = "0.0.1"
	typeVersion string = "alpha"
)


func main() {

	println("Hello, welcome to my system.")
	fmt.Println("I'm", author)
	fmt.Println("Version:", version, "-", typeVersion)

	println("This system consist to find the timeline about events.")
	fmt.Print("Search: ")
	var search string
	fmt.Scan(&search)


}