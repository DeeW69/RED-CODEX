package redproject

import (
    "encoding/json"
    "os"
)

type E_Abilities struct {
    Name         string `json:"name"`
    Description  string `json:"description"`
    Damages      int    `json:"damages"`
    ManaCost     int    `json:"mana_cost,omitempty"`
}

type elfe struct {
    Classe       string    `json:"name"`
    Description  string    `json:"description"`
    PV           int       `json:"pv"`
    Mana         int       `json:"mana"`
    Force        int       `json:"force"`
    Magic        int       `json:"magic"`
    Armure       int       `json:"armure"`
    MagicResist  int       `json:"magic_resist"`
    Vitesse      int       `json:"vitesse"`
    Abilities    []E_Abilities `json:"abilities"`
    NomPerso     string    // Ajouté dynamiquement ou fixé
}

// 🔮 Initialise le Magicien à partir du fichier JSON
func InitialiserElfe() (Personnage, error) {
    var elfe Personnage

    contenu, err := os.ReadFile("aa_perso.json") // Assure-toi que ce fichier est bien dans le dossier racine
    if err != nil {
        return elfe, err
    }

    err = json.Unmarshal(contenu, &elfe)
    if err != nil {
        return elfe, err
    }

    elfe.NomPerso = "Elwynn" // Nom imposé pour cette classe
    return elfe, nil
}
