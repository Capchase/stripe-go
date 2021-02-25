package stripe

import "encoding/json"

// ApplicationFeeParams is the set of parameters that can be used when refunding an application fee.
// For more details see https://stripe.com/docs/api#refund_application_fee.
type ApplicationFeeParams struct {
	Params `form:"*"`
}

// ApplicationFeeListParams is the set of parameters that can be used when listing application fees.
// For more details see https://stripe.com/docs/api#list_application_fees.
type ApplicationFeeListParams struct {
	ListParams   `form:"*"`
	Charge       *string           `form:"charge"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// ApplicationFee is the resource representing a Stripe application fee.
// For more details see https://stripe.com/docs/api#application_fees.
type ApplicationFee struct {
	APIResource
	Account                *Account            `json:"account,omitempty"`
	Amount                 *int64 `json:"amount,omitempty"`
	AmountRefunded         *int64 `json:"amount_refunded,omitempty"`
	Application            *string `json:"application,omitempty"`
	BalanceTransaction     *BalanceTransaction `json:"balance_transaction,omitempty"`
	Charge                 *Charge             `json:"charge,omitempty"`
	Created                *int64 `json:"created,omitempty"`
	Currency               Currency            `json:"currency,omitempty"`
	ID                     *string `json:"id,omitempty"`
	Livemode               *bool `json:"livemode,omitempty"`
	OriginatingTransaction *Charge             `json:"originating_transaction,omitempty"`
	Refunded               *bool `json:"refunded,omitempty"`
	Refunds                *FeeRefundList      `json:"refunds,omitempty"`
}

//ApplicationFeeList is a list of application fees as retrieved from a list endpoint.
type ApplicationFeeList struct {
	APIResource
	ListMeta
	Data []*ApplicationFee `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of an ApplicationFee.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (f *ApplicationFee) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		f.ID = &id
		return nil
	}

	type applicationFee ApplicationFee
	var v applicationFee
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*f = ApplicationFee(v)
	return nil
}
