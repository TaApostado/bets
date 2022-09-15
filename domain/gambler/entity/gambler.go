package entity

import (
	"errors"

	v "github.com/TaApostado/bets/domain/gambler/validator"
	"github.com/TaApostado/common/notification"
)

type Gambler struct {
	id          string
	credits     float32
	notificator notification.INotificator
}

func (gambler *Gambler) validate() error {
	_ = v.Validator().Execute(gambler)
	if gambler.notificator.HasNotifications() {
		return errors.New(gambler.notificator.Notifications())
	}
	return nil
}

func NewGambler(id string, credits float32) (*Gambler, error) {
	gambler := &Gambler{
		id:          id,
		credits:     credits,
		notificator: notification.NewNotificator(),
	}

	err := gambler.validate()
	if err != nil {
		return nil, err
	}
	return gambler, nil
}

func (gambler Gambler) Id() string {
	return gambler.id
}

func (gambler Gambler) Credits() float32 {
	return gambler.credits
}

func (gambler *Gambler) Deposit(amount float32) error {
	if amount < 0 {
		return errors.New("")
	}
	gambler.credits += amount
	return nil
}

func (gambler *Gambler) Withdraw(amount float32) error {
	if amount < 0 {
		return errors.New("")
	} else if amount > gambler.credits {
		return errors.New("")
	}
	gambler.credits -= amount
	return nil
}

// Notificator returns the validation error container (notification.INotificator)
func (gambler Gambler) Notificator() notification.INotificator {
	return gambler.notificator
}
