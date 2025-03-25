package sprite

import (
	"asteroid/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Bullet struct {
	Circle
}

func NewBullet(center utils.Vector2, radius int, speed float64, direction utils.Vector2) *Bullet {
	return &Bullet{
		Circle: Circle{
			Center:    center,
			Radius:    radius,
			Speed:     speed,
			Direction: direction,
		},
	}
}

func (b *Bullet) Update() {
	b.Center.Add(*b.Direction.Clone().Scale(b.Speed * dt))
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(b.Center.X), float32(b.Center.Y), float32(b.Radius), color.White, true)
}
