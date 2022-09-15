package entity

import (
	"errors"

	v "github.com/TaApostado/bets/domain/challenged/validator"
	"github.com/TaApostado/common/notification"
)

type Challenged struct {
	id          string
	credits     float32
	notificator notification.INotificator
}

func (challenged *Challenged) validate() error {
	_ = v.Validator().Execute(challenged)
	if challenged.notificator.HasNotifications() {
		return errors.New(challenged.notificator.Notifications())
	}
	return nil
}

func NewChallenged(id string, credits float32) (*Challenged, error) {
	challenged := &Challenged{
		id:          id,
		credits:     credits,
		notificator: notification.NewNotificator(),
	}

	err := challenged.validate()
	if err != nil {
		return nil, err
	}
	return challenged, nil
}

func (challenged Challenged) Id() string {
	return challenged.id
}

func (challenged Challenged) Credits() float32 {
	return challenged.credits
}

func (challenged *Challenged) Deposit(amount float32) error {
	if amount < 0 {
		return errors.New("")
	}
	challenged.credits += amount
	return nil
}

func (challenged *Challenged) Withdraw(amount float32) error {
	if amount < 0 {
		return errors.New("")
	} else if amount > challenged.credits {
		return errors.New("")
	}
	challenged.credits -= amount
	return nil
}

// Notificator returns the validation error container (notification.INotificator)
func (challenged Challenged) Notificator() notification.INotificator {
	return challenged.notificator
}
