package schema

type CartBody struct {
	UUID        string `json:"uuid"`
	UserID      string `json:"userid"`
	ProductUUID string `json:"product_uuid"`
	Quantity    uint   `json:"quantity"`
	Price       string `json:"price"`
}

type CartDeleteBody struct {
	UserID      string `json:"userid"`
	ProductUUID string `json:"product_uuid"`
	Quantity    uint   `json:"quantity"`
}
