package data

type Character struct {
	Id        string
	Name      string
	Class     string
	CurHP     int
	MaxHP     int
	CurSP     int
	MaxSP     int
	CurXP     int
	MaxXP     int
	Level     int
	Points    int
	Coins     int
	Atk       int
	Def       int
	InvSize   int
	Inventory []Item
	Equipment [4]Item
	Skills    [4]Skill
}

type Item struct {
	Id   string
	Name string
	Desc string
	Type string
	Icon string
	Mod  int
}

type Skill struct {
	Id   string
	Name string
	Desc string
	Icon string
	Cost int
}

type Merchant struct {
	Id       string
	Products []StoreItem
}
type StoreItem struct {
	Id       string
	Item     Item
	Quantity int
	Cost     int
}

type Blacksmith struct {
	Id    string
	Craft []ForgeItem
}
type ForgeItem struct {
	Id        string
	Item      Item
	Materials []Item
	Cost      int
}
