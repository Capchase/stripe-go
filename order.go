package stripe

import (
	"encoding/json"
)

// OrderStatus represents the statuses of an order object.
type OrderStatus string

// List of values that OrderStatus can take.
const (
	OrderStatusCanceled  OrderStatus = "canceled"
	OrderStatusCreated   OrderStatus = "created"
	OrderStatusFulfilled OrderStatus = "fulfilled"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusReturned  OrderStatus = "returned"
)

// OrderDeliveryEstimateType represents the type of delivery estimate for shipping methods
type OrderDeliveryEstimateType string

// List of values that OrderDeliveryEstimateType can take.
const (
	OrderDeliveryEstimateTypeExact OrderDeliveryEstimateType = "exact"
	OrderDeliveryEstimateTypeRange OrderDeliveryEstimateType = "range"
)

// OrderItemType represents the type of order item
type OrderItemType string

// List of values that OrderItemType can take.
const (
	OrderItemTypeCoupon   OrderItemType = "coupon"
	OrderItemTypeDiscount OrderItemType = "discount"
	OrderItemTypeSKU      OrderItemType = "sku"
	OrderItemTypeShipping OrderItemType = "shipping"
	OrderItemTypeTax      OrderItemType = "tax"
)

// OrderItemParentType represents the type of order item parent
type OrderItemParentType string

// List of values that OrderItemParentType can take.
const (
	OrderItemParentTypeCoupon   OrderItemParentType = "coupon"
	OrderItemParentTypeDiscount OrderItemParentType = "discount"
	OrderItemParentTypeSKU      OrderItemParentType = "sku"
	OrderItemParentTypeShipping OrderItemParentType = "shipping"
	OrderItemParentTypeTax      OrderItemParentType = "tax"
)

// OrderItemParent describes the parent of an order item.
type OrderItemParent struct {
	ID   *string `json:"id,omitempty"`
	SKU  *SKU                `json:"-,omitempty"`
	Type OrderItemParentType `json:"object,omitempty"`
}

// OrderParams is the set of parameters that can be used when creating an order.
type OrderParams struct {
	Params   `form:"*"`
	Coupon   *string            `form:"coupon"`
	Currency *string            `form:"currency"`
	Customer *string            `form:"customer"`
	Email    *string            `form:"email"`
	Items    []*OrderItemParams `form:"items"`
	Shipping *ShippingParams    `form:"shipping"`
}

// ShippingParams is the set of parameters that can be used for the shipping hash
// on order creation.
type ShippingParams struct {
	Address *AddressParams `form:"address"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// OrderUpdateParams is the set of parameters that can be used when updating an order.
type OrderUpdateParams struct {
	Params                 `form:"*"`
	Coupon                 *string                    `form:"coupon"`
	SelectedShippingMethod *string                    `form:"selected_shipping_method"`
	Shipping               *OrderUpdateShippingParams `form:"shipping"`
	Status                 *string                    `form:"status"`
}

// OrderUpdateShippingParams is the set of parameters that can be used for the shipping
// hash on order update.
type OrderUpdateShippingParams struct {
	Carrier        *string `form:"carrier"`
	TrackingNumber *string `form:"tracking_number"`
}

// Shipping describes the shipping hash on an order.
type Shipping struct {
	Address        *Address `json:"address,omitempty"`
	Carrier        *string `json:"carrier,omitempty"`
	Name           *string `json:"name,omitempty"`
	Phone          *string `json:"phone,omitempty"`
	TrackingNumber *string `json:"tracking_number,omitempty"`
}

// ShippingMethod describes a shipping method as available on an order.
type ShippingMethod struct {
	Amount           *int64 `json:"amount,omitempty"`
	ID               *string `json:"id,omitempty"`
	Currency         Currency          `json:"currency,omitempty"`
	DeliveryEstimate *DeliveryEstimate `json:"delivery_estimate,omitempty"`
	Description      *string `json:"description,omitempty"`
}

// DeliveryEstimate represent the properties available for a shipping method's
// estimated delivery.
type DeliveryEstimate struct {
	// If Type == Exact
	Date *string `json:"date,omitempty"`
	// If Type == Range
	Earliest *string `json:"earliest,omitempty"`
	Latest   *string `json:"latest,omitempty"`
	Type     OrderDeliveryEstimateType `json:"type,omitempty"`
}

// Order is the resource representing a Stripe charge.
// For more details see https://stripe.com/docs/api#orders.
type Order struct {
	APIResource
	Amount                 *int64 `json:"amount,omitempty"`
	AmountReturned         *int64 `json:"amount_returned,omitempty"`
	Application            *string `json:"application,omitempty"`
	ApplicationFee         *int64 `json:"application_fee,omitempty"`
	Charge                 *Charge           `json:"charge,omitempty"`
	Created                *int64 `json:"created,omitempty"`
	Currency               Currency          `json:"currency,omitempty"`
	Customer               Customer          `json:"customer,omitempty"`
	Email                  *string `json:"email,omitempty"`
	ID                     *string `json:"id,omitempty"`
	Items                  []*OrderItem      `json:"items,omitempty"`
	Livemode               *bool `json:"livemode,omitempty"`
	Metadata               map[string]string `json:"metadata,omitempty"`
	Returns                *OrderReturnList  `json:"returns,omitempty"`
	SelectedShippingMethod *string           `json:"selected_shipping_method,omitempty"`
	Shipping               *Shipping         `json:"shipping,omitempty"`
	ShippingMethods        []*ShippingMethod `json:"shipping_methods,omitempty"`
	Status                 *string `json:"status,omitempty"`
	StatusTransitions      StatusTransitions `json:"status_transitions,omitempty"`
	Updated                *int64 `json:"updated,omitempty"`
	UpstreamID             *string `json:"upstream_id,omitempty"`
}

// OrderList is a list of orders as retrieved from a list endpoint.
type OrderList struct {
	APIResource
	ListMeta
	Data []*Order `json:"data,omitempty"`
}

// OrderListParams is the set of parameters that can be used when listing orders.
type OrderListParams struct {
	ListParams        `form:"*"`
	Created           *int64                         `form:"created"`
	CreatedRange      *RangeQueryParams              `form:"created"`
	Customer          *string                        `form:"customer"`
	IDs               []*string                      `form:"ids"`
	Status            *string                        `form:"status"`
	StatusTransitions *StatusTransitionsFilterParams `form:"status_transitions"`
	UpstreamIDs       []*string                      `form:"upstream_ids"`
}

// StatusTransitionsFilterParams are parameters that can used to filter on status_transition when listing orders.
type StatusTransitionsFilterParams struct {
	Canceled       *int64            `form:"canceled"`
	CanceledRange  *RangeQueryParams `form:"canceled"`
	Fulfilled      *int64            `form:"fulfilled"`
	FulfilledRange *RangeQueryParams `form:"fulfilled"`
	Paid           *int64            `form:"paid"`
	PaidRange      *RangeQueryParams `form:"paid"`
	Returned       *int64            `form:"returned"`
	ReturnedRange  *RangeQueryParams `form:"returned"`
}

// StatusTransitions are the timestamps at which the order status was updated.
type StatusTransitions struct {
	Canceled  *int64 `json:"canceled,omitempty"`
	Fulfilled *int64 `json:"fulfilled,omitempty"`
	Paid      *int64 `json:"paid,omitempty"`
	Returned  *int64 `json:"returned,omitempty"`
}

// OrderPayParams is the set of parameters that can be used when paying orders.
type OrderPayParams struct {
	Params         `form:"*"`
	ApplicationFee *int64        `form:"application_fee"`
	Customer       *string       `form:"customer"`
	Email          *string       `form:"email"`
	Source         *SourceParams `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
}

// OrderItemParams is the set of parameters describing an order item on order creation or update.
type OrderItemParams struct {
	Amount      *int64  `form:"amount"`
	Currency    *string `form:"currency"`
	Description *string `form:"description"`
	Parent      *string `form:"parent"`
	Quantity    *int64  `form:"quantity"`
	Type        *string `form:"type"`
}

// OrderItem is the resource representing an order item.
type OrderItem struct {
	Amount      *int64 `json:"amount,omitempty"`
	Currency    Currency         `json:"currency,omitempty"`
	Description *string `json:"description,omitempty"`
	Parent      *OrderItemParent `json:"parent,omitempty"`
	Quantity    *int64 `json:"quantity,omitempty"`
	Type        OrderItemType    `json:"type,omitempty"`
}

// SetSource adds valid sources to a OrderParams object,
// returning an error for unsupported sources.
func (op *OrderPayParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	op.Source = source
	return err
}

// UnmarshalJSON handles deserialization of an OrderItemParent.
// This custom unmarshaling is needed because the resulting
// property may be an id or a full SKU struct if it was expanded.
func (p *OrderItemParent) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = &id
		return nil
	}

	type orderItemParent OrderItemParent
	var v orderItemParent
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var err error
	*p = OrderItemParent(v)

	switch p.Type {
	case OrderItemParentTypeSKU:
		// Currently only SKUs `parent` is returned as an object when expanded.
		// For other items only IDs are returned.
		if err = json.Unmarshal(data, &p.SKU); err != nil {
			return err
		}
		p.ID = p.SKU.ID
	}

	return nil
}

// UnmarshalJSON handles deserialization of an Order.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (o *Order) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		o.ID = &id
		return nil
	}

	type order Order
	var v order
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*o = Order(v)
	return nil
}
