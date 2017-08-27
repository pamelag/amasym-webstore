package shoppingcart

import (
	"errors"
	"reflect"
	"product"
	"util"
)

type ShoppingCart struct {
	ID string
	Contents []Item
	TotalAmount float64
	PromoCode string
	PricingRules []PricingRule
}

type Item struct {
	Product *product.Product
	Quantity int
	ItemTotal float64
	Discount float64
	RuleMatched bool
}

func New(pricingRules []PricingRule) *ShoppingCart{
	return &ShoppingCart{PricingRules: pricingRules}
}

func NewItem(product *product.Product, quantity int) Item {
	item := Item{}
	item.Product = product
	item.Quantity = quantity
	return item
}

func (cart *ShoppingCart) Add(args ...interface{}) (error) {
	
	if args == nil {
		return errors.New("No items to add")
	}
	
	for _, param := range args {
		
		switch reflect.TypeOf(param).String() {
			case "shoppingcart.Item" :
				cart.Contents = append(cart.Contents, param.(Item))
			case "string" :
				cart.PromoCode = param.(string)
		}
		
	}
	
	return nil
}


func (cart *ShoppingCart) Items() ([]Item, error) {
	
	if len(cart.Contents) < 1 {
		return nil, errors.New("The ShoppingCart is empty")
	}
	
	return cart.Contents, nil
}


func (cart *ShoppingCart) Total() float64 {
	
	for i := 0; i < len(cart.Contents); i++ {
		item := &cart.Contents[i]
		for j := 0; j < len(cart.PricingRules); j++ {
			cart.PricingRules[j].CalculateTotal(item, cart)
		}
		
	}
	
	cart.TotalAmount = util.Round(cart.TotalAmount)
	
	return cart.TotalAmount
}