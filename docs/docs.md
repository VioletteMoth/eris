# DOCS ERIS

---

### Variables

```
var ColorMachina = color.RGB(0, 255, 255)
var ColorDeus = color.RGB(255, 255, 0)
var ColorSarta = color.RGB(255, 0, 255)

var ColorError = color.RGB(255, 0, 0)
var ColorUnavailable = color.RGB(85, 85, 85)
```

Colors are used to precise purpose of affected text

---

### Functions

#### func AccessInventory

```
func AccessInventory(player, target *Character) bool
```

Displays the contents of the parameter player's inventory attribute and gives access to PickInvItem().

#### func AccessSkills

```
func AccessSkills(player, target *Character) bool
```

Displays the contents of the parameter player's skillList attribute and gives access to DescSkill() and UseSkill().

#### func Blacksmith

```
func Blacksmith(player *Character)
```

Displays the recipes of the Items and gives access to CraftItem().

#### func BuyProduct

```
func BuyProduct(player *Character, product *Product)
```

Displays if the parameter player's coin attribute is smaller than parameter product's price attribute, else displays if player's inventory attribute is full, else adds product's item attribute is added to player's inventory attribute. 

#### func Combat

```
func Combat(char *Character)
```

Displays start of fight and alls InitEnemy() to create var enemy Character, then checks for class advantage and creates var turn int. \
For parameter char curHP attribute > 0 and enemy curHP attribute > 0, calls CombatTurn() and increments turn.

#### func CombatTurn

```
func CombatTurn(char *Character, enemy *Character, advantage bool, turn int) bool 
```

Displays parameter turn, parameter enemy's class, name and health attributes, parameter char's class, name, curHP, maxHP, curSP, maxSP. \
Calls PlayerTurn() if parameter advantage is true, else calls EnemyTurn() and then PlayerTurn(). \
After each PlayerTurn(), calls IsWin() if parameter enemy's curHP <= 0. \ 
After each EnemyTurn(), calls IsDead() if parameter char's curHP <= 0.

#### func CraftItem

```
func CraftItem(player *Character, equipment Item, materials []Item)
```

Checks if parameter player's inventory attribut has the contents of parameter materials and removes them and decreases player's coin attribute to add parameter equipment to player's inventory attribute. \
Displays if the materials are missing or the obtained Item.

#### func DealDammage

```
func DealDammage(user, target *Character, dmg int) int
```

Returns an int equal to parameter dmg * int((parameter user's atk attribute / 10) + 1).

#### func DescSkill

```
func DescSkill(char *Character)
```

Displays the icon, name and description attributes of the items in parameter char's skill attribute.

#### func DisplayUpgrade

```
func DisplayUpgrade(c *Character)
```

Displays the upgrade with parameter c's points attribute, health, energy and attack choices, incrementing selected attribute.

#### func EnemyTurn

```
func EnemyTurn(char *Character, enemy *Character, turn int)
```

Calls var skillRush's effect() attribute if parameter turn is even.

#### func EquipmentMenu

```
func EquipmentMenu(player *Character) 
```

Displays the equipment and the items in parameter player's equipment attribute, and calls ManageEquipment().

#### func equipementName

```
func equipmentName(item Item) string
```

Returns parameter item's name attribute, and "None" if it is an empty string.

#### func GameHomeMenu

```
func GameMainMenu(player *Character, merchant *Merchant) bool
```

Creates merchant if nil and displays home menu, new game and resume and restarts game or calls GameMenu().

#### func GameMenu

```
func GameMenu(player *Character, merchant *Merchant)
```

Displays a menu and gives access to Character.DisplayInfo(), AccessInventory(), EquipmentMenu(), AccessSkills(), Merchant.DisplayInfo(), Blacksmith(), Combat, DisplayUpgrade().

#### func IsDead

```
func IsDead(char *Character)
```

Displays death, calls PlaySound() and sets parameter char's curHP attribute to half of the maxHP attribute.

#### func IsLevelUp

```
func IsLevelUp(c *Character) 
```

If parameter c's curXP attribute > maxXP attribute, substracts maxXP from curXP attribute and increments level attribute by 2 if it is multiple of 5, 1 else, displays level up and calls itself.

#### func IsWin

```
func IsWin(c *Character, e Character) bool
```

Displays win, increments parameter c's curXP attribute by parameter e's curXP attribute and c's coins attribute by e's coins attribute, displays e's curXP and coins attributes and calls IsLevelUp() before returning true.

#### func ManageEquipment

```
func ManageEquipment(player *Character, equipped *Item, slot string)
```

Displays the parameter slot, calls equipmentName() for parameter *equipped, and displays the items in parameter player's inventory attribute were the slot attribute == parameter slot. \
Displays that nothing is equipped if trying to unequip empty, if it can't unequip because player's inventory attribute is full, else, tha it has been returned to inventory. \
Equips selected item and returns equipped to inventory.

#### func PickInvItem

```
func PickInvItem(char, target *Character, item int) bool
```

isplays parameter char's inventory attribute at index parameter item-1 's icon, name and description attributes. \
If selected Item is an item, anf and target != nil, calls effect() attribute on target and removes it from cher's inventory attribute. Else, removes it from inventory.

#### func PlaySound

```
func PlaySound(path string)
```

Opens file.mp3 in sounds folder where name is parameter path, decodes the file, and plays it to completion.

#### func UseSkill

```
func UseSkill(char, target *Character, skill int) bool
```

Displays an error if parameter char's skills attribute name attribute is empty and return false, else calls char's skill effect() attribute and returns true.

---

### Types

#### type Character

```
type Character struct {
	name      string
	class     string
	curHP     int
	maxHP     int
	curSP     int
	maxSP     int
	curXP     int
	maxXP     int
	level     int
	points    int
	coins     int
	atk       float32
	skills    [4]Skill
	skillList []Skill
	invSize   int
	inventory []Item
	equipment Equipment
}
```

The type Character is used for both the player character and enemies.

#### func InitCharacter

```
func InitCharacter(nm, cl string) Character
```

Initiate and returns a starting character with chosen name and class.

#### func InitEnemy

```
func InitEnemy() Character
```

Initiates and returns a starting enemy

#### func (*Character) DisplayInfo

```
func (c Character) DisplayInfo()
```

Displays the name, class, level, points, curXP, maxXP, curHP, maxHP, curSP, maxSP and atk attributes with the lengh of inventory and number of slots filled.

#### type Equipment

```
type Equipment struct {
	weapon Item
	head   Item
	chest  Item
	feet   Item
}
```

Equipment holds 4 Items for the player to use. 

#### type Item

```
 type Item struct { 
	name   string
	desc   string
	icon   rune
	slot   string // Used for an equipment or material
	effect func(char *Character) // Used for an item
}
```

Item is used for items, equipment, and crafting materials.

#### type Merchant

```
type Merchant struct {
	objects   []Product
	weapons   []Product
	armor     []Product
	materials []Product
}
```

Merchant allows for easier display of the categories of items.

#### func InitMerchant()

```
func InitMerchant() Merchant
```

Initiates a fixed Merchant.

#### func (*Merchant) DisplayInfo

```
func (c *Merchant) DisplayInfo(player *Character)
```

Displays the list of item types of *Merchant to access a []Product.

#### func (*Merchant) DisplayProducts

```
func (c *Merchant) DisplayProducts(player *Character, title string, products products *[]Products)
```

Displays the list of items in a *[]Product.

#### type Product

```
type Product struct {
	item     Item
	price    int
	quantity int
}
```

The type Product is used to ease the use of the Merchant.

#### type Skill

```
type Skill struct {
	name   string
	desc   string
	icon   rune
	effect func(c, t *Character)
}
```
Skill is similar to an item of type Item with another parameter to the effect() attribute.
