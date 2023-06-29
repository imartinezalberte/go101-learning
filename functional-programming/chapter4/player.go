package chapter4

import (
	"errors"
	"math/rand"
)

type (
	Player string

	RandomizeFn func() int
)

const (
	FirstPlayer  Player = "first_player"
	SecondPlayer Player = "second_player"
)

var PlayerNotFoundErr = errors.New("player not found")

func SelectPlayer() (Player, error) {
	which := rand.Intn(2)
	switch which {
	case 0:
		return FirstPlayer, nil
	case 1:
		return SecondPlayer, nil
	}
	return "", PlayerNotFoundErr
}

// more testeable code
func WhichPlayer(randomize RandomizeFn) (Player, error) {
	switch randomize() {
	case 0:
		return FirstPlayer, nil
	case 1:
		return SecondPlayer, nil
	}
	return "", PlayerNotFoundErr
}
