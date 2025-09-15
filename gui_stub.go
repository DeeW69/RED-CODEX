//go:build !gui
// +build !gui

package main

import "fmt"

// LaunchGUI is a stub when the GUI build tag is not set.
func LaunchGUI() {
    fmt.Println("Interface graphique indisponible. Recompilez avec le tag '-tags gui' et assurez-vous qu'un environnement graphique est présent.")
}

