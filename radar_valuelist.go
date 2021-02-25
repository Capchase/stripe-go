package stripe

// RadarValueListItemType is the possible values for a type of value list items.
type RadarValueListItemType string

// List of values that RadarValueListItemType can take.
const (
	RadarValueListItemTypeCardBin             RadarValueListItemType = "card_bin"
	RadarValueListItemTypeCardFingerprint     RadarValueListItemType = "card_fingerprint"
	RadarValueListItemTypeCountry             RadarValueListItemType = "country"
	RadarValueListItemTypeEmail               RadarValueListItemType = "email"
	RadarValueListItemTypeIPAddress           RadarValueListItemType = "ip_address"
	RadarValueListItemTypeString              RadarValueListItemType = "string"
	RadarValueListItemTypeCaseSensitiveString RadarValueListItemType = "case_sensitive_string"
)

// RadarValueListParams is the set of parameters that can be used when creating a value list.
type RadarValueListParams struct {
	Params   `form:"*"`
	Alias    *string `form:"alias"`
	ItemType *string `form:"item_type"`
	Name     *string `form:"name"`
}

// RadarValueListListParams is the set of parameters that can be used when listing value lists.
type RadarValueListListParams struct {
	ListParams   `form:"*"`
	Alias        *string           `form:"alias"`
	Contains     *string           `form:"contains"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// RadarValueList is the resource representing a value list.
type RadarValueList struct {
	APIResource
	Alias     *string `json:"alias,omitempty"`
	Created   *int64 `json:"created,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Deleted   *bool `json:"deleted,omitempty"`
	ID        *string `json:"id,omitempty"`
	ItemType  RadarValueListItemType  `json:"item_type,omitempty"`
	ListItems *RadarValueListItemList `json:"list_items,omitempty"`
	Livemode  *bool `json:"livemode,omitempty"`
	Metadata  map[string]string       `json:"metadata,omitempty"`
	Name      *string `json:"name,omitempty"`
	Object    *string `json:"object,omitempty"`
	Updated   *int64 `json:"updated,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// RadarValueListList is a list of value lists as retrieved from a list endpoint.
type RadarValueListList struct {
	APIResource
	ListMeta
	Data []*RadarValueList `json:"data,omitempty"`
}
