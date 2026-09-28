package menus

import (
	"eris/character"
	"eris/combat"
	"eris/data"
	"eris/economy"
	"fmt"
)

func Menu(c *data.Character, zone string) {
	for {
		fmt.Println("\n	🙟🙦🙧  MENU 🙥🙤🙝")
		fmt.Println("1 > Status")
		fmt.Println("2 > Inventory")
		fmt.Println("3 > Equipment")
		fmt.Println("4 > Skills")
		fmt.Println("5 > Merchant")
		fmt.Println("6 > Blacksmith")
		fmt.Println("7 > Training")
		fmt.Println("0 > Quit")

		var choice int = -1
		for choice < 0 || choice > 8 {
			fmt.Print("> ")
			fmt.Scan(&choice)
		}
		switch choice {
		case 1:
			character.Display(*c)
		case 2:
			character.Inventory(c, nil)
		case 3:
			character.Equipment(c)
		case 4:
			character.Skills(c, nil)
		case 5:
			economy.Shop(c, zone)
		case 6:
			economy.Forge(c, zone)
		case 7:
			combat.Combat(c, zone)
		case 0:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
