package utils

import (
	"eris/data"
	"eris/sound"
	"fmt"
	"os"
	"os/exec"
)

func DealDammage(user, target *data.Character, dmg int) int {
	return dmg * (1 + ((user.Atk - target.Def) / 10))
}

func IsDead(player *data.Character) {
	ClearTerminal()
	fmt.Println("\n	🕱 THANATOS 🕱")
	sound.PlaySound("thanatos")
	player.CurHP = player.MaxHP / 2
	player.CurXP = 0

	var start int = -1
	fmt.Println("\nDon't calm down.")
	fmt.Println("0 > Rise again")
	for start != 0 {
		fmt.Print("> ")
		fmt.Scan(&start)
	}
}

func IsWin(p *data.Character, e data.Character) bool {
	fmt.Println("	🙟🙦  YOU WON 🙤🙝")
	p.CurXP += e.MaxXP
	p.Coins += e.Coins
	fmt.Println("+", e.CurXP, "XP / +", e.Coins, "🞠")
	isLevelUp(p)
	return true
}

func isLevelUp(p *data.Character) {
	if p.CurXP >= p.MaxXP {
		p.CurXP -= p.MaxXP
		p.MaxXP += p.MaxXP / 5
		p.Level += 1
		if p.Level%5 == 0 {
			p.Points += 2
		} else {
			p.Points += 1
		}
		fmt.Println("You leveled up !")
		isLevelUp(p)
	}
}

func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Print("\033[2J")
}
