package redproject

import (
    "encoding/json"
    "os"
)

type M_Abilities struct {
    Name         string `json:"name"`
    Description  string `json:"description"`
    Damages      int    `json:"damages"`
    ManaCost     int    `json:"mana_cost,omitempty"`
}

type Personnage struct {
    Classe       string    `json:"name"`
    Description  string    `json:"description"`
    PV           int       `json:"pv"`
    Mana         int       `json:"mana"`
    Force        int       `json:"force"`
    Magic        int       `json:"magic"`
    Armure       int       `json:"armure"`
    MagicResist  int       `json:"magic_resist"`
    Vitesse      int       `json:"vitesse"`
    Abilities    []M_Abilities `json:"abilities"`
    NomPerso     string    // Ajouté dynamiquement ou fixé
}

// 🔮 Initialise le Magicien à partir du fichier JSON
func InitialiserMagicien() (Personnage, error) {
    var magicien Personnage

    contenu, err := os.ReadFile("aa_perso.json") // Assure-toi que ce fichier est bien dans le dossier racine
    if err != nil {
        return magicien, err
    }

    err = json.Unmarshal(contenu, &magicien)
    if err != nil {
        return magicien, err
    }

    magicien.NomPerso = "Lyra" // Nom imposé pour cette classe
    return magicien, nil
}
