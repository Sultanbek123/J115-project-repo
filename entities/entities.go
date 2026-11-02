package entities

type Item struct {
	ID           int64
	Model        string
	Price        float64
	Weight       float32
	InvoicePrice float64
}

type Category struct {
	ID   int64
	Name string
}
