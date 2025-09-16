package redproject

import (
    "encoding/json"
    "os"
)

type G_Abilities struct {
    Name         string `json:"name"`
    Description string `json:"description"`
    Damages      int    `json:"damages"`
    ManaCost     int    `json:"mana_cost,omitempty"`
}

type grimline struct {
    Classe       string    `json:"name"`
    Description  string    `json:"description"`
    PV           int       `json:"pv"`
    Mana         int       `json:"mana"`
    Force        int       `json:"force"`
    Magic        int       `json:"magic"`
    Armure       int       `json:"armure"`
    MagicResist  int       `json:"magic_resist"`
    Vitesse      int       `json:"vitesse"`
    Abilities    []G_Abilities `json:"abilities"`
    NomPerso     string    // Ajouté dynamiquement ou fixé
}

// 🔮 Initialise le Magicien à partir du fichier JSON
func InitialiseGrimline() (Personnage, error) {
    var grimline Personnage

    contenu, err := os.ReadFile("aa_perso.json") // Assure-toi que ce fichier est bien dans le dossier racine
    if err != nil {
        return grimline, err
    }

    err = json.Unmarshal(contenu, &grimline)
    if err != nil {
        return grimline, err
    }

    grimline.NomPerso = "Nox" // Nom imposé pour cette classe
    return grimline, nil
}
