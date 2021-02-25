package stripe

// UsageRecordSummary represents a usage record summary.
// See https://stripe.com/docs/api#usage_records
type UsageRecordSummary struct {
	ID               *string `json:"id,omitempty"`
	Invoice          *string `json:"invoice,omitempty"`
	Livemode         *bool `json:"livemode,omitempty"`
	Object           *string `json:"object,omitempty"`
	Period           *Period `json:"period,omitempty"`
	SubscriptionItem *string `json:"subscription_item,omitempty"`
	TotalUsage       *int64 `json:"total_usage,omitempty"`
}

// UsageRecordSummaryListParams is the set of parameters that can be used when listing charges.
type UsageRecordSummaryListParams struct {
	ListParams       `form:"*"`
	SubscriptionItem *string `form:"-"` // Sent in with the URL
}

// UsageRecordSummaryList is a list of usage record summaries as retrieved from a list endpoint.
type UsageRecordSummaryList struct {
	APIResource
	ListMeta
	Data []*UsageRecordSummary `json:"data,omitempty"`
}
