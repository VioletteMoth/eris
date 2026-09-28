package menus

import (
	"eris/data"
	"fmt"
)

func GameMainMenu(player *data.Character) bool {
	for {
		fmt.Println("\n	🙟🙦🙧  HOME MENU 🙥🙤🙝")
		fmt.Println("1 > New Game")
		if player != nil {
			fmt.Println("2 > Resume ")
		}
		fmt.Println("0 > Quit Game")

		var choice int = -1
		for choice < 0 || choice > 3 {
			fmt.Print("> ")
			fmt.Scan(&choice)
		}

		switch choice {
		case 1:
			return true
		case 2:
			if player == nil {
				fmt.Println("No game in progress.")
				continue
			}
		case 0:
			return false
		}
	}
}
