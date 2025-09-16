package game

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// fightWave gère un affrontement contre un ennemi unique.
func fightWave(player *PlayerData, enemy Enemy) bool {
	fmt.Println("⚔️  Un combat commence !")
	fmt.Printf("➡️ Ennemi: %s (HP:%d ATK:%d)\n", enemy.Name, enemy.Health, enemy.Attack)

	gob, hasGob := player.Player.Companions["gobelin"]
	gobAtk := 0
	if hasGob && gob.Unlocked && gob.Attack > 0 {
		gobAtk = gob.Attack
		fmt.Printf("🟢 Le familier %s vous assiste (ATK +%d / tour)\n", gob.Name, gobAtk)
	}

	stats := &player.Player.Stats

	for enemy.Health > 0 && stats.Health > 0 {
		dmg := rand.Intn(5) + 3 // 3..7
		enemy.Health -= dmg
		fmt.Printf("%s frappe %s pour %d dégâts !\n", stats.Name, enemy.Name, dmg)

		if gobAtk > 0 && enemy.Health > 0 {
			enemy.Health -= gobAtk
			fmt.Printf("👺 %s frappe %s pour %d dégâts !\n", gob.Name, enemy.Name, gobAtk)
		}

		if enemy.Health > 0 {
			stats.Health -= enemy.Attack
			if stats.Health < 0 {
				stats.Health = 0
			}
			fmt.Printf("%s riposte pour %d dégâts ! (PV restants: %d/%d)\n", enemy.Name, enemy.Attack, stats.Health, stats.MaxHealth)
		}
	}

	if stats.Health <= 0 {
		fmt.Println("💀 Vous êtes mort...")
		return false
	}

	fmt.Printf("✅ %s a été vaincu !\n", enemy.Name)
	fmt.Printf("❤️ Santé actuelle : %d/%d\n", stats.Health, stats.MaxHealth)
	return true
}

// StartBattle lance un stage complet et applique les récompenses.
func StartBattle(session *Session, stageSlug string, reader *bufio.Reader) error {
	if session == nil || session.Player == nil {
		return errors.New("aucune session de jeu active")
	}
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}

	stage, ok := GetStage(stageSlug)
	if !ok {
		return fmt.Errorf("stage inconnu : %s", stageSlug)
	}

	player := session.Player
	stats := &player.Player.Stats
	if stats.Health <= 0 {
		stats.Health = stats.MaxHealth
	}

	fmt.Printf("\n🏰 Stage Zone %d (%s) - %s\n", stage.Zone, stage.Timing, stage.Name)
	if len(stage.Enemies) == 0 && stage.Boss == nil {
		fmt.Println("Pas d’ennemis ici. Exploration libre.")
		return nil
	}

	firstVictory := false

	for idx, enemy := range stage.Enemies {
		fmt.Printf("\n⚔️  Combat contre %s !\n", enemy.Name)
		if !fightWave(player, enemy) {
			fmt.Println("\n💀 Vous avez échoué au stage...")
			_ = session.Save()
			return nil
		}

		if !firstVictory {
			if stage.Slug == "cave" && session.UnlockCompanion("gobelin", Companion{
				Name:   "Gobelin",
				Attack: 2,
			}) {
				fmt.Println("✨ Un gobelin reconnaissant vous rejoint ! (ATK +2 / tour)")
			}
			firstVictory = true
		}

		drops := GenerateDrops(stage.Slug)
		if len(drops) > 0 {
			fmt.Print("🎁 Vous obtenez : ")
			for _, d := range drops {
				fmt.Printf("%dx %s ", d.Quantity, d.Item)
			}
			fmt.Println()
			session.AddDrops(drops)
		}

		goldGain := rand.Intn(16) + 5 // 5–20
		session.AddGold(goldGain)
		fmt.Printf("💰 Vous trouvez %d or ! (total: %d)\n", goldGain, session.CurrentGold())

		if idx < len(stage.Enemies)-1 {
		actionLoop:
			for {
				fmt.Print("\n➤ Action (C=Continuer / I=Inventaire / F=Forge / Q=Quitter) : ")
				choice, _ := reader.ReadString('\n')
				choice = strings.TrimSpace(strings.ToUpper(choice))

				switch choice {
				case "C":
					break actionLoop
				case "I":
					DisplayInventory(player)
				case "Q":
					fmt.Println("🚪 Vous quittez le stage et retournez au menu.")
					_ = session.Save()
					return nil
				case "F":
					LaunchForge(session)
				default:
					fmt.Println("❓ Saisie invalide. Tape C, I, F ou Q.")
				}
			}
		}
	}

	if stage.Boss != nil {
		boss := *stage.Boss
		fmt.Printf("\n👑 Boss final : %s (HP:%d ATK:%d)\n", boss.Name, boss.Health, boss.Attack)
		if !fightWave(player, boss) {
			fmt.Println("\n💀 Vous avez échoué contre le boss...")
			_ = session.Save()
			return nil
		}
	}

	if err := session.Save(); err != nil {
		fmt.Println("❌ Erreur sauvegarde :", err)
	} else {
		fmt.Println("💾 Progression sauvegardée.")
	}
	fmt.Printf("✅ Stage terminé : %s\n", stage.Name)
	return nil
}

// LaunchForge exécute l'interface Python du forgeron et applique le résultat.
func LaunchForge(session *Session) {
	fmt.Println("🛠️ Vous allez voir le forgeron...")
	cmd := exec.Command("python3", filepath.Join("game", "forgeron_ui.py"))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("❌ Erreur lors de l’ouverture du forgeron :", err)
		if len(output) > 0 {
			fmt.Println(string(output))
		}
		return
	}

	fmt.Println("📦 Données reçues du forgeron :", string(output))

	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		fmt.Println("❌ Erreur parsing JSON :", err)
		return
	}

	if session != nil && session.Player != nil {
		if gold, ok := result["remaining_gold"].(float64); ok {
			session.Player.Player.Stats.Gold = int(gold)
			fmt.Println("💰 Or du joueur mis à jour :", session.Player.Player.Stats.Gold)
		}
		if spell, ok := result["spell"].(string); ok && spell != "" {
			session.Player.Player.Stats.Spells = append(session.Player.Player.Stats.Spells, spell)
			fmt.Println("✨ Nouveau sort appris :", spell)
		}
		if err := session.Save(); err != nil {
			fmt.Println("❌ Impossible de sauvegarder après la forge :", err)
		}
	}
}

// Deprecated: utilisé par d'anciens appels, conservé pour compatibilité.
func StartBattleLegacy(player *PlayerData, stageName string, filename string) {
	session, err := NewSession(filename)
	if err != nil {
		fmt.Println("⚠️ Stage legacy indisponible :", err)
		return
	}
	session.Player = player
	_ = StartBattle(session, stageName, nil)
}
