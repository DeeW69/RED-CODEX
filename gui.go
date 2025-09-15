package main

import (
    "image/color"
    "log"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
    tileSize = 32
    mapWidth = 10
    mapHeight = 10
)

type Game struct{}

func (g *Game) Update() error { return nil }

func (g *Game) Draw(screen *ebiten.Image) {
    for y := 0; y < mapHeight; y++ {
        for x := 0; x < mapWidth; x++ {
            ebitenutil.DrawRect(screen, float64(x*tileSize), float64(y*tileSize), tileSize, tileSize, color.RGBA{0x80, 0xa0, 0x20, 0xff})
        }
    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return mapWidth * tileSize, mapHeight * tileSize
}

func LaunchGUI() {
    g := &Game{}
    w, h := g.Layout(0, 0)
    ebiten.SetWindowSize(w, h)
    ebiten.SetWindowTitle("RED-CODEX GUI")
    if err := ebiten.RunGame(g); err != nil {
        log.Fatal(err)
    }
}
