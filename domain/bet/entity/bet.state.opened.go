package entity

import errs "github.com/TaApostado/bets/domain/bet/errors"

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
