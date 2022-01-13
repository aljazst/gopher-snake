package snake_game


import (

	"math"
	"math/rand"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"golang.org/x/image/colornames"
)

func run() {
	rand.Seed(time.Now().UnixNano())

	sheet, anims, err := loadAnimationSheet("sheet.png", "sheet.csv", 12)

	if err != nil {

		panic(err)
	}

	cfg := pixelgl.WindowConfig{
		Title:  "Gopher snake!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync: true, // update the window only as often as the monitor refreshes https://github.com/faiface/pixel/wiki/Creating-a-Window
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	anim := &gopherAnim{
		sheet: sheet,
		anims: anims,
		rate:  1.0 / 10,
		dir:   +1,
		rect:      pixel.R(-6, -7, 6, 7),
	}
	
	
	// hardcoded level  https://pkg.go.dev/github.com/faiface/pixel#R
	border := &border{
		rect: pixel.R(-160/2, -120/2, 160/2, 120/2),
		//{rect: pixel.R(1, 0, 70, 2)},
		//{rect: pixel.R(1, 10, -50, 12)},
		
	}

	food := food{

	}
	
	border.color = randomNiceColor()
	

	canvas := pixelgl.NewCanvas(pixel.R(-160/2, -120/2, 160/2, 120/2))
	imd := imdraw.New(sheet)
	imd.Precision = 32


	last := time.Now()



	pic, err := loadPicture("cake.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())
	

	for !win.Closed() {
		rand.Seed(time.Now().UnixNano())
		dt := time.Since(last).Seconds()
		last = time.Now()
		



		if win.JustPressed(pixelgl.KeyEnter) {
			anim.rect = anim.rect.Moved(anim.rect.Center().Scaled(-1))
			anim.vel = pixel.ZV
		}

		if win.JustPressed(pixelgl.KeyJ) {
			food.pressed = true
		}

		ctrl := pixel.ZV

		if win.Pressed(pixelgl.KeyLeft) {
			ctrl.X--
		}
		if win.Pressed(pixelgl.KeyRight) {
			ctrl.X++
		}
		if win.JustPressed(pixelgl.KeyUp) {
			ctrl.Y++
		}
		if win.JustPressed(pixelgl.KeyDown) {
			ctrl.Y--
		}

		anim.update(dt,anim,ctrl, border)

		canvas.Clear(colornames.Black)
		imd.Clear()

		if anim.border {
			anim.rect = anim.rect.Moved(anim.rect.Center().Scaled(-1))
			anim.vel = pixel.ZV
		}

		//draw the border
		border.draw(imd)
		
		anim.draw(imd,anim)
		imd.Draw(canvas)
		sprite.Draw(canvas, pixel.IM.Scaled(pixel.ZV,0.3).Moved(pixel.V(food.randomPosition())).Rotated(pixel.ZV, math.Pi/2))
		

		win.Clear(colornames.White)
		// stretch the canvas to the window
		win.SetMatrix(pixel.IM.Scaled(pixel.ZV,
			math.Min(
				win.Bounds().W()/canvas.Bounds().W(),
				win.Bounds().H()/canvas.Bounds().H(),
			),
		).Moved(win.Bounds().Center()))
		canvas.Draw(win, pixel.IM.Moved(canvas.Bounds().Center()))

		win.Update()
	}

	
}

func randomNiceColor() pixel.RGBA {
	again:
		r := rand.Float64()
		g := rand.Float64()
		b := rand.Float64()
		len := math.Sqrt(r*r + g*g + b*b)
		if len == 0 {
			goto again
		}
		return pixel.RGB(r/len, g/len, b/len)
	}
func Engine() {
	pixelgl.Run(run)
}