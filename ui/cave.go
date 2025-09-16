package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func CaveEntranceAnimation() {
	fmt.Print("\033[2J\033[H") // Clear screen

	// Animation du château qui s'agrandit
	frames := []string{
		// Frame 1 - Petit
		`
		╔═════════════════╗
		║                 ║
		║      CAVE       ║
		║                 ║
		╚═════════════════╝
        `,

		// Frame 2 - Moyen
		`
		╔═══════════════════════╗
		║                       ║
		║                       ║
		║         CAVE          ║
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
		║               CAVE                ║
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
		║                    CAVE                      ║
		║                                              ║
		║                                              ║
		║                                              ║
		║                                              ║
		╚══════════════════════════════════════════════╝
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
					║                     CAVE                        ║
					║                                                 ║
					╚═════════════════════════════════════════════════╝
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

func askToEnterToTheCave() {
	fmt.Print("\033[2J\033[H")
	fmt.Println("\033[1;33m" +
		`
					╔═════════════════════════════════════════════════╗
					║                                                 ║
					║                     CAVE                        ║
					║                                                 ║
					╚═════════════════════════════════════════════════╝
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

func enterCave() {
	// Animation d'entrée
	frames := []string{
		`
					╔═════════════════════════════════════════════════╗
					║                                                 ║
					║                     CAVE                        ║
					║                                                 ║
					╚═════════════════════════════════════════════════╝
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

// Fonction pour être appelée depuis main.go
func StartCaveAnimation() {
	CaveEntranceAnimation()
}
