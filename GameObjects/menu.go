package gameobjects

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

/*Menu represents all properties of menu*/
type Menu struct {
	Title   string
	Options []string
	Atlas   *text.Atlas
	Text    []*text.Text
}

/*NewMenu returns a instance of a menu*/
func NewMenu() *Menu {
	title := "Invaders Words"
	options := []string{"Start", "Ranking", "Credits", "Quit"}

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := make([]*text.Text, len(options)+1)

	// set title text options
	txt[0] = text.New(pixel.V(318, 600), atlas)
	txt[0].Color = colornames.Cornflowerblue
	fmt.Fprintln(txt[0], title)

	fmt.Println(txt[0].BoundsOf(title))

	// set rest of menu options
	for i := 1; i < len(options)+1; i++ {
		height := float64(590 - (i * 60))
		// centering text on screen [center_pixel - (text_width * 2)]
		posx := 512 - (txt[0].BoundsOf(options[i-1]).Max.X-txt[0].BoundsOf(options[i-1]).Min.X)*2

		txt[i] = text.New(pixel.V(posx, height), atlas)
		txt[i].Color = colornames.White
		fmt.Fprintln(txt[i], options[i-1])
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
	// draw title
	m.Text[0].Draw(win, pixel.IM.Scaled(m.Text[0].Orig, 4))
	// draw menu options
	for i := 1; i < len(m.Options)+1; i++ {
		m.Text[i].Draw(win, pixel.IM.Scaled(m.Text[i].Orig, 3))
	}
}
