package errs

type ErrInvalidBetOperation struct{ message string }

func (err *ErrInvalidBetOperation) Error() string {
	return err.message
}

var (
	ErrCannotChangeName        = ErrInvalidBetOperation{"Name cannot be changed, because bet state must be suspended"}
	ErrCannotChangeValue       = ErrInvalidBetOperation{"Value cannot be changed, because bet state must be suspended"}
	ErrCannotChangeDescription = ErrInvalidBetOperation{"Description cannot be changed, because bet state must be suspended"}
	ErrCannotRegisterGambler   = ErrInvalidBetOperation{"Cannot register a gambler because bet is closed"}
	ErrCannotUnregisterGambler = ErrInvalidBetOperation{"Cannot unregister a gambler because bet is closed"}
	ErrCannotDeposit           = ErrInvalidBetOperation{"Cannot deposit credits because bet is closed"}
	ErrCannotWithdraw          = ErrInvalidBetOperation{"Cannot withdraw credits because bet is closed"}
)
