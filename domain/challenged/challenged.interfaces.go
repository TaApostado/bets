package challenged

import "github.com/TaApostado/common/notification"

type IChallenged interface {
	Id() string
	Credits() float32
	Deposit(float32) error
	Withdraw(float32) error
	Notificator() notification.INotificator
}
