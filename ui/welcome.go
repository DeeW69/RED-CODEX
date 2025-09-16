package ui

import (
	"fmt"
	"time"
)

func DisplayWelcomeArt() {
	fmt.Print("\033[2J\033[H") // Clear screen
	fmt.Println()
	fmt.Println("\033[1;36m" + `
    ╔══════════════════════════════╗
    ║                              ║
    ║  ██████╗ ███████╗██████╗     ║
    ║  ██╔══██╗██╔════╝██╔══██╗    ║
    ║  ██████╔╝█████╗  ██║  ██║    ║
    ║  ██╔══██╗██╔══╝  ██║  ██║    ║
    ║  ██║  ██║███████╗██████╔╝    ║
    ║  ╚═╝  ╚═╝╚══════╝╚═════╝     ║
    ║                              ║
    ║ RED PROJECT - Le Jeu de Rôle ║
    ║                              ║
    ╚══════════════════════════════╝
` + "\033[0m")
}

func DisplayMainMenu() {
	fmt.Println()
	fmt.Println("\033[1;33m" + "╔══════════════════════════════════════╗")
	fmt.Println("║             MENU PRINCIPAL           ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║  \033[1;32m1\033[1;33m. Nouvelle Partie                  ║")
	fmt.Println("║  \033[1;32m2\033[1;33m. Charger Partie                   ║")
	fmt.Println("║  \033[1;32m3\033[1;33m. Options                          ║")
	fmt.Println("║  \033[1;32m4\033[1;33m. Quitter                          ║")
	fmt.Println("╚══════════════════════════════════════╝\033[0m")
	fmt.Print("\n\033[1;37mChoisissez une option [1-4]: \033[0m")
}

func AnimatedWelcome() {
	colors := []string{"\033[1;31m", "\033[1;32m", "\033[1;33m", "\033[1;34m", "\033[1;35m", "\033[1;36m"}

	for i := 0; i < 2; i++ {
		for _, color := range colors {
			fmt.Print("\033[2J\033[H")
			fmt.Println()
			fmt.Println(color + `
    ╔══════════════════════════════╗
    ║                              ║
    ║  ██████╗ ███████╗██████╗     ║
    ║  ██╔══██╗██╔════╝██╔══██╗    ║
    ║  ██████╔╝█████╗  ██║  ██║    ║
    ║  ██╔══██╗██╔══╝  ██║  ██║    ║
    ║  ██║  ██║███████╗██████╔╝    ║
    ║  ╚═╝  ╚═╝╚══════╝╚═════╝     ║
    ║                              ║
    ║ RED PROJECT - Le Jeu de Rôle ║
    ║                              ║
    ╚══════════════════════════════╝
` + "\033[0m")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
