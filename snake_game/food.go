package snake_game

import ( 

	"github.com/faiface/pixel"
	//"github.com/faiface/pixel/imdraw"
	//"github.com/faiface/pixel/pixelgl"
	"math/rand"
	)

type food struct {
	sprite *pixel.Sprite
	rect  pixel.Rect
	position coordinates
}

/*
func (f *food) drawFood(canvas pixelgl.canvas, sprite pixel.Sprite) {

	sprite.Draw(canvas, pixel.IM.Scaled(pixel.ZV,0.3).Moved(canvas.Bounds().Center()))

}

*/
// return the coordinates of the food on display
func (f *food) foodPosition() (int, int) {
	return f.position.x, f.position.y
}



func randomPosition() (float64, float64){
	newX := rand.Intn(80 - (-80)) + (-80)
	newY := rand.Intn(80 - (-80)) + (-80)
	return  float64(newX), float64(newY)
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

