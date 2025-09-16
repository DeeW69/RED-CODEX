package tiles

import (
	"encoding/json"
	"fmt"
	"os"
)

type ForgeShopData struct {
	ForgeItems map[string]ShopItem `json:"forge_items"`
	ShopItems  map[string]ShopItem `json:"shop_items"`
}

type ShopItem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Type        string `json:"type"`
	Bonus       int    `json:"bonus,omitempty"`
	Effect      string `json:"effect,omitempty"`
}

type PlayerData struct {
	Name       string         `json:"name"`
	Level      int            `json:"level"`
	Experience int            `json:"experience"`
	Gold       int            `json:"gold"`
	Health     int            `json:"health"`
	MaxHealth  int            `json:"max_health"`
	Mana       int            `json:"mana"`
	MaxMana    int            `json:"max_mana"`
	Equipment  map[string]int `json:"equipment"`
	Inventory  map[string]int `json:"inventory"`
}

func LoadForgeShopData() (*ForgeShopData, error) {
	file, err := os.ReadFile("tiles/aa_forge_shop.json")
	if err != nil {
		return nil, fmt.Errorf("erreur lecture forge/shop: %v", err)
	}

	var data ForgeShopData
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, fmt.Errorf("erreur parsing JSON: %v", err)
	}

	return &data, nil
}

func LoadPlayerData() (*PlayerData, error) {
	file, err := os.ReadFile("data_players.json")
	if err != nil {
		return nil, fmt.Errorf("erreur lecture joueur: %v", err)
	}

	var playerData PlayerData
	err = json.Unmarshal(file, &playerData)
	if err != nil {
		return nil, fmt.Errorf("erreur parsing JSON: %v", err)
	}

	return &playerData, nil
}

func SavePlayerData(playerData *PlayerData) error {
	data, err := json.MarshalIndent(map[string]interface{}{
		"player": playerData,
	}, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur création JSON: %v", err)
	}

	err = os.WriteFile("data_players.json", data, 0644)
	if err != nil {
		return fmt.Errorf("erreur écriture fichier: %v", err)
	}

	return nil
}

func BuyItem(playerData *PlayerData, itemID string, itemData *ShopItem, itemType string) bool {
	if playerData.Gold < itemData.Price {
		fmt.Printf("\033[1;31mPas assez d'or! Il vous faut %d pièces.\033[0m\n", itemData.Price)
		return false
	}

	playerData.Gold -= itemData.Price

	if itemType == "forge" {
		// Pour les améliorations d'équipement
		playerData.Equipment[itemID] += itemData.Bonus
		playerData.Inventory[itemID]++
		fmt.Printf("\033[1;32m%s acheté! +%d %s\033[0m\n", itemData.Name, itemData.Bonus, getStatName(itemID))
	} else {
		// Pour les consommables
		playerData.Inventory[itemID]++
		fmt.Printf("\033[1;32m%s acheté! Ajouté à l'inventaire.\033[0m\n", itemData.Name)
	}

	return true
}

func getStatName(itemID string) string {
	switch itemID {
	case "sword_upgrade":
		return "attaque"
	case "armor_upgrade":
		return "défense"
	case "shield_upgrade":
		return "protection"
	default:
		return "stat"
	}
}
