package gameobjects

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

/*Menu represents all properties of menu*/
type Menu struct {
	Title   string
	Options []string
	Atlas   *text.Atlas
	Text    *text.Text
}

/*NewMenu returns a instance of a menu*/
func NewMenu() *Menu {
	title := "Invaders Words"
	options := []string{"Start", "Ranking", "Credits", "Quit"}

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(100, 500), atlas)

	fmt.Fprintln(txt, title)

	for i := 0; i < len(options); i++ {
		fmt.Fprintln(txt, options[i])
	}

	return &Menu{
		Title:   title,
		Options: options,
		Atlas:   atlas,
		Text:    txt,
	}
}

// Draw text on window
func (m *Menu) Draw(win *pixelgl.Window) {
	m.Text.Draw(win, pixel.IM)
}
