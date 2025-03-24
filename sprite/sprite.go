package sprite

import (
	"asteroid/utils"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	whiteImage = ebiten.NewImage(3, 3)

	// whiteSubImage is an internal sub image of whiteImage.
	// Use whiteSubImage at DrawTriangles instead of whiteImage in order to avoid bleeding edges.
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

const dt float64 = float64(1) / 60

func init() {
	whiteImage.Fill(color.White)
}

type Point struct {
	X, Y int
}

func (p *Point) Clamp(bounds image.Rectangle) {
	p.X = utils.Clamp(p.X, bounds.Min.X, bounds.Max.X)
	p.Y = utils.Clamp(p.Y, bounds.Min.Y, bounds.Max.Y)
}

type Collidable interface {
	GetHitboxCircule() (utils.Vector2, int)
	IsCollided(h Collidable) bool
}
