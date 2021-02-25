package stripe

import "encoding/json"

// ReversalParams is the set of parameters that can be used when reversing a transfer.
type ReversalParams struct {
	Params               `form:"*"`
	Amount               *int64  `form:"amount"`
	Description          *string `form:"description"`
	RefundApplicationFee *bool   `form:"refund_application_fee"`
	Transfer             *string `form:"-"` // Included in URL
}

// ReversalListParams is the set of parameters that can be used when listing reversals.
type ReversalListParams struct {
	ListParams `form:"*"`
	Transfer   *string `form:"-"` // Included in URL
}

// Reversal represents a transfer reversal.
type Reversal struct {
	APIResource
	Amount                   *int64 `json:"amount,omitempty"`
	BalanceTransaction       *BalanceTransaction `json:"balance_transaction,omitempty"`
	Created                  *int64 `json:"created,omitempty"`
	Currency                 Currency            `json:"currency,omitempty"`
	Description              *string `json:"description,omitempty"`
	DestinationPaymentRefund *Refund             `json:"destination_payment_refund,omitempty"`
	ID                       *string `json:"id,omitempty"`
	Metadata                 map[string]string   `json:"metadata,omitempty"`
	SourceRefund             *Refund             `json:"source_refund,omitempty"`
	Transfer                 *string `json:"transfer,omitempty"`
}

// ReversalList is a list of object for reversals.
type ReversalList struct {
	APIResource
	ListMeta
	Data []*Reversal `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of a Reversal.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *Reversal) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		r.ID = &id
		return nil
	}

	type reversal Reversal
	var v reversal
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*r = Reversal(v)
	return nil
}
