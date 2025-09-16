package game

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

// ====== Combat d’une “vague” (1 ennemi) avec aide éventuelle du gobelin ======
func fightWave(player *PlayerData, enemy Enemy) bool {
	fmt.Println("⚔️  Un combat commence !")
	fmt.Printf("➡️ Ennemi: %s (HP:%d ATK:%d)\n", enemy.Name, enemy.Health, enemy.Attack)

	// Le gobelin aide-t-il ?
	gob, hasGob := player.Player.Companions["gobelin"]
	gobAtk := 0
	if hasGob && gob.Unlocked && gob.Attack > 0 {
		gobAtk = gob.Attack
		fmt.Printf("🟢 Le familier %s vous assiste (ATK +%d / tour)\n", gob.Name, gobAtk)
	}

	for enemy.Health > 0 && player.Player.Stats.Health > 0 {
		// Joueur attaque
		dmg := rand.Intn(5) + 3 // 3..7
		enemy.Health -= dmg
		fmt.Printf("%s frappe %s pour %d dégâts !\n", player.Player.Stats.Name, enemy.Name, dmg)

		// Gobelin attaque (s’il est débloqué)
		if gobAtk > 0 && enemy.Health > 0 {
			enemy.Health -= gobAtk
			fmt.Printf("👺 %s frappe %s pour %d dégâts !\n", gob.Name, enemy.Name, gobAtk)
		}

		// Ennemi riposte
		if enemy.Health > 0 {
			player.Player.Stats.Health -= enemy.Attack
			fmt.Printf("%s riposte pour %d dégâts !\n", enemy.Name, enemy.Attack)
		}
	}

	if player.Player.Stats.Health <= 0 {
		fmt.Println("💀 Vous êtes mort...")
		return false
	}

	fmt.Printf("✅ %s a été vaincu !\n", enemy.Name)
	return true
}

func savePlayer(filename string, player *PlayerData) error {
	data, err := json.MarshalIndent(player, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// ====== Lancement d’un stage (interactif + gobelin) ======
func StartBattle(player *PlayerData, stageName string, filename string) {
	stage, ok := Stages[stageName]
	if !ok {
		fmt.Println("⚠️ Stage inconnu :", stageName)
		return
	}

	fmt.Printf("\n🏰 Stage Zone %d (%s)\n", stage.Zone, stage.Timing)
	if len(stage.Enemies) == 0 {
		fmt.Println("Pas d’ennemis ici. Exploration libre.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	firstVictory := false

	for i, e := range stage.Enemies {
		fmt.Printf("\n⚔️  Combat contre %s !\n", e.Name)
		if !fightWave(player, e) {
			fmt.Println("\n💀 Vous avez échoué au stage...")
			return
		}

		// -------- Déblocage du gobelin après la 1ère victoire en "cave" --------
		if !firstVictory && (stageName == "cave" || stageName == "caverne") {
			// Débloque le gobelin
			player.Player.Companions["gobelin"] = Companion{
				Name:     "Gobelin",
				Attack:   2, // dégâts par tour
				Unlocked: true,
			}
			firstVictory = true
			fmt.Println("✨ Un gobelin reconnaissant vous rejoint ! (ATK +2 / tour)")
		}
		// ----------------------------------------------------------------------

		// Loots de ce combat
		drops := GenerateDrops(stageName)
		if len(drops) > 0 {
			fmt.Print("🎁 Vous obtenez : ")
			for _, d := range drops {
				fmt.Printf("%dx %s ", d.Quantity, d.Item)
				if player.Player.Inventory.Drops == nil {
					player.Player.Inventory.Drops = make(map[string]int)
				}
				player.Player.Inventory.Drops[d.Item] += d.Quantity
			}
			fmt.Println()
		}
		// Gain d’or séparé
		goldGain := rand.Intn(16) + 5 // 5–20
		player.Player.Inventory.Gold += goldGain
		fmt.Printf("💰 Vous trouvez %d or ! (total: %d)\n", goldGain, player.Player.Inventory.Gold+player.Player.Stats.Gold)

		// Si ce n’était pas le dernier ennemi → proposer une action au joueur
		if i < len(stage.Enemies)-1 {
			for {
				fmt.Print("\n➤ Action (C=Continuer / I=Inventaire / F=Forge / Q=Quitter) : ")
				choice, _ := reader.ReadString('\n')
				choice = strings.TrimSpace(strings.ToUpper(choice))

				switch choice {
				case "C":
					goto NEXT_ENEMY
				case "I":
					DisplayInventory(player)
					// on reboucle pour reproposer l’action
				case "Q":
					fmt.Println("🚪 Vous quittez le stage et retournez au menu.")
					_ = savePlayer(filename, player)
					return
				case "F":
					fmt.Println("🛠️ Vous allez voir le forgeron...")

					// Appel du script Python
					cmd := exec.Command("python3", "forgeron_ui.py")
					output, err := cmd.CombinedOutput()

					if err != nil {
						fmt.Println("❌ Erreur lors de l’ouverture du forgeron :", err)
					} else {
						fmt.Println("📦 Données reçues du forgeron :", string(output))

						// Exemple de structure JSON attendue : {"action":"buy_spellbook","success":true,"remaining_gold":70}
						var result map[string]interface{}
						if err := json.Unmarshal(output, &result); err != nil {
							fmt.Println("❌ Erreur parsing JSON :", err)
						} else {
							// Mise à jour de l’or du joueur
							if gold, ok := result["remaining_gold"].(float64); ok {
								player.Player.Stats.Gold = int(gold) // <= au lieu de player.Gold
								fmt.Println("💰 Or du joueur mis à jour :", player.Player.Stats.Gold)
							}
							// Ajout d’un sort
							if spell, ok := result["spell"].(string); ok {
								player.Player.Stats.Spells = append(player.Player.Stats.Spells, spell) // <= dans Stats
								fmt.Println("✨ Nouveau sort appris :", spell)
							}
						}
					}

				default:
					fmt.Println("❓ Saisie invalide. Tape C, I ou Q.")
				}
			}
		}
	NEXT_ENEMY:
	}

	// // Or en fin de stage
	// rand.Seed(time.Now().UnixNano())
	// goldGain := rand.Intn(21) + 10 // 10–30
	// player.Player.Inventory.Gold += goldGain
	// fmt.Printf("💰 Vous gagnez %d or (total affiché: %d)\n",
	// 	goldGain, player.Player.Stats.Gold+player.Player.Inventory.Gold)

	// if err := savePlayer(filename, player); err != nil {
	// 	fmt.Println("❌ Erreur sauvegarde :", err)
	// }
	// fmt.Println("✅ Stage terminé :", stage.Name)
}
