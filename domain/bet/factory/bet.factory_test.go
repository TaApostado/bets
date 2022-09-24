package factory_test

import (
	"testing"

	errs "github.com/TaApostado/bets/domain/bet/errors"
	"github.com/TaApostado/bets/domain/bet/factory"
	domainChallenged "github.com/TaApostado/bets/domain/challenged"
	mockChallenged "github.com/TaApostado/bets/domain/challenged/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var challenged domainChallenged.IChallenged
var validName string = "Beber 2L de Agua"
var invalidName string = ""
var validValue float32 = 15
var invalidValue float32 = -1
var validDescription string = "Vou beber 2L de água em 2 minutos"
var invalidDescription string = ""

func setup(t *testing.T) {
	controller := gomock.NewController(t)
	challenged = mockChallenged.NewMockIChallenged(controller)
}

func TestBuildBetWithDefaultParameters(t *testing.T) {
	setup(t)
	f := factory.NewBetFactory(challenged, invalidName, invalidDescription, invalidValue)
	_, err := f.Build()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), errs.ErrInvalidBetName.Error())
	assert.Contains(t, err.Error(), errs.ErrInvalidBetDescription.Error())
	assert.Contains(t, err.Error(), errs.ErrInvalidBetValue.Error())

	f = factory.NewBetFactory(challenged, validName, validDescription, validValue)
	bet, err := f.Build()
	assert.Nil(t, err)
	assert.Equal(t, validName, bet.Name())
	assert.Equal(t, validDescription, bet.Description())
	assert.Equal(t, validValue, bet.Value())
}

func TestBuildOpenedBet(t *testing.T) {
	setup(t)
	f := factory.NewBetFactory(challenged, invalidName, validDescription, validValue)
	_, err := f.AsOpened().Build()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), errs.ErrInvalidBetName.Error())

	f = factory.NewBetFactory(challenged, validName, validDescription, validValue)
	bet, _ := f.AsOpened().Build()
	assert.Equal(t, validName, bet.Name())
	assert.Equal(t, validDescription, bet.Description())
	assert.Equal(t, validValue, bet.Value())
	assert.True(t, bet.IsOpen())
}

func TestBuildClosedBet(t *testing.T) {
	setup(t)
	f := factory.NewBetFactory(challenged, validName, invalidDescription, validValue)
	_, err := f.AsClosed().Build()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), errs.ErrInvalidBetDescription.Error())

	f = factory.NewBetFactory(challenged, validName, validDescription, validValue)
	bet, _ := f.AsClosed().Build()
	assert.Equal(t, validName, bet.Name())
	assert.Equal(t, validDescription, bet.Description())
	assert.Equal(t, validValue, bet.Value())
	assert.True(t, bet.IsClosed())
}

func TestBuildSuspendedBet(t *testing.T) {
	setup(t)
	f := factory.NewBetFactory(challenged, validName, validDescription, invalidValue)
	_, err := f.AsSuspended().Build()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), errs.ErrInvalidBetValue.Error())

	f = factory.NewBetFactory(challenged, validName, validDescription, validValue)
	bet, _ := f.AsSuspended().Build()
	assert.Equal(t, validName, bet.Name())
	assert.Equal(t, validDescription, bet.Description())
	assert.Equal(t, validValue, bet.Value())
	assert.True(t, bet.IsSuspended())
}
