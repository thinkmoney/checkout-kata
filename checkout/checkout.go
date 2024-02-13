package checkout

import "checkout-kata/pricing"

// CalculateBasket calculates the total price for a string of SKUs using loaded pricing rules
func CalculateBasket(inputSKUs string) int {
	co := NewCheckout()

	for _, item := range inputSKUs {
		sku := string(item)
		co.Scan(sku)
	}

	return co.GetTotalPrice()
}

// Checkout represents a supermarket checkout
type Checkout struct {
	basket map[string]int
}

// NewCheckout initializes a new Checkout instance
func NewCheckout() *Checkout {
	return &Checkout{
		basket: make(map[string]int),
	}
}

// Scan adds an item to the basket
func (co *Checkout) Scan(item string) {
	co.basket[item]++
}

// GetTotalPrice calculates the total price of all items in the basket
func (co *Checkout) GetTotalPrice() int {
	total := 0
	for item, quantity := range co.basket {
		if rule, ok := pricing.GetPricingRule(item); ok {
			total += pricing.CalculatePrice(quantity, rule)
		}
	}
	return total
}
