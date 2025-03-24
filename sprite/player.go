package sprite

import (
	"asteroid/utils"
	"errors"
	"image"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var ErrGunNotReady = errors.New("gun is not ready yet")

type GunConfig struct {
	Radius    int
	Speed     float64
	RateLimit time.Duration
}

type Player struct {
	Center   utils.Vector2
	Radius   int
	Bounds   image.Rectangle
	Rotation float64

	Speed         float64
	RotationSpeed float64

	gun GunConfig

	lastFired time.Time
}

func NewPlayer(center utils.Vector2, radius int, bounds image.Rectangle, speed float64, rotationSpeed float64, gun GunConfig) *Player {
	newBounds := image.Rectangle{
		Min: image.Point{X: bounds.Min.X + radius, Y: bounds.Min.Y + radius},
		Max: image.Point{X: bounds.Max.X - radius, Y: bounds.Max.Y - radius},
	}
	return &Player{
		Center:        center,
		Radius:        radius,
		Bounds:        newBounds,
		Speed:         speed,
		RotationSpeed: rotationSpeed,
		gun:           gun,
	}
}

func (p *Player) Update(keys []ebiten.Key) {
	for _, k := range keys {
		switch k {
		case ebiten.KeyW:
			p.Center.Y -= p.Speed
		case ebiten.KeyS:
			p.Center.Y += p.Speed
		case ebiten.KeyA:
			p.Rotation = math.Mod(p.Rotation+p.RotationSpeed*dt, 360)
		case ebiten.KeyD:
			p.Rotation = math.Mod(p.Rotation-p.RotationSpeed*dt, 360)
		}
	}
	newPos := &Point{X: int(math.Round(p.Center.X)), Y: int(math.Round(p.Center.Y))}
	newPos.Clamp(p.Bounds)
	p.Center.X, p.Center.Y = float64(newPos.X), float64(newPos.Y)
}

func (p *Player) Triangle() [3]*utils.Vector2 {
	forward := &utils.Vector2{X: 0, Y: -1}
	forward.Rotate(float64(p.Rotation)).Scale(float64(p.Radius))

	right := &utils.Vector2{X: 0, Y: 1}
	right.Rotate(float64(p.Rotation + 90)).Scale(float64(p.Radius) / 1.5)

	a := p.Center.Copy().Add(*forward)
	b := p.Center.Copy().Sub(*forward).Sub(*right)
	c := p.Center.Copy().Sub(*forward).Add(*right)

	return [3]*utils.Vector2{a, b, c}
}

func (p *Player) Draw(screen *ebiten.Image) {
	var vertices []ebiten.Vertex
	var indices []uint16
	var path vector.Path

	corners := p.Triangle()

	path.MoveTo(float32(corners[0].X), float32(corners[0].Y)) // Top vertex
	path.LineTo(float32(corners[1].X), float32(corners[1].Y)) // Bottom-left vertex
	path.LineTo(float32(corners[2].X), float32(corners[2].Y)) // Bottom-right vertex
	path.Close()

	vertices, indices = path.AppendVerticesAndIndicesForFilling(vertices, indices)

	op := &ebiten.DrawTrianglesOptions{}
	op.AntiAlias = true
	op.FillRule = ebiten.FillRuleNonZero
	screen.DrawTriangles(vertices, indices, whiteSubImage, op)
}

func (p *Player) GetHitboxCircule() (utils.Vector2, int) {
	return p.Center, p.Radius
}

func (p *Player) HitboxCollision(h Hitbox) bool {
	pPos, pRad := p.GetHitboxCircule()
	hPos, hRad := h.GetHitboxCircule()

	dist := utils.Distance(pPos.X, pPos.Y, hPos.X, hPos.Y)
	return dist < float64(pRad+hRad)
}

func (p *Player) Fire() (*Bullet, error) {
	if time.Since(p.lastFired) < p.gun.RateLimit {
		return nil, ErrGunNotReady
	}
	p.lastFired = time.Now()
	gunPos := p.Triangle()[0]
	dir := (&utils.Vector2{X: 0, Y: -1}).Rotate(p.Rotation)
	bullet := NewBullet(*gunPos, p.gun.Radius, p.gun.Speed, *dir)
	return bullet, nil
}
