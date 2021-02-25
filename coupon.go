package stripe

import "encoding/json"

// CouponDuration is the list of allowed values for the coupon's duration.
type CouponDuration string

// List of values that CouponDuration can take.
const (
	CouponDurationForever   CouponDuration = "forever"
	CouponDurationOnce      CouponDuration = "once"
	CouponDurationRepeating CouponDuration = "repeating"
)

// CouponAppliesToParams controls the products that a coupon applies to.
type CouponAppliesToParams struct {
	Products []*string `form:"products"`
}

// CouponParams is the set of parameters that can be used when creating a coupon.
// For more details see https://stripe.com/docs/api#create_coupon.
type CouponParams struct {
	Params           `form:"*"`
	AmountOff        *int64                 `form:"amount_off"`
	AppliesTo        *CouponAppliesToParams `form:"applies_to"`
	Currency         *string                `form:"currency"`
	Duration         *string                `form:"duration"`
	DurationInMonths *int64                 `form:"duration_in_months"`
	ID               *string                `form:"id"`
	MaxRedemptions   *int64                 `form:"max_redemptions"`
	Name             *string                `form:"name"`
	PercentOff       *float64               `form:"percent_off"`
	RedeemBy         *int64                 `form:"redeem_by"`
}

// CouponListParams is the set of parameters that can be used when listing coupons.
// For more detail see https://stripe.com/docs/api#list_coupons.
type CouponListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// CouponAppliesTo represents the products a coupon applies to.
type CouponAppliesTo struct {
	Products []string `json:"products,omitempty"`
}

// Coupon is the resource representing a Stripe coupon.
// For more details see https://stripe.com/docs/api#coupons.
type Coupon struct {
	APIResource
	AmountOff        *int64 `json:"amount_off,omitempty"`
	AppliesTo        *CouponAppliesTo  `json:"applies_to,omitempty"`
	Created          *int64 `json:"created,omitempty"`
	Currency         Currency          `json:"currency,omitempty"`
	Deleted          *bool `json:"deleted,omitempty"`
	Duration         CouponDuration    `json:"duration,omitempty"`
	DurationInMonths *int64 `json:"duration_in_months,omitempty"`
	ID               *string `json:"id,omitempty"`
	Livemode         *bool `json:"livemode,omitempty"`
	MaxRedemptions   *int64 `json:"max_redemptions,omitempty"`
	Metadata         map[string]string `json:"metadata,omitempty"`
	Name             *string `json:"name,omitempty"`
	Object           *string `json:"object,omitempty"`
	PercentOff       *float64 `json:"percent_off,omitempty"`
	RedeemBy         *int64 `json:"redeem_by,omitempty"`
	TimesRedeemed    *int64 `json:"times_redeemed,omitempty"`
	Valid            *bool `json:"valid,omitempty"`
}

// CouponList is a list of coupons as retrieved from a list endpoint.
type CouponList struct {
	APIResource
	ListMeta
	Data []*Coupon `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of a Coupon.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Coupon) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = &id
		return nil
	}

	type coupon Coupon
	var v coupon
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Coupon(v)
	return nil
}
