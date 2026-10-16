package utils

import (
	"eris/data"
	"fmt"
)

func UseSkill(id string, u, t *data.Character) {
	switch id {
	case "rush":
		func(u, t *data.Character) {
			dmg := DealDammage(u, t, 5)
			t.CurHP -= dmg
			fmt.Println(u.Name, "charged on", t.Name, "for", dmg, "dammage.")
		}(u, t)

	case "assurance":
		func(u, t *data.Character) {
			t.CurHP = u.MaxHP * int(float32(u.CurHP)/float32(t.MaxHP))
			fmt.Println(u.Name, "has brought", t.Name, "'s health to", t.CurHP, "points.")
		}(u, t)

	case "benediction":
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

	case "kick":
		func(u, t *data.Character) {
			dmg := DealDammage(u, t, 8)
			fmt.Println(u.Name, "kicked", t.Name, "in the knees for", dmg, "dammage.")
			t.CurHP -= dmg * int((t.Atk/10)+1)
		}(u, t)

	case "":
		func(u, t *data.Character) {
		}(u, t)
	}
}
