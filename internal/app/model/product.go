package model

type Product struct {
	ID        int
	SellerId  int
	OfferId   int
	Name      string
	Price     float64
	Quantity  int
	Available bool
}

type ProductFilter struct {
	SellerId int
	OfferId  int
	Name     string
}

type ProductListStat struct {
	Added   int
	Updated int
	Deleted int
	Failed  int
}
