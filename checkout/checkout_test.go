package checkout

import "testing"

func TestCalculateBasket(t *testing.T) {
	tests := []struct {
		inputSKUs     string
		expectedPrice int
	}{
		{"ABCDE", 115},
		{"AABBCC", 185},
		{"AAA", 130},
		{"BB", 45},
		{"EE", 0}, // Test case with no pricing rule for item E
	}

	for _, tt := range tests {
		totalPrice := CalculateBasket(tt.inputSKUs)
		if totalPrice != tt.expectedPrice {
			t.Fatalf("Expected price for input SKUs %s to be %d, but got %d", tt.inputSKUs, tt.expectedPrice, totalPrice)
		}
	}
}
