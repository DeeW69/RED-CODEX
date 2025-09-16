package game

import (
	"encoding/json"
	"fmt"
	"os"
)

// ================== Types de base ==================

type Weapon struct {
	ID        string         `json:"id"`
	Icon      string         `json:"icon"`
	Stats     map[string]int `json:"stats"`
	Level     int            `json:"level"`
	LevelStar int            `json:"level_star"`
	Quantity  int            `json:"quantity"`
}

type Potion struct {
	ID       string `json:"id"`
	Icon     string `json:"icon"`
	Effect   string `json:"effect"`
	Quantity int    `json:"quantity"`
}

type Companion struct {
	Name     string `json:"name"`
	Attack   int    `json:"attack"`   // dégâts par tour
	Unlocked bool   `json:"unlocked"` // débloqué ou non
}

// Un drop obtenu (utilisé par battle + stages)
type Drop struct {
	Item     string
	Quantity int
}
type Spell struct {
	Name  string `json:"name"`
	Power int    `json:"power"`
	Cost  int    `json:"cost"`
}

// ================== Données joueur (JSON) ==================

type PlayerData struct {
	Player struct {
		Stats struct {
			Name       string   `json:"name"`
			Level      int      `json:"level"`
			Experience int      `json:"experience"`
			Gold       int      `json:"gold"`
			Health     int      `json:"health"`
			MaxHealth  int      `json:"max_health"`
			Mana       int      `json:"mana"`
			MaxMana    int      `json:"max_mana"`
			Spells     []string `json:"spells"`
		} `json:"stats"`

		Equipment struct {
			Head struct {
				Name      string `json:"name"`
				LevelStar int    `json:"level_star"`
			} `json:"head"`
			Weapon struct {
				Name      string `json:"name"`
				LevelStar int    `json:"level_star"`
			} `json:"weapon"`
			Armor struct {
				Name      string `json:"name"`
				LevelStar int    `json:"level_star"`
			} `json:"armor"`
		} `json:"equipment"`

		Inventory struct {
			Potions []Potion       `json:"potions"`
			Weapons []Weapon       `json:"weapons"`
			Drops   map[string]int `json:"drops"`
			Gold    int            `json:"gold"`
		} `json:"inventory"`
		// AJOUT du GOBELIN
		Companions map[string]Companion `json:"companions,omitempty"`
	} `json:"player"`
}

// ================== Utilitaires ==================

func RenderStars(level int) string {
	colors := []string{
		"\033[37m", "\033[34m", "\033[33m", "\033[38;5;208m", "\033[31m", "\033[35m",
	}
	reset := "\033[0m"

	if level < 1 {
		level = 1
	}
	if level > 30 {
		level = 30
	}

	tier := (level - 1) / 5
	starsCount := (level-1)%5 + 1

	stars := ""
	for i := 1; i <= 5; i++ {
		if i <= starsCount {
			stars += "★"
		} else {
			stars += "☆"
		}
	}
	return fmt.Sprintf("%s%s%s", colors[tier], stars, reset)
}

// ================== I/O ==================

func LoadPlayer(filename string) (*PlayerData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var player PlayerData
	if err := json.Unmarshal(data, &player); err != nil {
		return nil, err
	}
	if player.Player.Inventory.Drops == nil {
		player.Player.Inventory.Drops = make(map[string]int)
	}
	if player.Player.Companions == nil {
		player.Player.Companions = make(map[string]Companion)
	}
	return &player, nil

}

func SavePlayer(filename string, player *PlayerData) error {
	data, err := json.MarshalIndent(player, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// ================== Affichage ==================

func DisplayInventory(p *PlayerData) {
	// Partie perso + potions
	fmt.Println("╔══════════════╗   ╔════════════════════════════════════════════════════════════╗")
	fmt.Printf("║ %-12s ║   ║                 Inventaire        Potions                  ║\n", p.Player.Stats.Name)
	fmt.Println("╠══════════════╣   ╠════════════════════════════════════════════════════════════╣")

	fmt.Printf("║ Tête : %-6s ║   ", p.Player.Equipment.Head.Name)
	for _, potion := range p.Player.Inventory.Potions {
		fmt.Printf("║   %s   ", potion.Icon)
	}
	fmt.Println("║")

	fmt.Printf("║     %s   ║   ", RenderStars(p.Player.Equipment.Head.LevelStar))
	for _, potion := range p.Player.Inventory.Potions {
		fmt.Printf("║ Effet: %-3s", potion.Effect)
	}
	fmt.Println("║")

	fmt.Printf("║ Arme : %-6s ║   ", p.Player.Equipment.Weapon.Name)
	for _, potion := range p.Player.Inventory.Potions {
		fmt.Printf("║ Nbs : x%-2d", potion.Quantity)
	}
	fmt.Println("║")

	fmt.Printf("║     %s   ║   ╚════════════════════════════════════════════════════════════╝\n",
		RenderStars(p.Player.Equipment.Weapon.LevelStar))

	fmt.Printf("║ Armure: %-6s ║\n", p.Player.Equipment.Armor.Name)
	fmt.Printf("║     %s   ║\n", RenderStars(p.Player.Equipment.Armor.LevelStar))
	fmt.Println("╚══════════════╝")

	// Armes
	fmt.Println("\n╔═══════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                            Inventaire        Armes                                 ║")
	fmt.Println("╠═══════════════════════════════════════════════════════════════════════════════════╣")
	for _, w := range p.Player.Inventory.Weapons {
		fmt.Printf("║  %s  ", w.Icon)
	}
	fmt.Println("║")
	for _, w := range p.Player.Inventory.Weapons {
		for stat, val := range w.Stats {
			fmt.Printf("║ +%d %s ", val, stat)
		}
	}
	fmt.Println("║")
	for _, w := range p.Player.Inventory.Weapons {
		fmt.Printf("║ Nv: %-3d %s", w.Level, RenderStars(w.LevelStar))
	}
	fmt.Println("║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════════════════════════╝")

	// Drops + Or
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║                Ressources & Drops            ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Printf("║ 💰 Or total : %d\n", p.Player.Stats.Gold+p.Player.Inventory.Gold)
	fmt.Println("║----------------------------------------------║")
	if len(p.Player.Inventory.Drops) == 0 {
		fmt.Println("║ Aucun drop encore récupéré.")
	} else {
		for item, qty := range p.Player.Inventory.Drops {
			fmt.Printf("║ %dx %s\n", qty, item)
		}
	}
	fmt.Println("╚══════════════════════════════════════════════╝")
}
