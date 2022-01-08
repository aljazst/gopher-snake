package snake_game

import ( 

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"math/rand"


	"image/color"
	)

type food struct {
	food rune
	rect  pixel.Rect
	color color.Color
	position coordinates
}

func (f *food) draw(imd *imdraw.IMDraw) {
	imd.Color = f.color
	imd.Push(f.rect.Min, f.rect.Max)
	imd.Rectangle(2)
}

// return the coordinates of the food on display
func (f *food) foodPosition() (int, int) {
	return f.position.x, f.position.y
}

func generateFood(coord coordinates) *food {
	return &food{
		food: randomFood(),
		position: coord,
	}

}

func (f *food) randomPosition() {
	newX := rand.Intn(100-1) +1
	newY := rand.Intn(100-1) +1
	f.position.x, f.position.y = newX, newY
}

// A rune is an alias to the int32 data type. It represents a Unicode code point.
func randomFood() rune {

	//food that gophers eat - gophers only feed on plants and are strict herbivores
	gopherFood := []rune{
		'🌱', //seeds
		'🥔', //potatoes
		'🍠', //sweet potato
		'🌿', //grass
		'🥦',
		'🌻', //flowers
		'💐',
	}

	return gopherFood[rand.Intn(len(gopherFood))]
	
}

