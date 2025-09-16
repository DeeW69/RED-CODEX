package game

type Parchemin struct {
    Nom         string
    Description string
    Recette     map[string]int // Nom du drop → quantité requise
}

var ParcheminFeu = Parchemin{
    Nom:         "Parchemin de Feu",
    Recette: map[string]int{
        "peau":      2,
        "venin":     3,
		"spore":     1,
    },
}
var ParcheminTerre = Parchemin{
	Nom:         "Parchemin de Terre",
	Recette: map[string]int{
		"peau":      2,
		"carapace":  3,
		"peau pâle": 1,
	},
}

var ParcheminEau = Parchemin{
	Nom:         "Parchemin d'Eau",
	Recette: map[string]int{
		"viande":     2,
		"carapace":   4,
		"champignon": 1,
	},
}

var ParcheminAir = Parchemin{
	Nom:         "Parchemin d'Air",
	Recette: map[string]int{
		"crocs":      2,
		"viande":     3,
		"peau pâle":  1,
	},
}

var ParcheminMagic = Parchemin{
	Nom:         "Parchemin de Magie",
	Recette: map[string]int{
		"viande avariée": 3,
		"charme":         4,
		"os":             2,
		"plumes":         1,
	},
}

var ParcheminSolide = Parchemin{
	Nom:         "Parchemin de Solide",
	Recette: map[string]int{
		"carapace":  4,
		"crocs":     3,
		"os": 	     2,
		"serre":     2,
	},
}

var ParcheminMiope = Parchemin{
	Nom:         "Parchemin de Myopie",
	Recette: map[string]int{
		"peau pâle": 3,
		"spore":     2,
		"serres":    2,
		"plumes":    2,
	},
}

var ParcheminViolent = Parchemin{
	Nom:         "Parchemin de Violence",
	Recette: map[string]int{
		"serre":         6,
		"croc":          8,
		"sang de titan": 2,
		"plasma":        1,
	},
}

var TheParchemins = []Parchemin{
    {
        Nom: "The Parchemins",
        Recette: map[string]int{
            "peau":           1,
            "venin":          1,
            "spore":          1,
            "carapace":       1,
            "peau pâle":      1,
            "viande":         1,
            "champignon":     1,
            "crocs":          1,
            "viande avariée": 1,
            "os":             1,
            "plumes":         1,
            "serres":         1,
            "charme":         1,
            "sang de titan":  1,
            "plasma":         1,
            "laine":          1,
            "chaire":         1,
            "cendres dorées": 1,
            "sang de tanos":  1,
        },
    },
}
