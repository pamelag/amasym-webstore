# amasym-webstore
##Amasym programming exercise built using Golang

#####This exercise is done entirely using Golang. It uses goland standard lib and the design decisions are as follows.

- ShoppingCart, Item and Product have been modeled to incorporate both behavior and data

- PricingRule has been defined as an Interface with implementations for special offers and promotions

- Repository interface has an in-memory implementation to hold Products



#####Setup Instructions

1. Install Go and Set GOPATH
2. Open a terminal and type 
```
cd amaysim-webstore/src/test

```
3. Once in test directory type
```
go test -v
```

3. You could change directory to shoppingcart and type "go test -v" to run cart_test and pricing_test

4. And you could also change directory to src and type 
```
go run main.go
or 
go run main.go -catalog

```

