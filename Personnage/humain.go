package redproject

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

type H_Abilities struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Degats      int    `json:"degats"`
    ManaCost    int    `json:"mana_cost,omitempty"`
}

type Humain struct {
    Classe       string    `json:"name"`
    Description  string    `json:"description"`
    PV           int       `json:"pv"`
    Mana         int       `json:"mana"`
    Force        int       `json:"force"`
    Magic        int       `json:"magic"`
    Armure       int       `json:"armure"`
    MagicResist  int       `json:"magic_resist"`
    Vitesse      int       `json:"vitesse"`
    Abilities    []H_Abilities `json:"abilities"`
    NomPerso     string    // Ajouté dynamiquement par l'utilisateur
}

func chargerClasseHumain(fichier string) (Humain, error) {
    var h Humain
    contenu, err := os.ReadFile(fichier)
    if err != nil {
        return h, err
    }
    err = json.Unmarshal(contenu, &h)
    return h, err
}

func CreerPersonnage() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("📝 Entrez le nom de votre personnage : ")
    nomPerso, _ := reader.ReadString('\n')
    nomPerso = strings.TrimSpace(nomPerso)

    humain, err := chargerClasseHumain("humain.json")
    if err != nil {
        fmt.Println("❌ Erreur de chargement :", err)
        return
    }

    humain.NomPerso = nomPerso

    fmt.Printf("\n👤 Bienvenue %s, vous incarnez la classe %s\n", humain.NomPerso, humain.Classe)
    fmt.Println(humain.Description)
    fmt.Printf("📊 Stats : PV=%d | Mana=%d | Force=%d | Armure=%d | Résistance Magique=%d | Vitesse=%d\n",
        humain.PV, humain.Mana, humain.Force, humain.Armure, humain.MagicResist, humain.Vitesse)

    fmt.Println("\n🧙 Compétences disponibles :")
    for _, a := range humain.Abilities {
        fmt.Printf("- %s : %s (Dégâts: %d", a.Name, a.Description, a.Degats)
        if a.ManaCost > 0 {
            fmt.Printf(", Coût Mana: %d", a.ManaCost)
        }
        fmt.Println(")")
    }
}
