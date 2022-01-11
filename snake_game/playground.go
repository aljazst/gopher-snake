package snake_game

import ( 

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"math/rand"


	"image/color"
	)

type border struct {
	rect  pixel.Rect
	color color.Color
	position coordinates
}

func (f *border) draw(imd *imdraw.IMDraw) {
	imd.Color = f.color
	imd.Push(f.rect.Min, f.rect.Max)
	imd.Rectangle(2)
}

// return the coordinates of the food on display
func (f *border) foodPosition() (int, int) {
	return f.position.x, f.position.y
}



