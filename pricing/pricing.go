package pricing

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// PricingRule represents the pricing rule for an item
type PricingRule struct {
	UnitPrice    int          `json:"unit_price"`
	SpecialPrice SpecialPrice `json:"special_price"`
}

// SpecialPrice represents a special pricing scheme
type SpecialPrice struct {
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

var PricingRules map[string]PricingRule

// LoadPricingRules loads pricing rules from a JSON file
func LoadPricingRules(filename string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := filepath.Join(cwd, filename)

	file, err := os.Open(filePath)
	if err != nil {
		// try directory above current one if running from test
		filePath = filepath.Join("../", filename)
	}
	file, err = os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&PricingRules)
	if err != nil {
		return err
	}

	return nil
}

// init initializes PricingRules map with pricing rules from JSON file
func init() {
	PricingRules = make(map[string]PricingRule)
	err := LoadPricingRules("pricing_rules.json")
	if err != nil {
		panic(err)
	}
}

// GetPricingRule returns the pricing rule for a given SKU
func GetPricingRule(sku string) (PricingRule, bool) {
	rule, ok := PricingRules[sku]
	return rule, ok
}

// CalculatePrice calculates the total price for given quantity and pricing rule
func CalculatePrice(quantity int, rule PricingRule) int {
	if rule.SpecialPrice.Quantity > 0 {
		specialOffers := quantity / rule.SpecialPrice.Quantity
		remainder := quantity % rule.SpecialPrice.Quantity
		return specialOffers*rule.SpecialPrice.Price + remainder*rule.UnitPrice
	}
	return quantity * rule.UnitPrice
}
