package stripe

import (
	"encoding/json"
)

// LineItemDiscount represent the details of one discount applied to a line item.
type LineItemDiscount struct {
	Amount   *int64 `json:"amount,omitempty"`
	Discount *Discount `json:"discount,omitempty"`
}

// LineItemTax represent the details of one tax rate applied to a line item.
type LineItemTax struct {
	Amount  *int64 `json:"amount,omitempty"`
	TaxRate *TaxRate `json:"tax_rate,omitempty"`
}

// LineItem is the resource representing a line item.
type LineItem struct {
	APIResource
	AmountSubtotal *int64 `json:"amount_subtotal,omitempty"`
	AmountTotal    *int64 `json:"amount_total,omitempty"`
	Currency       Currency            `json:"currency,omitempty"`
	Description    *string `json:"description,omitempty"`
	Discounts      []*LineItemDiscount `json:"discounts,omitempty"`
	Deleted        *bool `json:"deleted,omitempty"`
	ID             *string `json:"id,omitempty"`
	Object         *string `json:"object,omitempty"`
	Price          *Price              `json:"price,omitempty"`
	Quantity       *int64 `json:"quantity,omitempty"`
	Taxes          []*LineItemTax      `json:"taxes,omitempty"`
}

// LineItemList is a list of prices as returned from a list endpoint.
type LineItemList struct {
	APIResource
	ListMeta
	Data []*LineItem `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of a LineItem.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *LineItem) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = &id
		return nil
	}

	type price LineItem
	var v price
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = LineItem(v)
	return nil
}
