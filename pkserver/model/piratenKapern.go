package model

import (
	"bytes"
	"fmt"
	"math/rand"
	_ "reflect"
)

type PiratenKapern struct {
	Dice         [8]Die
	FortuneCards [35]FortuneCard
}

func NewPiratenKapern() *PiratenKapern {
	var dice [8]Die
	for i := 0; i < 8; i++ {
		die := NewDie(i + 1)
		dice[i] = *die
	}

	j := 0
	var fortuneCards [35]FortuneCard

	j = createCard(4, GOLD_COIN, j, &fortuneCards)
	j = createCard(4, DIAMONDS, j, &fortuneCards)
	j = createCard(4, SORCERESS, j, &fortuneCards)
	j = createCard(4, CAPTAIN, j, &fortuneCards)
	j = createCard(4, TREASURE_CHEST, j, &fortuneCards)
	j = createCard(4, MONKEY_BUSINESS, j, &fortuneCards)
	j = createCard(3, ONE_SKULL, j, &fortuneCards)
	j = createCard(2, TWO_SKULL, j, &fortuneCards)
	j = createCard(2, TWO_SABRE, j, &fortuneCards)
	j = createCard(2, THREE_SABRE, j, &fortuneCards)
	j = createCard(2, FOUR_SABRE, j, &fortuneCards)



	return &PiratenKapern{Dice: dice, FortuneCards: fortuneCards}
}

func createCard(count int, fortuneCardType FortuneCardType, j int, fortuneCards *[35]FortuneCard) int {
	for i := 0; i < count; i++ {
		fortuneCard := NewFortuneCard(j+1, fortuneCardType)
		fortuneCards[j] = *fortuneCard
		j++
	}
	return j
}

func (piratenKapern *PiratenKapern) Test() FortuneCard {
	var i = rand.Intn(35)
	fortuneCard := piratenKapern.FortuneCards[i]
	return fortuneCard
}

func (piratenKapern PiratenKapern) DrawFortuneCard(playRecord *PlayRecord) FortuneCard {
	var i = rand.Intn(35)
	fortuneCard := piratenKapern.FortuneCards[i]
	playRecord.FortuneCardType = fortuneCard.FortuneCardType
	return fortuneCard
}

type RollResult [8]DieSide

func (rollResult RollResult) String() string {
	b := new(bytes.Buffer)
	for i, dieSide := range rollResult {
		fmt.Fprintf(b, "%d=\"%s\" ", i+1, dieSide.String())
	}
	return b.String()
}

func (piratenKapern PiratenKapern) Roll(playRecord *PlayRecord) RollResult {
	var rollResult RollResult
	for i, die := range piratenKapern.Dice {
		dieSide := die.Roll()
		rollResult[i] = dieSide
	}
	playRecord.RollResult = rollResult
	return rollResult
}

func (piratenKapern PiratenKapern) ComputeScore(playRecord *PlayRecord) int {
	return 1000
}

func (piratenKapern PiratenKapern) UpdateScore(scoreRecord *ScoreRecord, score int) int {
	scoreRecord.Score = scoreRecord.Score + score
	return scoreRecord.Score
}
