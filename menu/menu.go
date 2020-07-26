package menu

import (
	"fmt"
	term "github.com/nsf/termbox-go"
	"go-search-history/survey"
	"gopkg.in/gookit/color.v1"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
)

var ColorText = color.S256(222,3)
var ColorAlertErros = color.Error // color.S256(3,160)
var MenuSelected = color.S256(222,3)
var White = color.White
var Blue = color.Blue

func Reset() {
	term.Sync() // cosmestic purpose
}

var arrayMenu = map[int] string{ 0: "0: Exit",
	1: "1: Create and write a file",
	2: "2: Open a new instance CMD",
	3: "3: Create a new folder",
	4: "4: Search",
}

var Clear map[string]func()

func init() {
	Clear = make(map[string]func())
	Clear["linux"] = func() {
		cmd := exec.Command("Clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	Clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	functionSelected, exists := Clear[runtime.GOOS]
	if exists {
		functionSelected()
	} else {
		panic("Your platform is unsupported! I can't Clear terminal screen :(")
	}
}

func initMenu(position int) int {
	color.Yellow.Println("From now one, you can test all the functionality about this system.\n")
	color.Yellow.Println("Select the apropriate option:\n")

	sliceMenu := make(map[int] string, len(arrayMenu))
	keys := make([]int,0)

	for key := range arrayMenu {
		if key != position{
			sliceMenu[key] = "  " +  arrayMenu[key]
		} else {
			sliceMenu[key] = "> " + arrayMenu[key]
		}
		keys = append(keys, key)
	}

	sort.Ints(keys)
	var positionMenuSelected int
	for k := len(keys)-1; k >= 0; k-- {
		if strings.ContainsAny(sliceMenu[k], ">") {
			MenuSelected.Println(sliceMenu[k])
			positionMenuSelected = k
		} else {
			color.White.Println(sliceMenu[k])
		}
	}

	fmt.Println()
	color.Yellow.Println("[↓] Down | [↑] Up | ESC: Quit")
	return positionMenuSelected
}

func ShowMenu() int {
	positionMenu := len(arrayMenu) - 1
	CallClear()
	initMenu(positionMenu)
	return positionMenu
}

func Menu() int {
	err := term.Init()

	if err != nil {
		panic(err)
	}

	defer term.Close()

	positionMenu := ShowMenu()

keyPressListenerLoop:
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				CallClear()
				color.Success.Println("Okay! Bye, bye!")
				os.Exit(0)
				break keyPressListenerLoop
			case term.KeyArrowDown:
				positionMenu -= 1
				if positionMenu >= 0 {
					if positionMenu > (len(arrayMenu) - 1) {
						positionMenu = len(arrayMenu) - 1
					}
					CallClear()
					positionMenu = initMenu(positionMenu)
				} else {
					positionMenu = 0
					CallClear()
					positionMenu = initMenu(positionMenu)
				}
			case term.KeyArrowUp:
				positionMenu += 1
				if positionMenu >= 0 && positionMenu <= (len(arrayMenu)-1) {
					if positionMenu <= 0 {
						positionMenu = 0
					}
					CallClear()
					positionMenu = initMenu(positionMenu)
				} else {
					positionMenu = len(arrayMenu) - 1
					CallClear()
					positionMenu = initMenu(positionMenu)
				}
			case term.KeyEnter:
				if positionMenu < 0 {
					positionMenu = 0
				} else if positionMenu >= len(arrayMenu) {
					positionMenu = len(arrayMenu) - 1
				}

				CallClear()
				return positionMenu
			default:
				Reset()
			}
		case term.EventError:

			panic(ev.Err)
		}
	}
	panic("Error! Exited the menu")
	return 0
}

func MainMenu(){
	var input = Menu()
	switch input {
	case 4:
		//TODO I need to save all the search into DB. I think to record everything in MongoDB
		//TODO Name of service: "Survey"
		color.Info.Println("Search (developing)")
		survey.Search()
		ShowMenu()
	case 3:
		color.Info.Println("Create a new folder")
		FolderMenu()
		MainMenu()
	case 2:
		color.Info.Println("Open a new instance CMD (developing)")
		ShowMenu()
	case 1:
		color.Info.Println("Create and write a file (developing)")
		ShowMenu()
	case 0:
		color.Success.Println("Okay! Bye, bye!")
		os.Exit(0)
	default:
		color.Error.Println("This option does not exist")
		os.Exit(-1)
	}

}
