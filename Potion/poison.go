package Potion

import (
	"encoding/json"
	"fmt"
	"os"
)

type PoisonPotion struct {
	Ico             string  `json:"ico"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Value           int     `json:"value"`
	Damage          int     `json:"damage"`
	Duration        int     `json:"duration"`
	Effect          string  `json:"effect"`
	DamagePerSecond float64 `json:"damage_per_second"`
}

func LoadPoisonPotion() (*PoisonPotion, error) {
	file, err := os.ReadFile("Potion/aa_potion.json")
	if err != nil {
		return nil, fmt.Errorf("cannot read JSON file: %v", err)
	}

	var data map[string]map[string]PoisonPotion
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	potion, exists := data["potions"]["poison"]
	if !exists {
		return nil, fmt.Errorf("poison potion not found in JSON")
	}

	return &potion, nil
}

func (p *PoisonPotion) Use() {
	fmt.Printf("Utilisation de %s: %s\n", p.Name, p.Description)
	fmt.Printf("Dégâts initiaux: %d, Dégâts par seconde: %.1f pendant %d secondes\n",
		p.Damage, p.DamagePerSecond, p.Duration)
}

func (p *PoisonPotion) GetTotalDamage() float64 {
	return float64(p.Damage) + (p.DamagePerSecond * float64(p.Duration))
}
