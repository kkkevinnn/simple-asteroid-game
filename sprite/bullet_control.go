package sprite

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type BulletControl struct {
	Bounds  image.Rectangle
	Bullets []*Bullet
}

func NewBulletControl(bounds image.Rectangle) *BulletControl {
	return &BulletControl{
		Bounds:  bounds,
		Bullets: make([]*Bullet, 0),
	}
}

func (bc *BulletControl) AddBullet(bullet *Bullet) {
	bc.Bullets = append(bc.Bullets, bullet)
}

func (bc *BulletControl) HitBullet(i int) {
	if i >= len(bc.Bullets) || bc.Bullets[i].IsDestoryed() {
		return
	}
	bc.Bullets[i].Destory()
}

func (bc *BulletControl) Clean() {
	mark := len(bc.Bullets)
	for i := len(bc.Bullets) - 1; i >= 0; i-- {
		if bc.Bullets[i].IsDestoryed() {
			mark--
			tmpBullet := bc.Bullets[i]
			bc.Bullets[i] = bc.Bullets[mark]
			bc.Bullets[mark] = tmpBullet
		}
	}

	bc.Bullets = bc.Bullets[:mark]
}

func (bc *BulletControl) Draw(screen *ebiten.Image) {
	for _, b := range bc.Bullets {
		b.Draw(screen)
	}
}

func (bc *BulletControl) Update() {
	for _, b := range bc.Bullets {
		b.Update()
	}
	// remove bullets that are out of bounds
	for _, b := range bc.Bullets {
		if b.Center.X < 0 || b.Center.X > float64(bc.Bounds.Max.X) ||
			b.Center.Y < 0 || b.Center.Y > float64(bc.Bounds.Max.Y) {
			b.Destory()
		}
	}
}
