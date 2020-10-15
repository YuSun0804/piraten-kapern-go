package model

type DieSide int32

const (
	SKULL   DieSide = 1
	MONKEY  DieSide = 2
	PARROT  DieSide = 3
	SWORD   DieSide = 4
	COIN    DieSide = 5
	DIAMOND DieSide = 6
)

type Die struct {
	Id        int
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
