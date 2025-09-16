package Equipement

import (
	"encoding/json"
	"fmt"
	"os"
)

type ArmeHumain struct {
	Nom         string `json:"nom"`
	Classe      string `json:"classe"`
	Type        string `json:"type"`
	Rarete      string `json:"rarete"`
	Degats      int    `json:"degats"`
	Mana        int    `json:"mana"`
	PV          int    `json:"pv"`
	Vitesse     int    `json:"vitesse"`
	Armure      int    `json:"armure"`
	MagicResist int    `json:"magic_resist"`
	Description string `json:"description"`
}

func ChargerArmesHumain() (map[string]ArmeHumain, error) {
	file, err := os.ReadFile("Equipement/aa_equipement.json")
	if err != nil {
		return nil, fmt.Errorf("cannot read JSON file: %v", err)
	}

	var data map[string]map[string]map[string]ArmeHumain
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	armes, exists := data["armes"]["humain"]
	if !exists {
		return nil, fmt.Errorf("armes humain not found in JSON")
	}

	return armes, nil
}

func AfficherArme(arme ArmeHumain) {
	couleurs := map[string]string{
		"GrisClair": "\033[1;37m",
		"Bleu":      "\033[1;34m",
		"Jaune":     "\033[1;33m",
		"Orange":    "\033[1;31m",
		"Rouge":     "\033[1;35m",
		"Violet":    "\033[1;95m",
	}

	etoileRarete := map[string]string{
		"GrisClair": "★",
		"Bleu":      "★★",
		"Jaune":     "★★★",
		"Orange":    "★★★★",
		"Rouge":     "★★★★★",
		"Violet":    "★★★★★★",
	}

	couleur := couleurs[arme.Rarete]
	if couleur == "" {
		couleur = "\033[1;37m"
	}
	etoiles := etoileRarete[arme.Rarete]
	if etoiles == "" {
		etoiles = "?"
	}

	fmt.Printf("%s╔════════════════════════════════════════════╗\n", couleur)
	fmt.Println("║", arme.Classe)
	fmt.Printf("╠════════════════════════════════════════════╣\n")
	fmt.Printf("║ Description: %-23s \n", etoiles)
	fmt.Printf("║ Rareté: %-28s  \n", arme.Rarete)
	fmt.Printf("║ Dégats: %-19d \n", arme.Degats)
	fmt.Printf("║ Vitesse: %-27d \n", arme.Vitesse)
	fmt.Printf("║ Mana: %-24d \n", arme.Mana)
	fmt.Printf("╚═══════════════════════════════════════════╝\n")
}
