package errs

type ErrInvalidChallenged struct{ message string }

func (err *ErrInvalidChallenged) Error() string {
	return err.message
}

var (
	ErrInvalidChallengedId      = ErrInvalidChallenged{"Id is required and must be uuidV4"}
	ErrInvalidChallengedCredits = ErrInvalidChallenged{"Credits must be greater or equal 0"}
)
