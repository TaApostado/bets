package validator

import domain "github.com/TaApostado/bets/domain/bet"

type IBetValidator interface {
	Execute(bet domain.IBet) error
}

func Validator() IBetValidator {
	return ozzoBetValidator{}
}
