package stripe

// ReportTypeListParams is the set of parameters that can be used when listing report types.
type ReportTypeListParams struct {
	ListParams `form:"*"`
}

// ReportTypeParams is the set of parameters that can be used when retrieving a report type.
type ReportTypeParams struct {
	Params `form:"*"`
}

// ReportType is the resource representing a report type.
type ReportType struct {
	APIResource
	DefaultColumns     []string `json:"default_columns,omitempty"`
	Created            *int64 `json:"created,omitempty"`
	DataAvailableEnd   *int64 `json:"data_available_end,omitempty"`
	DataAvailableStart *int64 `json:"data_available_start,omitempty"`
	ID                 *string `json:"id,omitempty"`
	Name               *string `json:"name,omitempty"`
	Object             *string `json:"object,omitempty"`
	Updated            *int64 `json:"updated,omitempty"`
	Version            *int64 `json:"version,omitempty"`
}

// ReportTypeList is a list of report types as retrieved from a list endpoint.
type ReportTypeList struct {
	APIResource
	ListMeta
	Data []*ReportType `json:"data,omitempty"`
}
