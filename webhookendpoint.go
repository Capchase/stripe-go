package stripe

import "encoding/json"

// WebhookEndpointParams is the set of parameters that can be used when creating a webhook endpoint.
// For more details see https://stripe.com/docs/api#create_webhook_endpoint.
type WebhookEndpointParams struct {
	Params        `form:"*"`
	Connect       *bool     `form:"connect"`
	Description   *string   `form:"description"`
	Disabled      *bool     `form:"disabled"`
	EnabledEvents []*string `form:"enabled_events"`
	URL           *string   `form:"url"`

	// This parameter is only available on creation.
	// We recommend setting the API version that the library is pinned to. See apiversion in stripe.go
	APIVersion *string `form:"api_version"`
}

// WebhookEndpointListParams is the set of parameters that can be used when listing webhook endpoints.
// For more detail see https://stripe.com/docs/api#list_webhook_endpoints.
type WebhookEndpointListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// WebhookEndpoint is the resource representing a Stripe webhook endpoint.
// For more details see https://stripe.com/docs/api#webhook_endpoints.
type WebhookEndpoint struct {
	APIResource
	APIVersion    *string `json:"api_version,omitempty"`
	Application   *string `json:"application,omitempty"`
	Connect       *bool `json:"connect,omitempty"`
	Created       *int64 `json:"created,omitempty"`
	Deleted       *bool `json:"deleted,omitempty"`
	Description   *string `json:"description,omitempty"`
	EnabledEvents []string          `json:"enabled_events,omitempty"`
	ID            *string `json:"id,omitempty"`
	Livemode      *bool `json:"livemode,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty"`
	Object        *string `json:"object,omitempty"`
	Secret        *string `json:"secret,omitempty"`
	Status        *string `json:"status,omitempty"`
	URL           *string `json:"url,omitempty"`
}

// WebhookEndpointList is a list of webhook endpoints as retrieved from a list endpoint.
type WebhookEndpointList struct {
	APIResource
	ListMeta
	Data []*WebhookEndpoint `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of a WebhookEndpoint.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *WebhookEndpoint) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = &id
		return nil
	}

	type endpoint WebhookEndpoint
	var v endpoint
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = WebhookEndpoint(v)
	return nil
}
