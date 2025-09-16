package main

import (
	"RED_Project/game"
	"RED_Project/ui"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const saveFilePath = "data_players.json"

type ConsoleGame struct {
	session *game.Session
	reader  *bufio.Reader
}

func NewConsoleGame(session *game.Session) *ConsoleGame {
	return &ConsoleGame{
		session: session,
		reader:  bufio.NewReader(os.Stdin),
	}
}

func (g *ConsoleGame) Run() {
	ui.AnimatedWelcome()
	time.Sleep(1 * time.Second)
	ui.DisplayWelcomeArt()
	time.Sleep(1 * time.Second)

	for {
		ui.DisplayMainMenu()
		if err := g.waitForMenuInput(); err != nil {
			fmt.Println("\033[1;31mErreur de lecture :\033[0m", err)
			return
		}
	}
}

func (g *ConsoleGame) waitForMenuInput() error {
	input, err := g.reader.ReadString('\n')
	if err != nil {
		return err
	}
	g.handleMenuChoice(input)
	return nil
}

func (g *ConsoleGame) handleMenuChoice(choice string) {
	cleanedInput := strings.TrimSpace(choice)

	switch cleanedInput {
	case "1":
		fmt.Println("\n\033[1;32m✓ Chargement du jeu...\033[0m")
		g.startNewGame()
	case "2":
		fmt.Println("\n\033[1;33m🔍 Recherche des sauvegardes...\033[0m")
		g.loadGame()
	case "3":
		fmt.Println("\n\033[1;36m⚙️  Entrée dans les options...\033[0m")
		g.showOptions()
	case "4":
		fmt.Println("\n\033[1;31m❌ Vous quittez le jeu ? (OUI/NON)\033[0m")
		g.confirmQuit()
	case "5":
		fmt.Println("\n\033[1;36m🎒 Ouverture de l'inventaire...\033[0m")
		g.showInventory(true)
	default:
		fmt.Printf("\n\033[1;31m✗ Choix invalide: '%s'. Veuillez choisir 1-5.\033[0m\n", cleanedInput)
		time.Sleep(2 * time.Second)
	}
}

func (g *ConsoleGame) startNewGame() {
	if err := g.session.Reload(); err != nil {
		fmt.Println("\033[1;31m✗ Impossible de charger la sauvegarde :\033[0m", err)
		return
	}
	g.session.ResetForNewGame()
	_ = g.session.Save()

	ui.ShowIntro()
	ui.StartCastleAnimation()

	fmt.Println("\033[1;32m🎮 Nouvelle partie lancée !\033[0m")
	g.gameLoop()
}

func (g *ConsoleGame) loadGame() {
	if err := g.session.Reload(); err != nil {
		fmt.Println("\033[1;33mAucune sauvegarde trouvée.\033[0m")
		time.Sleep(2 * time.Second)
		return
	}

	fmt.Printf("\033[1;32mSauvegarde chargée. Bienvenue %s !\033[0m\n", g.session.Player.Player.Stats.Name)
	g.gameLoop()
}

func (g *ConsoleGame) showOptions() {
	fmt.Println("\n\033[1;36m════════════ OPTIONS ════════════\033[0m")
	fmt.Println("• 🔊 Volume: 80%")
	fmt.Println("• ⚔️  Difficulté: Normal")
	fmt.Println("• 🖥️  Résolution: 1920x1080")
	fmt.Println("• 🎨 Thème: Sombre")
	fmt.Println("• 🎵 Musique: Activée")
	fmt.Println("\033[1;36m══════════════════════════════════\033[0m")

	fmt.Print("\n\033[1;37mAppuyez sur Entrée pour retourner au menu...\033[0m")
	g.reader.ReadString('\n')
}

func (g *ConsoleGame) confirmQuit() {
	for {
		fmt.Print("\033[1;31mConfirmez (OUI/NON): \033[0m")
		confirmation, err := g.reader.ReadString('\n')
		if err != nil {
			fmt.Println("\n\033[1;31m✗ Lecture impossible.\033[0m")
			return
		}
		confirmation = strings.TrimSpace(strings.ToUpper(confirmation))

		switch confirmation {
		case "OUI", "O", "Y", "YES":
			fmt.Println("\n\033[1;35m👋 À bientôt !\033[0m")
			time.Sleep(1 * time.Second)
			os.Exit(0)
		case "NON", "N", "NO":
			fmt.Println("\n\033[1;32m✓ Retour au menu principal.\033[0m")
			time.Sleep(1 * time.Second)
			return
		default:
			fmt.Println("\n\033[1;31m✗ Réponse non reconnue.\033[0m")
		}
	}
}

func (g *ConsoleGame) showInventory(pause bool) {
	if g.session.Player == nil {
		fmt.Println("\033[1;31m✗ Aucune donnée joueur chargée.\033[0m")
		return
	}

	game.DisplayInventory(g.session.Player)

	if pause {
		fmt.Print("\n\033[1;37mAppuyez sur Entrée pour revenir...\033[0m")
		g.reader.ReadString('\n')
	}
}

func (g *ConsoleGame) chooseStage() (string, bool) {
	stages := game.ListStages()
	if len(stages) == 0 {
		fmt.Println("\033[1;31mAucun stage disponible pour le moment.\033[0m")
		return "", false
	}

	fmt.Println("\n\033[1;36m═══════════ STAGES DISPONIBLES ═══════════\033[0m")
	for i, stage := range stages {
		fmt.Printf("%d. Zone %d - %s (%s)\n", i+1, stage.Zone, stage.Name, stage.Timing)
	}
	fmt.Println("0. Retour")
	fmt.Print("Votre choix : ")

	input, _ := g.reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" || input == "0" {
		return "", false
	}

	idx, err := strconv.Atoi(input)
	if err != nil || idx < 1 || idx > len(stages) {
		fmt.Println("\033[1;31mChoix invalide.\033[0m")
		return "", false
	}

	return stages[idx-1].Slug, true
}

func (g *ConsoleGame) gameLoop() {
	for {
		fmt.Println("\n\033[1;36m═══════════════════════════\033[0m")
		fmt.Println("📜 Commandes disponibles :")
		fmt.Println("• Q = Quitter vers le menu")
		fmt.Println("• C = Choisir un stage")
		fmt.Println("• I = Inventaire")
		fmt.Println("• S = Faire du shopping")
		fmt.Println("• F = Améliorer l'équipement")
		fmt.Println("\033[1;36m═══════════════════════════\033[0m")
		fmt.Print("Votre choix : ")

		input, err := g.reader.ReadString('\n')
		if err != nil {
			fmt.Println("\033[1;31mErreur de lecture. Retour au menu.\033[0m")
			return
		}
		choice := strings.TrimSpace(strings.ToUpper(input))

		switch choice {
		case "I":
			g.showInventory(true)
		case "Q":
			fmt.Println("\033[1;31mRetour au menu principal...\033[0m")
			if err := g.session.Save(); err != nil {
				fmt.Println("\033[1;31mErreur de sauvegarde :\033[0m", err)
			}
			time.Sleep(1 * time.Second)
			return
		case "C":
			if stageSlug, ok := g.chooseStage(); ok {
				if err := game.StartBattle(g.session, stageSlug, g.reader); err != nil {
					fmt.Println("\033[1;31mErreur :\033[0m", err)
				}
			}
		case "S":
			fmt.Println("🛍️ Le marchand prépare encore son échoppe...")
			time.Sleep(1 * time.Second)
		case "F":
			game.LaunchForge(g.session)
		default:
			fmt.Println("\033[1;31m✗ Commande inconnue.\033[0m")
		}
	}
}

func main() {
	session, err := game.NewSession(saveFilePath)
	if err != nil {
		fmt.Println("\033[1;31mImpossible de lancer la session :\033[0m", err)
		return
	}

	consoleGame := NewConsoleGame(session)
	consoleGame.Run()
}
