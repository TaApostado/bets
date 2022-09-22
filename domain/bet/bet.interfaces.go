package bet

import (
	"github.com/TaApostado/bets/domain/challenged"
	"github.com/TaApostado/bets/domain/gambler"
	"github.com/TaApostado/common/notification"
)

type IState interface {
	ChangeName(name string) error
	ChangeDescription(description string) error
	ChangeValue(value float32) error

	Deposit(float32) error
	Withdraw(float32) error
}

type IBet interface {
	Id() string
	Name() string
	Description() string
	Challenged() challenged.IChallenged
	Value() float32
	Credits() float32
	IsOpen() bool
	IsSuspended() bool
	IsClosed() bool
	Gamblers() []gambler.IGambler

	Open() error
	Close() error
	Suspend() error

	ChangeName(name string) error
	ChangeDescription(description string) error
	ChangeValue(value float32) error

	Deposit(float32) error
	Withdraw(float32) error

	Notificator() notification.INotificator
}
