package main

import (
	"product"
	"inmem"
	"fmt"
	"flag"
	"log"
)

var productStore product.Repository

func init () {
	
	productStore = inmem.GetInstance()
	productStore.Save(product.New(product.ULTSMALL, "Unlimited 1GB", 24.90))
	productStore.Save(product.New(product.ULTMEDIUM, "Unlimited 2GB", 29.90))
	productStore.Save(product.New(product.ULTLARGE, "Unlimited 5GB", 44.90))
	productStore.Save(product.New(product.ONEGBDATAPACK, "1 GB Data-pack", 9.90))
}


func main() {
	
	displayCatalog := flag.Bool("catalog", true, "Product Catalog")
	
	if *displayCatalog {
		
		var products []product.Product
		var err error
		
		if products, err = productStore.FetchAll(); err != nil {
			log.Fatal("Something went wrong")
		}
		
		for _, product := range products {
			fmt.Println(product.Code, " | ", product.Name, " | ", product.UnitPrice)
		}
		
		
	} else {
		// default block
		fmt.Println("Shopping Cart Programming Exercise")
	}
	
	
	
}
