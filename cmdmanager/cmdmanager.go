package cmdmanager

import (
	"fmt"
)

type CMDManager struct {
}

func (cm *CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm with enter")

	var prices []string
	for {
		var price string
		fmt.Println("Print: ")
		fmt.Scanln(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cm *CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() *CMDManager {
	return &CMDManager{}
}
