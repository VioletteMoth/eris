package character

import "eris/data"

func InitEnemy() data.Character {
	enemy := data.Character{
		Name:  "Dunmy",
		Class: "None",
		CurXP: 32,
		Coins: 13,
		CurHP: 15,
		MaxHP: 15,
		CurSP: 4,
		Atk:   1,
	}
	return enemy
}
