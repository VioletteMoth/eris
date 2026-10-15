package story

import (
	"eris/character"
	"eris/menus"
	"eris/utils"
	"fmt"
	"strings"
	"unicode"
)

func StartStory() {
	utils.ClearTerminal()
	fmt.Println("		🙟🙦🙧  E R I S 🙥🙤🙝")
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
	fmt.Println("Ah yes, " + name + ".")

	fmt.Println("\nAs your hands enter your field of view you see...")
	var class int = -1
	fmt.Println("1 > Wires sprouting out of metal plates.")
	fmt.Println("2 > A warm, soft and reassuring glow.")
	fmt.Println("3 > Cuts and wounds all over.")
	fmt.Println("0 > ...")
	for class < 0 && class < 4 {
		fmt.Print("> ")
		fmt.Scan(&class)
	}
	player := character.InitCharacter(name, class)
	if player.Class != "" {
		fmt.Println("That's right, you're a " + player.Class + ".")
	} else {
		fmt.Println("I'm so sorry for you.")
	}

	fmt.Println("\nNow is time to go.")
	var start int = -1
	fmt.Println("0 > Get up")
	for start != 0 {
		fmt.Print("> ")
		fmt.Scan(&start)
	}
	menus.Menu(&player, "zone1")

	if !menus.HomeMenu(&player) {
		return
	}
}
