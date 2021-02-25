package stripe

import "encoding/json"

// DiscountParams is the set of parameters that can be used when deleting a discount.
type DiscountParams struct {
	Params `form:"*"`
}

// Discount is the resource representing a Stripe discount.
// For more details see https://stripe.com/docs/api#discounts.
type Discount struct {
	APIResource
	CheckoutSession *CheckoutSession `json:"checkout_session,omitempty"`
	Coupon          *Coupon          `json:"coupon,omitempty"`
	Customer        *string `json:"customer,omitempty"`
	Deleted         *bool `json:"deleted,omitempty"`
	End             *int64 `json:"end,omitempty"`
	ID              *string `json:"id,omitempty"`
	Invoice         *string `json:"invoice,omitempty"`
	InvoiceItem     *string `json:"invoice_item,omitempty"`
	Object          *string `json:"object,omitempty"`
	PromotionCode   *PromotionCode   `json:"promotion_code,omitempty"`
	Start           *int64 `json:"start,omitempty"`
	Subscription    *string `json:"subscription,omitempty"`
}

// UnmarshalJSON handles deserialization of a Discount.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *Discount) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = &id
		return nil
	}

	type discount Discount
	var v discount
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = Discount(v)
	return nil
}
