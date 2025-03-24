package sprite

import (
	"asteroid/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Bullet struct {
	Center    utils.Vector2
	Radius    int
	Speed     float64
	Direction utils.Vector2
}

func NewBullet(center utils.Vector2, radius int, speed float64, direction utils.Vector2) *Bullet {
	return &Bullet{
		Center:    center,
		Radius:    radius,
		Speed:     speed,
		Direction: direction,
	}
}

func (b *Bullet) Update() {
	b.Center.Add(*b.Direction.Copy().Scale(b.Speed * dt))
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(b.Center.X), float32(b.Center.Y), float32(b.Radius), color.White, true)
}

func (b *Bullet) GetHitboxCircule() (p utils.Vector2, r int) {
	return b.Center, b.Radius
}

func (b *Bullet) HitboxCollision(h Collidable) bool {
	bPos, bRad := b.GetHitboxCircule()
	hPos, hRad := h.GetHitboxCircule()

	dist := utils.Distance(bPos.X, bPos.Y, hPos.X, hPos.Y)
	return dist < float64(bRad+hRad)
}
