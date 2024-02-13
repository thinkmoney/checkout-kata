package main

import (
	"checkout-kata/checkout"
	"fmt"
)

func main() {
	fmt.Println("Hello, thinkmoney!")
	baskets := []struct {
		inputSKUs string
	}{
		{"AABBCCDD"},
		{"EE"},
		{"ABCD"},
		{"AAAAA"},
		{"BBB"},
		{"CCCCC"},
	}

	for _, basket := range baskets {
		totalPrice := checkout.CalculateBasket(basket.inputSKUs)
		fmt.Printf("Basket: %s, Total Price: %d\n", basket.inputSKUs, totalPrice)
	}
}
