package character

import (
	"eris/data"
	"eris/skills"
	"eris/utils"
	"fmt"
	"strconv"
)

func Skills(user, target *data.Character) bool {
	for {
		utils.ClearTerminal()
		fmt.Println("	🙟  SKILLS 🙝")
		for i, skill := range user.Skills {
			fmt.Print(strconv.Itoa(i+1) + "[" + string(skill.Icon) + " " + skill.Name + "] ")
		}
		fmt.Print("\n")
		fmt.Println("0 > Back")

		var skill int = -1
		for skill < 0 || 4 < skill {
			fmt.Print("> ")
			fmt.Scan(&skill)
		}
		switch {
		case 0 < skill && skill <= 4:
			if pickSkill(user, target, user.Skills[skill-1]) {
				return true
			}
		case skill == 0:
			fmt.Print("\n")
			return false
		}
	}
}

func pickSkill(user, target *data.Character, skill data.Skill) bool {
	for {
		fmt.Println(string(skill.Icon) + " " + skill.Name + " : " + skill.Desc)

		aim := "Away"
		if target != nil {
			aim = target.Name
		}
		fmt.Println("1 > Self	2 > " + aim)
		fmt.Println("0 > Back")

		var choice int = -1
		for choice < 0 || 2 < choice {
			fmt.Print("> ")
			fmt.Scan(&choice)
		}
		if user.CurSP >= skill.Cost {
			switch choice {
			case 0:
				fmt.Print("\n")
				return false
			case 1:
				fmt.Print(user.Name, " cast ", string(skill.Icon), " ", skill.Name, ", ")
				skills.UseFunc(skill.Id, user, user)
				fmt.Print("\n")
				return true
			case 2:
				if target == nil {
					fmt.Print(user.Name, " cast ", string(skill.Icon), " ", skill.Name, " away.")
				} else {
					fmt.Print(user.Name, " cast ", string(skill.Icon), " ", skill.Name, " on ", target.Name, ", ")
					skills.UseFunc(skill.Id, user, target)
				}
				fmt.Print("\n")
				return true
			}
		} else {
			fmt.Println("You don't have enough energy")
			return false
		}
	}
}
