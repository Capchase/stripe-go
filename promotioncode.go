package stripe

import "encoding/json"

// PromotionCodeRestrictionsParams is the set of parameters for restrictions on a promotion code.
type PromotionCodeRestrictionsParams struct {
	FirstTimeTransaction  *bool   `form:"first_time_transaction"`
	MinimumAmount         *int64  `form:"minimum_amount"`
	MinimumAmountCurrency *string `form:"minimum_amount_currency"`
}

// PromotionCodeParams is the set of parameters that can be used when creating a promotion code.
type PromotionCodeParams struct {
	Params         `form:"*"`
	Active         *bool                            `form:"active"`
	Code           *string                          `form:"code"`
	Coupon         *string                          `form:"coupon"`
	Customer       *string                          `form:"customer"`
	ExpiresAt      *int64                           `form:"expires_at"`
	MaxRedemptions *int64                           `form:"max_redemptions"`
	Restrictions   *PromotionCodeRestrictionsParams `form:"restrictions"`
}

// PromotionCodeListParams is the set of parameters that can be used when listing promotion codes.
type PromotionCodeListParams struct {
	ListParams   `form:"*"`
	Active       *bool             `form:"active"`
	Code         *string           `form:"code"`
	Coupon       *string           `form:"coupon"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Customer     *string           `form:"customer"`
}

// PromotionCodeRestrictions is the set of restrictions associated with a promotion code.
type PromotionCodeRestrictions struct {
	FirstTimeTransaction  *bool `json:"first_time_transaction,omitempty"`
	MinimumAmount         *int64 `json:"minimum_amount,omitempty"`
	MinimumAmountCurrency Currency `json:"minimum_amount_currency,omitempty"`
}

// PromotionCode is the resource representing a Stripe promotion code.
type PromotionCode struct {
	APIResource
	Active         *bool `json:"active,omitempty"`
	Code           *string `json:"code,omitempty"`
	Coupon         *Coupon                    `json:"coupon,omitempty"`
	Created        *int64 `json:"created,omitempty"`
	Customer       *Customer                  `json:"customer,omitempty"`
	ExpiresAt      *int64 `json:"expires_at,omitempty"`
	ID             *string `json:"id,omitempty"`
	Livemode       *bool `json:"livemode,omitempty"`
	MaxRedemptions *int64 `json:"max_redemptions,omitempty"`
	Metadata       map[string]string          `json:"metadata,omitempty"`
	Object         *string `json:"object,omitempty"`
	Restrictions   *PromotionCodeRestrictions `json:"restrictions,omitempty"`
	TimesRedeemed  *int64 `json:"times_redeemed,omitempty"`
}

// PromotionCodeList is a list of promotion codes as retrieved from a list endpoint.
type PromotionCodeList struct {
	APIResource
	ListMeta
	Data []*PromotionCode `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of a PromotionCode.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *PromotionCode) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = &id
		return nil
	}

	type pc PromotionCode
	var v pc
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = PromotionCode(v)
	return nil
}
