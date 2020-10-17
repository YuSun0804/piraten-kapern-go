package model

import "math/rand"

type DieSide int32

const (
	SKULL   DieSide = 1
	MONKEY  DieSide = 2
	PARROT  DieSide = 3
	SWORD   DieSide = 4
	COIN    DieSide = 5
	DIAMOND DieSide = 6
)

func (dieSide DieSide) String() string {
	switch dieSide {
	case SKULL:
		return "SKULL"
	case MONKEY:
		return "MONKEY"
	case PARROT:
		return "PARROT"
	case SWORD:
		return "SWORD"
	case COIN:
		return "COIN"
	case DIAMOND:
		return "DIAMOND"
	default:
		return "UNKNOWN"
	}
}

type Die struct {
	Id       int
	DieSides []DieSide
}

func NewDie(id int) *Die {
	var dieSides []DieSide
	dieSides = append(dieSides, SKULL)
	dieSides = append(dieSides, MONKEY)
	dieSides = append(dieSides, PARROT)
	dieSides = append(dieSides, SWORD)
	dieSides = append(dieSides, COIN)
	dieSides = append(dieSides, DIAMOND)
	return &Die{Id: id, DieSides: dieSides}
}

func (die Die) Roll() DieSide {
	var i = rand.Intn(6)
	return die.DieSides[i]
}
