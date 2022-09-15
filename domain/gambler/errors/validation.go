package errs

type ErrInvalidGambler struct{ message string }

func (err *ErrInvalidGambler) Error() string {
	return err.message
}

var (
	ErrInvalidGamblerId      = ErrInvalidGambler{"Id is required and must be uuidV4"}
	ErrInvalidGamblerCredits = ErrInvalidGambler{"Credits must be greater or equal 0"}
)
