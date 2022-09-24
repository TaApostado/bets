package entity

import (
	errs "github.com/TaApostado/bets/domain/bet/errors"
	"github.com/TaApostado/bets/domain/gambler"
)

type Closed struct {
	bet *Bet
}

func (state Closed) ChangeName(name string) error {
	return &errs.ErrCannotChangeName
}

func (state Closed) ChangeDescription(description string) error {
	return &errs.ErrCannotChangeDescription
}

func (state Closed) ChangeValue(value float32) error {
	return &errs.ErrCannotChangeValue
}

func (state Closed) Deposit(amount float32) error {
	return &errs.ErrCannotDeposit
}

func (state Closed) Withdraw(amount float32) error {
	return &errs.ErrCannotWithdraw
}

func (state Closed) AddGambler(gambler.IGambler) error {
	return &errs.ErrCannotRegisterGamblerBetIsntOpen
}

func (state Closed) RemoveGambler(gambler.IGambler) error {
	return &errs.ErrCannotUnregisterGamblerBetIsClosed
}
