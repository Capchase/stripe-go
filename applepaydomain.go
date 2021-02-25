package stripe

// ApplePayDomainParams is the set of parameters that can be used when creating an ApplePayDomain object.
type ApplePayDomainParams struct {
	Params     `form:"*"`
	DomainName *string `form:"domain_name"`
}

// ApplePayDomain is the resource representing a Stripe ApplePayDomain object
type ApplePayDomain struct {
	APIResource
	Created    *int64 `json:"created,omitempty"`
	Deleted    *bool `json:"deleted,omitempty"`
	DomainName *string `json:"domain_name,omitempty"`
	ID         *string `json:"id,omitempty"`
	Livemode   *bool `json:"livemode,omitempty"`
}

// ApplePayDomainListParams are the parameters allowed during ApplePayDomain listing.
type ApplePayDomainListParams struct {
	ListParams `form:"*"`
}

// ApplePayDomainList is a list of ApplePayDomains as returned from a list endpoint.
type ApplePayDomainList struct {
	APIResource
	ListMeta
	Data []*ApplePayDomain `json:"data,omitempty"`
}
