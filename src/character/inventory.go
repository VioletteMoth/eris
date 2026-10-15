package character

import (
	"eris/data"
	"eris/items"
	"eris/utils"
	"fmt"
	"strconv"
)

func Inventory(char, target *data.Character) bool {
	for {
		utils.ClearTerminal()
		fmt.Println("	🙟  INVENTORY 🙝")
		for i, item := range char.Inventory {
			fmt.Print(strconv.Itoa(i+1) + "[" + string(item.Icon) + " " + item.Name + "]	")
		}
		fmt.Print("\n")
		fmt.Println("0 > Back")

		var item int = -1
		for item < 0 || len(char.Inventory) < item {
			fmt.Print("> ")
			fmt.Scan(&item)
		}
		switch {
		case 0 < item && item <= len(char.Inventory):
			if pickItem(char, target, item) {
				return true
			}
		case item == 0:
			fmt.Print("\n")
			return false
		}
	}
}

func pickItem(char, target *data.Character, index int) bool {
	switch char.Inventory[index-1].Type {
	case "":
		for {
			itm := char.Inventory[index-1]
			fmt.Println(string(itm.Icon) + " " + itm.Name + " : " + itm.Desc)
			aim := "Throw"
			if target != nil {
				aim = target.Name
			}
			fmt.Println("1 > Self	2 > " + aim)
			fmt.Println("0 > Back")

			var choice int = -1
			for choice < 0 || 2 < choice {
				fmt.Print("> ")
				fmt.Scan(&choice)
			}

			switch choice {
			case 0:
				fmt.Print("\n")
				return false
			case 1:
				fmt.Print(char.Name, " used ", string(itm.Icon), " ", itm.Name, ", ")
				items.UseFunc(itm.Id, char, char)
				char.Inventory[index-1] = data.GetItem("item", "empty")
				fmt.Print("\n")
				return true
			case 2:
				if target == nil {
					fmt.Print(char.Name, " threw ", string(itm.Icon), " ", itm.Name, " on the ground.")
				} else {
					fmt.Print(char.Name, " threw ", string(itm.Icon), " ", itm.Name, " on ", target.Name, ", ")
					items.UseFunc(itm.Id, char, target)
				}
				char.Inventory[index-1] = data.GetItem("item", "empty")
				fmt.Print("\n")
				return true
			}
		}
	default:
		itm := char.Inventory[index-1]
		fmt.Println(string(itm.Icon) + " " + itm.Name + " : " + itm.Desc + "(" + itm.Type + ")")
		if target == nil {
			fmt.Println("1 > Throw")
			fmt.Println("0 > Back")

			var choice int = -1
			for choice != 0 && choice != 1 {
				fmt.Print("> ")
				fmt.Scan(&choice)
			}
			switch choice {
			case 1:
				fmt.Print(char.Name, " threw ", string(itm.Icon), " ", itm.Name, " on the ground.")
				char.Inventory[index-1] = data.GetItem("item", "empty")
				fmt.Print("\n")
				return false
			case 0:
				return false
			}
		}
		fmt.Println("0 > Back")

		var choice int = -1
		for choice != 0 {
			fmt.Print("> ")
			fmt.Scan(&choice)
		}
		return false
	}
}
