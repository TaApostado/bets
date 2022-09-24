package entity

import (
	"errors"

	domain "github.com/TaApostado/bets/domain/bet"
	v "github.com/TaApostado/bets/domain/bet/validator"
	challenged "github.com/TaApostado/bets/domain/challenged"
	gambler "github.com/TaApostado/bets/domain/gambler"
	"github.com/TaApostado/common/notification"
)

type Bet struct {
	id, name, description string
	challenged            challenged.IChallenged
	value, credits        float32
	state                 domain.IState
	gamblers              []gambler.IGambler

	opened, closed, suspended domain.IState
	notificator               notification.INotificator
}

func (b *Bet) validate() error {
	v.Validator().Execute(b)
	if b.notificator.HasNotifications() {
		return errors.New(b.notificator.Notifications())
	}
	return nil
}

func NewBet(
	id, name, description string, challenged challenged.IChallenged,
	value float32) (*Bet, error) {
	bet := Bet{
		id:          id,
		name:        name,
		description: description,
		challenged:  challenged,
		value:       value,
		credits:     float32(0.0),
		gamblers:    []gambler.IGambler{},
		notificator: notification.NewNotificator(),
	}

	err := bet.validate()
	if err != nil {
		return nil, err
	}

	opened := &Opened{bet: &bet}
	closed := &Closed{bet: &bet}
	suspended := &Suspended{bet: &bet}

	bet.opened = opened
	bet.state = bet.opened
	bet.closed = closed
	bet.suspended = suspended
	return &bet, nil
}

// Id returns bet id
func (bet *Bet) Id() string {
	return bet.id
}

// Name returns bet name
func (bet *Bet) Name() string {
	return bet.name
}

// Description returns bet description
func (bet *Bet) Description() string {
	return bet.description
}

// Value returns bet value
func (bet *Bet) Value() float32 {
	return bet.value
}

// Value returns bet value
func (bet *Bet) Credits() float32 {
	return bet.credits
}

// Challenged returns the creator of bet
func (bet *Bet) Challenged() challenged.IChallenged {
	return bet.challenged
}

// Gamblers returns a list of gamblers to bet
func (bet *Bet) Gamblers() []gambler.IGambler {
	return bet.gamblers
}

// Notificator returns the validation error container (notification.INotificator)
func (bet *Bet) Notificator() notification.INotificator {
	return bet.notificator
}

func (bet *Bet) Open() error {
	if bet.IsClosed() {
		return errors.New("cannot open the bet because it has already been closed")
	}
	bet.state = bet.opened
	return nil
}

func (bet *Bet) IsOpen() bool {
	return bet.state == bet.opened
}

func (bet *Bet) Close() error {
	if bet.IsOpen() || bet.IsSuspended() {
		bet.state = bet.closed
		return nil
	}
	return errors.New("cannot close the bet because it has already been closed")
}

func (bet *Bet) IsClosed() bool {
	return bet.state == bet.closed
}

func (bet *Bet) Suspend() error {
	if bet.IsClosed() {
		return errors.New("cannot suspend the bet because it has already been closed")
	}
	bet.state = bet.suspended
	return nil
}

func (bet *Bet) IsSuspended() bool {
	return bet.state == bet.suspended
}

func (bet *Bet) ChangeName(name string) error {
	oldName := bet.name
	err := bet.state.ChangeName(name)
	if err != nil {
		bet.name = oldName
		return err
	}

	if err = bet.validate(); err != nil {
		bet.name = oldName
		return err
	}
	return nil
}

func (bet *Bet) ChangeDescription(description string) error {
	oldDescription := bet.description
	err := bet.state.ChangeDescription(description)
	if err != nil {
		bet.description = oldDescription
		return err
	}

	if err = bet.validate(); err != nil {
		bet.description = oldDescription
		return err
	}
	return nil
}

func (bet *Bet) ChangeValue(value float32) error {
	oldValue := bet.value
	err := bet.state.ChangeValue(value)
	if err != nil {
		bet.value = oldValue
		return err
	}

	if err = bet.validate(); err != nil {
		bet.value = oldValue
		return err
	}
	return nil
}

func (bet *Bet) AddGambler(gambler gambler.IGambler) error {
	return bet.state.AddGambler(gambler)
}

func (bet *Bet) RemoveGambler(gambler gambler.IGambler) error {
	return bet.state.RemoveGambler(gambler)
}

func (bet *Bet) Deposit(amount float32) error {
	err := bet.state.Deposit(amount)
	if err != nil {
		return err
	}

	if err = bet.validate(); err != nil {
		return err
	}
	return nil
}

func (bet *Bet) Withdraw(amount float32) error {
	err := bet.state.Withdraw(amount)
	if err != nil {
		return err
	}

	if err = bet.validate(); err != nil {
		return err
	}
	return nil
}
