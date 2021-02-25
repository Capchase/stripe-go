package stripe

// AddressParams describes the common parameters for an Address.
type AddressParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
}

// Address describes common properties for an Address hash.
type Address struct {
	City       *string `json:"city,omitempty"`
	Country    *string `json:"country,omitempty"`
	Line1      *string `json:"line1,omitempty"`
	Line2      *string `json:"line2,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	State      *string `json:"state,omitempty"`
}
