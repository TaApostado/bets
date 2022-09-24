package entity

import (
	errs "github.com/TaApostado/bets/domain/bet/errors"
	gambler "github.com/TaApostado/bets/domain/gambler"
)

type Opened struct {
	bet *Bet
}

func (state Opened) ChangeName(name string) error {
	return &errs.ErrCannotChangeName
}

func (state Opened) ChangeDescription(description string) error {
	return &errs.ErrCannotChangeDescription
}

func (state Opened) ChangeValue(value float32) error {
	return &errs.ErrCannotChangeValue
}

func (state Opened) Deposit(amount float32) error {
	return &errs.ErrCannotDeposit
}

func (state Opened) Withdraw(amount float32) error {
	return &errs.ErrCannotWithdraw
}

func (state Opened) AddGambler(gambler gambler.IGambler) error {
	for _, g := range state.bet.Gamblers() {
		if g.Id() == gambler.Id() {
			return nil
		}
	}
	state.bet.gamblers = append(state.bet.Gamblers(), gambler)
	return nil
}

func (state Opened) RemoveGambler(gambler gambler.IGambler) error {
	for index, g := range state.bet.Gamblers() {
		if g.Id() == gambler.Id() {
			state.bet.gamblers[index] = state.bet.gamblers[len(state.bet.gamblers)-1]
			state.bet.gamblers = state.bet.gamblers[:len(state.bet.gamblers)-1]
			return nil
		}
	}
	return nil
}
