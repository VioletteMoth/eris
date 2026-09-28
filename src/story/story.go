package story

import (
	"eris/character"
	"eris/menus"
	"fmt"
	"strings"
	"unicode"
)

func StartStory() {
	fmt.Println("		🙟🙦🙧  E R I S 🙥🙤🙝")
	for {
		fmt.Println("\nYou wake up on the side of a street bathed in shadow. You struggle to get your mind right...")
		fmt.Println("First, what is your name ?")

		var name string
		for name == "" {
			fmt.Print("> ")
			fmt.Scan(&name)
			tname := ""
			for _, r := range name {
				if 'a' <= r && r <= 'z' {
					tname += string(r)
				} else if 'A' <= r && r <= 'Z' {
					tname += string(rune(r - 'A' + 'a'))
				}
			}
			name = tname
			name = strings.ToLower(name)
			letters := []rune(name)
			if len(letters) != 0 {
				letters[0] = unicode.ToUpper(letters[0])
			}
			name = string(letters)
		}

		fmt.Println("Ah yes, " + name) //Ah oui, <nom>

		fmt.Println("\nAs your hands enter your field of view you see...")
		var clChoice int
		var clSelect string
		fmt.Println("1 > Wires sprouting out of metal plates.")
		fmt.Println("2 > A warm, soft and reassuring glow.")
		fmt.Println("3 > Cuts and wounds all over.")

		for clChoice != 1 && clChoice != 2 && clChoice != 3 {
			fmt.Print("> ")
			fmt.Scan(&clChoice)
		}

		switch clChoice {
		case 1:
			clSelect = "Machina"
		case 2:
			clSelect = "Deus"
		case 3:
			clSelect = "Sarta"
		}
		fmt.Println("That's right, you're a " + clSelect)

		player := character.InitCharacter(name, clSelect)

		menus.Menu(&player, "zone1")

		if !menus.HomeMenu(&player) {
			return

		}
	}
}
