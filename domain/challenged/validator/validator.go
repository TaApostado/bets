package challenged

import domain "github.com/TaApostado/bets/domain/challenged"

type IChallengedValidator interface {
	Execute(bet domain.IChallenged) error
}

func Validator() IChallengedValidator {
	return ozzoChallengedValidator{}
}
