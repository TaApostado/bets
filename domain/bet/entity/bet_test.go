package entity

import (
	"errors"
	"fmt"
	"testing"

	errs "github.com/TaApostado/bets/domain/bet/errors"
	bet "github.com/TaApostado/bets/domain/bet/mock"
	challenged "github.com/TaApostado/bets/domain/challenged/mock"
	gambler_interface "github.com/TaApostado/bets/domain/gambler"
	gambler "github.com/TaApostado/bets/domain/gambler/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var invalidIds = [...]string{"", "1234", "abs"}
var invalidNames = [...]string{"Beber", "5", "565", "B", ""}
var invalidDescriptions = [...]string{"Chute ao gol", "Teste"}
var invalidValues = [...]float32{-1, 0}

var validId = uuid.New().String()

const validName = "Beber 2L de Água"
const validValue = float32(100)
const validDescription = "Beber 2L de água em 5 minutos"

func TestNewBet(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)

	expectedErrorInvalidId := fmt.Sprintf("bet: %s", errs.ErrInvalidBetId.Error())
	for _, invalidId := range invalidIds {
		_, err := NewBet(
			invalidId, validName,
			validDescription, challenged,
			validValue,
		)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErrorInvalidId, err.Error())
	}

	expectedErrorInvalidName := fmt.Sprintf("bet: %s", errs.ErrInvalidBetName.Error())
	for _, invalidName := range invalidNames {
		_, err := NewBet(
			validId, invalidName,
			validDescription, challenged,
			validValue,
		)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErrorInvalidName, err.Error())
	}

	expectedErrorInvalidDescription := fmt.Sprintf("bet: %s", errs.ErrInvalidBetDescription.Error())
	for _, invalidDescription := range invalidDescriptions {
		_, err := NewBet(
			validId, validName,
			invalidDescription, challenged,
			validValue,
		)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErrorInvalidDescription, err.Error())
	}

	expectedErrorInvalidValue := fmt.Sprintf("bet: %s", errs.ErrInvalidBetValue.Error())
	for _, invalidValue := range invalidValues {
		_, err := NewBet(
			validId, validName,
			validDescription, challenged,
			invalidValue,
		)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErrorInvalidValue, err.Error())
	}

	_, err := NewBet(
		validId, invalidNames[0],
		validDescription, challenged,
		invalidValues[0],
	)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), expectedErrorInvalidValue)
	assert.Contains(t, err.Error(), expectedErrorInvalidName)
}

func TestId(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	bet, _ := NewBet(
		validId, validName,
		validDescription, challenged,
		validValue,
	)

	assert.Equal(t, validId, bet.Id())
}

func TestName(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	bet, _ := NewBet(
		validId, validName,
		validDescription, challenged,
		validValue,
	)

	assert.Equal(t, validName, bet.Name())
}

func TestDescription(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	bet, _ := NewBet(
		validId, validName,
		validDescription, challenged,
		validValue,
	)

	assert.Equal(t, validDescription, bet.Description())
}

func TestValue(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	bet, _ := NewBet(
		validId, validName,
		validDescription, challenged,
		validValue,
	)

	assert.Equal(t, validValue, bet.Value())
}

func TestCredits(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	bet, err := NewBet(
		validId, validName,
		validDescription, challenged,
		validValue,
	)
	fmt.Println(err)
	assert.Equal(t, float32(0), bet.Credits())
}

func TestChallenged(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	bet, _ := NewBet(
		validId, validName,
		validDescription, challenged,
		validValue,
	)

	assert.Equal(t, challenged, bet.Challenged())
}

func TestGamblers(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	bet, _ := NewBet(
		validId, validName,
		validDescription, challenged,
		validValue,
	)

	assert.Equal(t, []gambler_interface.IGambler{}, bet.Gamblers())

	gamblerMock := gambler.NewMockIGambler(controller)
	gamblers := []gambler_interface.IGambler{gamblerMock}
	bet.gamblers = gamblers

	assert.Equal(t, gamblers, bet.Gamblers())
}

func TestOpen(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	bet, _ := NewBet(
		validId, validName,
		validDescription, challenged,
		validValue,
	)

	err := bet.Open()
	assert.Nil(t, err)

	bet.state = bet.suspended
	err = bet.Open()
	assert.Nil(t, err)

	bet.state = bet.closed
	err = bet.Open()
	assert.NotNil(t, err)
	assert.Equal(t, "cannot open the bet because it has already been closed", err.Error())
}

func TestClose(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	bet, _ := NewBet(
		validId, validName,
		validDescription, challenged,
		validValue,
	)

	err := bet.Close()
	assert.Nil(t, err)

	bet.state = bet.suspended
	err = bet.Close()
	assert.Nil(t, err)

	bet.state = bet.closed
	err = bet.Close()
	assert.NotNil(t, err)
	assert.Equal(t, "cannot close the bet because it has already been closed", err.Error())
}

func TestSuspend(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	bet, _ := NewBet(
		validId, validName,
		validDescription, challenged,
		validValue,
	)

	err := bet.Suspend()
	assert.Nil(t, err)

	bet.state = bet.suspended
	err = bet.Suspend()
	assert.Nil(t, err)

	bet.state = bet.closed
	err = bet.Suspend()
	assert.NotNil(t, err)
	assert.Equal(t, "cannot suspend the bet because it has already been closed", err.Error())
}

func TestChangeName(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	state := bet.NewMockIState(controller)
	state.EXPECT().ChangeName(gomock.Any()).Return(nil).MaxTimes(2)

	entity, _ := NewBet(validId, validName, validDescription, challenged, validValue)
	entity.state = state

	entity.name = validName
	err := entity.ChangeName(validName)
	assert.Nil(t, err)
	assert.Equal(t, validName, entity.name)

	entity.name = invalidNames[0]
	err = entity.ChangeName(invalidNames[0])
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("bet: %s", &errs.ErrInvalidBetName), err.Error())

	state.EXPECT().ChangeName(gomock.Any()).Return(errors.New("message...")).AnyTimes()

	entity.name = validName
	err = entity.ChangeName(validName)
	assert.NotNil(t, err)
	assert.Equal(t, "message...", err.Error())
	assert.Equal(t, entity.name, validName)

	entity.name = validName
	err = entity.ChangeName(invalidNames[0])
	assert.NotNil(t, err)
	assert.Equal(t, "message...", err.Error())
	assert.Equal(t, entity.name, validName)
}

func TestChangeDescription(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	state := bet.NewMockIState(controller)
	state.EXPECT().ChangeDescription(gomock.Any()).Return(nil).MaxTimes(2)

	entity, _ := NewBet(validId, validName, validDescription, challenged, validValue)
	entity.state = state

	entity.description = validDescription
	err := entity.ChangeDescription(validDescription)
	assert.Nil(t, err)
	assert.Equal(t, validDescription, entity.description)

	entity.description = invalidDescriptions[0]
	err = entity.ChangeDescription(invalidDescriptions[0])
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("bet: %s", &errs.ErrInvalidBetDescription), err.Error())

	state.EXPECT().ChangeDescription(gomock.Any()).Return(errors.New("message...")).AnyTimes()

	entity.description = validDescription
	err = entity.ChangeDescription(validDescription)
	assert.NotNil(t, err)
	assert.Equal(t, "message...", err.Error())
	assert.Equal(t, entity.description, validDescription)

	entity.description = validDescription
	err = entity.ChangeDescription(invalidDescriptions[0])
	assert.NotNil(t, err)
	assert.Equal(t, "message...", err.Error())
	assert.Equal(t, entity.description, validDescription)
}

func TestChangeValue(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	state := bet.NewMockIState(controller)
	state.EXPECT().ChangeValue(gomock.Any()).Return(nil).MaxTimes(2)

	entity, _ := NewBet(validId, validName, validDescription, challenged, validValue)
	entity.state = state

	entity.value = validValue
	err := entity.ChangeValue(validValue)
	assert.Nil(t, err)
	assert.Equal(t, validValue, entity.value)

	entity.value = invalidValues[0]
	err = entity.ChangeValue(invalidValues[0])
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("bet: %s", &errs.ErrInvalidBetValue), err.Error())

	state.EXPECT().ChangeValue(gomock.Any()).Return(errors.New("message...")).AnyTimes()

	entity.value = validValue
	err = entity.ChangeValue(validValue)
	assert.NotNil(t, err)
	assert.Equal(t, "message...", err.Error())
	assert.Equal(t, entity.value, validValue)

	entity.value = validValue
	err = entity.ChangeValue(invalidValues[0])
	assert.NotNil(t, err)
	assert.Equal(t, "message...", err.Error())
	assert.Equal(t, entity.value, validValue)
}

func TestDeposit(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	state := bet.NewMockIState(controller)
	state.EXPECT().Deposit(gomock.Any()).Return(nil).MaxTimes(2)

	entity, _ := NewBet(validId, validName, validDescription, challenged, validValue)
	entity.state = state

	err := entity.Deposit(100)
	assert.Nil(t, err)

	entity.credits = -100
	err = entity.Deposit(-100)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("bet: %s", &errs.ErrInvalidBetCredits), err.Error())

	state.EXPECT().Deposit(gomock.Any()).Return(errors.New("message...")).AnyTimes()
	err = entity.Deposit(100)
	assert.NotNil(t, err)
	assert.Equal(t, "message...", err.Error())
}

func TestWithdraw(t *testing.T) {
	controller := gomock.NewController(t)
	challenged := challenged.NewMockIChallenged(controller)
	state := bet.NewMockIState(controller)
	state.EXPECT().Withdraw(gomock.Any()).Return(nil).MaxTimes(2)

	entity, _ := NewBet(validId, validName, validDescription, challenged, validValue)
	entity.state = state

	err := entity.Withdraw(100)
	assert.Nil(t, err)

	entity.credits = -100
	err = entity.Withdraw(-100)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("bet: %s", &errs.ErrInvalidBetCredits), err.Error())

	state.EXPECT().Withdraw(gomock.Any()).Return(errors.New("message...")).AnyTimes()
	err = entity.Withdraw(100)
	assert.NotNil(t, err)
	assert.Equal(t, "message...", err.Error())
}
