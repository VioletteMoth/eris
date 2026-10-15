package combat

import (
	"eris/character"
	"eris/data"
	"eris/utils"
	"fmt"
	"strconv"
)

func Combat(c *data.Character, zone string) {
	utils.ClearTerminal()
	enemy := data.GetEnemy(zone)
	e := &enemy

	if e.Name == "" {
		fmt.Println("\nFor no one is here, there is nothing to fight")
		return
	}
	fmt.Println("	🙟🙦  COMBAT START 🙤🙝")
	var advantage bool = true
	if c.Class == "Machina" && e.Class == "Sarta" {
		advantage = false
	}
	if c.Class == "Deus" && e.Class == "Machina" {
		advantage = false
	}
	if c.Class == "Sarta" && e.Class == "Deus" {
		advantage = false
	}
	var turn int = 1
	for c.CurHP > 0 && e.CurHP > 0 {
		if CombatTurn(c, e, advantage, turn) {
			return
		}
		turn++
		fmt.Print("\n")
	}
}

func CombatTurn(char, enemy *data.Character, advantage bool, turn int) bool {
	utils.ClearTerminal()
	fmt.Println("	🙟  TURN", turn, "🙝")
	switch enemy.Class {
	case "None":
		fmt.Print("◌ ")
	case "Machina":
		fmt.Print("⚙  ")
	case "Deus":
		fmt.Print("✧ ")
	case "Sarta":
		fmt.Print("☺ ")
	}
	fmt.Print(enemy.Name + "	")
	fmt.Print(strconv.Itoa(enemy.CurHP) + "/" + strconv.Itoa(enemy.MaxHP) + "HP ")
	fmt.Println(strconv.Itoa(enemy.CurSP) + "/" + strconv.Itoa(enemy.MaxSP) + "SP")
	fmt.Print("\n")

	switch char.Class {
	case "Machina":
		fmt.Print("⚙  ")
	case "Deus":
		fmt.Print("✧ ")
	case "Sarta":
		fmt.Print("☺ ")
	}
	fmt.Print(char.Name + "	")
	fmt.Print(strconv.Itoa(char.CurHP) + "/" + strconv.Itoa(char.MaxHP) + "HP ")
	fmt.Println(strconv.Itoa(char.CurSP) + "/" + strconv.Itoa(char.MaxSP) + "SP")

	if advantage {
		playerTurn(char, enemy)
		if enemy.CurHP <= 0 {
			return utils.IsWin(char, *enemy)
		}
	}
	enemyTurn(char, enemy, turn)
	if char.CurHP <= 0 {
		utils.IsDead(char)
		return true
	}
	if !advantage {
		playerTurn(char, enemy)
		if enemy.CurHP <= 0 {
			return utils.IsWin(char, *enemy)
		}
	}
	return false
}

func playerTurn(char, enemy *data.Character) {
	for {
		fmt.Println("0 > Wait	1 > Charge	2 > Item	3 > Skill")

		var choice int = -1
		for choice < 0 || choice > 3 {
			fmt.Print("> ")
			fmt.Scan(&choice)
		}
		switch choice {
		case 1:
			char.CurSP += 2
			if char.CurSP > char.MaxSP {
				char.CurSP = char.MaxSP
			}
			return
		case 2:
			if character.Inventory(char, enemy) {
				return
			}
		case 3:
			if character.Skills(char, enemy) {
				return
			}
		case 0:
			return
		}
	}
}
func enemyTurn(c, e *data.Character, turn int) {
	if turn%2 == 0 {
		c.CurHP -= utils.DealDammage(e, c, 5)
	}
}
