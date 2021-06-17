package database

import (
	"strings"
	"fmt"
	"regexp"
)

func AddProduce(produce *Produce) error {
	produceNameReg := regexp.MustCompile(`[a-zA-Z0-9 ]+`).MatchString
	priceString := strings.Split(fmt.Sprintf("%v", produce.UnitPrice), ".")
	decimalPlace := ""
	if len(priceString) > 1 {
		decimalPlace = priceString[1]
	}
	if produceNameReg(produce.Name) && validateProduceCode(produce.ProduceCode) && len(decimalPlace) < 3{
		ProduceList = append(ProduceList, *produce)
	}else{
		return fmt.Errorf("Produce is not in a valid format")
	}
	
	return nil
}

func DeleteProduce(produceCode string) error {
	if !validateProduceCode(produceCode) {
		return fmt.Errorf("Produce code: %s is not in a valid format", produceCode)
	}
	for i, produce := range ProduceList {
		if produce.ProduceCode == produceCode {
			ProduceList = append(ProduceList[:i], ProduceList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("ProduceCode %s not found", produceCode)
}

func FetchProduceAll() []Produce {
	return ProduceList
}

func FetchProduce(produceCode string) (Produce, error) {
	if !validateProduceCode(produceCode) {
		return Produce{}, fmt.Errorf("Produce code: %s is not in a valid format", produceCode)
	}
	for _, produce := range ProduceList{
		if produce.ProduceCode == produceCode {
			return produce, nil
		}
	}
	return Produce{}, fmt.Errorf("Produce code: %s is not found", produceCode)
}

func validateProduceCode(produceCode string) bool {
	code := strings.Split(produceCode, "-")
	produceCodeReg := regexp.MustCompile(`[a-zA-Z0-9]+`).MatchString
	if len(code) != 4 {
		fmt.Errorf("ProduceCode %s is not valid", produceCode)
		return false
	}
	for _, codeGroup := range code{
		if !produceCodeReg(codeGroup) || (len(codeGroup) != 4 ) {
			fmt.Errorf("ProduceCode %s is not valid", produceCode)
			return false 
		}
	}
	return true
}
func InitializeDatabase() {
	AddProduce(&Produce{Name: "Lettuce", ProduceCode: "A12T-4GH7-QPL9-3N4M", UnitPrice: 3.46})
	AddProduce(&Produce{Name: "Peach", ProduceCode: "E5T6-9UI3-TH15-QR88", UnitPrice: 2.99})
	AddProduce(&Produce{Name: "Green Pepper", ProduceCode: "YRT6-72AS-K736-L4AR", UnitPrice: 0.79})
	AddProduce(&Produce{Name: "Gala Apple", ProduceCode: "TQ4C-VV6T-75ZX-1RMR", UnitPrice: 3.59})
}
