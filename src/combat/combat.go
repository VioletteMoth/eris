package combat

import (
	"eris/character"
	"eris/sound"
	"fmt"
	"strconv"
)

func Combat(c, e *character.Char) {
	fmt.Print("\n")
	fmt.Println("	🙟🙦  COMBAT START 🙤🙝")
	if e == nil {
		e = data.GetEnemy()
	}
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
		if CombatTurn(c, &e, advantage, turn) {
			return
		}
		turn++
		fmt.Print("\n")
	}
}

func CombatTurn(char, enemy *character.Char, advantage bool, turn int) bool {
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

func PlayerTurn(char, enemy *Character) {
	for {
		fmt.Println("0[WAIT] 1[CHARGE] 2[ITEM] 3[SKILL]")

		var choice int = -1
		for choice < 0 || choice > 3 {
			fmt.Print("> ")
			fmt.Scan(&choice)
		}
		switch choice {
		case 1:
			char.curSP += 2
			if char.curSP > char.maxSP {
				char.curSP = char.maxSP
			}
			return
		case 2:
			if AccessInventory(char, enemy) {
				return
			}
		case 3:
			if AccessSkills(char, enemy) {
				return
			}
		case 0:
			return
		}
	}
}

func EnemyTurn(char, enemy *character.Char, turn int) {
	if turn%2 == 0 {
		skills.SkillRush.effect(enemy, char)
	}
}

func IsDead(c *character.Char) {
	fmt.Println("\n	🕱 THANATOS 🕱")
	sound.PlaySound("thanatos")
	c.CurHP = c.MaxHP / 2
	c.CurXP = 0
}

func IsWin(c *character.Char, e character.Char) bool {
	fmt.Println("	🙟🙦  YOU WON 🙤🙝")
	c.CurXP += e.CurXP
	c.Coins += e.Coins
	fmt.Println("+", e.CurXP, "XP / +", e.Coins, "🞠")
	IsLevelUp(c)
	return true
}

func IsLevelUp(c *character.Char) {
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

func DealDammage(user, target *character.Char, dmg int) int {
	return int(float32(dmg) * (1 + (user.Atk / 10)))
}
