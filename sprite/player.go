package sprite

import (
	"asteroid/utils"
	"errors"
	"image"
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

type MoveDirection int

const (
	MoveForward MoveDirection = iota
	MoveBackward
)

type RotateDirection int

const (
	RotateAntiClockwise RotateDirection = iota
	RotateClockwise
)

type Player struct {
	Circle

	Bounds        image.Rectangle
	RotationSpeed float64

	Gun GunConfig

	lastFired time.Time
}

func NewPlayer(center utils.Vector2, radius int, bounds image.Rectangle, speed float64, rotationSpeed float64, gun GunConfig) *Player {
	newBounds := image.Rectangle{
		Min: image.Point{X: bounds.Min.X + radius, Y: bounds.Min.Y + radius},
		Max: image.Point{X: bounds.Max.X - radius, Y: bounds.Max.Y - radius},
	}
	p := Player{
		Circle: Circle{
			Center:    center,
			Radius:    radius,
			Speed:     speed,
			Direction: utils.Vector2{X: 0, Y: -1},
		},
		Bounds:        newBounds,
		RotationSpeed: rotationSpeed,
		Gun:           gun,
	}

	return &p
}

func (p *Player) Update(keys []ebiten.Key) {
	for _, k := range keys {
		switch k {
		case ebiten.KeyW:
			p.Move(MoveForward, p.Speed*dt)
		case ebiten.KeyS:
			p.Move(MoveBackward, p.Speed*dt)
		case ebiten.KeyA:
			p.Rotate(RotateAntiClockwise, p.RotationSpeed*dt)
		case ebiten.KeyD:
			p.Rotate(RotateClockwise, p.RotationSpeed*dt)
		}
	}
	p.Center.Clamp(p.Bounds)
}

func (p *Player) Triangle() [3]*utils.Vector2 {
	forward := p.Direction.Clone().Scale(float64(p.Radius))

	right := p.Direction.Clone().Reverse().Rotate(90).Scale(float64(p.Radius) / 1.5)

	a := p.Center.Clone().Add(*forward)
	b := p.Center.Clone().Sub(*forward).Sub(*right)
	c := p.Center.Clone().Sub(*forward).Add(*right)

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

func (p *Player) Fire() (*Bullet, error) {
	if time.Since(p.lastFired) < p.Gun.RateLimit {
		return nil, ErrGunNotReady
	}
	p.lastFired = time.Now()
	gunPos := p.Triangle()[0]
	dir := p.Direction.Clone()
	bullet := NewBullet(*gunPos, p.Gun.Radius, p.Gun.Speed, *dir)
	return bullet, nil
}

func (p *Player) Move(direction MoveDirection, distance float64) {
	var v *utils.Vector2
	switch direction {
	case MoveForward:
		v = p.Direction.Clone().Scale(distance)
	case MoveBackward:
		v = p.Direction.Clone().Reverse().Scale(distance)
	default:
		return
	}

	p.Center.Add(*v)
}

func (p *Player) Rotate(direction RotateDirection, deg float64) {
	// for drawing, y-axis is inverted. Thus the expected values are inverted
	switch direction {
	case RotateAntiClockwise:
		p.Direction.Rotate(-deg)
	case RotateClockwise:
		p.Direction.Rotate(deg)
	default:
		return
	}
}
