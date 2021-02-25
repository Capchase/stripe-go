package stripe

// RadarValueListItemParams is the set of parameters that can be used when creating a value list item.
type RadarValueListItemParams struct {
	Params         `form:"*"`
	Value          *string `form:"value"`
	RadarValueList *string `form:"value_list"`
}

// RadarValueListItemListParams is the set of parameters that can be used when listing value list items.
type RadarValueListItemListParams struct {
	ListParams     `form:"*"`
	Created        *int64            `form:"created"`
	CreatedRange   *RangeQueryParams `form:"created"`
	RadarValueList *string           `form:"value_list"`
	Value          *string           `form:"value"`
}

// RadarValueListItem is the resource representing a value list item.
type RadarValueListItem struct {
	APIResource
	Created        *int64 `json:"created,omitempty"`
	CreatedBy      *string `json:"created_by,omitempty"`
	Deleted        *bool `json:"deleted,omitempty"`
	ID             *string `json:"id,omitempty"`
	Livemode       *bool `json:"livemode,omitempty"`
	Name           *string `json:"name,omitempty"`
	Object         *string `json:"object,omitempty"`
	Value          *string `json:"value,omitempty"`
	RadarValueList *string `json:"value_list,omitempty"`
}

// RadarValueListItemList is a list of value list items as retrieved from a list endpoint.
type RadarValueListItemList struct {
	APIResource
	ListMeta
	Data []*RadarValueListItem `json:"data,omitempty"`
}
