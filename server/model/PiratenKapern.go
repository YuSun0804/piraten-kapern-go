package model

import (
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

func (piratenKapern PiratenKapern) DrawFortuneCard() FortuneCard {
	var i = rand.Intn(35) + 1
	fortuneCard := piratenKapern.FortuneCards[i]
	return fortuneCard
}
