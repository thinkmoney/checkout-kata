package pricing

import "testing"

func TestCalculatePrice(t *testing.T) {
	// Test for item with special price
	ruleA := PricingRule{UnitPrice: 50, SpecialPrice: SpecialPrice{Quantity: 3, Price: 130}}
	if price := CalculatePrice(4, ruleA); price != 180 {
		t.Errorf("Expected price for 4 items with unit price 50 and SpecialPrice 130 to be 170, but got %d", price)
	}

	// Test for item with special price
	ruleB := PricingRule{UnitPrice: 30, SpecialPrice: SpecialPrice{Quantity: 2, Price: 45}}
	if price := CalculatePrice(3, ruleB); price != 75 {
		t.Errorf("Expected price for 3 items with special price (2 for 45) to be 75, but got %d", price)
	}

	// Test for item with no special price
	ruleC := PricingRule{UnitPrice: 20, SpecialPrice: SpecialPrice{}}
	if price := CalculatePrice(4, ruleC); price != 80 {
		t.Errorf("Expected price for 4 items with price 20 to be 80, but got %d", price)
	}
}
