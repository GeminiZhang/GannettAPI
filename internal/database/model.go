package database

var (
	// ProduceList contains array of produce data
	ProduceList []Produce
)

// Produce for groceries
type Produce struct {
	Name        string  `json:"Name"`
	ProduceCode string  `json:"Produce Code"`
	UnitPrice   float32 `json:"Unit Price"`
}
