package data

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"

	"github.com/tidwall/gjson"
)

func GetItem(kind, id string) Item {
	file := GetJson("items")

	var ans Item
	ans.Id = gjson.GetBytes(file, string(kind+".#(id=="+id+").id")).String()
	ans.Name = gjson.GetBytes(file, string(kind+".#(id=="+id+").name")).String()
	ans.Desc = gjson.GetBytes(file, string(kind+".#(id=="+id+").desc")).String()
	ans.Type = gjson.GetBytes(file, string(kind+".#(id=="+id+").type")).String()
	ans.Icon = gjson.GetBytes(file, string(kind+".#(id=="+id+").icon")).String()
	ans.Mod = ConvertInt(file, kind, id, "mod")
	return ans
}

func GetSkill(id string) Skill {
	file := GetJson("skills")

	var ans Skill
	ans.Id = gjson.GetBytes(file, string("#(id=="+id+").id")).String()
	ans.Name = gjson.GetBytes(file, string("#(id=="+id+").name")).String()
	ans.Desc = gjson.GetBytes(file, string("#(id=="+id+").desc")).String()
	ans.Icon = gjson.GetBytes(file, string("#(id=="+id+").icon")).String()
	ans.Cost = ConvertInt(file, "", id, "cost")
	return ans
}

func GetEnemy(zone string) Character {
	file := GetJson("enemies")

	var list []string
	err := json.Unmarshal([]byte(gjson.GetBytes(file, string("zone.#(id=="+zone+").list")).String()), &list)
	fmt.Println(gjson.GetBytes(file, string("zone.#(id=="+zone+").list")).String())
	fmt.Println(list)
	if err != nil {
		return Character{Name: ""}
	}

	id := list[rand.Intn(len(list))]

	var ans Character
	ans.Id = gjson.GetBytes(file, string("#(id=="+id+").id")).String()
	ans.Name = gjson.GetBytes(file, string("enemy.#(id=="+id+").name")).String()
	ans.Class = gjson.GetBytes(file, string("enemy.#(id=="+id+").class")).String()
	ans.MaxHP = ConvertInt(file, "enemy", id, "health")
	ans.CurHP = ans.MaxHP
	ans.MaxSP = ConvertInt(file, "enemy", id, "energy")
	ans.CurSP = ans.MaxSP
	ans.CurXP = ConvertInt(file, "enemy", id, "exp")
	ans.Coins = ConvertInt(file, "enemy", id, "money")
	return ans
}

func GetMerc(id string) Merchant {
	file := GetJson("shop")

	var ans Merchant
	ans.Id = gjson.GetBytes(file, string("#(id=="+id+").id")).String()
	for i, c := range gjson.GetBytes(file, string("#(id=="+id+").inventory.")).Array() {
		ans.Products[i].Item = GetItem("", c.String())
		ans.Products[i].Cost = ConvertInt(file, "", id, "cost")
		ans.Products[i].Quantity = ConvertInt(file, "", id, "quantity")
	}
	return ans
}

func GetSmith(id string) Blacksmith {
	file := GetJson("forge")

	var ans Blacksmith
	ans.Id = gjson.GetBytes(file, string("#(id=="+id+").id")).String()
	for i, c := range gjson.GetBytes(file, string("#(id=="+id+").inventory.")).Array() {
		ans.Craft[i].Item = GetItem("equipment", c.String())
		for j, d := range gjson.GetBytes(file, string("#(id=="+id+").inventory.")).Array() {
			ans.Craft[i].Materials[j] = GetItem("materials", d.String())
		}
		ans.Craft[i].Cost = ConvertInt(file, "", id, "cost")
	}
	return ans
}

func GetJson(path string) []byte {
	filename, err := os.Open("data/files/" + path + ".json")
	if err != nil {
		ColorError.Println("ERROR:", err)
	}
	defer filename.Close()

	data, err := io.ReadAll(filename)
	if err != nil {
		ColorError.Println("ERROR:", err)
	}
	return data
}

func ConvertInt(file []byte, kind, id, arg string) int {
	var chk string
	var err error
	var ans int
	chk = gjson.GetBytes(file, string(kind+".#(id=="+id+")."+arg)).String()
	if chk != "" {
		ans, err = strconv.Atoi(chk)
		if err != nil {
			ColorError.Println("ERROR:", err)
		}
	}
	return ans
}
