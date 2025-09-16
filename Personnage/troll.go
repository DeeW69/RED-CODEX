package redproject

import (
    "encoding/json"
    "os"
)

type T_Abilities struct {
    Name         string `json:"name"`
    Description  string `json:"description"`
    Damages      int    `json:"damages"`
    ManaCost     int    `json:"mana_cost,omitempty"`
}

type troll struct {
    Classe       string    `json:"name"`
    Description  string    `json:"description"`
    PV           int       `json:"pv"`
    Mana         int       `json:"mana"`
    Force        int       `json:"force"`
    Magic        int       `json:"magic"`
    Armure       int       `json:"armure"`
    MagicResist  int       `json:"magic_resist"`
    Vitesse      int       `json:"vitesse"`
    Abilities    []T_Abilities `json:"abilities"`
    NomPerso     string    // Ajouté dynamiquement ou fixé
}

// 🔮 Initialise le Troll à partir du fichier JSON
func InitialiserTroll() (Personnage, error) {
    var troll Personnage

    contenu, err := os.ReadFile("aa_perso.json") // Assure-toi que ce fichier est bien dans le dossier racine
    if err != nil {
        return troll, err
    }

    err = json.Unmarshal(contenu, &troll)
    if err != nil {
        return troll, err
    }

    troll.NomPerso = "Brak" // Nom imposé pour cette classe
    return troll, nil
}
