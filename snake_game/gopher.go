package snake_game

import ("fmt"

"math"

_ "image/png"

"github.com/faiface/pixel"
)

var pathHolderX bool = false
var pathHolderY bool = false

type direction int // make a new type direction, We could just use int, but naming it direction is more underestandable I guess.

const (
	GOPHER_UP direction = iota // The IOTA keyword represent integer constant starting from zero and increments them. We could alsp write: USNAKE_UP = 0, SNAKE_DOWN = 1 ... and so on, but iota does that for us.
	GOPHER_DOWN 		 		// we could also write iota in each line, would be the same output https://golangbyexample.com/iota-in-golang/
	GOPHER_RIGHT 
	GOPHER_LEFT 
)

type gopherAnim struct {
	sheet pixel.Picture
	anims map[string][]pixel.Rect
	rate  float64

	direction   direction
	counter float64
	dir     float64

	frame pixel.Rect

	sprite *pixel.Sprite

	rect   pixel.Rect

	vel    pixel.Vec

	border bool
}
func (ga *gopherAnim) resetVel()  {
	ga.vel.X = 0
	ga.vel.Y = 0
}


func (ga *gopherAnim) update( dt float64,phys *gopherAnim,ctrl pixel.Vec, bor *border, food *food) {
	
	ga.counter += dt

	switch {
	case ctrl.X < 0:
		ga.direction = GOPHER_LEFT
		ga.resetVel()
		ga.vel.X = -10
	case ctrl.X > 0:
		ga.direction = GOPHER_RIGHT
		ga.resetVel()
		ga.vel.X = +10
	case ctrl.Y < 0:
		ga.direction = GOPHER_DOWN
		ga.resetVel()
		ga.vel.Y = -10
	case ctrl.Y > 0:
		ga.direction = GOPHER_UP
		ga.resetVel()
		ga.vel.Y = +10

	}

	ga.rect = ga.rect.Moved(ga.vel.Scaled(dt))

	//fmt.Println("ga.rect.Max.X: ",ga.rect.Max.X, "bor.rect.Max.X: ", bor.rect.Max.X)

	//fmt.Println("ga.rect.Min.X: ",ga.rect.Min.X, "bor.rect.Min.X: ", bor.rect.Min.X)
	ga.border = false
	if ga.rect.Max.X >= bor.rect.Max.X || ga.rect.Min.X <= bor.rect.Min.X {
		ga.border = true
	}
	if ga.rect.Min.Y <= bor.rect.Min.Y || ga.rect.Max.Y >= bor.rect.Max.Y {
		ga.border = true
	}
	
	// adding +5 or -15 to the offset. +5 if comming at the emoji from the left and -15 if comming at it from the right. It is how it is.

	if math.Round(ga.rect.Max.X)+5 == food.position.X || math.Round(ga.rect.Max.X)-15 == food.position.X || pathHolderY {
		pathHolderY = true
		if math.Round(ga.rect.Max.Y)+5 == food.position.Y || math.Round(ga.rect.Max.Y)-15 == food.position.Y {
			food.collected = true
			pathHolderX = false
			pathHolderY = false
			fmt.Println("WINWINWINWINWINWINWINWINWINWINWINWINWINWINWIN")
		}
	}
	if math.Round(ga.rect.Max.Y)+5 == food.position.Y || math.Round(ga.rect.Max.Y)-15 == food.position.Y || pathHolderX {
		pathHolderX = true
		if math.Round(ga.rect.Max.X)+5 == food.position.X || math.Round(ga.rect.Max.X)-15 == food.position.X {
			food.collected = true
			pathHolderY = false
			pathHolderX = false
			fmt.Println("WINWINWINWINWINWINWINWINWINWINWINWINWINWINWIN")
		}
	}

	fmt.Println("Gopher: Min X: ", math.Round(ga.rect.Min.X) , "Min Y: ",math.Round(ga.rect.Min.Y),   " Max X:", math.Round(ga.rect.Max.X), " Max Y: ", math.Round(ga.rect.Max.Y))

	// determine the correct animation frame
	switch ga.direction {

	case GOPHER_RIGHT:
		i := int(math.Floor(ga.counter / ga.rate))
		ga.frame = ga.anims["Run"][i%len(ga.anims["Run"])]
	case GOPHER_LEFT:
		i := int(math.Floor(ga.counter / ga.rate))
		ga.frame = ga.anims["Run"][i%len(ga.anims["Run"])]
	case GOPHER_UP:
		i := int(math.Floor(ga.counter / ga.rate))
		ga.frame = ga.anims["Run"][i%len(ga.anims["Run"])]
	case GOPHER_DOWN:
		i := len(ga.anims["Jump"]) - 1
		ga.frame = ga.anims["Jump"][i]
	
	default: 
		ga.frame = ga.anims["Front"][0]
	}


	// set the facing direction of the gopher
	if phys.vel.X != 0 {
		if phys.vel.X > 0 {
			ga.dir = +1
		} else {
			ga.dir = -1
		}
	} 
}
func (ga *gopherAnim) drawGopher(t pixel.Target, phys *gopherAnim) {
	if ga.sprite == nil {
		ga.sprite = pixel.NewSprite(nil, pixel.Rect{})
	}
	//fmt.Println("frame; ",ga.frame )
	// draw the correct frame with the correct position and direction
	ga.sprite.Set(ga.sheet, ga.frame)
	ga.sprite.Draw(t, pixel.IM.
		ScaledXY(pixel.ZV, pixel.V(
			phys.rect.W()/ga.sprite.Frame().W(),
			phys.rect.H()/ga.sprite.Frame().H(),
		)).
		ScaledXY(pixel.ZV,pixel.V(-ga.dir, 1.0)).
		Moved(phys.rect.Center()),
	)
}

/*
//main gopher is the "snake" head.
func (gopher *gopher) mainGopher() *coordinates {
	return &gopher.position[len(gopher.position)-1]
}
*/

func foodFound(tmp coordinates, coord coordinates) bool {
	return tmp.x == coord.x && tmp.y == coord.y
}

func test()  {
	fmt.Println("hel")
}