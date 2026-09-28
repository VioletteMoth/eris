package story

import (
	"eris/menus"
	"fmt"
	"os"
	"strings"
	"unicode"
)
'''
func StartStory() {
	fmt.Println("		🙟🙦🙧  E R I S 🙥🙤🙝")
	if _, err := os.Stat("save.json"); err == nil {
		newGame := menus.GameMainMenu(nil)
		if !newGame {
			return
		}
	}
	for {
		fmt.Println("\nYou wake up on the side of a street bathed in shadow. You struggle to get your mind right...")
		// Vous vous réveillez au
		// bord d'une route drapée d'ombres. Vous essayez devous remémmorerles récents évènements et...
		fmt.Println("First, what is your name ?") // D'abord, quel est votre nom ?'
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
			// Première lettre en majuscule, reste en minuscules.
			name = strings.ToLower(name)
			letters := []rune(name)
			if len(letters) != 0 {
				letters[0] = unicode.ToUpper(letters[0])
			}
			name = string(letters)
		}

		fmt.Println("Ah yes, " + name) //Ah oui, <nom>

		fmt.Print("\n")
		fmt.Println("As your hands enter your field of view you see...") //Alors que vos mains entrent votre champ de vision vous voyez...
		var clChoice int
		var clSelect string
		fmt.Println("1 > Wires sprouting out of metal plates.") // Des cables surgissant de plaques de métal
		fmt.Println("2 > A warm, soft and reassuring glow.")    // Une douce brillance, chaude et rassurante
		fmt.Println("3 > Cuts and wounds all over.")            // Des plaies et blessures de partout

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
		fmt.Println("That's right, you're a " + clSelect) // c'est vrai, t'es des <clSelect>

		player := InitCharacter(name, clSelect)
		merchant := InitMerchant()

		GameMenu(&player, &merchant)

		if !GameMainMenu(&player, &merchant) {
			return

		}
	}
}