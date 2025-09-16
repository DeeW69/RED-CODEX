package main

import (
	"RED_Project/game"
	"RED_Project/ui"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// ==================== MENU HANDLER ====================

func handleMenuChoice(choice string) {
	// Nettoyer l'input (enlever les espaces et retours à la ligne)
	cleanedInput := strings.TrimSpace(choice)

	switch cleanedInput {
	case "1":
		fmt.Println("\n\033[1;32m✓ Chargement du jeu...\033[0m")
		startNewGame()

	case "2":
		fmt.Println("\n\033[1;33m🔍 Recherche des sauvegardes...\033[0m")
		loadGame()

	case "3":
		fmt.Println("\n\033[1;36m⚙️  Entrée dans les options...\033[0m")
		showOptions()

	case "4":
		fmt.Println("\n\033[1;31m❌ Vous quittez le jeu ? (OUI/NON)\033[0m")
		confirmQuit()

	case "5":
		fmt.Println("\n\033[1;36m🎒 Ouverture de l'inventaire...\033[0m")
		showInventory()

	default:
		fmt.Printf("\n\033[1;31m✗ Choix invalide: '%s'. Veuillez choisir 1-5.\033[0m\n", cleanedInput)
		time.Sleep(2 * time.Second)
		// Re-afficher le menu
		ui.DisplayMainMenu()
		waitForMenuInput()
	}
}

func waitForMenuInput() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	handleMenuChoice(input)
}

// ==================== GAME HANDLERS ====================

func startNewGame() {
	fmt.Println("\n\033[1;32m✓ Chargement du jeu...\033[0m")
	time.Sleep(1 * time.Second)
	//intro avant le chateau
	ui.ShowIntro()
	// Animation d'entrée dans le château
	ui.StartCastleAnimation()

	fmt.Println("\033[1;32m🎮 Nouvelle partie lancée !\033[0m")

	// On démarre la boucle de jeu (sans retourner au menu)
	gameLoop()
}

func loadGame() {
	// Simulation de recherche
	fmt.Print("\033[1;33m")
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(400 * time.Millisecond)
	}
	fmt.Println("\033[0m")
	fmt.Println("\033[1;33mAucune sauvegarde trouvée.\033[0m")
	time.Sleep(2 * time.Second)

	ui.DisplayMainMenu()
	waitForMenuInput()
}

func showOptions() {
	fmt.Println("\n\033[1;36m════════════ OPTIONS ════════════\033[0m")
	fmt.Println("• 🔊 Volume: 80%")
	fmt.Println("• ⚔️  Difficulté: Normal")
	fmt.Println("• 🖥️  Résolution: 1920x1080")
	fmt.Println("• 🎨 Thème: Sombre")
	fmt.Println("• 🎵 Musique: Activée")
	fmt.Println("\033[1;36m══════════════════════════════════\033[0m")

	fmt.Print("\n\033[1;37mAppuyez sur Entrée pour retourner au menu...\033[0m")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	ui.DisplayMainMenu()
	waitForMenuInput()
}

func confirmQuit() {
	fmt.Print("\033[1;31mConfirmez (OUI/NON): \033[0m")

	reader := bufio.NewReader(os.Stdin)
	confirmation, _ := reader.ReadString('\n')
	confirmation = strings.TrimSpace(strings.ToUpper(confirmation))

	switch confirmation {
	case "OUI", "O", "Y", "YES":
		fmt.Println("\n\033[1;35m👋 À bientôt !\033[0m")
		time.Sleep(1 * time.Second)
		os.Exit(0)

	case "NON", "N", "NO":
		fmt.Println("\n\033[1;32m✓ Retour au menu principal.\033[0m")
		time.Sleep(1 * time.Second)
		ui.DisplayMainMenu()
		waitForMenuInput()

	default:
		fmt.Println("\n\033[1;31m✗ Réponse non reconnue.\033[0m")
		confirmQuit()
	}
}

// ==================== INVENTAIRE HANDLER ====================

func showInventory() {
	player, err := game.LoadPlayer("data_players.json") // fichier JSON à la racine
	if err != nil {
		fmt.Println("\033[1;31m✗ Erreur chargement inventaire:", err, "\033[0m")
		return
	}

	game.DisplayInventory(player)

	fmt.Print("\n\033[1;37mAppuyez sur Entrée pour revenir au jeu...\033[0m")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func gameLoop() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n\033[1;36m═══════════════════════════\033[0m")
		fmt.Println("📜 Commandes disponibles :")
		fmt.Println("• Q = Quitter vers le menu")
		fmt.Println("• C = Continuer l'aventure")
		fmt.Println("\033[1;36m═══════════════════════════\033[0m")
		fmt.Println("• I = Inventaire")
		fmt.Println("• S = Faire du shopping")
		fmt.Println("• F = Ameliorer l'équipement")
		fmt.Println("\033[1;36m═══════════════════════════\033[0m")
		fmt.Print("Votre choix : ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(strings.ToUpper(input))

		switch choice {
		case "I":
			showInventory() // ouvre l'inventaire JSON
		case "Q":
			fmt.Println("\033[1;31mRetour au menu principal...\033[0m")
			time.Sleep(1 * time.Second)
			ui.DisplayMainMenu()
			waitForMenuInput()
			return
		case "C":
			fmt.Println("\033[1;32m⚔️ Vous avancez dans la caverne...\033[0m")
			player, err := game.LoadPlayer("data_players.json")
			if err != nil {
				fmt.Println("Erreur:", err)
				return
			}
			game.StartBattle(player, "cave", "data_players.json")

		case "F":
			fmt.Println("🛠️ Vous allez voir le forgeron...")

			// Appelle le script Python (forgeron_ui.py)
			cmd := exec.Command("python3", "game/forgeron_ui.py")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Erreur lors de l’ouverture du forgeron :", err)
			} else {
				fmt.Println("Résultat du forgeron :", string(output))
			}
		default:
			fmt.Println("\033[1;31m✗ Commande inconnue.\033[0m")
		}
	}
}

// ==================== MAIN ====================
func main() {
	// Afficher l'animation de bienvenue
	ui.AnimatedWelcome()
	time.Sleep(1 * time.Second)

	// Afficher l'écran de bienvenue fixe
	ui.DisplayWelcomeArt()
	time.Sleep(1 * time.Second)

	// Afficher le menu principal et attendre l'input
	ui.DisplayMainMenu()
	waitForMenuInput()
}
