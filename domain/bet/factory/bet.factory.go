package factory

import (
	domain "github.com/TaApostado/bets/domain/bet"
	entity "github.com/TaApostado/bets/domain/bet/entity"
	"github.com/TaApostado/bets/domain/challenged"
	"github.com/TaApostado/bets/domain/gambler"
	"github.com/google/uuid"
)

type state string

const (
	opened    state = "opened"
	closed    state = "closed"
	suspended state = "suspended"
)

type BetFactory struct {
	challenged  challenged.IChallenged
	name        string
	description string
	value       float32
	gamblers    []gambler.IGambler
	state       state
}

func NewBetFactory(challenged challenged.IChallenged, name, description string, value float32) *BetFactory {
	return &BetFactory{
		challenged:  challenged,
		name:        name,
		description: description,
		value:       value,
		gamblers:    []gambler.IGambler{},
	}
}

func (factory *BetFactory) AsOpened() *BetFactory {
	factory.state = opened
	return factory
}

func (factory *BetFactory) AsSuspended() *BetFactory {
	factory.state = suspended
	return factory
}

func (factory *BetFactory) AsClosed() *BetFactory {
	factory.state = closed
	return factory
}

func (factory *BetFactory) Build() (domain.IBet, error) {
	bet, err := entity.NewBet(
		uuid.NewString(), factory.name,
		factory.description, factory.challenged,
		factory.value)

	if err != nil {
		return nil, err
	}

	if factory.state == opened {
		err = bet.Open()
	} else if factory.state == suspended {
		err = bet.Suspend()
	} else if factory.state == closed {
		err = bet.Close()
	}

	if err != nil {
		return nil, err
	}
	return bet, nil
}
