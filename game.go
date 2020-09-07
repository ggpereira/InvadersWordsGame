package main

import (
	gameobjects "InvadersWords/gameObjects"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Entire game inside this
func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Invaders Words",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	// create screen
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// player, err := GameObjects.NewPlayer("teste")
	// if err != nil {
	// 	panic(err)
	// }

	menu := gameobjects.NewMenu()

	for !win.Closed() {
		win.Clear(colornames.Black)

		// player.Sprite.Draw(win, pixel.IM.Moved(pixel.V(player.PosX, player.PosY)))
		menu.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
