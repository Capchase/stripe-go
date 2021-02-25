package stripe

// Possible values for the action parameter on usage record creation.
const (
	UsageRecordActionIncrement string = "increment"
	UsageRecordActionSet       string = "set"
)

// UsageRecord represents a usage record.
// See https://stripe.com/docs/api#usage_records
type UsageRecord struct {
	APIResource
	ID               *string `json:"id,omitempty"`
	Livemode         *bool `json:"livemode,omitempty"`
	Quantity         *int64 `json:"quantity,omitempty"`
	SubscriptionItem *string `json:"subscription_item,omitempty"`
	Timestamp        *int64 `json:"timestamp,omitempty"`
}

// UsageRecordParams create a usage record for a specified subscription item
// and date, and fills it with a quantity.
type UsageRecordParams struct {
	Params           `form:"*"`
	Action           *string `form:"action"`
	Quantity         *int64  `form:"quantity"`
	SubscriptionItem *string `form:"-"` // passed in the URL
	Timestamp        *int64  `form:"timestamp"`
}
