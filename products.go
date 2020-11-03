package products

type ProductID int

type CategorieID int

type Product struct {
	ID         ProductID
	Name       string
	categories []CategorieID
}

type ListProduct struct {
	list map[ProductID]Product
}

func (lp *ListProduct) Init(){

}

func (lp *ListProduct) GetListJSON(){

}

