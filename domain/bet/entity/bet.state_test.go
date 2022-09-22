package entity

import (
	"testing"

	"github.com/TaApostado/bets/domain/bet"
	errs "github.com/TaApostado/bets/domain/bet/errors"
	"github.com/stretchr/testify/assert"
)

var closed bet.IState
var opened bet.IState
var suspended bet.IState
var entity *Bet

const name string = "Some Name..."
const description string = "Some Description..."
const value float32 = 100
const credits float32 = 0

func setup(t *testing.T) {
	entity = &Bet{
		name:        name,
		description: description,
		value:       value,
		credits:     credits,
	}
	closed = Closed{bet: entity}
	opened = Opened{bet: entity}
	suspended = Suspended{bet: entity}
}

func TestChangeNameWhenBetIsClose(t *testing.T) {
	setup(t)
	err := closed.ChangeName("some...")
	assert.NotNil(t, err)
	assert.Equal(t, errs.ErrCannotChangeName.Error(), err.Error())
}

func TestChangeDescriptionWhenBetIsClose(t *testing.T) {
	setup(t)
	err := closed.ChangeDescription("some...")
	assert.NotNil(t, err)
	assert.Equal(t, errs.ErrCannotChangeDescription.Error(), err.Error())
}

func TestChangeValueWhenBetIsClose(t *testing.T) {
	setup(t)
	err := closed.ChangeValue(18491)
	assert.NotNil(t, err)
	assert.Equal(t, errs.ErrCannotChangeValue.Error(), err.Error())
}

func TestDepositWhenBetIsClose(t *testing.T) {
	setup(t)
	err := closed.Deposit(1894)
	assert.NotNil(t, err)
	assert.Equal(t, errs.ErrCannotDeposit.Error(), err.Error())
}

func TestWithdrawWhenBetIsClose(t *testing.T) {
	setup(t)
	err := closed.Withdraw(48544)
	assert.NotNil(t, err)
	assert.Equal(t, errs.ErrCannotWithdraw.Error(), err.Error())
}

func TestChangeNameWhenBetIsOpen(t *testing.T) {
	setup(t)
	err := opened.ChangeName("some...")
	assert.NotNil(t, err)
	assert.Equal(t, errs.ErrCannotChangeName.Error(), err.Error())
}

func TestChangeDescriptionWhenBetIsOpen(t *testing.T) {
	setup(t)
	err := opened.ChangeDescription("some...")
	assert.NotNil(t, err)
	assert.Equal(t, errs.ErrCannotChangeDescription.Error(), err.Error())
}

func TestChangeValueWhenBetIsOpen(t *testing.T) {
	setup(t)
	err := opened.ChangeValue(18491)
	assert.NotNil(t, err)
	assert.Equal(t, errs.ErrCannotChangeValue.Error(), err.Error())
}

func TestDepositWhenBetIsOpen(t *testing.T) {
	setup(t)
	err := opened.Deposit(1894)
	assert.NotNil(t, err)
	assert.Equal(t, errs.ErrCannotDeposit.Error(), err.Error())
}

func TestWithdrawWhenBetIsOpen(t *testing.T) {
	setup(t)
	err := opened.Withdraw(1894)
	assert.NotNil(t, err)
	assert.Equal(t, errs.ErrCannotWithdraw.Error(), err.Error())
}

func TestChangeNameWhenBetIsSuspended(t *testing.T) {
	setup(t)
	err := suspended.ChangeName(name)
	assert.Equal(t, entity.name, name)
	assert.Nil(t, err)

	invalidName := ""
	err = suspended.ChangeName(invalidName)
	assert.Equal(t, entity.name, invalidName)
	assert.Nil(t, err)
}

func TestChangeDescriptionWhenBetIsSuspended(t *testing.T) {
	setup(t)
	err := suspended.ChangeDescription(description)
	assert.Equal(t, entity.description, description)
	assert.Nil(t, err)

	invalidDescription := ""
	err = suspended.ChangeDescription(invalidDescription)
	assert.Equal(t, entity.description, invalidDescription)
	assert.Nil(t, err)
}

func TestChangeValueWhenBetIsSuspended(t *testing.T) {
	setup(t)
	err := suspended.ChangeValue(1000)
	assert.Nil(t, err)
	assert.Equal(t, entity.value, float32(1000))

	invalidValue := float32(-1000)
	err = suspended.ChangeValue(invalidValue)
	assert.Equal(t, entity.value, invalidValue)
	assert.Nil(t, err)
}

func TestDepositWhenBetIsSuspend(t *testing.T) {
	setup(t)
	actualCredits := entity.credits + 9098
	err := suspended.Deposit(9098)
	assert.Equal(t, actualCredits, entity.credits)
	assert.Nil(t, err)

	actualCredits -= 1000
	err = suspended.Deposit(-1000)
	assert.Equal(t, actualCredits, entity.credits)
	assert.Nil(t, err)
}

func TestWithdrawWhenBetIsSuspend(t *testing.T) {
	setup(t)
	actualCredits := entity.credits - 151561
	err := suspended.Withdraw(151561)
	assert.Nil(t, err)
	assert.Equal(t, actualCredits, entity.credits)

	actualCredits += 1000
	err = suspended.Withdraw(-1000)
	assert.Equal(t, actualCredits, entity.credits)
	assert.Nil(t, err)
}
