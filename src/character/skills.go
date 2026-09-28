package character

import (
	"eris/data"
	"fmt"
	"strconv"
)

func Skills(user, target *data.Character) bool {
	if target != nil {
		for {
			fmt.Println("\n	🙟  SKILLS 🙝")
			for i, skill := range user.Skills {
				fmt.Print(strconv.Itoa(i+1) + "[" + string(skill.Icon) + " " + skill.Name + "] ")
			}
			fmt.Print("\n")
			fmt.Println("0[BACK]	5[INFO]")

			var skill int = -1
			for skill < 0 || 5 < skill {
				fmt.Print("> ")
				fmt.Scan(&skill)
			}
			switch {
			case 0 < skill && skill <= 4:
				if UseSkill(user, target, skill) {
					return true
				}
			case skill == 5:
				DescSkill(user)
			case skill == 0:
				fmt.Print("\n")
				return false
			}
		}
	} else {
		for {
			fmt.Println("\n	🙟  SKILLS 🙝")
			for _, skill := range user.Skills {
				fmt.Print("[" + string(skill.Icon) + " " + skill.Name + "] ")
			}
			fmt.Print("\n")
			fmt.Println("0[BACK]	1[INFO]")

			var skill int = -1
			for skill != 0 && 1 != skill {
				fmt.Print("> ")
				fmt.Scan(&skill)
			}
			switch {
			case skill == 1:
				DescSkill(user)
			case skill == 0:
				fmt.Print("\n")
				return false
			}
		}
	}
}

func DescSkill(user *data.Character) {
	for _, skill := range user.Skills {
		fmt.Println(string(skill.Icon) + " " + skill.Name + " : " + skill.Desc)
	}
}

func UseSkill(user, target *data.Character, skill int) bool {
	if user.Skills[skill-1].Name != "" {
		fmt.Print("\n")
		return true
	} else {
		fmt.Println("There is no skill in that slot.")
		fmt.Print("\n")
		return false
	}
}
