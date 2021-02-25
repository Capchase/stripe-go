package stripe

// Country is the list of supported countries
type Country string

// VerificationFieldsList lists the fields needed for an account verification.
// For more details see https://stripe.com/docs/api#country_spec_object-verification_fields.
type VerificationFieldsList struct {
	AdditionalFields []string `json:"additional,omitempty"`
	Minimum          []string `json:"minimum,omitempty"`
}

// CountrySpec is the resource representing the rules required for a Stripe account.
// For more details see https://stripe.com/docs/api/#country_specs.
type CountrySpec struct {
	APIResource
	DefaultCurrency                Currency                                        `json:"default_currency,omitempty"`
	ID                             *string `json:"id,omitempty"`
	SupportedBankAccountCurrencies map[Currency][]Country                          `json:"supported_bank_account_currencies,omitempty"`
	SupportedPaymentCurrencies     []Currency                                      `json:"supported_payment_currencies,omitempty"`
	SupportedPaymentMethods        []string                                        `json:"supported_payment_methods,omitempty"`
	SupportedTransferCountries     []string                                        `json:"supported_transfer_countries,omitempty"`
	VerificationFields             map[AccountBusinessType]*VerificationFieldsList `json:"verification_fields,omitempty"`
}

// CountrySpecParams are the parameters allowed during CountrySpec retrieval.
type CountrySpecParams struct {
	Params `form:"*"`
}

// CountrySpecList is a list of country specs as retrieved from a list endpoint.
type CountrySpecList struct {
	APIResource
	ListMeta
	Data []*CountrySpec `json:"data,omitempty"`
}

// CountrySpecListParams are the parameters allowed during CountrySpec listing.
type CountrySpecListParams struct {
	ListParams `form:"*"`
}
