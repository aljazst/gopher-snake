package snake_game

type coordinates struct {
	x, y int
}


func (c *coordinates) getCoordinates() (int, int) {
	return c.x, c.y
}

func (c *coordinates) setCoordinates(x int, y int) {
	c.x = x
	c.y = y
}