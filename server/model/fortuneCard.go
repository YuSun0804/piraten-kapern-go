package model

type FortuneCardType int32

const (
	GOLD_COIN       FortuneCardType = 1
	DIAMONDS        FortuneCardType = 2
	SORCERESS       FortuneCardType = 3
	CAPTAIN         FortuneCardType = 4
	TREASURE_CHEST  FortuneCardType = 5
	MONKEY_BUSINESS FortuneCardType = 6
	ONE_SKULL       FortuneCardType = 7
	TWO_SKULL       FortuneCardType = 8
	TWO_SABRE       FortuneCardType = 9
	THREE_SABRE     FortuneCardType = 10
	FOUR_SABRE      FortuneCardType = 11
)

func (fortuneCardType FortuneCardType) String() string {
	switch fortuneCardType {
	case GOLD_COIN:
		return "GOLD_COIN"
	case DIAMONDS:
		return "DIAMONDS"
	case SORCERESS:
		return "SORCERESS"
	case CAPTAIN:
		return "CAPTAIN"
	case TREASURE_CHEST:
		return "TREASURE_CHEST"
	case MONKEY_BUSINESS:
		return "MONKEY_BUSINESS"
	case ONE_SKULL:
		return "ONE_SKULL"
	case TWO_SKULL:
		return "TWO_SKULL"
	case TWO_SABRE:
		return "TWO_SABRE"
	case THREE_SABRE:
		return "THREE_SABRE"
	case FOUR_SABRE:
		return "FOUR_SABRE"
	default:
		return "UNKNOWN"
	}
}

type FortuneCard struct {
	Id              int
	FortuneCardType FortuneCardType
}

func NewFortuneCard(id int, fortuneCardType FortuneCardType) *FortuneCard {
	return &FortuneCard{Id: id, FortuneCardType: fortuneCardType}
}
