package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Equipment struct {
	Head  string
	Chest string
	Feet  string
}

type Character struct {
	Name              string
	Class             string
	Level             int
	XP                int
	NextLevelXP       int
	MaxHP             int
	CurrentHP         int
	MaxMana           int
	Mana              int
	Gold              int
	Inventory         []string
	InventoryCapacity int
	InventoryUpgrades int
	Skills            []string
	Equipment         Equipment
}

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
	Turn      int
}

func InitCharacter() Character {
	return Character{
		Level:             1,
		XP:                0,
		NextLevelXP:       100,
		Gold:              100,
		InventoryCapacity: 10,
		Skills:            []string{"Coup de poing"},
	}
}

func CharacterCreation() Character {
	c := InitCharacter()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez votre nom: ")
	name, _ := reader.ReadString('\n')
	c.Name = strings.TrimSpace(name)
	for {
		fmt.Println("Choisissez votre classe: 1) Humain 2) Elfe 3) Nain")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			c.Class = "Humain"
			c.MaxHP = 100
			c.MaxMana = 30
		case "2":
			c.Class = "Elfe"
			c.MaxHP = 80
			c.MaxMana = 50
		case "3":
			c.Class = "Nain"
			c.MaxHP = 120
			c.MaxMana = 20
		default:
			fmt.Println("Choix invalide")
			continue
		}
		break
	}
	c.CurrentHP = c.MaxHP
	c.Mana = c.MaxMana
	return c
}

func (c *Character) DisplayInfo() {
	fmt.Printf("Nom: %s\nClasse: %s\nNiveau: %d\nXP: %d/%d\nPV: %d/%d\nMana: %d/%d\nOr: %d\n", c.Name, c.Class, c.Level, c.XP, c.NextLevelXP, c.CurrentHP, c.MaxHP, c.Mana, c.MaxMana, c.Gold)
	fmt.Println("Compétences:", strings.Join(c.Skills, ", "))
	fmt.Printf("Équipement: Tête:%s Torse:%s Pieds:%s\n", c.Equipment.Head, c.Equipment.Chest, c.Equipment.Feet)
	fmt.Println("Inventaire:")
	for i, item := range c.Inventory {
		fmt.Printf("%d: %s\n", i+1, item)
	}
	fmt.Printf("Capacité: %d/%d\n", len(c.Inventory), c.InventoryCapacity)
}

func (c *Character) AddInventory(item string) bool {
	if len(c.Inventory) >= c.InventoryCapacity {
		fmt.Println("Inventaire plein")
		return false
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Println(item, "ajouté à l'inventaire.")
	return true
}

func (c *Character) RemoveInventory(index int) {
	if index < 0 || index >= len(c.Inventory) {
		return
	}
	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
}

func (c *Character) TakePotion() {
	heal := 50
	c.CurrentHP += heal
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Printf("Vous récupérez %d PV. (%d/%d)\n", heal, c.CurrentHP, c.MaxHP)
}

func (c *Character) TakeManaPotion() {
	restore := 30
	c.Mana += restore
	if c.Mana > c.MaxMana {
		c.Mana = c.MaxMana
	}
	fmt.Printf("Vous récupérez %d Mana. (%d/%d)\n", restore, c.Mana, c.MaxMana)
}

func (c *Character) IsDead() bool {
	if c.CurrentHP <= 0 {
		fmt.Println("Vous êtes mort... mais vous vous relevez!")
		c.CurrentHP = c.MaxHP / 2
		return true
	}
	return false
}

func (c *Character) PoisonPotion(m *Monster) {
	fmt.Println("Vous utilisez une potion de poison!")
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		m.CurrentHP -= 10
		if m.CurrentHP < 0 {
			m.CurrentHP = 0
		}
		fmt.Printf("Le %s subit 10 dégâts de poison (%d/%d)\n", m.Name, m.CurrentHP, m.MaxHP)
		if m.CurrentHP == 0 {
			break
		}
	}
}

func (c *Character) SpellBook() {
	for _, s := range c.Skills {
		if s == "Boule de feu" {
			fmt.Println("Vous connaissez déjà Boule de feu.")
			return
		}
	}
	c.Skills = append(c.Skills, "Boule de feu")
	fmt.Println("Vous apprenez Boule de feu!")
}

func (c *Character) UpgradeInventorySlot() {
	if c.InventoryUpgrades >= 3 {
		fmt.Println("Capacité maximale déjà atteinte.")
		return
	}
	c.InventoryCapacity += 10
	c.InventoryUpgrades++
	fmt.Printf("Capacité augmentée à %d\n", c.InventoryCapacity)
}

func InitGoblin() Monster {
	return Monster{Name: "Gobelin", MaxHP: 60, CurrentHP: 60, Attack: 10}
}

func (m *Monster) GoblinPattern() int {
	m.Turn++
	dmg := m.Attack
	if m.Turn%3 == 0 {
		dmg = m.Attack * 2
		fmt.Println("Le gobelin devient furieux!")
	}
	return dmg
}

func SpendGold(c *Character, cost int) bool {
	if c.Gold < cost {
		fmt.Println("Pas assez d'or.")
		return false
	}
	c.Gold -= cost
	return true
}

func AccessInventory(c *Character, m *Monster, reader *bufio.Reader) {
	if len(c.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	for {
		fmt.Println("Inventaire:")
		for i, item := range c.Inventory {
			fmt.Printf("%d) %s\n", i+1, item)
		}
		fmt.Println("0) Retour")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "0" {
			return
		}
		idx, err := strconv.Atoi(input)
		if err != nil || idx < 1 || idx > len(c.Inventory) {
			fmt.Println("Choix invalide.")
			continue
		}
		item := c.Inventory[idx-1]
		fmt.Println("1) Utiliser 2) Supprimer 0) Annuler")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)
		switch action {
		case "1":
			switch item {
			case "Potion de vie":
				c.TakePotion()
				c.RemoveInventory(idx - 1)
			case "Potion de poison":
				if m != nil {
					c.PoisonPotion(m)
					c.RemoveInventory(idx - 1)
				} else {
					fmt.Println("Pas d'ennemi à empoisonner.")
				}
			case "Livre de Sort: Boule de feu":
				c.SpellBook()
				c.RemoveInventory(idx - 1)
			case "Augmentation d’inventaire":
				c.UpgradeInventorySlot()
				c.RemoveInventory(idx - 1)
			case "Potion de mana":
				c.TakeManaPotion()
				c.RemoveInventory(idx - 1)
			case "Matériaux":
				fmt.Println("Ce sont des matériaux de fabrication.")
			default:
				fmt.Println("Objet inconnu.")
			}
			if m != nil {
				return
			}
		case "2":
			c.RemoveInventory(idx - 1)
			fmt.Println("Objet supprimé.")
			if m != nil {
				return
			}
		default:
			continue
		}
	}
}

func DrinkPotionInBattle(c *Character, reader *bufio.Reader) {
	indices := []int{}
	for i, item := range c.Inventory {
		if item == "Potion de vie" || item == "Potion de mana" {
			indices = append(indices, i)
		}
	}
	if len(indices) == 0 {
		fmt.Println("Vous n'avez pas de potion.")
		return
	}
	fmt.Println("Choisissez une potion:")
	for i, idx := range indices {
		fmt.Printf("%d) %s\n", i+1, c.Inventory[idx])
	}
	fmt.Println("0) Annuler")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "0" {
		return
	}
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(indices) {
		fmt.Println("Choix invalide.")
		return
	}
	invIdx := indices[choice-1]
	item := c.Inventory[invIdx]
	switch item {
	case "Potion de vie":
		c.TakePotion()
	case "Potion de mana":
		c.TakeManaPotion()
	}
	c.RemoveInventory(invIdx)
}

func GenerateLoot() []string {
	lootTable := []string{"Matériaux", "Potion de vie", "Potion de mana", "Peau de gobelin"}
	n := rand.Intn(2) + 1
	loot := make([]string, 0, n)
	for i := 0; i < n; i++ {
		item := lootTable[rand.Intn(len(lootTable))]
		loot = append(loot, item)
	}
	return loot
}

func PlayerAttackMenu(c *Character, m *Monster, reader *bufio.Reader) {
	for {
		fmt.Println("Choisissez votre attaque:")
		fmt.Println("1) Attaque de base (5 dégâts)")
		fmt.Println("2) Coup de poing (8 dégâts)")
		fireball := false
		for _, s := range c.Skills {
			if s == "Boule de feu" {
				fireball = true
			}
		}
		if fireball {
			fmt.Println("3) Boule de feu (18 dégâts, 5 mana)")
		}
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			m.CurrentHP -= 5
			fmt.Printf("Vous infligez 5 dégâts. (%d/%d)\n", m.CurrentHP, m.MaxHP)
			return
		case "2":
			m.CurrentHP -= 8
			fmt.Printf("Coup de poing! (%d/%d)\n", m.CurrentHP, m.MaxHP)
			return
		case "3":
			if fireball {
				if c.Mana >= 5 {
					c.Mana -= 5
					m.CurrentHP -= 18
					fmt.Printf("Boule de feu! (%d/%d)\n", m.CurrentHP, m.MaxHP)
					return
				}
				fmt.Println("Pas assez de mana.")
			}
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func Battle(c *Character) {
	g := InitGoblin()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Un gobelin apparaît!")
	for {
		fmt.Printf("\n%s PV:%d/%d Mana:%d/%d\n", c.Name, c.CurrentHP, c.MaxHP, c.Mana, c.MaxMana)
		fmt.Printf("%s PV:%d/%d\n", g.Name, g.CurrentHP, g.MaxHP)
		fmt.Println("1) Attaquer 2) Boire potion 3) Utiliser Inventaire 4) Quitter")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			PlayerAttackMenu(c, &g, reader)
		case "2":
			DrinkPotionInBattle(c, reader)
		case "3":
			AccessInventory(c, &g, reader)
		case "4":
			fmt.Println("Vous fuyez le combat.")
			return
		default:
			fmt.Println("Choix invalide.")
			continue
		}
		if g.CurrentHP <= 0 {
			fmt.Println("Le gobelin est vaincu!")
			c.Gold += 10
			GainXP(c, 50)
			loot := GenerateLoot()
			fmt.Println("Butin obtenu:")
			for _, item := range loot {
				c.AddInventory(item)
			}
			return
		}
		dmg := g.GoblinPattern()
		c.CurrentHP -= dmg
		if c.CurrentHP < 0 {
			c.CurrentHP = 0
		}
		fmt.Printf("Le gobelin vous inflige %d dégâts! (%d/%d)\n", dmg, c.CurrentHP, c.MaxHP)
		if c.CurrentHP <= 0 {
			c.IsDead()
			fmt.Println("Le combat est terminé.")
			return
		}
	}
}

func CountItem(c *Character, item string) int {
	count := 0
	for _, it := range c.Inventory {
		if it == item {
			count++
		}
	}
	return count
}

func RemoveItem(c *Character, item string, n int) {
	removed := 0
	for i := 0; i < len(c.Inventory) && removed < n; {
		if c.Inventory[i] == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			removed++
		} else {
			i++
		}
	}
}

func BlacksmithMenu(c *Character) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- Forgeron ---")
		fmt.Printf("Or: %d\n", c.Gold)
		fmt.Println("Matériaux:", CountItem(c, "Matériaux"))
		fmt.Println("Peau de gobelin:", CountItem(c, "Peau de gobelin"))
		fmt.Println("1) Chapeau de l’aventurier (+10 PV max)")
		fmt.Println("2) Tunique de l’aventurier (+25 PV max)")
		fmt.Println("3) Bottes de l’aventurier (+15 PV max)")
		fmt.Println("4) Armure en peau de gobelin (+20 PV max)")
		fmt.Println("0) Retour")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		var bonus int
		var slot string
		var name string
		switch choice {
		case "1":
			bonus = 10
			slot = "Head"
			name = "Chapeau de l’aventurier"
		case "2":
			bonus = 25
			slot = "Chest"
			name = "Tunique de l’aventurier"
		case "3":
			bonus = 15
			slot = "Feet"
			name = "Bottes de l’aventurier"
		case "4":
			if CountItem(c, "Peau de gobelin") < 3 {
				fmt.Println("Pas assez de Peau de gobelin.")
				continue
			}
			if !SpendGold(c, 5) {
				continue
			}
			RemoveItem(c, "Peau de gobelin", 3)
			c.Equipment.Chest = "Armure en peau de gobelin"
			c.MaxHP += 20
			c.CurrentHP += 20
			fmt.Println("Armure en peau de gobelin fabriquée et équipée! PV max +20")
			continue
		case "0":
			return
		default:
			fmt.Println("Choix invalide.")
			continue
		}
		if CountItem(c, "Matériaux") < 2 {
			fmt.Println("Pas assez de matériaux.")
			continue
		}
		if !SpendGold(c, 5) {
			continue
		}
		RemoveItem(c, "Matériaux", 2)
		switch slot {
		case "Head":
			c.Equipment.Head = name
		case "Chest":
			c.Equipment.Chest = name
		case "Feet":
			c.Equipment.Feet = name
		}
		c.MaxHP += bonus
		c.CurrentHP += bonus
		fmt.Printf("%s fabriqué et équipé! PV max +%d\n", name, bonus)
	}
}

func MerchantMenu(c *Character) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- Marchand ---")
		fmt.Printf("Or: %d\n", c.Gold)
		fmt.Println("1) Potion de vie (3 or)")
		fmt.Println("2) Potion de poison (6 or)")
		fmt.Println("3) Livre de Sort : Boule de feu (25 or)")
		fmt.Println("4) Matériaux (1-7 or)")
		fmt.Println("5) Augmentation d’inventaire (30 or)")
		fmt.Println("6) Potion de mana (5 or)")
		fmt.Println("0) Retour")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			if SpendGold(c, 3) {
				if !c.AddInventory("Potion de vie") {
					c.Gold += 3
				}
			}
		case "2":
			if SpendGold(c, 6) {
				if !c.AddInventory("Potion de poison") {
					c.Gold += 6
				}
			}
		case "3":
			if SpendGold(c, 25) {
				if !c.AddInventory("Livre de Sort: Boule de feu") {
					c.Gold += 25
				}
			}
		case "4":
			cost := rand.Intn(7) + 1
			if SpendGold(c, cost) {
				if !c.AddInventory("Matériaux") {
					c.Gold += cost
				}
			}
		case "5":
			if SpendGold(c, 30) {
				if !c.AddInventory("Augmentation d’inventaire") {
					c.Gold += 30
				}
			}
		case "6":
			if SpendGold(c, 5) {
				if !c.AddInventory("Potion de mana") {
					c.Gold += 5
				}
			}
		case "0":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func InnMenu(c *Character) {
	const cost = 10
	if !SpendGold(c, cost) {
		return
	}
	c.CurrentHP = c.MaxHP
	c.Mana = c.MaxMana
	fmt.Println("Vous vous reposez à l'auberge. PV et Mana restaurés.")
}

func GainXP(c *Character, amount int) {
	c.XP += amount
	fmt.Printf("Vous gagnez %d XP.\n", amount)
	for c.XP >= c.NextLevelXP {
		c.XP -= c.NextLevelXP
		c.Level++
		c.NextLevelXP += 100
		c.MaxHP += 10
		c.CurrentHP = c.MaxHP
		c.MaxMana += 5
		c.Mana = c.MaxMana
		fmt.Printf("Niveau %d atteint! PV max +10, Mana max +5\n", c.Level)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := CharacterCreation()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- Menu Principal ---")
		fmt.Println("1) Afficher infos personnage")
		fmt.Println("2) Inventaire")
		fmt.Println("3) Marchand")
		fmt.Println("4) Forgeron")
		fmt.Println("5) Entraînement (combat gobelin)")
		fmt.Println("6) Auberge")
		fmt.Println("7) Interface graphique")
		fmt.Println("0) Quitter")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			c.DisplayInfo()
		case "2":
			AccessInventory(&c, nil, reader)
		case "3":
			MerchantMenu(&c)
		case "4":
			BlacksmithMenu(&c)
		case "5":
			Battle(&c)
		case "6":
			InnMenu(&c)
		case "7":
			LaunchGUI()
		case "0":
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
