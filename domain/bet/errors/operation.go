package errs

type ErrInvalidBetOperation struct{ message string }

func (err *ErrInvalidBetOperation) Error() string {
	return err.message
}

var (
	ErrCannotChangeName                         = ErrInvalidBetOperation{"Name cannot be changed, because bet state must be suspended"}
	ErrCannotChangeValue                        = ErrInvalidBetOperation{"Value cannot be changed, because bet state must be suspended"}
	ErrCannotChangeDescription                  = ErrInvalidBetOperation{"Description cannot be changed, because bet state must be suspended"}
	ErrCannotRegisterGamblerBetIsntOpen         = ErrInvalidBetOperation{"Cannot register a gambler because bet isn`t open"}
	ErrCannotRegisterGamblerInsufficientCredits = ErrInvalidBetOperation{"Cannot register a gambler, insufficient credits for bet"}
	ErrCannotUnregisterGamblerBetIsClosed       = ErrInvalidBetOperation{"Cannot unregister a gambler because bet isn`t open or suspended"}
	ErrCannotDeposit                            = ErrInvalidBetOperation{"Cannot deposit credits because bet is closed"}
	ErrCannotWithdraw                           = ErrInvalidBetOperation{"Cannot withdraw credits because bet is closed"}
)
