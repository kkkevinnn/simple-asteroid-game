package sprite

import "asteroid/utils"

type Circle struct {
	Center    utils.Vector2
	Radius    int
	Speed     float64
	Direction utils.Vector2
}

func (c *Circle) GetHitboxCircule() (p utils.Vector2, r int) {
	return c.Center, c.Radius
}

func (c *Circle) IsCollided(h Collidable) bool {
	bPos, bRad := c.GetHitboxCircule()
	hPos, hRad := h.GetHitboxCircule()

	dist := utils.Distance(bPos.X, bPos.Y, hPos.X, hPos.Y)
	return dist <= float64(bRad+hRad)
}
