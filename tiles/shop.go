package tiles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ShowShop() {
	shopData, err := LoadForgeShopData()
	if err != nil {
		fmt.Printf("\033[1;31mErreur shop: %v\033[0m\n", err)
		return
	}

	playerData, err := LoadPlayerData()
	if err != nil {
		fmt.Printf("\033[1;31mErreur joueur: %v\033[0m\n", err)
		return
	}

	for {
		fmt.Print("\033[2J\033[H")
		fmt.Println("\033[1;32m" + `
        ╔══════════════════════════════╗
        ║         🛒 SHOP 🛒           ║
        ╚══════════════════════════════╝
        ` + "\033[0m")

		fmt.Printf("\033[1;33m💰 Or: %d pièces\033[0m\n\n", playerData.Gold)
		fmt.Println("\033[1;36m══════════════════════════════════════\033[0m")

		// Afficher les items du shop
		i := 1
		items := make([]string, 0)
		for id, item := range shopData.ShopItems {
			fmt.Printf("%d. %s - %d pièces\n", i, item.Name, item.Price)
			fmt.Printf("   %s\n", item.Description)
			fmt.Printf("   Dans l'inventaire: %d\n\n", playerData.Inventory[id])
			items = append(items, id)
			i++
		}

		fmt.Println("\033[1;36m══════════════════════════════════════\033[0m")
		fmt.Println("0. Retour au vieil homme")
		fmt.Print("\n\033[1;37mVotre choix: \033[0m")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil || choice < 0 || choice > len(items) {
			if choice == 0 {
				return
			}
			fmt.Println("\033[1;31mChoix invalide!\033[0m")
			time.Sleep(1 * time.Second)
			continue
		}

		if choice > 0 {
			itemID := items[choice-1]
			item := shopData.ShopItems[itemID]

			if BuyItem(playerData, itemID, &item, "shop") {
				SavePlayerData(playerData)
			}

			fmt.Print("\n\033[1;37mAppuyez sur Entrée pour continuer...\033[0m")
			reader.ReadString('\n')
		}
	}
}
