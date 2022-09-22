package validator

import (
	"fmt"
	"math"
	"strings"

	domain "github.com/TaApostado/bets/domain/bet"
	errs "github.com/TaApostado/bets/domain/bet/errors"
	"github.com/TaApostado/common/notification"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

/*
Bet validator implemented with ozzo-validation library (IBetValidator)

Documentation of library available: https://github.com/go-ozzo/ozzo-validation
*/
type ozzoBetValidator struct{}

const delimiter = ";"

func formatError(err error) string {
	return err.Error() + delimiter
}

func (ozzoBetValidator) Execute(bet domain.IBet) error {

	errs := validation.Errors{
		"id": validation.Validate(
			bet.Id(), validation.Required.Error(formatError(&errs.ErrInvalidBetId)),
			is.UUIDv4.Error(formatError(&errs.ErrInvalidBetId))),

		"name": validation.Validate(
			bet.Name(), validation.Required.Error(formatError(&errs.ErrInvalidBetName)),
			validation.Length(10, 50).Error(formatError(&errs.ErrInvalidBetName))),

		"description": validation.Validate(
			bet.Description(),
			validation.Required.Error(formatError(&errs.ErrInvalidBetDescription)),
			validation.Length(15, 200).Error(formatError(&errs.ErrInvalidBetDescription))),

		"value": validation.Validate(
			bet.Value(),
			validation.Required.Error(formatError(&errs.ErrInvalidBetValue)),
			validation.Min(float32(1)).Error(formatError(&errs.ErrInvalidBetValue)),
			validation.Max(math.MaxFloat32).Error(formatError(&errs.ErrInvalidBetValue))),

		"credits": validation.Validate(
			bet.Credits(),
			validation.Min(float32(0.0)).Error(formatError(&errs.ErrInvalidBetCredits)),
			validation.Max(math.MaxFloat32).Error(formatError(&errs.ErrInvalidBetCredits))),

		"challenged": validation.Validate(
			bet.Challenged(),
			validation.NotNil.Error(formatError(&errs.ErrChallengedIsRequired))),
	}

	for _, err := range errs {
		if err != nil && err.Error() != "" {
			message := strings.ReplaceAll(err.Error(), ";", "")
			bet.Notificator().AddNotification(
				*notification.NewNotification("bet", message),
			)
		}
	}

	if bet.Notificator().HasNotifications() {
		return fmt.Errorf(bet.Notificator().Notifications())
	}
	return nil
}
