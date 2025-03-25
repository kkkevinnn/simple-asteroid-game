package sprite

import (
	"asteroid/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Asteroid struct {
	Circle
}

func NewAsteroid(center utils.Vector2, radius int, speed float64, direction utils.Vector2) *Asteroid {
	return &Asteroid{
		Circle: Circle{
			Center:    center,
			Radius:    radius,
			Speed:     speed,
			Direction: direction,
		},
	}
}

func (a *Asteroid) Update() {
	a.Center.Add(*a.Direction.Clone().Scale(a.Speed * dt))
}

func (a *Asteroid) Draw(screen *ebiten.Image) {
	vector.StrokeCircle(screen, float32(a.Center.X), float32(a.Center.Y), float32(a.Radius), 2.0, color.White, true)
}
