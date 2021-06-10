package database

import (
	"fmt"
)

func AddProduce(produce *Produce) error {
	ProduceList = append(ProduceList, *produce)
	return nil
}

func DeleteProduce(produceCode string) error {
	for i, produce := range ProduceList {
		if produce.ProduceCode == produceCode {
			ProduceList = append(ProduceList[:i], ProduceList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("ProduceCode %s not found", produceCode)
}

func FetchProduce() []Produce {
	return ProduceList
}

func InitializeDatabase() {
	AddProduce(&Produce{Name: "Lettuce", ProduceCode: "A12T-4GH7-QPL9-3N4M", UnitPrice: 3.46})
	AddProduce(&Produce{Name: "Peach", ProduceCode: "E5T6-9UI3-TH15-QR88", UnitPrice: 2.99})
	AddProduce(&Produce{Name: "Green Pepper", ProduceCode: "YRT6-72AS-K736-L4AR", UnitPrice: 0.79})
	AddProduce(&Produce{Name: "Gala Apple", ProduceCode: "TQ4C-VV6T-75ZX-1RMR", UnitPrice: 3.59})
}
