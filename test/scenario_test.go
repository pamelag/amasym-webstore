package test

import (
	"../product"
	"../inmem"
	ShoppingCart "../shoppingcart"
	"testing"
	"log"
	"bytes"
	"strconv"
	"strings"
	"fmt"
)

var small, medium, large, _1GB *product.Product
var err error
var productStore product.Repository
var pricingRules []ShoppingCart.PricingRule

func init () {
	
	createProductStore()
	initializeCatalog()
	initializePricingRules()
}

func createProductStore() {
	
	productStore = inmem.GetInstance()
	productStore.Save(product.New(product.ULTSMALL, "Unlimited 1GB", 24.90))
	productStore.Save(product.New(product.ULTMEDIUM, "Unlimited 2GB", 29.90))
	productStore.Save(product.New(product.ULTLARGE, "Unlimited 5GB", 44.90))
	productStore.Save(product.New(product.ONEGBDATAPACK, "1 GB Data-pack", 9.90))
}

func initializeCatalog() {
	
	if small, err = productStore.Query("ult_small"); err != nil {
		log.Fatal(err)
	}
	
	if medium, err = productStore.Query("ult_medium"); err != nil {
		log.Fatal(err)
	}
	
	if large, err = productStore.Query("ult_large"); err != nil {
		log.Fatal(err)
	}
	
	if _1GB, err = productStore.Query("1gb"); err != nil {
		log.Fatal(err)
	}
}

func initializePricingRules() {
	
	_3For2DealUnlimited1GBPricingRule := ShoppingCart.Create3For2DealUnlimited1GBRule()
	_3Unlimited5GBPricingRule := ShoppingCart.Create3Unlimited5GBRule()
	_DefaultPricingRule := ShoppingCart.CreateDefaultRule()
	_1GBDataPackPricingRule := ShoppingCart.Create1GBDataPackRule()
	_PromoCodePricingRule := ShoppingCart.CreatePromoCodeRule()
	
	pricingRules = make([]ShoppingCart.PricingRule, 0)
	pricingRules = append(pricingRules, _3For2DealUnlimited1GBPricingRule)
	pricingRules = append(pricingRules, _3Unlimited5GBPricingRule)
	pricingRules = append(pricingRules, _1GBDataPackPricingRule)
	pricingRules = append(pricingRules, _DefaultPricingRule)
	pricingRules = append(pricingRules, _PromoCodePricingRule)
}

func printItems(contents []ShoppingCart.Item, promo string, cartItems bool) {
	var buffer bytes.Buffer
	
	if !cartItems {
		buffer.WriteString("Items Added : ")
	} else {
		buffer.WriteString("Actual Cart Items : ")
	}
	
	
	for i := 0; i < len(contents); i++ {
		buffer.WriteString(strconv.Itoa(contents[i].Quantity))
		buffer.WriteString(" X ")
		buffer.WriteString(contents[i].Product.Name)
		if i < len(contents) - 1 {
			buffer.WriteString(", ")
		}
		
	}
	
	if len(strings.TrimSpace(promo)) > 1 {
		if len(contents) > 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(promo)
		buffer.WriteString(" Promo Applied")
	}
	fmt.Println(buffer.String())
}

func printTotal(total float64, label string) {
	fmt.Println(label + " Cart Total : $", total)
}

func printExpectedItems() {
	
	var buffer bytes.Buffer
	
	buffer.WriteString("Expected Cart Items : ")
	
	buffer.WriteString("1 X ")
	buffer.WriteString(small.Name)
	buffer.WriteString(", ")
	buffer.WriteString("2 X ")
	buffer.WriteString(medium.Name)
	buffer.WriteString(", ")
	buffer.WriteString("2 X ")
	buffer.WriteString(_1GB.Name)
	
	fmt.Println(buffer.String())
	
}

func TestShoppingScenario_1(t *testing.T) {
	
	expectedTotal := 94.7
	
	cart := ShoppingCart.New(pricingRules)
	
	smallItem := ShoppingCart.NewItem(small, 3)
	largeItem := ShoppingCart.NewItem(large, 1)

	cart.Add(smallItem)
	cart.Add(largeItem)
	
	
	itemsAdded, err := cart.Items()
	
	if err != nil {
		log.Println("Error while fetching cart items", err)
	}
	
	printItems(itemsAdded, "", false)
	
	
	cartItems, err := cart.Items()
	
	if err != nil {
		log.Println("Error while fetching cart items", err)
	}
	
	actualTotal := cart.Total()
	
	if actualTotal != expectedTotal {
		t.Fatalf("Expected %f but got %f", expectedTotal, actualTotal)
	}
	
	printItems(cartItems, "", true)
	printTotal(expectedTotal, "Expected")
	printTotal(actualTotal, "Actual")

}


// scenario 2 -
// 2 x Unlimited 1 GB
// 4 x Unlimited 5 GB
func TestShoppingScenario_2(t *testing.T) {
	
	expectedTotal := 209.4
	
	
	cart := ShoppingCart.New(pricingRules)
	
	smallItem := ShoppingCart.NewItem(small, 2)
	largeItem := ShoppingCart.NewItem(large, 4)
	
	cart.Add(smallItem)
	cart.Add(largeItem)
	
	itemsAdded, err := cart.Items()
	
	if err != nil {
		log.Println("Error while fetching cart items", err)
	}
	
	printItems(itemsAdded, "", false)
	
	actualTotal := cart.Total()
	
	cartItems, err := cart.Items()
	
	if err != nil {
		log.Println("Error while fetching cart items", err)
	}
	
	if actualTotal != expectedTotal {
		t.Fatalf("Expected %f but got %f", expectedTotal, actualTotal)
	}
	
	printItems(cartItems, "", true)
	printTotal(expectedTotal, "Expected")
	printTotal(actualTotal, "Actual")
}

// scenario 3 -
// 1 x Unlimited 1 GB
// 2 X Unlimited 2 GB

func TestShoppingScenario_3(t *testing.T) {
	
	expectedTotal := 84.7
	expectedNumberOfItems := 3
	
	cart := ShoppingCart.New(pricingRules)
	
	smallItem := ShoppingCart.NewItem(small, 1)
	mediumItem := ShoppingCart.NewItem(medium, 2)
	
	cart.Add(smallItem)
	cart.Add(mediumItem)
	
	itemsAdded, err := cart.Items()
	
	if err != nil {
		log.Println("Error while fetching cart items", err)
	}
	
	printItems(itemsAdded, "", false)
	
	actualTotal := cart.Total()
	
	cartItems, err := cart.Items()
	
	if err != nil {
		log.Println("Error while fetching cart items", err)
	}
	
	
	actualNumberOfItems := len(cartItems)
	
	if actualTotal != expectedTotal {
		t.Fatalf("Expected %f but got %f", expectedTotal, actualTotal)
	}
	
	if actualNumberOfItems != expectedNumberOfItems {
		t.Fatalf("Expected %i but got %i", expectedNumberOfItems, actualNumberOfItems)
	}
	
	printExpectedItems()
	printItems(cartItems, "", true)
	
	printTotal(expectedTotal, "Expected")
	printTotal(expectedTotal, "Actual")
}


// scenario 4 -
// 1 x Unlimited 1 GB
// 1 x 1 GB Data-pack + 'DemoPromo' Promo Applied
func TestShoppingScenario_4(t *testing.T) {
	
	expectedTotal := 31.32
	
	cart := ShoppingCart.New(pricingRules)
	
	smallItem := ShoppingCart.NewItem(small, 1)
	_1GBItem := ShoppingCart.NewItem(_1GB, 1)
	
	promoCode := product.PROMOCODE
	
	param := make([]interface{}, 0)
	param = append(param, smallItem)
	param = append(param, _1GBItem)
	param = append(param, promoCode)
	
	cart.Add(param...)
	
	itemsAdded, err := cart.Items()
	
	if err != nil {
		log.Println("Error while fetching cart items", err)
	}
	
	printItems(itemsAdded, product.PROMOCODE, false)
	
	actualTotal := cart.Total()
	
	cartItems, err := cart.Items()
	
	if err != nil {
		log.Println("Error while fetching cart items", err)
	}
	
	if actualTotal != expectedTotal {
		t.Fatalf("Expected %f but got %f", expectedTotal, actualTotal)
	}
	
	printItems(cartItems, "", true)
	printTotal(expectedTotal, "Expected")
	printTotal(expectedTotal, "Actual")
	
}