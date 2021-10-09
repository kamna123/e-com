package schema

type Order struct {
	UUID       string      `json:"uuid"`
	Lines      []OrderLine `json:"lines"`
	TotalPrice uint        `json:"total_price"`
	Status     string      `json:"status"`
	UserID     string      `json:"userid"`
}

type OrderBodyParam struct {
	UserID string               `json:"userid" validate:"required"`
	Lines  []OrderLineBodyParam `json:"lines,omitempty" validate:"required"`
}

type OrderQueryParam struct {
	UserID string `json:"userid" validate:"required"`
	Status string `json:"status,omitempty" form:"active"`
}
