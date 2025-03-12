package models

type CarCategory string

const (
	Economy CarCategory = "ECONOMY"
	Premium CarCategory = "PREMIUM"
	Luxury  CarCategory = "LUXURY"
)

type Vehicle struct {
	ID       string
	Model    string
	Category CarCategory
}
