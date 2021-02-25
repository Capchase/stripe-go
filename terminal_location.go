package stripe

// TerminalLocationParams is the set of parameters that can be used when creating or updating a terminal location.
type TerminalLocationParams struct {
	Params      `form:"*"`
	Address     *AccountAddressParams `form:"address"`
	DisplayName *string               `form:"display_name"`
}

// TerminalLocationListParams is the set of parameters that can be used when listing temrinal locations.
type TerminalLocationListParams struct {
	ListParams `form:"*"`
}

// TerminalLocation is the resource representing a Stripe terminal location.
type TerminalLocation struct {
	APIResource
	Address     *AccountAddressParams `json:"address,omitempty"`
	Deleted     *bool `json:"deleted,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	ID          *string `json:"id,omitempty"`
	Livemode    *bool `json:"livemode,omitempty"`
	Metadata    map[string]string     `json:"metadata,omitempty"`
	Object      *string `json:"object,omitempty"`
}

// TerminalLocationList is a list of terminal readers as retrieved from a list endpoint.
type TerminalLocationList struct {
	APIResource
	ListMeta
	Data []*TerminalLocation `json:"data,omitempty"`
}
