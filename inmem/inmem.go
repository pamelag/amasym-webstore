package inmem

import (
	"sync"
	"errors"
	"product"
)

var instance *ProductDataStore
var once sync.Once

type ProductDataStore struct {
	ProductTable map[string]*product.Product
	lock  sync.RWMutex
}

func GetInstance() product.Repository {
	once.Do(func() {
		prods := make(map[string]*product.Product, 0)
		instance = &ProductDataStore{ProductTable : prods}
	})
	return instance
}


func (pr *ProductDataStore) FetchAll() ([]product.Product, error) {
	pr.lock.RLock()
	defer pr.lock.RUnlock()
	
	productList := make([]product.Product, 0)
	for _, value := range pr.ProductTable {
		productList = append(productList, *value)
	}
	return productList, nil
}


func (pr *ProductDataStore) Save(product *product.Product) error {
	pr.lock.Lock()
	defer pr.lock.Unlock()
	
	pr.ProductTable[product.Code] = product
	return nil
}


func (pr *ProductDataStore) Query(code string) (*product.Product, error) {
	pr.lock.RLock()
	defer pr.lock.RUnlock()
	
	if code == "" {
		return nil, errors.New("No product code given")
	}
	if val, ok := pr.ProductTable[code]; ok {
		return val, nil
	} else {
		return nil, errors.New("Product with this code doesn't exist")
	}
}
