package skills

import (
	"eris/data"
	"eris/utils"
	"fmt"
)

func UseFunc(id string, u, t *data.Character) {
	switch id {
	case "rush":
		func(u, t *data.Character) {
			dmg := utils.DealDammage(u, t, 5)
			t.CurHP -= dmg
			fmt.Println(u.Name, "charged on", t.Name, "for", dmg, "dammage.")
		}(u, t)

	case "assurance":
		func(u, t *data.Character) {
			u.CurSP--
			t.CurHP = u.MaxHP * int(float32(u.CurHP)/float32(t.MaxHP))
			fmt.Println(u.Name, "has brought", t.Name, "'s health to", t.CurHP, "points.")
		}(u, t)

	case "benediction":
		func(u, t *data.Character) {
			if u.CurSP < 1 {
				fmt.Println("You don't have enough energy.")
				return
			}
			u.CurSP--
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
			if u.CurSP < 1 {
				fmt.Println("You don't have enough energy.")
				return
			}
			u.CurSP--
			dmg := utils.DealDammage(u, t, 8)
			fmt.Println(u.Name, "kicked", t.Name, "in the knees for", dmg, "dammage.")
			t.CurHP -= dmg * int((t.Atk/10)+1)
		}(u, t)

	case "":
		func(u, t *data.Character) {
		}(u, t)
	}
}
