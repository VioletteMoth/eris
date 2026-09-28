package economy

import (
	"eris/data"
	"fmt"
)

func Forge(char *data.Character, blacksmith *data.Blacksmith) {
	for {
		fmt.Println("\n	🙟🙦  FORGE 🙤🙝")
		fmt.Println("🞠", char.Coins)
		for i, k := range blacksmith.Craft {
			fmt.Println(i, ">", k.Item.Name)
			for _, l := range k.Materials {
				fmt.Println("	-", l)
			}
			fmt.Println("	", k.Cost, "🞠")
		}
		fmt.Println("0 > Back")

		var choice int = -1
		for choice < 0 || len(blacksmith.Craft) < choice {
			fmt.Print(">")
			fmt.Scan(&choice)
		}
		switch choice {
		case 0:
			return
		default:
			CraftItem(char, blacksmith.Craft[choice-1])
		}
	}
}
func CraftItem(char *data.Character, item data.ForgeItem) {
	if char.Coins < 1 {
		fmt.Println("Not enough coins. Crafting costs 1 coins.")
		return
	}

	inventory := make([]data.Item, len(char.Inventory))
	copy(inventory, char.Inventory)

	for _, material := range item.Materials {
		found := false
		for i := range inventory {
			if inventory[i].Name == material.Name {
				inventory[i] = data.GetItem("item", "empty")
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Missing", material.Name)
			return
		}
	}
	for i := range inventory {
		if inventory[i].Name == "" {
			inventory[i] = item.Item

			copy(char.Inventory, inventory)
			char.Coins -= item.Cost

			fmt.Println("Forged", item.Item.Name)
			fmt.Println("Materials consumed. You spent 5 🞠.")
			return
		}
	}
	fmt.Println("But your bag is full.")
}
