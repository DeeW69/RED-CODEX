package Potion

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ExpPotion struct {
	Ico         string `json:"ico"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       int    `json:"value"`
	Duration    int    `json:"duration"`
	Effect      string `json:"effect"`
	Boost       int    `json:"boost"`
}

func LoadExpPotion() (*ExpPotion, error) {
	file, err := os.ReadFile("Potion/aa_potion.json")
	if err != nil {
		return nil, fmt.Errorf("cannot read JSON file: %v", err)
	}

	var data map[string]map[string]ExpPotion
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	potion, exists := data["potions"]["exp"]
	if !exists {
		return nil, fmt.Errorf("exp potion not found in JSON")
	}

	return &potion, nil
}

func (p *ExpPotion) Use() {
	fmt.Printf("Utilisation de %s: %s\n", p.Name, p.Description)
	fmt.Printf("Boost d'expérience: +%d%% pendant %d secondes\n", p.Boost, p.Duration)
}

func (p *ExpPotion) GetDuration() time.Duration {
	return time.Duration(p.Duration) * time.Second
}
