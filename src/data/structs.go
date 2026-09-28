package data

type Character struct {
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
	Atk       float32
	InvSize   int
	Inventory []Item
	Equipment [4]Item
	Skills    [4]Skill
}

type Item struct {
	Name string
	Desc string
	Type string
	Icon string
	Mod  int
}

type Skill struct {
	Name string
	Desc string
	Icon string
	Cost int
}

type Merchant struct {
	Products []StoreItem
}
type StoreItem struct {
	Item     Item
	Quantity int
	Cost     int
}

type Blacksmith struct {
	Craft []ForgeItem
}
type ForgeItem struct {
	Item      Item
	Materials []Item
	Cost      int
}
