package utils

import (
	"eris/data"
	"fmt"
)

func UseItem(id string, u, t *data.Character) {
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

	case "skillAssurance":
		func(u, t *data.Character) {
			t.CurHP = u.MaxHP * int(float32(u.CurHP)/float32(t.MaxHP))
			fmt.Println(u.Name, "has brought", t.Name, "'s health to", t.CurHP, "points.")
		}(u, t)

	case "skillBlessing":
		func(u, t *data.Character) {
			heal := u.MaxHP / 4
			u.CurHP += heal
			if u.CurHP >= u.MaxHP {
				u.CurHP = u.MaxHP
				fmt.Println(u.Name, "healed themselves fully.")
				return
			}
			fmt.Println(u.Name, "healed themselves", heal, "health.")
		}(u, t)

	case "skillKick":
		func(u, t *data.Character) {
			dmg := DealDammage(u, t, 8)
			fmt.Println(u.Name, "kicked", t.Name, "in the knees for", dmg, "dammage.")
			t.CurHP -= dmg * int((t.Atk/10)+1)
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
