package main

import (
	"fmt"
	"go-search-history/menu"
	"gopkg.in/gookit/color.v1"
)

var ColorText = color.S256(222,3)
var ColorAlertErros = color.Error // color.S256(3,160)
var MenuSelected = color.S256(222,3)
var White = color.White
var Blue = color.Blue

func welcome(){
	color.Blue.Println("********* Welcome to system CMD test *********")
	color.Blue.Print("|")
	White.Print("     This CMD was created for local test    ")
	color.Blue.Println("|")
	color.Blue.Println("**********************************************")
}

func getText() string{
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	return input
}

func main() {
	//TODO I need to find the folder and open, after I want to create a new folder. So, in the future I'd like to create a folder with the name of person that
	//TODO access the system.
	//TODO I want to recreate a new form to study the history, the people could find a country to initialize your research.
	welcome()
	menu.MainMenu()
}