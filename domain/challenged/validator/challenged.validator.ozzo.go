package challenged

import (
	"fmt"
	"math"
	"strings"

	core "github.com/TaApostado/bets/domain/challenged"
	errs "github.com/TaApostado/bets/domain/challenged/errors"
	"github.com/TaApostado/common/notification"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

/*
Challenged validator implemented with ozzo-validation library (IChallengedValidator)

Documentation of library available: https://github.com/go-ozzo/ozzo-validation
*/
type ozzoChallengedValidator struct{}

const delimiter = ";"

func formatError(err error) string {
	return err.Error() + delimiter
}

func (ozzoChallengedValidator) Execute(challenged core.IChallenged) error {

	errs := validation.Errors{
		"id": validation.Validate(
			challenged.Id(), validation.Required.Error(formatError(&errs.ErrInvalidChallengedId)),
			is.UUIDv4.Error(formatError(&errs.ErrInvalidChallengedId))),

		"credits": validation.Validate(
			challenged.Credits(),
			validation.Required.Error(formatError(&errs.ErrInvalidChallengedCredits)),
			validation.Min(float32(0)).Error(formatError(&errs.ErrInvalidChallengedCredits)),
			validation.Max(math.MaxFloat32).Error(formatError(&errs.ErrInvalidChallengedCredits))),
	}

	for _, err := range errs {
		if err != nil && err.Error() != "" {
			message := strings.ReplaceAll(err.Error(), ";", "")
			challenged.Notificator().AddNotification(
				*notification.NewNotification("challenged", message),
			)
		}
	}

	if challenged.Notificator().HasNotifications() {
		return fmt.Errorf(challenged.Notificator().Notifications())
	}
	return nil
}
