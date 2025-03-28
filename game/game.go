package game

import (
	"asteroid/assets/fonts"
	"asteroid/constant"
	"asteroid/sprite"
	"asteroid/utils"

	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"slices"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Game States
type gameState int

const (
	StatePlaying gameState = iota
	StateGameOver
)

const (
	gameOverFontSizeLarge = 48
	gameOverFontSizeSmall = 24
)

var pressStart2pFont *text.GoTextFaceSource

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2pRegular_ttf))
	if err != nil {
		log.Fatalf("load font error: %v", err)
	}
	pressStart2pFont = s
}

type Game struct {
	player            sprite.Player
	asteroidCtrl      sprite.AsteroidControl
	bulletCtrl        sprite.BulletControl
	keys              []ebiten.Key
	state             gameState
	gameOverFontLarge *text.GoTextFace
	gameOverFontSmall *text.GoTextFace
}

func NewGame() *Game {
	game := &Game{}
	game.Reset()

	return game
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	switch g.state {
	case StatePlaying:
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
			g.state = StateGameOver
			return nil
		}
		g.CheckBulletCollidedWithAsteroid()

		g.bulletCtrl.Clean()
		g.asteroidCtrl.Clean()

		if slices.Contains(g.keys, ebiten.KeySpace) {
			g.Fire()
		}
	case StateGameOver:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.Reset()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%.2f, %.2f", g.player.Center.X, g.player.Center.Y))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.2f", ebiten.ActualFPS()), constant.SCREEN_WIDTH-70, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %.2f", ebiten.ActualTPS()), constant.SCREEN_WIDTH-70, 0)

	// Always draw game elements
	g.player.Draw(screen)
	g.asteroidCtrl.Draw(screen)
	g.bulletCtrl.Draw(screen)

	// Draw Game Over overlay if needed
	if g.state == StateGameOver {
		g.drawGameOverOverlay(screen)
	}
}

func (g *Game) drawGameOverOverlay(screen *ebiten.Image) {
	overlayWidth := float64(constant.SCREEN_WIDTH) * 0.6
	overlayHeight := float64(constant.SCREEN_HEIGHT) * 0.6
	x := (float64(constant.SCREEN_WIDTH) - overlayWidth) / 2
	y := (float64(constant.SCREEN_HEIGHT) - overlayHeight) / 2

	bgColor := color.RGBA{R: 20, G: 20, B: 20, A: 200}
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(overlayWidth), float32(overlayHeight), bgColor, true)

	gameOverText := "GAME OVER"
	restartText := "Press Enter to Restart"
	textColor := color.White
	restartTextColor := color.Gray{Y: 180}

	wl, hl := text.Measure(gameOverText, g.gameOverFontLarge, g.gameOverFontLarge.Metrics().CapHeight)
	ws, _ := text.Measure(restartText, g.gameOverFontSmall, g.gameOverFontSmall.Metrics().CapHeight)

	gameOverX := x + (overlayWidth-wl)/2
	gameOverY := y + overlayHeight*0.4

	restartX := x + (overlayWidth-ws)/2
	restartY := gameOverY + hl + overlayHeight*0.1

	// Draw text
	gameOverTextOp := &text.DrawOptions{}
	gameOverTextOp.GeoM.Translate(gameOverX, gameOverY)
	gameOverTextOp.ColorScale.ScaleWithColor(textColor)
	text.Draw(screen, gameOverText, g.gameOverFontLarge, gameOverTextOp)

	restartTextOp := &text.DrawOptions{}
	restartTextOp.GeoM.Translate(restartX, restartY)
	restartTextOp.ColorScale.ScaleWithColor(restartTextColor)
	text.Draw(screen, restartText, g.gameOverFontSmall, restartTextOp)
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
	g.state = StatePlaying

	g.gameOverFontLarge = &text.GoTextFace{
		Source: pressStart2pFont,
		Size:   gameOverFontSizeLarge,
	}

	g.gameOverFontSmall = &text.GoTextFace{
		Source: pressStart2pFont,
		Size:   gameOverFontSizeSmall,
	}
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
