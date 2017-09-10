# go-webstore
## go programming exercise built using Golang

This exercise is done entirely using Golang. It uses standard lib only and my design decisions are as follows :

- ShoppingCart, Item and Product have been modeled as domain entities with factory methods to create new instances

- The PricingRule has been defined as an interface with implementations for special offers and promotions

- Product has a Repository interface with an in-memory implementation to hold Products



##### Setup Instructions

1. Install Go

2. Open a terminal and type 
```
go get go-webstore

```
3. Once in test directory type
```
go test -v
```

3. You could change directory to shoppingcart and type "go test -v" to run cart_test and pricing_test


