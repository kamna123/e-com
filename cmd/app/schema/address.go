package schema

type Address struct {
	UUID   string `json:"uuid"`
	UserID string `json:"user_id"`

	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	Country       string `json:"country"`
	ZipCode       string `json:"zip_code"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	PhoneNumber   string `json:"phone_number"`
}

type AddressBodyParam struct {
	Code string `json:"code,omitempty" validate:"required"`
	Name string `json:"name,omitempty" validate:"required"`
}

type AddressQueryParam struct {
	Code   string `json:"code,omitempty" form:"code"`
	Active string `json:"active" form:"active"`
}
