package character

import (
	"eris/data"
	"fmt"
)

func Upgrade(c *data.Character) {
	for {
		fmt.Println("\n	🙟  UPGARDES 🙝")
		fmt.Println("🟁", c.Points)
		fmt.Println("1 > Health")
		fmt.Println("2 > Energy")
		fmt.Println("3 > Attack")
		//fmt.Println("4 > Skills")
		fmt.Println("0 > Back")

		var choice int = -1
		for choice < 0 || choice > 3 {
			fmt.Print("> ")
			fmt.Scan(&choice)
		}
		if c.Points <= 0 && choice != 0 {
			choice = -1
		}
		switch choice {
		case 1:
			c.MaxHP += (c.MaxHP / 10) + 5
			c.CurHP = c.MaxHP
			c.Points--
			fmt.Println("Health upgraded !")
		case 2:
			c.MaxSP += 1
			c.CurSP = c.MaxSP
			c.Points--
			fmt.Println("Energy upgraded !")
		case 3:
			c.Atk += 1
			c.Points--
			fmt.Println("Attack upgraded !")
		case 0:
			return
		case -1:
			fmt.Println("You don't have enough points.")
		default:
			fmt.Println("Invalid choice")
		}
	}
}
