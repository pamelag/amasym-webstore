package product

const (
	ULTSMALL = "ult_small"
	ULTMEDIUM = "ult_medium"
	ULTLARGE = "ult_large"
	ONEGBDATAPACK = "1gb"
	PROMOCODE = "DemoPromo"
)

type Product struct {
	Code string
	Name string
	UnitPrice float64
}

func New(code string, name string, price float64) *Product{
	return &Product{Code: code, Name: name, UnitPrice: price}
}

type Repository interface {
	FetchAll() ([]Product, error)
	Save(product *Product) error
	Query(code string) (*Product, error)
}