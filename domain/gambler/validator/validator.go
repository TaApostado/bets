package validator

import domain "github.com/TaApostado/bets/domain/gambler"

type IGamblerValidator interface {
	Execute(bet domain.IGambler) error
}

func Validator() IGamblerValidator {
	return ozzoGamblerValidator{}
}
