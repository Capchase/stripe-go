package stripe

// TerminalReaderParams is the set of parameters that can be used for creating or updating a terminal reader.
type TerminalReaderParams struct {
	Params           `form:"*"`
	Label            *string `form:"label"`
	Location         *string `form:"location"`
	RegistrationCode *string `form:"registration_code"`
}

// TerminalReaderGetParams is the set of parameters that can be used to get a terminal reader.
type TerminalReaderGetParams struct {
	Params `form:"*"`
}

// TerminalReaderListParams is the set of parameters that can be used when listing temrinal readers.
type TerminalReaderListParams struct {
	ListParams `form:"*"`
	DeviceType *string `form:"device_type"`
	Location   *string `form:"location"`
	Status     *string `form:"status"`
}

// TerminalReader is the resource representing a Stripe terminal reader.
type TerminalReader struct {
	APIResource
	Deleted         *bool `json:"deleted,omitempty"`
	DeviceSwVersion *string `json:"device_sw_version,omitempty"`
	DeviceType      *string `json:"device_type,omitempty"`
	ID              *string `json:"id,omitempty"`
	IPAddress       *string `json:"ip_address,omitempty"`
	Label           *string `json:"label,omitempty"`
	Livemode        *bool `json:"livemode,omitempty"`
	Location        *string `json:"location,omitempty"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	Object          *string `json:"object,omitempty"`
	SerialNumber    *string `json:"serial_number,omitempty"`
	Status          *string `json:"status,omitempty"`
}

// TerminalReaderList is a list of terminal readers as retrieved from a list endpoint.
type TerminalReaderList struct {
	APIResource
	ListMeta
	Data     []*TerminalReader `json:"data,omitempty"`
	Location *string           `json:"location,omitempty"`
	Status   *string           `json:"status,omitempty"`
}
