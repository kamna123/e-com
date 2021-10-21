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
	UserID string `json:"user_id" validate:"required" form:"userid"`
	Status string `json:"status,omitempty" form:"status"`
}
type RazorPayOrderParam struct {
	Amount   uint   `json:"amount" validate:"required"`
	Currency string `json:"currency"`
}

type RazorPayResp struct {
	ID string `json:"id"`
}
