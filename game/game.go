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
	player       sprite.Player
	asteroidCtrl sprite.AsteroidControl
	bulletCtrl   sprite.BulletControl
	keys         []ebiten.Key
}

func NewGame() *Game {
	game := &Game{}
	game.Reset()

	return game
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go g.updatePlayer(wg)
	wg.Add(1)
	go g.updateAsteroids(wg)
	wg.Add(1)
	go g.updateBullets(wg)
	wg.Wait()

	// collision detection
	if g.IsPlayerCollidedWithAsteroid() {
		g.Reset()
		return nil
	}
	g.CheckBulletCollidedWithAsteroid()

	g.bulletCtrl.Clean()
	g.asteroidCtrl.Clean()

	if slices.Contains(g.keys, ebiten.KeySpace) {
		g.Fire()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%.2f, %.2f", g.player.Center.X, g.player.Center.Y))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.2f", ebiten.ActualFPS()), constant.SCREEN_WIDTH-70, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %.2f", ebiten.ActualTPS()), constant.SCREEN_WIDTH-70, 0)
	g.player.Draw(screen)
	g.asteroidCtrl.Draw(screen)
	g.bulletCtrl.Draw(screen)
}

func (g *Game) updatePlayer(wg *sync.WaitGroup) {
	defer wg.Done()
	g.player.Update(g.keys)
}

func (g *Game) updateAsteroids(wg *sync.WaitGroup) {
	defer wg.Done()
	g.asteroidCtrl.Update()
}

func (g *Game) updateBullets(wg *sync.WaitGroup) {
	defer wg.Done()
	g.bulletCtrl.Update()
}

func (g *Game) Fire() {
	bullet, err := g.player.Fire()
	if err != nil {
		if err == sprite.ErrGunNotReady {
			return
		}
		log.Fatal(err)
	}
	g.bulletCtrl.AddBullet(bullet)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT
}

func (g *Game) Reset() {
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
	asteroidCtrl := sprite.NewAsteroidControl(
		constant.ASTEROID_MIN_RADIUS,
		constant.ASTEROID_KINDS,
		bounds,
		constant.ASTEROID_SPAWN_RATE,
	)

	g.player = *player
	g.asteroidCtrl = *asteroidCtrl
	g.bulletCtrl = sprite.BulletControl{Bounds: bounds}
}

func (g *Game) IsPlayerCollidedWithAsteroid() bool {
	for _, a := range g.asteroidCtrl.Asteroids {
		if g.player.IsCollided(a) {
			log.Printf("Player(%.2f, %.2f) collided with Asteroid(%.2f, %.2f)", g.player.Center.X, g.player.Center.Y, a.Center.X, a.Center.Y)
			return true
		}
	}
	return false
}

func (g *Game) CheckBulletCollidedWithAsteroid() {
	for i, b := range g.bulletCtrl.Bullets {
		for j, a := range g.asteroidCtrl.Asteroids {
			if b.IsDestoryed() || a.IsDestoryed() {
				continue
			}

			if b.IsCollided(a) {
				log.Printf("Bullet(%.2f, %.2f) collided with Asteroid(%.2f, %.2f)", b.Center.X, b.Center.Y, a.Center.X, a.Center.Y)
				g.bulletCtrl.HitBullet(i)
				g.asteroidCtrl.HitAsteroid(j)
			}
		}
	}
}
