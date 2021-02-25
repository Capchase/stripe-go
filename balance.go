package stripe

// BalanceSourceType is the list of allowed values for the balance amount's source_type field keys.
type BalanceSourceType string

// List of values that BalanceSourceType can take.
const (
	BalanceSourceTypeBankAccount BalanceSourceType = "bank_account"
	BalanceSourceTypeCard        BalanceSourceType = "card"
	BalanceSourceTypeFPX         BalanceSourceType = "fpx"
)

// BalanceTransactionStatus is the list of allowed values for the balance transaction's status.
type BalanceTransactionStatus string

// BalanceParams is the set of parameters that can be used when retrieving a balance.
// For more details see https://stripe.com/docs/api#balance.
type BalanceParams struct {
	Params `form:"*"`
}

// Amount is a structure wrapping an amount value and its currency.
type Amount struct {
	Currency    Currency                    `json:"currency,omitempty"`
	SourceTypes map[BalanceSourceType]int64 `json:"source_types,omitempty"`
	Value       *int64 `json:"amount,omitempty"`
}

// BalanceDetails is the resource representing details about a specific product's balance.
type BalanceDetails struct {
	Available []*Amount `json:"available,omitempty"`
}

// Balance is the resource representing your Stripe balance.
// For more details see https://stripe.com/docs/api/#balance.
type Balance struct {
	APIResource
	Available        []*Amount       `json:"available,omitempty"`
	ConnectReserved  []*Amount       `json:"connect_reserved,omitempty"`
	InstantAvailable []*Amount       `json:"instant_available,omitempty"`
	Issuing          *BalanceDetails `json:"issuing,omitempty"`
	Livemode         *bool `json:"livemode,omitempty"`
	Object           *string `json:"object,omitempty"`
	Pending          []*Amount       `json:"pending,omitempty"`
}
