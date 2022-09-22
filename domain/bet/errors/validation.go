package errs

type ErrInvalidBet struct{ message string }

func (err *ErrInvalidBet) Error() string {
	return err.message
}

var (
	ErrInvalidBetId          = ErrInvalidBet{"Id is required and must be uuidV4"}
	ErrInvalidBetName        = ErrInvalidBet{"Name is required and the lenght must be between 10 and 50"}
	ErrInvalidBetDescription = ErrInvalidBet{"Description is required and the lenght must be between 15 and 200"}
	ErrInvalidBetValue       = ErrInvalidBet{"Value must be greater 0"}
	ErrInvalidBetCredits     = ErrInvalidBet{"Credits must be greater 0"}
	ErrChallengedIsRequired  = ErrInvalidBet{"Challanged is required"}
)
