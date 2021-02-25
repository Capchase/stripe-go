package stripe

import "encoding/json"

// CapabilityDisabledReason describes why a capability is disabled.
type CapabilityDisabledReason string

// List of values that CapabilityDisabledReason can take.
const (
	CapabilityDisabledReasonPendingOnboarding        CapabilityDisabledReason = "pending.onboarding"
	CapabilityDisabledReasonPendingReview            CapabilityDisabledReason = "pending.review"
	CapabilityDisabledReasonRejectedFraud            CapabilityDisabledReason = "rejected_fraud"
	CapabilityDisabledReasonRejectedListed           CapabilityDisabledReason = "rejected.listed"
	CapabilityDisabledReasonRejectedOther            CapabilityDisabledReason = "rejected.other"
	CapabilityDisabledReasonRequirementsFieldsNeeded CapabilityDisabledReason = "requirement.fields_needed"
)

// CapabilityStatus describes the different statuses for a capability's status.
type CapabilityStatus string

// List of values that CapabilityStatus can take.
const (
	CapabilityStatusActive      CapabilityStatus = "active"
	CapabilityStatusInactive    CapabilityStatus = "inactive"
	CapabilityStatusPending     CapabilityStatus = "pending"
	CapabilityStatusUnrequested CapabilityStatus = "unrequested"
)

// CapabilityParams is the set of parameters that can be used when updating a capability.
// For more details see https://stripe.com/docs/api/capabilities/update
type CapabilityParams struct {
	Params    `form:"*"`
	Account   *string `form:"-"` // Included in URL
	Requested *bool   `form:"requested"`
}

// CapabilityListParams is the set of parameters that can be used when listing capabilities.
// For more detail see https://stripe.com/docs/api/capabilities/list
type CapabilityListParams struct {
	ListParams `form:"*"`
	Account    *string `form:"-"` // Included in URL
}

// CapabilityRequirements represents information that needs to be collected for a capability.
type CapabilityRequirements struct {
	CurrentDeadline     *int64 `json:"current_deadline,omitempty"`
	CurrentlyDue        []string                    `json:"currently_due,omitempty"`
	DisabledReason      CapabilityDisabledReason    `json:"disabled_reason,omitempty"`
	Errors              []*AccountRequirementsError `json:"errors,omitempty"`
	EventuallyDue       []string                    `json:"eventually_due,omitempty"`
	PastDue             []string                    `json:"past_due,omitempty"`
	PendingVerification []string                    `json:"pending_verification,omitempty"`
}

// Capability is the resource representing a Stripe capability.
// For more details see https://stripe.com/docs/api/capabilities
type Capability struct {
	APIResource
	Account      *Account                `json:"account,omitempty"`
	ID           *string `json:"id,omitempty"`
	Object       *string `json:"object,omitempty"`
	Requested    *bool `json:"requested,omitempty"`
	RequestedAt  *int64 `json:"requested_at,omitempty"`
	Requirements *CapabilityRequirements `json:"requirements,omitempty"`
	Status       CapabilityStatus        `json:"status,omitempty"`
}

// CapabilityList is a list of capabilities as retrieved from a list endpoint.
type CapabilityList struct {
	APIResource
	ListMeta
	Data []*Capability `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of a Capability.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Capability) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = &id
		return nil
	}

	type capability Capability
	var v capability
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Capability(v)
	return nil
}
