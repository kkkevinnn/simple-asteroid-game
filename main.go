package main

import (
	"fmt"
	"image"
	"log"
	"slices"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"asteroid/constant"
	"asteroid/sprite"
	"asteroid/utils"
)

type Game struct {
	mouse     sprite.Point
	player    sprite.Player
	asteroids []*sprite.Asteriod
	bullets   []*sprite.Bullet
	keys      []ebiten.Key
}

var game *Game

func NewGame() *Game {
	center := utils.Vector2{
		X: constant.SCREEN_WIDTH / 2,
		Y: constant.SCREEN_HEIGHT / 2,
	}
	bounds := image.Rect(0, 0, constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT)
	gun := sprite.GunConfig{
		Radius:    constant.BULLET_RADIUS,
		Speed:     constant.BULLET_SPEED,
		RateLimit: constant.PLAYER_FIRE_RATE,
	}
	player := sprite.NewPlayer(center, constant.PLAYER_RADUIS, bounds, constant.PLAYER_MOVE_SPEED, constant.PLAYER_ROTATION_SPEED, gun)
	return &Game{
		player: *player,
	}
}

func (g *Game) Update() error {
	g.mouse.X, g.mouse.Y = ebiten.CursorPosition()
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	for _, k := range g.keys {
		switch k {
		case ebiten.KeySpace:
			bullet := g.player.Fire()
			g.bullets = append(g.bullets, bullet)
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		g.player.Update(g.keys)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, a := range g.asteroids {
			a.Update()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// update bullets position
		for _, b := range g.bullets {
			b.Update()
		}
		// remove bullets that are out of bounds
		g.bullets = slices.DeleteFunc(g.bullets, func(el *sprite.Bullet) bool {
			if el.Center.X < 0 || el.Center.X > constant.SCREEN_WIDTH || el.Center.Y < 0 || el.Center.Y > constant.SCREEN_HEIGHT {
				return true
			}
			return false
		})
	}()

	wg.Wait()
	// collision detection

	// spwan asteroids

	// fire bullets

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%d, %d", g.mouse.X, g.mouse.Y))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.2f", ebiten.ActualFPS()), constant.SCREEN_WIDTH-70, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %.2f", ebiten.ActualTPS()), constant.SCREEN_WIDTH-70, 0)
	g.player.Draw(screen)
	for _, a := range g.asteroids {
		a.Draw(screen)
	}
	for _, b := range g.bullets {
		b.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Geometry Matrix")
	game = NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
