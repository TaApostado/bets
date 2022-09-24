package entity

import (
	errs "github.com/TaApostado/bets/domain/bet/errors"
	gambler "github.com/TaApostado/bets/domain/gambler"
)

type Suspended struct {
	bet *Bet
}

func (state Suspended) ChangeName(name string) error {
	state.bet.name = name
	return nil
}

func (state Suspended) ChangeDescription(description string) error {
	state.bet.description = description
	return nil
}

func (state Suspended) ChangeValue(value float32) error {
	state.bet.value = value
	return nil
}

func (state Suspended) Deposit(amount float32) error {
	state.bet.credits += amount
	return nil
}

func (state Suspended) Withdraw(amount float32) error {
	state.bet.credits -= amount
	return nil
}

func (state Suspended) AddGambler(gambler gambler.IGambler) error {
	return &errs.ErrCannotRegisterGamblerBetIsntOpen
}

func (state Suspended) RemoveGambler(gambler gambler.IGambler) error {
	for index, g := range state.bet.Gamblers() {
		if g.Id() == gambler.Id() {
			state.bet.gamblers[index] = state.bet.gamblers[len(state.bet.gamblers)-1]
			state.bet.gamblers = state.bet.gamblers[:len(state.bet.gamblers)-1]
			return nil
		}
	}
	return nil
}
