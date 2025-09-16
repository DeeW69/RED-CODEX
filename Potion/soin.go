package Potion

import (
	"encoding/json"
	"fmt"
	"os"
)

type SoinPotion struct {
	Ico           string `json:"ico"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Value         int    `json:"value"`
	RestoreAmount int    `json:"restore_amount"`
	Effect        string `json:"effect"`
}

func LoadSoinPotion() (*SoinPotion, error) {
	file, err := os.ReadFile("Potion/aa_potion.json")
	if err != nil {
		return nil, fmt.Errorf("cannot read JSON file: %v", err)
	}

	var data map[string]map[string]SoinPotion
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	potion, exists := data["potions"]["soin"]
	if !exists {
		return nil, fmt.Errorf("soin potion not found in JSON")
	}

	return &potion, nil
}

func (p *SoinPotion) Use(currentHealth int) int {
	fmt.Printf("Utilisation de %s: %s\n", p.Name, p.Description)
	fmt.Printf("Vie restaurée: %d points\n", p.RestoreAmount)
	return currentHealth + p.RestoreAmount
}
