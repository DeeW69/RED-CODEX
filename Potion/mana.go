package Potion

import (
	"encoding/json"
	"fmt"
	"os"
)

type ManaPotion struct {
	Ico           string `json:"ico"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Value         int    `json:"value"`
	RestoreAmount int    `json:"restore_amount"`
	Effect        string `json:"effect"`
}

func LoadManaPotion() (*ManaPotion, error) {
	file, err := os.ReadFile("Potion/aa_potion.json")
	if err != nil {
		return nil, fmt.Errorf("cannot read JSON file: %v", err)
	}

	var data map[string]map[string]ManaPotion
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	potion, exists := data["potions"]["mana"]
	if !exists {
		return nil, fmt.Errorf("mana potion not found in JSON")
	}

	return &potion, nil
}

func (p *ManaPotion) Use(currentMana int) int {
	fmt.Printf("Utilisation de %s: %s\n", p.Name, p.Description)
	fmt.Printf("Mana restauré: %d points\n", p.RestoreAmount)
	return currentMana + p.RestoreAmount
}
