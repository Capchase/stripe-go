package stripe

import "encoding/json"

// OrderReturnParams is the set of parameters that can be used when returning orders.
type OrderReturnParams struct {
	Params `form:"*"`
	Items  []*OrderItemParams `form:"items"`
	Order  *string            `form:"-"` // Included in the URL
}

// OrderReturn is the resource representing an order return.
// For more details see https://stripe.com/docs/api#order_returns.
type OrderReturn struct {
	APIResource
	Amount   *int64 `json:"amount,omitempty"`
	Created  *int64 `json:"created,omitempty"`
	Currency Currency     `json:"currency,omitempty"`
	ID       *string `json:"id,omitempty"`
	Items    []*OrderItem `json:"items,omitempty"`
	Livemode *bool `json:"livemode,omitempty"`
	Order    *Order       `json:"order,omitempty"`
	Refund   *Refund      `json:"refund,omitempty"`
}

// OrderReturnList is a list of order returns as retrieved from a list endpoint.
type OrderReturnList struct {
	APIResource
	ListMeta
	Data []*OrderReturn `json:"data,omitempty"`
}

// OrderReturnListParams is the set of parameters that can be used when listing order returns.
type OrderReturnListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Order        *string           `form:"order"`
}

// UnmarshalJSON handles deserialization of an OrderReturn.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *OrderReturn) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		r.ID = &id
		return nil
	}

	type orderReturn OrderReturn
	var v orderReturn
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*r = OrderReturn(v)
	return nil
}
