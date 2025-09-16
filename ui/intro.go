package ui

import (
	"fmt"
	"time"
)

func ShowIntro() {
	intro := []string{
		"Réveille-toi... Ou pas. DE toute façon c’est déjà trop tard.",
		"Tu viens d’être ressuscité par erreur. Félicitations, glitch magique.",
		"Tu ne sais pas où tu es, ni pourquoi tu sens le champignon moisi.",
		"Mais laisse-moi t’expliquer ce que tu viens de rater pendant les cent dernières années.",
		"Autrefois, Elyndor était un royaume paisible...Feu, Eau, Terre, Air…",
		"et une cinquième qu’on a oubliée parce qu’elle était trop bizarre.",
		"Mais Zagreus les vola... Ce qui emporta, avec lui, le monde dans le chaos.",
		"Et toi ? Tu n’as pas été choisi. Juste une arme rouillée, une potion douteuse avec toi.",
		"Ta mission ? Explorer des lieux absurdes, recruter Lyra, Brak, Nox et Elwyn. Alors lève-toi, aventurier par accident.",
		"Le monde est foutu, mais il n’a pas encore vu ce que tu vaux.",
	}

	fmt.Println("\n\033[1;35m═══════════════════════════════════════════════\033[0m")
	for _, line := range intro {
		fmt.Println(line)
		time.Sleep(1500 * time.Millisecond)
	}
	fmt.Println("\033[1;35m═══════════════════════════════════════════════\033[0m")
	fmt.Println("\033[1;36mVous venez de revenir d’un endroit que personne ne quitte.\nLe sol tremble sous vos pas incertains.\nUn château se dresse au loin, sombre et silencieux.\nVous ignorez pourquoi, mais chaque fibre de votre corps vous pousse à avancer.\033[0m")
	time.Sleep(2000 * time.Millisecond)

}
