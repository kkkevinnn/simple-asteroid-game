package game

import (
	"asteroid/sprite"
	"asteroid/utils"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

func newTestGame() *Game {
	game := &Game{}
	game.Reset()
	return game
}

func TestGameState_PlayingToGameOverOnCollision(t *testing.T) {
	assert := assert.New(t)
	g := newTestGame()
	assert.Equal(StatePlaying, g.state, "Initial state should be StatePlaying")

	// Place an asteroid directly on the player
	playerPos := g.player.Center
	g.asteroidCtrl.AddAsteroid(
		sprite.NewAsteroid(playerPos, 10, 0, *utils.NewVector2(1, 0)),
	)

	// Update should detect collision and change state
	err := g.Update()
	assert.NoError(err, "Update should not return an error")

	assert.Equal(StateGameOver, g.state, "State should be StateGameOver after collision")
}

func TestGameState_ResetSetsStateToPlaying(t *testing.T) {
	assert := assert.New(t)
	g := newTestGame()
	g.state = StateGameOver
	g.Reset()

	assert.Equal(StatePlaying, g.state, "State should be StatePlaying after Reset")
}

func TestGameState_UpdatesPausedOnGameOver(t *testing.T) {
	assert := assert.New(t)
	g := newTestGame()

	initialPlayerPos := g.player.Center
	g.state = StateGameOver
	g.keys = []ebiten.Key{ebiten.KeyArrowUp}

	for i := 0; i < 10; i++ {
		err := g.Update()
		assert.NoError(err, "Update should not return an error during GameOver state")
	}

	assert.Equal(initialPlayerPos, g.player.Center, "Player position should remain unchanged in StateGameOver")
}
