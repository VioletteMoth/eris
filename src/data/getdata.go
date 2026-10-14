package data

import (
	"io"
	"os"
	"strconv"

	"github.com/tidwall/gjson"
)

func GetItem(kind, id string) Item {
	file := GetJson("items")

	var ans Item
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
	ans.Name = gjson.GetBytes(file, string("#(id=="+id+").name")).String()
	ans.Desc = gjson.GetBytes(file, string("#(id=="+id+").desc")).String()
	ans.Icon = gjson.GetBytes(file, string("#(id=="+id+").icon")).String()
	ans.Cost = ConvertInt(file, "", id, "cost")
	return ans
}

func GetEnemy(id string) Character {
	file := GetJson("enemies")

	var ans Character
	ans.Name = gjson.GetBytes(file, string("#(id=="+id+").name")).String()
	ans.Class = gjson.GetBytes(file, string("#(id=="+id+").desc")).String()
	ans.CurHP = ConvertInt(file, "", id, "max_hp")
	ans.MaxHP = ConvertInt(file, "", id, "max_hp")
	return ans
}

func GetMerc(id string) Merchant {
	file := GetJson("shop")

	var ans Merchant
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
