package schema

type CartBody struct {
	UUID      string `json:"uuid"`
	UserID    string `json:"userid"`
	ProductID string `json:"product_id"`
	Quantity  string `json:"quantity"`
	Price     string `json:"price"`
}

type CartDeleteBody struct {
	UserID    string `json:"userid"`
	ProductID string `json:"product_id"`
	Quantity  string `json:"quantity"`
}
