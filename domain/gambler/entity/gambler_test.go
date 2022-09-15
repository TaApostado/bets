package entity

import (
	"fmt"
	"testing"

	errs "github.com/TaApostado/bets/domain/gambler/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var invalidIds = [...]string{"", "123", "abc", "123-as"}
var invalidCredits = [...]float32{-1}

func TestNewGambler(t *testing.T) {
	validId := uuid.NewString()
	validCredits := float32(58)

	expectedErrInvalidId := fmt.Sprintf("gambler: %s", errs.ErrInvalidGamblerId.Error())
	for _, invalidId := range invalidIds {
		_, err := NewGambler(invalidId, validCredits)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErrInvalidId, err.Error())
	}

	expectedErrInvalidCredits := fmt.Sprintf("gambler: %s", errs.ErrInvalidGamblerCredits.Error())
	for _, invalidCredit := range invalidCredits {
		_, err := NewGambler(validId, invalidCredit)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErrInvalidCredits, err.Error())
	}

	_, err := NewGambler(invalidIds[0], invalidCredits[0])
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), expectedErrInvalidId)
	assert.Contains(t, err.Error(), expectedErrInvalidCredits)

	_, err = NewGambler(validId, validCredits)
	assert.Nil(t, err)
}

func TestId(t *testing.T) {
	id := uuid.NewString()
	credits := float32(58)
	gambler, _ := NewGambler(id, credits)
	assert.Equal(t, id, gambler.Id())
}

func TestCredits(t *testing.T) {
	id := uuid.NewString()
	credits := float32(89)
	gambler, _ := NewGambler(id, credits)
	assert.Equal(t, credits, gambler.Credits())
}

func TestDeposit(t *testing.T) {
	id := uuid.NewString()
	credits := float32(89)
	gambler, _ := NewGambler(id, credits)
	assert.Equal(t, credits, gambler.Credits())

	err := gambler.Deposit(100)
	assert.Nil(t, err)
	credits += 100
	assert.Equal(t, credits, gambler.Credits())

	err = gambler.Deposit(-1)
	assert.NotNil(t, err)
	assert.Equal(t, credits, gambler.Credits())
}

func TestWithdraw(t *testing.T) {
	id := uuid.NewString()
	credits := float32(89)
	gambler, _ := NewGambler(id, credits)
	assert.Equal(t, credits, gambler.Credits())

	err := gambler.Withdraw(50)
	assert.Nil(t, err)
	credits -= 50
	assert.Equal(t, credits, gambler.Credits())

	err = gambler.Withdraw(-1)
	assert.NotNil(t, err)
	assert.Equal(t, credits, gambler.Credits())

	err = gambler.Withdraw(credits + 1)
	assert.NotNil(t, err)
	assert.Equal(t, credits, gambler.Credits())
}
