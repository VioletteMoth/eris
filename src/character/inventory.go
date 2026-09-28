package character

import (
	"eris/data"
	"fmt"
	"strconv"
)

func Inventory(char, target *data.Character) bool {
	for {
		fmt.Println("\n	🙟  INVENTORY 🙝")
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
			if PickInvItem(char, target, item) {
				return true
			}
		case item == 0:
			fmt.Print("\n")
			return false
		}
	}
}

func PickInvItem(char, target *data.Character, index int) bool {
	if char.Inventory[index-1].Type == "" {
		for {
			itm := char.Inventory[index-1]
			fmt.Println(string(itm.Icon) + " " + itm.Name + " : " + itm.Desc)
			fmt.Println("1[USE]	2[THROW]")
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
				char.Inventory[index-1] = data.GetItem("item", "empty")
				fmt.Print("\n")
				return true
			case 2:
				if target == nil {
					fmt.Print(char.Name, " threw ", string(itm.Icon), " ", itm.Name, " on the ground.")
				} else {
					fmt.Print(char.Name, " threw ", string(itm.Icon), " ", itm.Name, " on ", target.Name, ", ")
				}
				char.Inventory[index-1] = data.GetItem("item", "empty")
				fmt.Print("\n")
				return true
			}
		}
	} else {
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
