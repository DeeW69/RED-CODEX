package game

import (
	"RED_Project/tiles"
)

var ForgeShop *tiles.ForgeShopData
var Player *tiles.PlayerData

func InitGameData() error {
	var err error
	ForgeShop, err = tiles.LoadForgeShopData()
	if err != nil {
		return err
	}

	Player, err = tiles.LoadPlayerData()
	if err != nil {
		return err
	}
	return nil
}
