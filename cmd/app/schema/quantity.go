package schema

type Quantity struct {
	UUID        string `json:"uuid"`
	ProductUUID string `json:"product_uuid"`
	Quantity    uint   `json:"quantity"`
}

type QuantityQueryParam struct {
	ProductUUID string `json:"product_uuid,omitempty" form:"code"`
}

type QuantityBodyParam struct {
	ProductUUID string `json:"product_uuid,omitempty" validate:"required"`
	Quantity    uint   `json:"quantity" validate:"required"`
}
