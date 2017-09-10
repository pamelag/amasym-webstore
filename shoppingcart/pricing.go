package shoppingcart

import (
	"product"
)


//-----------------------------PricingRule--------------------------//

type PricingRule interface {
	CalculateTotal(item *Item, cart *ShoppingCart)
}


//-------------------------------DefaultRule--------------------//

type _DefaultRule struct {
	name string
	total float64
}

func (_defaultRule *_DefaultRule) CalculateTotal(item *Item, cart *ShoppingCart) {
	
	if !item.RuleMatched {
		item.ItemTotal = float64(item.Quantity) * item.Product.UnitPrice
	}
}

func CreateDefaultRule() PricingRule {
	return &_DefaultRule{name: "_DefaultRule"}
}


//-------------------------------3For2DealUnlimited1GB--------------------//

type _3For2DealUnlimited1GB struct {
	name string
	total float64
	next PricingRule
}


func (_3for2U1Rule *_3For2DealUnlimited1GB) CalculateTotal(item *Item, cart *ShoppingCart) {
	
	if _3for2U1Rule.isSatisfied(item.Product.Code, item.Quantity) {
		
		item.RuleMatched = true
		
		var chargeableQuantity int
		quotient := item.Quantity / 3 // 6/3 = 2 7/3 = 2
		modulo := item.Quantity % 3 // 6/3 = 0, 5/3 = #%$#
		
		if modulo == 0 {
			chargeableQuantity = 2 * quotient
		}
		
		if modulo > 0 {
			chargeableQuantity = 2 * quotient + modulo
		}
		
		item.ItemTotal = float64(chargeableQuantity) * item.Product.UnitPrice
		
	}
}



func (_3fo2U1Rule *_3For2DealUnlimited1GB) isSatisfied(productCode string, quantity int) bool {
	return productCode == product.ULTSMALL && quantity > 2
}



func Create3For2DealUnlimited1GBRule() PricingRule {
	return &_3For2DealUnlimited1GB{name: "_3For2DealUnlimited1GB"}
}



//-------------------------------3Unlimited5GB-----------------------------//

type _3Unlimited5GB struct {
	name string
	total float64
}


func (_3Ult5GBRule *_3Unlimited5GB) CalculateTotal(item *Item, cart *ShoppingCart) {
	
	if _3Ult5GBRule.isSatisfied(item.Product.Code, item.Quantity) {
		
		item.RuleMatched = true
		discountedPrice := 39.90
		
		item.ItemTotal = float64(item.Quantity) * discountedPrice
		
	}
}


func (_3Ult5GBRule *_3Unlimited5GB) isSatisfied(productCode string, quantity int) bool {
	return productCode == product.ULTLARGE && quantity > 3
}


func Create3Unlimited5GBRule() PricingRule {
	return &_3Unlimited5GB{name: "_3Unlimited5GB"}
}



//-------------------------------1GBDataPack-----------------------------//

type _1GBDataPack struct {
	name string
	total float64
}


func (_1GBDPRule *_1GBDataPack) CalculateTotal(item *Item, cart *ShoppingCart) {
	
	if _1GBDPRule.isSatisfied(item.Product.Code, item.Quantity) {
		_1GB := product.New(product.ONEGBDATAPACK, "1 GB Data-pack", 9.90)
		item := NewItem(_1GB, item.Quantity)
		item.RuleMatched = true
		cart.Add(item)
	}
}



func (_1GBDPRule *_1GBDataPack) isSatisfied(productCode string, quantity int) bool {
	return productCode == product.ULTMEDIUM && quantity > 0
}



func Create1GBDataPackRule() PricingRule {
	return &_1GBDataPack{name: "_1GBDataPack"}
}



//-------------------------------PromoCode-----------------------------//

type _PromoCode struct {
	name string
	total float64
}


func (promoCode *_PromoCode) CalculateTotal(item *Item, cart *ShoppingCart) {
	
	if promoCode.isSatisfied(cart.PromoCode) {
		item.ItemTotal = item.ItemTotal - (item.ItemTotal * 0.1)
	}
	
	cart.TotalAmount = cart.TotalAmount + item.ItemTotal
}



func (promoCode *_PromoCode) isSatisfied(label string) bool {
	return label == product.PROMOCODE
}



func CreatePromoCodeRule() PricingRule {
	return &_PromoCode{name: "_PromoCode"}
}
