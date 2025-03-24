package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"asteroid/constant"
	"asteroid/game"
)

var g *game.Game

func main() {
	ebiten.SetWindowSize(constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Geometry Matrix")
	g = game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
