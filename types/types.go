package types

type ItemNotFound struct{}

func (nti ItemNotFound) Error() string {
	return "Item not found"
}
