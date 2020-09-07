package gameobjects

import (
	"InvadersWords/utils"
	"path/filepath"

	"github.com/faiface/pixel"
)

// Player represents a object contains state of player
type Player struct {
	Name   string
	Sprite *pixel.Sprite
	PosX   float64
	PosY   float64
	Score  int
}

// NewPlayer return a initialized player
func NewPlayer(nickname string) (*Player, error) {
	abs, err := filepath.Abs("./source/sprites/spaceship.png")
	if err != nil {
		return nil, err
	}

	pic, err := utils.LoadPicture(abs)
	if err != nil {
		return nil, err
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	return &Player{
		Name:   nickname,
		Sprite: sprite,
		PosX:   500,
		PosY:   50,
		Score:  0,
	}, nil
}
