package entity

import (
	"fmt"
	"testing"

	errs "github.com/TaApostado/bets/domain/challenged/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var invalidIds = [...]string{"", "123", "abc", "123-as"}
var invalidCredits = [...]float32{-1}

func TestNewChallenged(t *testing.T) {
	validId := uuid.NewString()
	validCredits := float32(58)

	expectedErrInvalidId := fmt.Sprintf("challenged: %s", errs.ErrInvalidChallengedId.Error())
	for _, invalidId := range invalidIds {
		_, err := NewChallenged(invalidId, validCredits)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErrInvalidId, err.Error())
	}

	expectedErrInvalidCredits := fmt.Sprintf("challenged: %s", errs.ErrInvalidChallengedCredits.Error())
	for _, invalidCredit := range invalidCredits {
		_, err := NewChallenged(validId, invalidCredit)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErrInvalidCredits, err.Error())
	}

	_, err := NewChallenged(invalidIds[0], invalidCredits[0])
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), expectedErrInvalidId)
	assert.Contains(t, err.Error(), expectedErrInvalidCredits)

	_, err = NewChallenged(validId, validCredits)
	assert.Nil(t, err)
}

func TestId(t *testing.T) {
	id := uuid.NewString()
	credits := float32(58)
	challenged, _ := NewChallenged(id, credits)
	assert.Equal(t, id, challenged.Id())
}

func TestCredits(t *testing.T) {
	id := uuid.NewString()
	credits := float32(89)
	challenged, _ := NewChallenged(id, credits)
	assert.Equal(t, credits, challenged.Credits())
}

func TestDeposit(t *testing.T) {
	id := uuid.NewString()
	credits := float32(89)
	challenged, _ := NewChallenged(id, credits)
	assert.Equal(t, credits, challenged.Credits())

	err := challenged.Deposit(100)
	assert.Nil(t, err)
	credits += 100
	assert.Equal(t, credits, challenged.Credits())

	err = challenged.Deposit(-1)
	assert.NotNil(t, err)
	assert.Equal(t, credits, challenged.Credits())
}

func TestWithdraw(t *testing.T) {
	id := uuid.NewString()
	credits := float32(89)
	challenged, _ := NewChallenged(id, credits)
	assert.Equal(t, credits, challenged.Credits())

	err := challenged.Withdraw(50)
	assert.Nil(t, err)
	credits -= 50
	assert.Equal(t, credits, challenged.Credits())

	err = challenged.Withdraw(-1)
	assert.NotNil(t, err)
	assert.Equal(t, credits, challenged.Credits())

	err = challenged.Withdraw(credits + 1)
	assert.NotNil(t, err)
	assert.Equal(t, credits, challenged.Credits())
}
