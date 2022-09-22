package entity

type Suspended struct {
	bet *Bet
}

func (state Suspended) ChangeName(name string) error {
	state.bet.name = name
	return nil
}

func (state Suspended) ChangeDescription(description string) error {
	state.bet.description = description
	return nil
}

func (state Suspended) ChangeValue(value float32) error {
	state.bet.value = value
	return nil
}

func (state Suspended) Deposit(amount float32) error {
	state.bet.credits += amount
	return nil
}

func (state Suspended) Withdraw(amount float32) error {
	state.bet.credits -= amount
	return nil
}
