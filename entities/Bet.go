package entities

// Bet represents one bet
type Bet struct {
	ID int
	OwnerName string
	Values map [int] int
	Place int
}