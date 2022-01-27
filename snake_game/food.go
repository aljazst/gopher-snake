package snake_game

import ( 

	"github.com/faiface/pixel"
	//"github.com/faiface/pixel/imdraw"
	//"github.com/faiface/pixel/pixelgl"
	"math/rand"
	"fmt"
	)

	
type food struct {
	sprite *pixel.Sprite
	
	emojiImg pixel.Picture
	//position coordinates
	position    pixel.Vec
	frame []pixel.Rect
	pressed bool
	collected bool
}
var newX, newY int = 0,0

/*
func (f *food) drawFood(canvas pixelgl.canvas, sprite pixel.Sprite) {

	sprite.Draw(canvas, pixel.IM.Scaled(pixel.ZV,0.3).Moved(canvas.Bounds().Center()))

}


// return the coordinates of the food on display
func (f *food) foodPosition() (int, int) {
	return f.position.x, f.position.y
}*/



func( f *food) randomPosition() (float64, float64){
	

	if(f.pressed || f.collected){
		//(rand.Intn(max - min) + min)
		newX = rand.Intn(60 - (-60)) + (-60)
		newY = rand.Intn(60 - (-60)) + (-60)
		f.pressed = false
		f.collected = false
	}
	f.position.X = float64(newX) 
	f.position.Y = float64(newY)
	fmt.Println("X:", f.position.X, " Y: ", f.position.Y)
	return  float64(newX), float64(newY)
}

func (fo *food) drawEmoji(t pixel.Target, food *food) {
	if fo.sprite == nil {
		fo.sprite = pixel.NewSprite(nil, pixel.Rect{})
	}
	if(fo.pressed || fo.collected){
	fo.sprite = pixel.NewSprite(food.emojiImg, food.emojiImg.Bounds())
	fo.sprite.Set(food.emojiImg, fo.frame[rand.Intn(len(fo.frame))])
}
	fo.sprite.Draw(t, pixel.IM.Scaled(pixel.ZV,0.15).Moved(pixel.V(food.randomPosition())))

}


/*
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
	
}*/

