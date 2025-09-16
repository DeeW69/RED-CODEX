package Potion

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type GoldPotion struct {
	Ico         string  `json:"ico"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       int     `json:"value"`
	Duration    int     `json:"duration"`
	Effect      string  `json:"effect"`
	Multiplier  float64 `json:"multiplier"`
}

func LoadGoldPotion() (*GoldPotion, error) {
	file, err := os.ReadFile("Potion/aa_potion.json")
	if err != nil {
		return nil, fmt.Errorf("cannot read JSON file: %v", err)
	}

	var data map[string]map[string]GoldPotion
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	potion, exists := data["potions"]["gold"]
	if !exists {
		return nil, fmt.Errorf("gold potion not found in JSON")
	}

	return &potion, nil
}

func (p *GoldPotion) Use() {
	fmt.Printf("Utilisation de %s: %s\n", p.Name, p.Description)
	fmt.Printf("Multiplicateur d'or: %.1fx pendant %d secondes\n", p.Multiplier, p.Duration)
}

func (p *GoldPotion) GetDuration() time.Duration {
	return time.Duration(p.Duration) * time.Second
}
