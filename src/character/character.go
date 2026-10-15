package character

import (
	"eris/data"
	"eris/utils"
	"fmt"
	"strconv"
)

func InitCharacter(name string, class int) data.Character {
	c := data.Character{
		Id:        "player",
		Name:      name,
		Level:     0,
		CurXP:     0,
		MaxXP:     15,
		Points:    0,
		Coins:     100,
		CurHP:     5,
		MaxHP:     12,
		CurSP:     4,
		MaxSP:     4,
		Atk:       0,
		Def:       0,
		InvSize:   1,
		Inventory: []data.Item{data.GetItem("item", "elixir"), data.GetItem("item", "elixir"), data.GetItem("item", "poison")},
		Skills:    [4]data.Skill{data.GetSkill("rush")},
	}
	switch class {
	case 1:
		c.Class = "Machina"
		c.Skills[1] = data.GetSkill("assurance")
	case 2:
		c.Class = "Deus"
		c.Skills[1] = data.GetSkill("benediction")
	case 3:
		c.Class = "Sarta"
		c.Skills[1] = data.GetSkill("kick")
	case 0:
		c.Class = ""
	}
	return c
}

func Display(c data.Character) {
	utils.ClearTerminal()
	fmt.Println("	🙟  INFORMATION 🙝")
	fmt.Println("Name:		" + c.Name)
	fmt.Println("Class:		" + c.Class)
	fmt.Print("Level:		", c.Level)
	fmt.Println("(" + strconv.Itoa(c.Points) + ")")
	fmt.Print("Exp:		", c.CurXP)
	fmt.Println(" /", c.MaxXP)
	fmt.Print("Health:		", c.CurHP)
	fmt.Println(" /", c.MaxHP)
	fmt.Print("Energy:		", c.CurSP)
	fmt.Println(" /", c.MaxSP)
	fmt.Println("Attack:		x", 1+(c.Atk/10))
	lenInv := 0
	for _, item := range c.Inventory {
		if item.Name != "" {
			lenInv++
		}
	}
	fmt.Print("Inventory:	", lenInv)
	fmt.Println(" /", c.InvSize*10)
	var choice int = -1
	if c.Points <= 0 {
		data.ColorUnavailable.Println("1 > Upgrade")
	} else {
		fmt.Println("1 > Upgrade")
	}
	fmt.Println("0 > Back")
	for choice != 0 || (choice == 1 && c.Points >= 0) {
		fmt.Print("> ")
		fmt.Scan(&choice)
	}
	switch choice {
	case 1:
		Upgrade(&c)
	case 0:
		return
	}
}
