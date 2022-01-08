package snake_game

import (
	"fmt"
	"errors"
)

type direction1 int // make a new type direction, We could just use int, but naming it direction is more underestandable I guess.

const (
	SNAKE_UP direction1 = iota // The IOTA keyword represent integer constant starting from zero and increments them. We could alsp write: USNAKE_UP = 0, SNAKE_DOWN = 1 ... and so on, but iota does that for us.
	SNAKE_DOWN 		 		// we could also write iota in each line, would be the same output https://golangbyexample.com/iota-in-golang/
	SNAKE_RIGHT 
	SNAKE_LEFT 
)

// a struct in go is like a struct in c. OPP languages have classes, but this does the trick too. https://golangbyexample.com/struct-in-golang-complete-guide/
type snake struct {
	name string
	color string
	bodyLength int
	direction direction1 
	position []coordinates

}


func newSnake(name string) *snake {
	s := snake{name: name}
	s.color = "red"
	s.bodyLength = 3

	return &s
}


func (s *snake) head() *coordinates  {
	return &s.position[len(s.position)-1]
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }
    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}

func Print1() {
    fmt.Println(snake{name:"Gonzales", color: "red", bodyLength: 5})

}