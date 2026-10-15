package items

import (
	"eris/data"
	"fmt"
)

func UseFunc(id string, u, t *data.Character) {
	switch id {
	case "elixir":
		func(u, t *data.Character) {
			t.CurHP += 10
			if t.CurHP > t.MaxHP {
				t.CurHP = t.MaxHP
				fmt.Println("+", t.CurHP-t.MaxHP, " HP")
				return
			}
			fmt.Println("+10 HP")
		}(u, t)

	case "poison":
		func(u, t *data.Character) {
			t.CurHP -= t.MaxHP / int(10/(1+(t.Def/10)))
			fmt.Println("-", t.MaxHP/int(10/(1+(t.Def/10))), " HP")
		}(u, t)

	case "skillBlessing":
		func(u, t *data.Character) {
			fmt.Print("Bless you !")
			t.CurHP += t.MaxHP / 6
			if t.CurHP >= t.MaxHP {
				t.CurHP = t.MaxHP
				fmt.Println("+", t.CurHP-t.MaxHP, " HP")
				return
			}
			fmt.Println("+10 HP")
		}(u, t)

	case "skillKick":
		func(u, t *data.Character) {
			fmt.Print("Kicking them!")
			t.CurHP -= 8 * (1 + (u.Atk-t.Def)/10)
		}(u, t)

	case "invSpace":
		func(u, t *data.Character) {
			fmt.Print("Inventory upgraded")
			u.InvSize++
		}(u, t)

	case "":
		func(u, t *data.Character) {
		}(u, t)
	}
}
