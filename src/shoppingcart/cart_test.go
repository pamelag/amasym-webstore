package shoppingcart

import (
	"testing"
	"../product"
)

var pricingRules []PricingRule

func setUp() {
	
	_3For2DealUnlimited1GBPricingRule := Create3For2DealUnlimited1GBRule()
	_3Unlimited5GBPricingRule := Create3Unlimited5GBRule()
	_DefaultPricingRule := CreateDefaultRule()
	_1GBDataPackPricingRule := Create1GBDataPackRule()
	_PromoCodePricingRule := CreatePromoCodeRule()
	
	pricingRules = make([]PricingRule, 0)
	pricingRules = append(pricingRules, _3For2DealUnlimited1GBPricingRule)
	pricingRules = append(pricingRules, _3Unlimited5GBPricingRule)
	pricingRules = append(pricingRules, _1GBDataPackPricingRule)
	pricingRules = append(pricingRules, _DefaultPricingRule)
	pricingRules = append(pricingRules, _PromoCodePricingRule)
}

func TestShoppingCart_Add(t *testing.T) {
	
	setUp()
	
	expectedResult := 2
	
	cart := New(pricingRules)
	
	small := product.New(product.ULTSMALL, "Unlimited 1GB", 24.9)
	large := product.New(product.ULTLARGE, "Unlimited 5GB", 44.9)
	
	smallItem := NewItem(small, 3)
	largeItem := NewItem(large, 1)
	
	cart.Add(smallItem)
	cart.Add(largeItem)
	
	actualResult := len(cart.Contents)
	
	if actualResult != expectedResult {
		t.Fatalf("Expected %i but got %i", expectedResult, actualResult)
	}
}



func TestShoppingCart_Items(t *testing.T) {
	
	setUp()
	
	expectedResult := 2
	
	cart := New(pricingRules)
	
	small := product.New(product.ULTSMALL, "Unlimited 1GB", 24.9)
	large := product.New(product.ULTLARGE, "Unlimited 5GB", 44.9)
	
	smallItem := NewItem(small, 3)
	largeItem := NewItem(large, 1)
	
	cart.Add(smallItem)
	cart.Add(largeItem)
	
	items, err := cart.Items()
	
	if err != nil {
		t.FailNow()
	}
	
	actualResult := len(items)
	
	if actualResult != expectedResult {
		t.Fatalf("Expected %i but got %i", expectedResult, actualResult)
	}
}


func TestShoppingCart_Total(t *testing.T) {
	
	setUp()
	
	expectedResult := 69.8
	
	cart := New(pricingRules)
	
	small := product.New(product.ULTSMALL, "Unlimited 1GB", 24.9)
	large := product.New(product.ULTLARGE, "Unlimited 5GB", 44.9)
	
	smallItem := NewItem(small, 1)
	largeItem := NewItem(large, 1)
	
	cart.Add(smallItem)
	cart.Add(largeItem)
	
	
	actualResult := cart.Total()
	
	if actualResult != expectedResult {
		t.Fatalf("Expected %f but got %f", expectedResult, actualResult)
	}
}
