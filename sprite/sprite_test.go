package sprite_test

import (
	"asteroid/sprite"
	"asteroid/utils"
)

type mockCollidable struct {
	Center utils.Vector2
	Radius int
}

func (m mockCollidable) GetHitboxCircule() (utils.Vector2, int) {
	return m.Center, m.Radius
}

func (m mockCollidable) IsCollided(h sprite.Collidable) bool {
	return false
}
