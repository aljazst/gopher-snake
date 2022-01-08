package snake_game

import (
	"fmt"
)

func PrintCoordinatestest() {
   // test := newSnake("olala")

	tmp := new(coordinates)
	tmp.setCoordinates(3,44)

	cordX,cordY := tmp.getCoordinates()

	fmt.Println(cordX,cordY)
}
