package shoppingcart

import (
	"testing"
	"../product"
)

func Test_DefaultRule_CalculateTotal(t *testing.T) {
	
	expectedResult := 49.8
	
	small := product.New(product.ULTSMALL, "Unlimited 1GB", 24.9)
	
	defaultRule := CreateDefaultRule()
	pricingRules := make([]PricingRule, 0)
	pricingRules = append(pricingRules, defaultRule)
	
	cart := New(pricingRules)
	
	smallItem := NewItem(small, 2)
	
	defaultRule.CalculateTotal(&smallItem, cart)
	
	actualResult := smallItem.ItemTotal
	
	if actualResult != expectedResult {
		t.Fatalf("Expected %f but got %f", expectedResult, actualResult)
	}
	
}


func Test_3For2DealUnlimited1GB_CalculateTotal(t *testing.T) {
	
	expectedResult := 49.8
	
	small := product.New(product.ULTSMALL, "Unlimited 1GB", 24.9)
	
	_3For2DealUnlimited1GBRule := Create3For2DealUnlimited1GBRule()
	pricingRules := make([]PricingRule, 0)
	pricingRules = append(pricingRules, _3For2DealUnlimited1GBRule)
	
	cart := New(pricingRules)
	
	smallItem := NewItem(small, 3)
	
	_3For2DealUnlimited1GBRule.CalculateTotal(&smallItem, cart)
	
	actualResult := smallItem.ItemTotal
	
	if actualResult != expectedResult {
		t.Fatalf("Expected %f but got %f", expectedResult, actualResult)
	}
	
}


func Test_3Unlimited5GB_CalculateTotal(t *testing.T) {
	
	expectedResult := 159.6
	
	large := product.New(product.ULTLARGE, "Unlimited 5GB", 44.9)
	
	_3Unlimited5GBRule := Create3Unlimited5GBRule()
	pricingRules := make([]PricingRule, 0)
	pricingRules = append(pricingRules, _3Unlimited5GBRule)
	
	cart := New(pricingRules)
	
	largeItem := NewItem(large, 4)
	
	_3Unlimited5GBRule.CalculateTotal(&largeItem, cart)
	
	actualResult := largeItem.ItemTotal
	
	if actualResult != expectedResult {
		t.Fatalf("Expected %f but got %f", expectedResult, actualResult)
	}
	
}


func Test_1GBDataPack_CalculateTotal(t *testing.T) {
	
	expectedResult := product.ONEGBDATAPACK
	
	medium := product.New(product.ULTMEDIUM, "Unlimited 2GB", 29.9)
	
	_1GBDataPackRule := Create1GBDataPackRule()
	pricingRules := make([]PricingRule, 0)
	pricingRules = append(pricingRules, _1GBDataPackRule)
	
	cart := New(pricingRules)
	
	mediumItem := NewItem(medium, 1)
	
	_1GBDataPackRule.CalculateTotal(&mediumItem, cart)
	
	actualResult := cart.Contents[0].Product.Code
	
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}


func Test_PromoCode_CalculateTotal(t *testing.T) {
	
	expectedResult := 40.41
	
	large := product.New(product.ULTLARGE, "Unlimited 5GB", 44.9)
	
	defaultRule := CreateDefaultRule()
	promoCodeRule := CreatePromoCodeRule()
	pricingRules := make([]PricingRule, 0)
	pricingRules = append(pricingRules, defaultRule)
	pricingRules = append(pricingRules, promoCodeRule)
	
	cart := New(pricingRules)
	cart.PromoCode = product.PROMOCODE
	
	largeItem := NewItem(large, 1)
	
	defaultRule.CalculateTotal(&largeItem, cart)
	promoCodeRule.CalculateTotal(&largeItem, cart)
	
	actualResult := largeItem.ItemTotal
	
	if actualResult != expectedResult {
		t.Fatalf("Expected %f but got %f", expectedResult, actualResult)
	}
	
}