package economy

import (
	"eris/data"
	"fmt"
)

func Shop(char *data.Character, merchant *data.Merchant) {
	for {
		fmt.Println("\n	🙟🙦  SHOP 🙤🙝")
		fmt.Println("🞠", char.Coins)
		for i, k := range merchant.Products {
			if char.Coins < k.Cost || k.Quantity <= 0 {
				data.ColorUnavailable.Println(i, ">", k.Cost, "🞠	", k.Item.Name, "(", k.Quantity, ")")
			} else {
				fmt.Println(i, ">", k.Cost, "🞠	", k.Item.Name, "(", k.Quantity, ")")
			}
		}
		fmt.Println("0 > Back")

		var choice int = -1
		for choice < 0 || len(merchant.Products) < choice {
			fmt.Print(">")
			fmt.Scan(&choice)
		}
		switch choice {
		case 0:
			return
		default:
			BuyItem(char, &merchant.Products[choice-1])
		}

	}
}

func BuyItem(char *data.Character, product *data.StoreItem) {
	if product.Quantity <= 0 {
		fmt.Println("Out of stock.")
		return
	}
	if char.Coins < product.Cost {
		fmt.Println("Not enough coins.")
		return
	}

	for i, itm := range char.Inventory {
		if itm.Name == "" {
			char.Inventory[i] = product.Item

			char.Coins -= product.Cost
			product.Quantity--

			fmt.Println("Purchase successful!")
			fmt.Println(product.Item.Name, "was added to your inventory.")
			return
		}
	}
	if len(char.Inventory) != 10*char.InvSize {
		char.Inventory = append(char.Inventory, product.Item)

		char.Coins -= product.Cost
		product.Quantity--

		fmt.Println("Purchase successful!")
		fmt.Println(product.Item.Name, "was added to your inventory.")
		return
	}
	fmt.Println("Not enough space in inventory.")
}
