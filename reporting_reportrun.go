package stripe

// ReportRunStatus is the possible values for status on a report run.
type ReportRunStatus string

// List of values that ReportRunStatus can take.
const (
	ReportRunStatusFailed    ReportRunStatus = "failed"
	ReportRunStatusPending   ReportRunStatus = "pending"
	ReportRunStatusSucceeded ReportRunStatus = "succeeded"
)

// ReportRunParametersParams is the set of parameters that can be used when creating a report run.
type ReportRunParametersParams struct {
	Columns           []*string `form:"columns"`
	ConnectedAccount  *string   `form:"connected_account"`
	Currency          *string   `form:"currency"`
	IntervalEnd       *int64    `form:"interval_end"`
	IntervalStart     *int64    `form:"interval_start"`
	Payout            *string   `form:"payout"`
	ReportingCategory *string   `form:"reporting_category"`
	Timezone          *string   `form:"timezone"`
}

// ReportRunParams is the set of parameters that can be used when creating a report run.
type ReportRunParams struct {
	Params     `form:"*"`
	Parameters *ReportRunParametersParams `form:"parameters"`
	ReportType *string                    `form:"report_type"`
}

// ReportRunListParams is the set of parameters that can be used when listing report runs.
type ReportRunListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// ReportRunParameters describes the parameters hash on a report run.
type ReportRunParameters struct {
	Columns           []string `json:"columns,omitempty"`
	ConnectedAccount  *string `json:"connected_account,omitempty"`
	Currency          Currency `json:"currency,omitempty"`
	IntervalEnd       *int64 `json:"interval_end,omitempty"`
	IntervalStart     *int64 `json:"interval_start,omitempty"`
	Payout            *string `json:"payout,omitempty"`
	ReportingCategory *string `json:"reporting_category,omitempty"`
	Timezone          *string `json:"timezone,omitempty"`
}

// ReportRun is the resource representing a report run.
type ReportRun struct {
	APIResource
	Created     *int64 `json:"created,omitempty"`
	Error       *string `json:"error,omitempty"`
	ID          *string `json:"id,omitempty"`
	Livemode    *bool `json:"livemode,omitempty"`
	Object      *string `json:"object,omitempty"`
	Parameters  *ReportRunParameters `json:"parameters,omitempty"`
	ReportType  *string `json:"report_type,omitempty"`
	Result      *File                `json:"result,omitempty"`
	Status      ReportRunStatus      `json:"status,omitempty"`
	SucceededAt *int64 `json:"succeeded_at,omitempty"`
}

// ReportRunList is a list of report runs as retrieved from a list endpoint.
type ReportRunList struct {
	APIResource
	ListMeta
	Data []*ReportRun `json:"data,omitempty"`
}
