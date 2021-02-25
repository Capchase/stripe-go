package stripe

import "encoding/json"

// IssuingCardholderRequirementsDisabledReason is the possible values for the disabled reason on an
// issuing cardholder.
type IssuingCardholderRequirementsDisabledReason string

// List of values that IssuingCardholderRequirementsDisabledReason can take.
const (
	IssuingCardholderRequirementsDisabledReasonListed         IssuingCardholderRequirementsDisabledReason = "listed"
	IssuingCardholderRequirementsDisabledReasonRejectedListed IssuingCardholderRequirementsDisabledReason = "rejected.listed"
	IssuingCardholderRequirementsDisabledReasonUnderReview    IssuingCardholderRequirementsDisabledReason = "under_review"
)

// IssuingCardholderSpendingControlsSpendingLimitInterval is the list of possible values for the interval
// for a spending limit on an issuing cardholder.
type IssuingCardholderSpendingControlsSpendingLimitInterval string

// List of values that IssuingCardShippingStatus can take.
const (
	IssuingCardholderSpendingControlsSpendingLimitIntervalAllTime          IssuingCardholderSpendingControlsSpendingLimitInterval = "all_time"
	IssuingCardholderSpendingControlsSpendingLimitIntervalDaily            IssuingCardholderSpendingControlsSpendingLimitInterval = "daily"
	IssuingCardholderSpendingControlsSpendingLimitIntervalMonthly          IssuingCardholderSpendingControlsSpendingLimitInterval = "monthly"
	IssuingCardholderSpendingControlsSpendingLimitIntervalPerAuthorization IssuingCardholderSpendingControlsSpendingLimitInterval = "per_authorization"
	IssuingCardholderSpendingControlsSpendingLimitIntervalWeekly           IssuingCardholderSpendingControlsSpendingLimitInterval = "weekly"
	IssuingCardholderSpendingControlsSpendingLimitIntervalYearly           IssuingCardholderSpendingControlsSpendingLimitInterval = "yearly"
)

// IssuingCardholderStatus is the possible values for status on an issuing cardholder.
type IssuingCardholderStatus string

// List of values that IssuingCardholderStatus can take.
const (
	IssuingCardholderStatusActive   IssuingCardholderStatus = "active"
	IssuingCardholderStatusInactive IssuingCardholderStatus = "inactive"
)

// IssuingCardholderType is the type of an issuing cardholder.
type IssuingCardholderType string

// List of values that IssuingCardholderType can take.
const (
	IssuingCardholderTypeCompany    IssuingCardholderType = "company"
	IssuingCardholderTypeIndividual IssuingCardholderType = "individual"
)

// IssuingCardholderBillingParams is the set of parameters that can be used for billing with the Issuing APIs.
type IssuingCardholderBillingParams struct {
	Address *AddressParams `form:"address"`
}

// IssuingCardholderCompanyParams represents additional information about a company cardholder.
type IssuingCardholderCompanyParams struct {
	TaxID *string `form:"tax_id"`
}

// IssuingCardholderIndividualDOBParams represents the date of birth of the
// cardholder individual.
type IssuingCardholderIndividualDOBParams struct {
	Day   *int64 `form:"day"`
	Month *int64 `form:"month"`
	Year  *int64 `form:"year"`
}

// IssuingCardholderIndividualVerificationDocumentParams represents an
// identifying document, either a passport or local ID card.
type IssuingCardholderIndividualVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// IssuingCardholderIndividualVerificationParams represents government-issued ID
// document for this cardholder.
type IssuingCardholderIndividualVerificationParams struct {
	Document *IssuingCardholderIndividualVerificationDocumentParams `form:"document"`
}

// IssuingCardholderIndividualParams represents additional information about an
// `individual` cardholder.
type IssuingCardholderIndividualParams struct {
	DOB          *IssuingCardholderIndividualDOBParams          `form:"dob"`
	FirstName    *string                                        `form:"first_name"`
	LastName     *string                                        `form:"last_name"`
	Verification *IssuingCardholderIndividualVerificationParams `form:"verification"`
}

// IssuingCardholderSpendingControlsSpendingLimitParams is the set of parameters that can be used to
// represent a given spending limit for an issuing cardholder.
type IssuingCardholderSpendingControlsSpendingLimitParams struct {
	Amount     *int64    `form:"amount"`
	Categories []*string `form:"categories"`
	Interval   *string   `form:"interval"`
}

// IssuingCardholderSpendingControlsParams is the set of parameters that can be used to configure
// the spending controls for an issuing cardholder
type IssuingCardholderSpendingControlsParams struct {
	AllowedCategories      []*string                                               `form:"allowed_categories"`
	BlockedCategories      []*string                                               `form:"blocked_categories"`
	SpendingLimits         []*IssuingCardholderSpendingControlsSpendingLimitParams `form:"spending_limits"`
	SpendingLimitsCurrency *string                                                 `form:"spending_limits_currency"`
}

// IssuingCardholderParams is the set of parameters that can be used when creating or updating an issuing cardholder.
type IssuingCardholderParams struct {
	Params           `form:"*"`
	Billing          *IssuingCardholderBillingParams          `form:"billing"`
	Company          *IssuingCardholderCompanyParams          `form:"company"`
	Email            *string                                  `form:"email"`
	Individual       *IssuingCardholderIndividualParams       `form:"individual"`
	Name             *string                                  `form:"name"`
	PhoneNumber      *string                                  `form:"phone_number"`
	SpendingControls *IssuingCardholderSpendingControlsParams `form:"spending_controls"`
	Status           *string                                  `form:"status"`
	Type             *string                                  `form:"type"`
}

// IssuingCardholderListParams is the set of parameters that can be used when listing issuing cardholders.
type IssuingCardholderListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Email        *string           `form:"email"`
	PhoneNumber  *string           `form:"phone_number"`
	Status       *string           `form:"status"`
	Type         *string           `form:"type"`
}

// IssuingCardholderBilling is the resource representing the billing hash with the Issuing APIs.
type IssuingCardholderBilling struct {
	Address *Address `json:"address,omitempty"`
}

// IssuingCardholderCompany represents additional information about a company cardholder.
type IssuingCardholderCompany struct {
	TaxIDProvided *bool `json:"tax_id_provided,omitempty"`
}

// IssuingCardholderIndividualDOB represents the date of birth of the issuing card hoder
// individual.
type IssuingCardholderIndividualDOB struct {
	Day   *int64 `json:"day,omitempty"`
	Month *int64 `json:"month,omitempty"`
	Year  *int64 `json:"year,omitempty"`
}

// IssuingCardholderIndividualVerificationDocument represents an identifying
// document, either a passport or local ID card.
type IssuingCardholderIndividualVerificationDocument struct {
	Back  *File `json:"back,omitempty"`
	Front *File `json:"front,omitempty"`
}

// IssuingCardholderIndividualVerification represents the Government-issued ID
// document for this cardholder
type IssuingCardholderIndividualVerification struct {
	Document *IssuingCardholderIndividualVerificationDocument `json:"document,omitempty"`
}

// IssuingCardholderIndividual represents additional information about an
// individual cardholder.
type IssuingCardholderIndividual struct {
	DOB          *IssuingCardholderIndividualDOB          `json:"dob,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
	Verification *IssuingCardholderIndividualVerification `json:"verification,omitempty"`
}

// IssuingCardholderRequirements contains the verification requirements for the cardholder.
type IssuingCardholderRequirements struct {
	DisabledReason IssuingCardholderRequirementsDisabledReason `json:"disabled_reason,omitempty"`
	PastDue        []string                                    `json:"past_due,omitempty"`
}

// IssuingCardholderSpendingControlsSpendingLimit is the resource representing a spending limit
// for an issuing cardholder.
type IssuingCardholderSpendingControlsSpendingLimit struct {
	Amount     *int64 `json:"amount,omitempty"`
	Categories []string                                               `json:"categories,omitempty"`
	Interval   IssuingCardholderSpendingControlsSpendingLimitInterval `json:"interval,omitempty"`
}

// IssuingCardholderSpendingControls is the resource representing spending controls
// for an issuing cardholder.
type IssuingCardholderSpendingControls struct {
	AllowedCategories      []string                                          `json:"allowed_categories,omitempty"`
	BlockedCategories      []string                                          `json:"blocked_categories,omitempty"`
	SpendingLimits         []*IssuingCardholderSpendingControlsSpendingLimit `json:"spending_limits,omitempty"`
	SpendingLimitsCurrency Currency                                          `json:"spending_limits_currency,omitempty"`
}

// IssuingCardholder is the resource representing a Stripe issuing cardholder.
type IssuingCardholder struct {
	APIResource
	Billing          *IssuingCardholderBilling          `json:"billing,omitempty"`
	Company          *IssuingCardholderCompany          `json:"company,omitempty"`
	Created          *int64 `json:"created,omitempty"`
	Email            *string `json:"email,omitempty"`
	ID               *string `json:"id,omitempty"`
	Individual       *IssuingCardholderIndividual       `json:"individual,omitempty"`
	Livemode         *bool `json:"livemode,omitempty"`
	Metadata         map[string]string                  `json:"metadata,omitempty"`
	Name             *string `json:"name,omitempty"`
	Object           *string `json:"object,omitempty"`
	PhoneNumber      *string `json:"phone_number,omitempty"`
	Requirements     *IssuingCardholderRequirements     `json:"requirements,omitempty"`
	SpendingControls *IssuingCardholderSpendingControls `json:"spending_controls,omitempty"`
	Status           IssuingCardholderStatus            `json:"status,omitempty"`
	Type             IssuingCardholderType              `json:"type,omitempty"`
}

// IssuingCardholderList is a list of issuing cardholders as retrieved from a list endpoint.
type IssuingCardholderList struct {
	APIResource
	ListMeta
	Data []*IssuingCardholder `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of an IssuingCardholder.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingCardholder) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = &id
		return nil
	}

	type issuingCardholder IssuingCardholder
	var v issuingCardholder
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingCardholder(v)
	return nil
}
