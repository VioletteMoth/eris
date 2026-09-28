package combat

import (
	"eris/character"
	"eris/data"
	"eris/sound"
	"fmt"
	"strconv"
)

func Combat(c *data.Character, zone string) {
	enemy := data.GetEnemy("zone")
	e := &enemy
	fmt.Print("\n")
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
	fmt.Print(enemy.Name)
	switch {
	case enemy.CurHP <= enemy.MaxHP/4:
		fmt.Println(" █")
	case enemy.CurHP <= enemy.MaxHP/2:
		fmt.Println(" ██")
	case enemy.CurHP <= enemy.MaxHP/4*3:
		fmt.Println(" ███")
	case enemy.CurHP <= enemy.MaxHP:
		fmt.Println(" ████")
	}
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
		PlayerTurn(char, enemy)
		if enemy.CurHP <= 0 {
			return IsWin(char, *enemy)
		}
	}
	EnemyTurn(char, enemy, turn)
	if char.CurHP <= 0 {
		IsDead(char)
		return true
	}
	if !advantage {
		PlayerTurn(char, enemy)
		if enemy.CurHP <= 0 {
			return IsWin(char, *enemy)
		}
	}
	return false
}

func PlayerTurn(char, enemy *data.Character) {
	for {
		fmt.Println("0[WAIT] 1[CHARGE] 2[ITEM] 3[SKILL]")

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

func EnemyTurn(c, e *data.Character, turn int) {
	if turn%2 == 0 {
		c.CurHP -= DealDammage(e, c, 5)
	}
}

func IsDead(c *data.Character) {
	fmt.Println("\n	🕱 THANATOS 🕱")
	sound.PlaySound("thanatos")
	c.CurHP = c.MaxHP / 2
	c.CurXP = 0
}

func IsWin(c *data.Character, e data.Character) bool {
	fmt.Println("	🙟🙦  YOU WON 🙤🙝")
	c.CurXP += e.CurXP
	c.Coins += e.Coins
	fmt.Println("+", e.CurXP, "XP / +", e.Coins, "🞠")
	IsLevelUp(c)
	return true
}

func IsLevelUp(c *data.Character) {
	if c.CurXP >= c.MaxXP {
		c.CurXP -= c.MaxXP
		c.MaxXP += c.MaxXP / 4
		c.Level += 1
		if c.Level%5 == 0 {
			c.Points += 2
		} else {
			c.Points += 1
		}
		fmt.Println("You leveled up !")
		IsLevelUp(c)
	}
}

func DealDammage(user, target *data.Character, dmg int) int {
	return int(float32(dmg) * (1 + (user.Atk / 10)))
}
