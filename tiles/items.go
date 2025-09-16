// File: tiles/items.go
package tiles

type WeaponData struct {
    Name   string
    Damage int
}

type ArmorData struct {
    Name    string
    Defense int
}

type PotionData struct {
    Name     string
    Effect   string
    Duration int
}

var (
    WeaponList = []*WeaponData{
        {Name: "Épée de feu", Damage: 25},
        {Name: "Hache de glace", Damage: 30},
        {Name: "Dague empoisonnée", Damage: 18},
    }

    ArmorList = []*ArmorData{
        {Name: "Armure de fer", Defense: 20},
        {Name: "Cape magique", Defense: 15},
        {Name: "Bouclier de titan", Defense: 35},
    }

    PotionList = []*PotionData{
        {Name: "Potion de soin", Effect: "Restaure la vie", Duration: 5},
        {Name: "Potion de vitesse", Effect: "Augmente la vitesse", Duration: 10},
        {Name: "Potion d'invisibilité", Effect: "Devient invisible", Duration: 8},
    }
)
