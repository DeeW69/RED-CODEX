package game

import (
	"math/rand"
	"sort"
)

// ================== Structures ==================

type Enemy struct {
	Name   string
	Health int
	Attack int
}

type Stage struct {
	ID      int      `json:"id"`
	Slug    string   `json:"slug"`
	Zone    int      `json:"zone"`
	Timing  string   `json:"timing"`
	Name    string   `json:"name"`
	Enemies []Enemy  `json:"enemies"`
	Boss    *Enemy   `json:"boss"`
	Drops   []string `json:"drops"`
}

// ================== Stages ==================

var Stages = map[string]Stage{
	"cave": {
		ID:     1,
		Slug:   "cave",
		Zone:   20,
		Timing: "Crépuscule",
		Name:   "Caverne",
		Enemies: []Enemy{
			{Name: "Loup des cavernes", Health: 15, Attack: 3},
			{Name: "Chauve-souris", Health: 8, Attack: 2},
			{Name: "Ours grotteux", Health: 30, Attack: 6},
		},
		Boss: &Enemy{
			Name:   "Magicien",
			Health: 70,
			Attack: 30,
		},
		Drops: []string{"peau", "crocs", "viande"},
	},
	"village": {
		ID:     2,
		Slug:   "village",
		Zone:   30,
		Timing: "Nuit",
		Name:   "Village abandonné",
		Enemies: []Enemy{
			{Name: "Chien errant", Health: 10, Attack: 2},
			{Name: "Rat des égouts", Health: 12, Attack: 3},
			{Name: "Sanglier fou", Health: 20, Attack: 4},
		},
		Drops: []string{"os", "viande avariée"},
	},
	"forest": {
		ID:     3,
		Slug:   "forest",
		Zone:   40,
		Timing: "Brume",
		Name:   "Forêt",
		Enemies: []Enemy{
			{Name: "Renard rusé", Health: 14, Attack: 3},
			{Name: "Ours brun", Health: 28, Attack: 6},
		},
		Boss: &Enemy{
			Name:   "Troll",
			Health: 100,
			Attack: 40,
		},
		Drops: []string{"crocs", "plumes"},
	},
	// Ajoute ici les autres stages comme "desert", "mountain", etc.
}

// ================== Fonctions ==================

// GetStage retourne les informations d'un stage à partir de son identifiant textuel.
func GetStage(slug string) (Stage, bool) {
	stage, ok := Stages[slug]
	if !ok {
		return Stage{}, false
	}
	if stage.Slug == "" {
		stage.Slug = slug
	}
	return stage, true
}

// ListStages renvoie les stages triés par zone puis identifiant.
func ListStages() []Stage {
	stages := make([]Stage, 0, len(Stages))
	for slug, stage := range Stages {
		if stage.Slug == "" {
			stage.Slug = slug
		}
		stages = append(stages, stage)
	}
	sort.Slice(stages, func(i, j int) bool {
		if stages[i].Zone == stages[j].Zone {
			return stages[i].ID < stages[j].ID
		}
		return stages[i].Zone < stages[j].Zone
	})
	return stages
}

// GenerateDrops calcule les récompenses potentielles d'un stage.
func GenerateDrops(stageSlug string) []Drop {
	stage, ok := GetStage(stageSlug)
	if !ok || len(stage.Drops) == 0 {
		return nil
	}

	nb := rand.Intn(3) + 1 // 1–3 objets
	out := make([]Drop, 0, nb)
	for i := 0; i < nb; i++ {
		item := stage.Drops[rand.Intn(len(stage.Drops))]
		qty := rand.Intn(3) + 1 // 1–3
		out = append(out, Drop{Item: item, Quantity: qty})
	}
	return out
}
