package character

import (
	"eris/data"
	"fmt"
)

func Equipment(char *data.Character) {
	for {
		fmt.Println("\n	🙟 EQUIPMENT 🙝")
		fmt.Println("1 > Weapon:", char.Equipment[0].Name)
		fmt.Println("2 > Head:", char.Equipment[1].Name)
		fmt.Println("3 > Chest:", char.Equipment[2].Name)
		fmt.Println("4 > Feet:", char.Equipment[3].Name)
		fmt.Println("0 > Back")
		fmt.Print("> ")

		var choice string
		if _, err := fmt.Scan(&choice); err != nil {
			return
		}
		switch choice {
		case "0":
			return
		case "1":
			ManageEquipment(char, &char.Equipment[0], "weapon")
		case "2":
			ManageEquipment(char, &char.Equipment[1], "head")
		case "3":
			ManageEquipment(char, &char.Equipment[2], "chest")
		case "4":
			ManageEquipment(char, &char.Equipment[3], "feet")
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func ManageEquipment(char *data.Character, equipped *data.Item, slot string) {
	for {
		fmt.Println("\nSlot:", slot)
		fmt.Println("Equipped:", equipped.Name)
		fmt.Println("Choose an inventory item to equip:")

		var choices []int
		for i, item := range char.Inventory {
			if item.Type == slot {
				choices = append(choices, i)
				fmt.Println(len(choices), ">", item.Name)
			}
		}
		unequipChoice := len(choices) + 1
		if equipped.Name == "" {
			data.ColorUnavailable.Println(unequipChoice, "> Unequip")
		} else {
			fmt.Println(unequipChoice, "> Unequip")
		}
		fmt.Println("0 > Back")

		var choice int = -1
		for choice < 0 || choice > len(choices)+1 {
			fmt.Print("> ")
			fmt.Scan(&choice)
		}

		if choice == 0 {
			return
		} else if choice == unequipChoice {
			if equipped.Name == "" {
				fmt.Println("Nothing is equipped in this slot.")
				continue
			}

			freeSlot := -1
			for i, itm := range char.Inventory {
				if itm.Name == "" {
					freeSlot = i
					break
				}
			}
			if freeSlot == -1 && len(char.Inventory) <= 10*char.InvSize {
				freeSlot = len(char.Inventory)
				char.Inventory = append(char.Inventory, data.GetItem("item", "empty"))
			}
			if freeSlot == -1 {
				fmt.Println("Inventory full. Cannot unequip this item.")
				continue
			}

			char.Inventory[freeSlot] = *equipped
			fmt.Println(equipped.Name, "returned to your inventory.")
			*equipped = data.GetItem("item", "empty")
			return
		} else if choice < 1 || choice > len(choices) {
			fmt.Println("Invalid choice.")
			continue
		}

		index := choices[choice-1]
		selected := char.Inventory[index]

		// Vérifie que l'objet existe et correspond au slot
		if selected.Name == "" || selected.Type != slot {
			fmt.Println("This item cannot be equipped in this slot.")
			continue
		}

		previous := *equipped

		if previous.Name == "" {
			previous = data.GetItem("item", "empty")
		}

		char.Inventory[index] = previous
		*equipped = selected

		fmt.Println(selected.Name, "equipped.")

		if previous.Name != "" {
			fmt.Println(previous.Name, "returned to your inventory.")
		}
		return
	}
}
