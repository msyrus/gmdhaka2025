package service

type Product struct {
	products ProductRepository
}

func NewProduct(products ProductRepository) *Product {
	return &Product{
		products: products,
	}
}
