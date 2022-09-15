package validator

import (
	"fmt"
	"math"
	"strings"

	core "github.com/TaApostado/bets/domain/gambler"
	errs "github.com/TaApostado/bets/domain/gambler/errors"
	"github.com/TaApostado/common/notification"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

/*
Bet validator implemented with ozzo-validation library (IBetValidator)

Documentation of library available: https://github.com/go-ozzo/ozzo-validation
*/
type ozzoGamblerValidator struct{}

const delimiter = ";"

func formatError(err error) string {
	return err.Error() + delimiter
}

func (ozzoGamblerValidator) Execute(gambler core.IGambler) error {

	errs := validation.Errors{
		"id": validation.Validate(
			gambler.Id(), validation.Required.Error(formatError(&errs.ErrInvalidGamblerId)),
			is.UUIDv4.Error(formatError(&errs.ErrInvalidGamblerId))),

		"credits": validation.Validate(
			gambler.Credits(),
			validation.Required.Error(formatError(&errs.ErrInvalidGamblerCredits)),
			validation.Min(float32(0)).Error(formatError(&errs.ErrInvalidGamblerCredits)),
			validation.Max(math.MaxFloat32).Error(formatError(&errs.ErrInvalidGamblerCredits))),
	}

	for _, err := range errs {
		if err != nil && err.Error() != "" {
			message := strings.ReplaceAll(err.Error(), ";", "")
			gambler.Notificator().AddNotification(
				*notification.NewNotification("gambler", message),
			)
		}
	}

	if gambler.Notificator().HasNotifications() {
		return fmt.Errorf(gambler.Notificator().Notifications())
	}
	return nil
}
