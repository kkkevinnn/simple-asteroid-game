package sprite

import "github.com/hajimehoshi/ebiten/v2"

type Asteriod struct {
	Pos    Point
	Radius float64
}

func (a *Asteriod) Update() {
}

func (a *Asteriod) Draw(screen *ebiten.Image) {
}
