package game

import (
	"math/rand"
	"time"
)

// Session centralise la gestion du joueur et du fichier de sauvegarde.
type Session struct {
	Player   *PlayerData
	savePath string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewSession charge la sauvegarde existante et prépare une session de jeu.
func NewSession(savePath string) (*Session, error) {
	player, err := LoadPlayer(savePath)
	if err != nil {
		return nil, err
	}
	session := &Session{
		Player:   player,
		savePath: savePath,
	}
	session.ensureCollections()
	return session, nil
}

// Reload recharge les données depuis le fichier de sauvegarde.
func (s *Session) Reload() error {
	player, err := LoadPlayer(s.savePath)
	if err != nil {
		return err
	}
	s.Player = player
	s.ensureCollections()
	return nil
}

// Save persiste l'état actuel du joueur sur disque.
func (s *Session) Save() error {
	if s == nil || s.Player == nil {
		return nil
	}
	return SavePlayer(s.savePath, s.Player)
}

// ResetForNewGame réinitialise les jauges principales pour démarrer une partie propre.
func (s *Session) ResetForNewGame() {
	if s == nil || s.Player == nil {
		return
	}
	stats := &s.Player.Player.Stats
	stats.Health = stats.MaxHealth
	stats.Mana = stats.MaxMana
	s.ensureCollections()
}

// AddDrops ajoute les drops remportés à l'inventaire.
func (s *Session) AddDrops(drops []Drop) {
	if s == nil || s.Player == nil || len(drops) == 0 {
		return
	}
	inv := &s.Player.Player.Inventory
	if inv.Drops == nil {
		inv.Drops = make(map[string]int)
	}
	for _, drop := range drops {
		inv.Drops[drop.Item] += drop.Quantity
	}
}

// AddGold crédite l'or gagné.
func (s *Session) AddGold(amount int) {
	if s == nil || s.Player == nil || amount == 0 {
		return
	}
	s.Player.Player.Inventory.Gold += amount
}

// UnlockCompanion débloque un compagnon et renvoie true si c'est la première fois.
func (s *Session) UnlockCompanion(key string, companion Companion) bool {
	if s == nil || s.Player == nil {
		return false
	}
	companions := s.Player.Player.Companions
	if companions == nil {
		companions = make(map[string]Companion)
		s.Player.Player.Companions = companions
	}
	existing, ok := companions[key]
	if ok && existing.Unlocked {
		return false
	}
	companion.Unlocked = true
	companions[key] = companion
	return true
}

// CurrentGold retourne la quantité totale d'or affichable.
func (s *Session) CurrentGold() int {
	if s == nil || s.Player == nil {
		return 0
	}
	return s.Player.Player.Stats.Gold + s.Player.Player.Inventory.Gold
}

func (s *Session) ensureCollections() {
	if s == nil || s.Player == nil {
		return
	}
	if s.Player.Player.Inventory.Drops == nil {
		s.Player.Player.Inventory.Drops = make(map[string]int)
	}
	if s.Player.Player.Companions == nil {
		s.Player.Player.Companions = make(map[string]Companion)
	}
}
