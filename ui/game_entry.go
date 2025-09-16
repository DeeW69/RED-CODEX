package ui

import (
	"RED_Project/tiles"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func CastleEntranceAnimation() {
	fmt.Print("\033[2J\033[H") // Clear screen

	// Animation du château qui s'agrandit
	frames := []string{
		// Frame 1 - Petit
		`
		╔═════════════════╗
		║                 ║
		║     CHÂTEAU     ║
		║                 ║
		╚═════════════════╝
        `,

		// Frame 2 - Moyen
		`
		╔═══════════════════════╗
		║                       ║
		║                       ║
		║        CHÂTEAU        ║
		║                       ║
		║                       ║
		╚═══════════════════════╝
        `,

		// Frame 3 - Grand
		`
		╔═══════════════════════════════════╗
		║                                   ║
		║                                   ║
		║                                   ║
		║              CHÂTEAU              ║
		║                                   ║
		║                                   ║
		║                                   ║
		╚═══════════════════════════════════╝
        `,

		// Frame 4 - Très grand
		`
		╔══════════════════════════════════════════════╗
		║                                              ║
		║                                              ║
		║                                              ║
		║                                              ║
		║                   CHÂTEAU                    ║
		║                                              ║
		║                                              ║
		║                                              ║
		║                                              ║
		╚══════════════════════════════════════════════╝
        `,

		// Frame 5 - Devant la porte
		`
					╔═════════════════════════════════════════════════╗
					║                                                 ║
					║                   CHÂTEAU                       ║
					║                                                 ║
					╠═══════════╦════════════════════════╦════════════╣
					║           ║                        ║            ║
					║   SHOP    ║                        ║    FORGE   ║
					║           ║                        ║            ║
					╚═══════════╩════════════════════════╩════════════╝
        `,
	}

	// Afficher l'animation d'agrandissement
	fmt.Println("\033[1;33m") // Couleur jaune/or
	for i, frame := range frames {
		fmt.Print("\033[2J\033[H")
		fmt.Println(frame)

		messages := []string{
			"Approche du château...",
			"Le château se révèle...",
			"Majestueux et imposant...",
			"Devant les portes ancestrales...",
			"Le héros arrive devant l'entrée...",
		}

		if i < len(messages) {
			fmt.Printf("\n\033[1;37m%s\033[0m\n", messages[i])
		}

		time.Sleep(600 * time.Millisecond)
	}

	// Message de bienvenue
	time.Sleep(1 * time.Second)
	fmt.Print("\033[2J\033[H")
	fmt.Println("\033[1;33m" +
		`
					╔═════════════════════════════════════════════════╗
					║                                                 ║
					║                   CHÂTEAU                       ║
					║                                                 ║
					╠═══════════╦════════════════════════╦════════════╣
					║           ║                        ║            ║
					║   SHOP    ║                        ║    FORGE   ║
					║           ║                        ║            ║
					╚═══════════╩════════════════════════╩════════════╝
        ` + "\033[0m")

	fmt.Println("\n\033[1;36m══════════════════════════════════════════════════╗")
	fmt.Println("║                                                      ║")
	fmt.Println("║  \"Bienvenue dans cette aventure chère voyageur       ║")
	fmt.Println("║   égaré ! Les portes de ce château renferment        ║")
	fmt.Println("║   des mystères et des dangers insoupçonnés...\"       ║")
	fmt.Println("║                                                      ║")
	fmt.Println("╚══════════════════════════════════════════════════════╝\033[0m")

	time.Sleep(3 * time.Second)

	// Demander si le joueur veut entrer
	askToEnter()
}

func askToEnter() {
	fmt.Print("\033[2J\033[H")
	fmt.Println("\033[1;33m" +
		`
					╔═════════════════════════════════════════════════╗
					║                                                 ║
					║                   CHÂTEAU                       ║
					║                                                 ║
					╠═══════════╦════════════════════════╦════════════╣
					║           ║                        ║            ║
					║   SHOP    ║                        ║    FORGE   ║
					║           ║                        ║            ║
					╚═══════════╩════════════════════════╩════════════╝
        ` + "\033[0m")

	fmt.Println("\n\033[1;35m══════════════════════════════════════════════════╗")
	fmt.Println("║                                                      ║")
	fmt.Println("║           Voulez-vous entrer ?                       ║")
	fmt.Println("║                                                      ║")
	fmt.Println("║           \033[1;32mOUI\033[1;35m - Affronter votre destin              ║")
	fmt.Println("║           \033[1;31mNON\033[1;35m - Revenir sur vos pas                ║")
	fmt.Println("║                                                      ║")
	fmt.Println("╚══════════════════════════════════════════════════════╝\033[0m")
	fmt.Print("\n\033[1;37mVotre choix (OUI/NON): \033[0m")

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToUpper(response))

	switch response {
	case "OUI", "O", "Y", "YES":
		enterCastle()
	case "NON", "N", "NO":
		refuseEntry()
	default:
		fmt.Println("\n\033[1;31mRéponse non comprise. Veuillez choisir.\033[0m")
		time.Sleep(1 * time.Second)
		askToEnter()
	}
}

func enterCastle() {
	// Animation d'entrée
	frames := []string{
		`
					╔═════════════════════════════════════════════════╗
					║                                                 ║
					║                   CHÂTEAU                       ║
					║                                                 ║
					╠═══════════╦════════════════════════╦════════════╣
					║           ║                        ║            ║
					║   SHOP    ║                        ║    FORGE   ║
					║           ║                        ║            ║
					║           ║                        ║            ║
					╚═══════════╬═══════╬╬╬╬╬╬╬╬╬════════╬════════════╝
					            ║       ╬  CAVE ╬        ║            
					            ║       ╬╬╬╬╬╬╬╬╬        ║            
						        ╚════════════════════════╝
        `,
	}

	fmt.Println("\033[1;32m") // Vert pour l'entrée
	for i, frame := range frames {
		fmt.Print("\033[2J\033[H")
		fmt.Println(frame)

		messages := []string{
			"La porte s'ouvre lentement...",
			"Un vieux sage vous accueille...",
			"L'aventure commence véritablement!",
		}

		fmt.Printf("\n\033[1;32m%s\033[0m\n", messages[i])
		time.Sleep(1000 * time.Millisecond)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("\n\033[1;36m✓ Vous avez choisi l'aventure !\033[0m")
}

func refuseEntry() {
	fmt.Print("\033[2J\033[H")
	fmt.Println("\033[1;33m" +
		`
					╔═════════════════════════════════════════════════╗
					║                                                 ║
					║                   CHÂTEAU                       ║
					║                                                 ║
					╠═══════════╦════════════════════════╦════════════╣
					║           ║                        ║            ║
					║   SHOP    ║                        ║    FORGE   ║
					║           ║                        ║            ║
					╚═══════════╩════════════════════════╩════════════╝
        ` + "\033[0m")

	fmt.Println("\n\033[1;31m══════════════════════════════════════════════════╗")
	fmt.Println("║                                                      ║")
	fmt.Println("║  \"Certaines opportunités ne se présentent qu'une     ║")
	fmt.Println("║   fois... Peut-être une autre fois, voyageur.\"       ║")
	fmt.Println("║                                                      ║")
	fmt.Println("╚══════════════════════════════════════════════════════╝\033[0m")

	time.Sleep(3 * time.Second)
	fmt.Println("\n\033[1;35mRetour au menu principal...\033[0m")
	time.Sleep(2 * time.Second)
}

// Fonction pour être appelée depuis main.go
func StartCastleAnimation() {
	CastleEntranceAnimation()
}

func visitForge() {
	// Remplacer le contenu par :
	tiles.ShowForge()
}

func visitShop() {
	// Remplacer le contenu par :
	tiles.ShowShop()
}
