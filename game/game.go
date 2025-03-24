package game

import (
	"asteroid/constant"
	"asteroid/sprite"
	"asteroid/utils"
	"fmt"
	"image"
	"log"
	"slices"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	mouse     sprite.Point
	player    sprite.Player
	asteroids []*sprite.Asteriod
	bullets   []*sprite.Bullet
	keys      []ebiten.Key
}

func NewGame() *Game {
	center := utils.Vector2{
		X: constant.SCREEN_WIDTH / 2,
		Y: constant.SCREEN_HEIGHT / 2,
	}
	bounds := image.Rect(0, 0, constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT)
	rate, err := time.ParseDuration(constant.PLAYER_FIRE_RATE)
	if err != nil {
		log.Fatal(err)
	}
	gun := sprite.GunConfig{
		Radius:    constant.BULLET_RADIUS,
		Speed:     constant.BULLET_SPEED,
		RateLimit: rate,
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
			g.Fire()
		}
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go g.updatePlayer(wg)

	wg.Add(1)
	go g.updateAsteroids(wg)

	wg.Add(1)
	go g.updateBullets(wg)

	wg.Wait()
	// collision detection

	// spwan asteroids

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

func (g *Game) updatePlayer(wg *sync.WaitGroup) {
	defer wg.Done()
	g.player.Update(g.keys)
}

func (g *Game) updateAsteroids(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, a := range g.asteroids {
		a.Update()
	}
}

func (g *Game) updateBullets(wg *sync.WaitGroup) {
	defer wg.Done()
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
}

func (g *Game) Fire() {
	bullet, err := g.player.Fire()
	if err != nil {
		if err == sprite.ErrGunNotReady {
			return
		}
		log.Fatal(err)
	}
	g.bullets = append(g.bullets, bullet)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT
}
