package model

type Product struct {
	SellerId  int
	OfferId   int
	Name      string
	Price     float64
	Quantity  int
	Available bool
}
