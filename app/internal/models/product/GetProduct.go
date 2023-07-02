package product

type GetProductRequest struct {
	ID int `schema:"id"`
}

type GetProductResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
