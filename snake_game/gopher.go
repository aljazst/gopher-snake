package snake_game

import ("fmt"
"encoding/csv"
"image"

"io"
"math"

"os"
"strconv"


_ "image/png"

"github.com/faiface/pixel"


"github.com/pkg/errors"


)

// function shamelessly stolen from https://github.com/faiface/pixel-examples/tree/master/platformer
func loadAnimationSheet(sheetPath, descPath string, frameWidth float64) (sheet pixel.Picture, anims map[string][]pixel.Rect, err error) {
	// total hack, nicely format the error at the end, so I don't have to type it every time
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "error loading animation sheet")
		}
	}()

	// open and load the spritesheet
	sheetFile, err := os.Open(sheetPath)
	if err != nil {
		return nil, nil, err
	}
	defer sheetFile.Close()
	sheetImg, _, err := image.Decode(sheetFile)
	if err != nil {
		return nil, nil, err
	}
	sheet = pixel.PictureDataFromImage(sheetImg)

	// create a slice of frames inside the spritesheet
	var frames []pixel.Rect
	for x := 0.0; x+frameWidth <= sheet.Bounds().Max.X; x += frameWidth {
		frames = append(frames, pixel.R(
			x,
			0,
			x+frameWidth,
			sheet.Bounds().H(),
		))
	}
	descFile, err := os.Open(descPath)
	if err != nil {
		return nil, nil, err
	}
	defer descFile.Close()

	anims = make(map[string][]pixel.Rect)

	// load the animation information, name and interval inside the spritesheet
	desc := csv.NewReader(descFile)
	for {
		anim, err := desc.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}

		name := anim[0]
		start, _ := strconv.Atoi(anim[1])
		end, _ := strconv.Atoi(anim[2])

		anims[name] = frames[start : end+1]
	}

	return sheet, anims, nil
}

type gopher struct {
	name string
	bodyLength int
	direction direction 
	position []coordinates
}

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
}
func (ga *gopherAnim) resetVel()  {
	ga.vel.X = 0
	ga.vel.Y = 0
}


func (ga *gopherAnim) update( dt float64,phys *gopherAnim,ctrl pixel.Vec) {
	
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

	// determine the correct animation frame
	switch ga.direction {
	//case GOPHER_LEFT:
	//	ga.frame = ga.anims["Front"][0]
	case GOPHER_RIGHT:
		i := int(math.Floor(ga.counter / ga.rate))
		ga.frame = ga.anims["Run"][i%len(ga.anims["Run"])]
		fmt.Println("1",ga.frame)
	case GOPHER_LEFT:
		i := int(math.Floor(ga.counter / ga.rate))
		ga.frame = ga.anims["Run"][i%len(ga.anims["Run"])]
		fmt.Println("2",ga.frame)
	case GOPHER_UP:
		//speed := phys.vel.Y
		i := len(ga.anims["Jump"]) - 1
		ga.frame = ga.anims["Jump"][i]
		fmt.Println("3",ga.frame)
	case GOPHER_DOWN:
		i := len(ga.anims["Jump"]) - 1
		
		ga.frame = ga.anims["Jump"][i]
		fmt.Println("4",ga.frame)
		
	default: 
		ga.frame = ga.anims["Front"][0]
		fmt.Println("5",ga.frame)
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
func (ga *gopherAnim) draw(t pixel.Target, phys *gopherAnim) {
	if ga.sprite == nil {
		ga.sprite = pixel.NewSprite(nil, pixel.Rect{})
	}
	// draw the correct frame with the correct position and direction
	ga.sprite.Set(ga.sheet, ga.frame)
	ga.sprite.Draw(t, pixel.IM.
		ScaledXY(pixel.ZV, pixel.V(
			phys.rect.W()/ga.sprite.Frame().W(),
			phys.rect.H()/ga.sprite.Frame().H(),
		)).
		ScaledXY(pixel.ZV,pixel.V(-ga.dir, 1)).
		Moved(phys.rect.Center()),
	)
}


//main gopher is the "snake" head.
func (gopher *gopher) mainGopher() *coordinates {
	return &gopher.position[len(gopher.position)-1]
}

func (gop *gopher) update(tmp float64, ctrl pixel.Vec, food []food) {

	

}

func foodFound(tmp coordinates, coord coordinates) bool {
	return tmp.x == coord.x && tmp.y == coord.y
}

func test()  {
	fmt.Println("hel")
}