package menu

import (
	"fmt"
	term "github.com/nsf/termbox-go"
	"gopkg.in/gookit/color.v1"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

var arrayMenuFolder = map[int] string{
	3: "3: Search folder",
	2: "2: Create folder",
	1: "1: Delete folder",
	0: "0: Back to main menu",
}

func SearchDirectory(folder string) string{
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("You cannot visualize this path")
	}

	if runtime.GOOS == "windows" {
		dir = dir + "\\" + folder
	} else {
		dir = dir + "/" + folder
	}

	_, err = os.Stat(dir)
	if err != nil {
		ColorAlertErros.Println("Error: folder not found!", err)
		time.Sleep(2 * time.Second)
		//main()
		//log.Fatalln("Error: folder not found!", err)
	}

	ColorText.Println("Path:", dir)
	var files []string

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !strings.Contains(path, ".") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("Error: ", err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
	//TODO open the current directory, after open this only folder, you can create folder.
	//TODO I'd like to create a list of folders and access the folder. I'd like to select one folder to create a new file ou visualize the files into.

	//TODO I need to recreate a new menu here, because I think it is important that someone select what function the person would like to do.
	//TODO Call menu folder HERE
	FolderMenu()
	/*fmt.Print("Which folder do you want to enter?: ")
	var input string
	fmt.Scanln(&input)
	SearchDirectory(input)*/
	return "0" //input
}


func initMenuFolder(position int) int {
	color.Yellow.Println("Select the apropriate option:\n")

	sliceMenu := make(map[int] string, len(arrayMenuFolder))
	keys := make([]int,0)

	for key := range arrayMenuFolder {
		if key != position{
			sliceMenu[key] = "  " +  arrayMenuFolder[key]
		} else {
			sliceMenu[key] = "> " + arrayMenuFolder[key]
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
	color.Yellow.Println("[↓] Down | [↑] Up | ESC: Back to main menu")
	return positionMenuSelected
}

func initShowMenuFolder() int {
	positionMenu := len(arrayMenuFolder) - 1
	CallClear()
	initMenuFolder(positionMenu)
	return positionMenu
}

func menuActionFolder() int {
	err := term.Init()

	if err != nil {
		panic(err)
	}

	defer term.Close()

	positionMenu := initShowMenuFolder()

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
						if positionMenu > (len(arrayMenuFolder) - 1) {
							positionMenu = len(arrayMenuFolder) - 1
						}
						CallClear()
						positionMenu = initMenuFolder(positionMenu)
					} else {
						positionMenu = 0
						CallClear()
						positionMenu = initMenuFolder(positionMenu)
					}
				case term.KeyArrowUp:
					positionMenu += 1
					if positionMenu >= 0 && positionMenu <= (len(arrayMenuFolder)-1) {
						if positionMenu <= 0 {
							positionMenu = 0
						}
						CallClear()
						positionMenu = initMenuFolder(positionMenu)
					} else {
						positionMenu = len(arrayMenuFolder) - 1
						CallClear()
						positionMenu = initMenuFolder(positionMenu)
					}
				case term.KeyEnter:
					if positionMenu < 0 {
						positionMenu = 0
					} else if positionMenu >= len(arrayMenuFolder) {
						positionMenu = len(arrayMenuFolder) - 1
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

func FolderMenu(){
	var input = menuActionFolder()

	fmt.Println("You selected:", input)
	switch input {
	case 3:
		color.Info.Println("Search folder")
		FolderMenu()
	case 2:
		color.Info.Println("Create folder")
		SearchDirectory("")
		FolderMenu()
	case 1:
		color.Info.Println("Delete folder")
		FolderMenu()
	case 0:
		color.Info.Println("Back to main menu")
		MainMenu()
	default:
		color.Error.Println("This option does not exist")
		os.Exit(-1)
	}

}
