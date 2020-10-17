package model

type Player struct {
	PlayerIndex int
	PlayerName  string
}

type PlayRecord struct {
	Player          Player
	TurnNum         int
	FortuneCardType FortuneCardType
	RollResult      RollResult
}

type ScoreRecord struct {
	Player Player
	Score  int
}
