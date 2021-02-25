package stripe

// TopupParams is the set of parameters that can be used when creating or updating a top-up.
// For more details see https://stripe.com/docs/api#create_topup and https://stripe.com/docs/api#update_topup.
type TopupParams struct {
	Params              `form:"*"`
	Amount              *int64        `form:"amount"`
	Currency            *string       `form:"currency"`
	Description         *string       `form:"description"`
	Source              *SourceParams `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	StatementDescriptor *string       `form:"statement_descriptor"`
	TransferGroup       *string       `form:"transfer_group"`
}

// SetSource adds valid sources to a TopupParams object,
// returning an error for unsupported sources.
func (p *TopupParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
}

// TopupListParams is the set of parameters that can be used when listing top-ups.
// For more details see https://stripe.com/docs/api#list_topups.
type TopupListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// TopupList is a list of top-ups as retrieved from a list endpoint.
type TopupList struct {
	APIResource
	ListMeta
	Data []*Topup `json:"data,omitempty"`
}

// Topup is the resource representing a Stripe top-up.
// For more details see https://stripe.com/docs/api#topups.
type Topup struct {
	APIResource
	Amount                   *int64 `json:"amount,omitempty"`
	BalanceTransaction       *BalanceTransaction `json:"balance_transaction,omitempty"`
	Created                  *int64 `json:"created,omitempty"`
	Currency                 Currency            `json:"currency,omitempty"`
	Description              *string `json:"description,omitempty"`
	ExpectedAvailabilityDate *int64 `json:"expected_availability_date,omitempty"`
	FailureCode              *string `json:"failure_code,omitempty"`
	FailureMessage           *string `json:"failure_message,omitempty"`
	ID                       *string `json:"id,omitempty"`
	Livemode                 *bool `json:"livemode,omitempty"`
	Metadata                 map[string]string   `json:"metadata,omitempty"`
	Object                   *string `json:"object,omitempty"`
	Source                   *PaymentSource      `json:"source,omitempty"`
	StatementDescriptor      *string `json:"statement_descriptor,omitempty"`
	Status                   *string `json:"status,omitempty"`
	TransferGroup            *string `json:"transfer_group,omitempty"`

	// The following property is deprecated
	ArrivalDate *int64 `json:"arrival_date,omitempty"`
}
