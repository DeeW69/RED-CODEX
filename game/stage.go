package game

import (
    "math/rand"
    "time"
)

// ================== Structures ==================

type Enemy struct {
    Name   string
    Health int
    Attack int
}

type Stage struct {
    ID      int      `json:"id"`
    Zone    int
    Timing  string
    Name    string
    Enemies []Enemy
    Boss    *Enemy
    Drops   []string
}

// ================== Stages ==================

var Stages = map[string]Stage{
    "20": {
        ID:    1,
        Zone:  20,
        Timing: "Crépuscule",
        Name:  "Caverne",
        Enemies: []Enemy{
            {"Loup des cavernes", 15, 3},
            {"Chauve-souris", 8, 2},
            {"Ours grotteux", 30, 6},
        },
        Boss: &Enemy{
            Name:   "Magicien",
            Health: 70,
            Attack: 30,
        },
        Drops: []string{"peau", "crocs", "viande"},
    },
    "30": {
        ID:    2,
        Zone:  30,
        Timing: "Nuit",
        Name:  "Village abandonné",
        Enemies: []Enemy{
            {"Chien errant", 10, 2},
            {"Rat des égouts", 12, 3},
            {"Sanglier fou", 20, 4},
        },
        Boss: nil,
        Drops: []string{"os", "viande avariée"},
    },
    "40": {
        ID:    3,
        Zone:  40,
        Timing: "Brume",
        Name:  "Forêt",
        Enemies: []Enemy{
            {"Renard rusé", 14, 3},
            {"Ours brun", 28, 6},
        },
        Boss: &Enemy{
            Name:   "Troll",
            Health: 100,
            Attack: 40,
        },
        Drops: []string{"crocs", "plumes"},
    },
    // Ajoute ici les autres stages comme "60", "80", "100", etc.
    // en suivant le même modèle : boss un stage sur deux, drops, ennemis...
}

// ================== Fonctions ==================

func GenerateDrops(stageName string) []Drop {
    stage, ok := Stages[stageName]
    if !ok || len(stage.Drops) == 0 {
        return nil
    }
    rand.Seed(time.Now().UnixNano())

    nb := rand.Intn(3) + 1 // 1–3 objets
    out := make([]Drop, 0, nb)
    for i := 0; i < nb; i++ {
        item := stage.Drops[rand.Intn(len(stage.Drops))]
        qty := rand.Intn(3) + 1 // 1–3
        out = append(out, Drop{Item: item, Quantity: qty})
    }
    return out
}
