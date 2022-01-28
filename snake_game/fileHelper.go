package snake_game

import ( 
	"image"
	"os"
	"strconv"
	_ "image/png"
	"encoding/csv"
	"github.com/faiface/pixel"

	"io"
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

func loadGopherEmoji(imagePath string, frameWidth, frameHeight float64) (emojiImg pixel.Picture, frame []pixel.Rect, err error) {
	// total hack, nicely format the error at the end, so I don't have to type it every time
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "error loading moji")
		}
	}()

	// open and load the spritesheet
	sheetFile, err := os.Open(imagePath)
	if err != nil {
		return nil, nil, err
	}
	defer sheetFile.Close()

	sheetImg, _, err := image.Decode(sheetFile)
	if err != nil {
		return nil, nil, err
	}
	emojiImg = pixel.PictureDataFromImage(sheetImg)

	// create a slice of frames inside the spritesheet
	var frames []pixel.Rect
	for x := 0.0; x+frameWidth <= emojiImg.Bounds().Max.X; x += frameWidth {
		for y := 0.0; y+frameHeight <= emojiImg.Bounds().Max.Y; y += frameHeight {
			frames = append(frames, pixel.R(
				x,
				y,
				x+frameWidth,
				y+frameHeight,
			))
		}
	}
	
	/*
	length, height := emojiImg.Bounds().Max.X, emojiImg.Bounds().Max.Y


	fmt.Println("length of image: ", length , " and height of image: ",height)

	for i := range frames {
        f := frames[i]
        fmt.Println("frame: ",f) 
	}
	fmt.Println("num of frames: ",len(frames)) //number of images

	*/
	 
	return emojiImg, frames, nil
}



	func loadPicture(path string) (pixel.Picture, error) {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		img, _, err := image.Decode(file)
		if err != nil {
			return nil, err
		}
		return pixel.PictureDataFromImage(img), nil
	}