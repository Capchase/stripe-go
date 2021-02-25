package stripe

import (
	"encoding/json"
)

// DisputeReason is the list of allowed values for a discount's reason.
type DisputeReason string

// List of values that DisputeReason can take.
const (
	DisputeReasonCreditNotProcessed   DisputeReason = "credit_not_processed"
	DisputeReasonDuplicate            DisputeReason = "duplicate"
	DisputeReasonFraudulent           DisputeReason = "fraudulent"
	DisputeReasonGeneral              DisputeReason = "general"
	DisputeReasonProductNotReceived   DisputeReason = "product_not_received"
	DisputeReasonProductUnacceptable  DisputeReason = "product_unacceptable"
	DisputeReasonSubscriptionCanceled DisputeReason = "subscription_canceled"
	DisputeReasonUnrecognized         DisputeReason = "unrecognized"
)

// DisputeStatus is the list of allowed values for a discount's status.
type DisputeStatus string

// List of values that DisputeStatus can take.
const (
	DisputeStatusChargeRefunded       DisputeStatus = "charge_refunded"
	DisputeStatusLost                 DisputeStatus = "lost"
	DisputeStatusNeedsResponse        DisputeStatus = "needs_response"
	DisputeStatusUnderReview          DisputeStatus = "under_review"
	DisputeStatusWarningClosed        DisputeStatus = "warning_closed"
	DisputeStatusWarningNeedsResponse DisputeStatus = "warning_needs_response"
	DisputeStatusWarningUnderReview   DisputeStatus = "warning_under_review"
	DisputeStatusWon                  DisputeStatus = "won"
)

// DisputeParams is the set of parameters that can be used when updating a dispute.
// For more details see https://stripe.com/docs/api#update_dispute.
type DisputeParams struct {
	Params   `form:"*"`
	Evidence *DisputeEvidenceParams `form:"evidence"`
	Submit   *bool                  `form:"submit"`
}

// DisputeEvidenceParams is the set of parameters that can be used when submitting
// evidence for disputes.
type DisputeEvidenceParams struct {
	AccessActivityLog            *string `form:"access_activity_log"`
	BillingAddress               *string `form:"billing_address"`
	CancellationPolicy           *string `form:"cancellation_policy"`
	CancellationPolicyDisclosure *string `form:"cancellation_policy_disclosure"`
	CancellationRebuttal         *string `form:"cancellation_rebuttal"`
	CustomerCommunication        *string `form:"customer_communication"`
	CustomerEmailAddress         *string `form:"customer_email_address"`
	CustomerName                 *string `form:"customer_name"`
	CustomerPurchaseIP           *string `form:"customer_purchase_ip"`
	CustomerSignature            *string `form:"customer_signature"`
	DuplicateChargeDocumentation *string `form:"duplicate_charge_documentation"`
	DuplicateChargeExplanation   *string `form:"duplicate_charge_explanation"`
	DuplicateChargeID            *string `form:"duplicate_charge_id"`
	ProductDescription           *string `form:"product_description"`
	Receipt                      *string `form:"receipt"`
	RefundPolicy                 *string `form:"refund_policy"`
	RefundPolicyDisclosure       *string `form:"refund_policy_disclosure"`
	RefundRefusalExplanation     *string `form:"refund_refusal_explanation"`
	ServiceDate                  *string `form:"service_date"`
	ServiceDocumentation         *string `form:"service_documentation"`
	ShippingAddress              *string `form:"shipping_address"`
	ShippingCarrier              *string `form:"shipping_carrier"`
	ShippingDate                 *string `form:"shipping_date"`
	ShippingDocumentation        *string `form:"shipping_documentation"`
	ShippingTrackingNumber       *string `form:"shipping_tracking_number"`
	UncategorizedFile            *string `form:"uncategorized_file"`
	UncategorizedText            *string `form:"uncategorized_text"`
}

// DisputeListParams is the set of parameters that can be used when listing disputes.
// For more details see https://stripe.com/docs/api#list_disputes.
type DisputeListParams struct {
	ListParams    `form:"*"`
	Charge        *string           `form:"charge"`
	Created       *int64            `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	PaymentIntent *string           `form:"payment_intent"`
}

// Dispute is the resource representing a Stripe dispute.
// For more details see https://stripe.com/docs/api#disputes.
type Dispute struct {
	APIResource
	Amount              *int64 `json:"amount,omitempty"`
	BalanceTransactions []*BalanceTransaction `json:"balance_transactions,omitempty"`
	Charge              *Charge               `json:"charge,omitempty"`
	Created             *int64 `json:"created,omitempty"`
	Currency            Currency              `json:"currency,omitempty"`
	Evidence            *DisputeEvidence      `json:"evidence,omitempty"`
	EvidenceDetails     *EvidenceDetails      `json:"evidence_details,omitempty"`
	ID                  *string `json:"id,omitempty"`
	IsChargeRefundable  *bool `json:"is_charge_refundable,omitempty"`
	Livemode            *bool `json:"livemode,omitempty"`
	Metadata            map[string]string     `json:"metadata,omitempty"`
	PaymentIntent       *PaymentIntent        `json:"payment_intent,omitempty"`
	Reason              DisputeReason         `json:"reason,omitempty"`
	Status              DisputeStatus         `json:"status,omitempty"`
}

// DisputeList is a list of disputes as retrieved from a list endpoint.
type DisputeList struct {
	APIResource
	ListMeta
	Data []*Dispute `json:"data,omitempty"`
}

// EvidenceDetails is the structure representing more details about
// the dispute.
type EvidenceDetails struct {
	DueBy           *int64 `json:"due_by,omitempty"`
	HasEvidence     *bool `json:"has_evidence,omitempty"`
	PastDue         *bool `json:"past_due,omitempty"`
	SubmissionCount *int64 `json:"submission_count,omitempty"`
}

// DisputeEvidence is the structure that contains various details about
// the evidence submitted for the dispute.
// Almost all fields are strings since there structures (i.e. address)
// do not typically get parsed by anyone and are thus presented as-received.
type DisputeEvidence struct {
	AccessActivityLog            *string `json:"access_activity_log,omitempty"`
	BillingAddress               *string `json:"billing_address,omitempty"`
	CancellationPolicy           *File  `json:"cancellation_policy,omitempty"`
	CancellationPolicyDisclosure *string `json:"cancellation_policy_disclosure,omitempty"`
	CancellationRebuttal         *string `json:"cancellation_rebuttal,omitempty"`
	CustomerCommunication        *File  `json:"customer_communication,omitempty"`
	CustomerEmailAddress         *string `json:"customer_email_address,omitempty"`
	CustomerName                 *string `json:"customer_name,omitempty"`
	CustomerPurchaseIP           *string `json:"customer_purchase_ip,omitempty"`
	CustomerSignature            *File  `json:"customer_signature,omitempty"`
	DuplicateChargeDocumentation *File  `json:"duplicate_charge_documentation,omitempty"`
	DuplicateChargeExplanation   *string `json:"duplicate_charge_explanation,omitempty"`
	DuplicateChargeID            *string `json:"duplicate_charge_id,omitempty"`
	ProductDescription           *string `json:"product_description,omitempty"`
	Receipt                      *File  `json:"receipt,omitempty"`
	RefundPolicy                 *File  `json:"refund_policy,omitempty"`
	RefundPolicyDisclosure       *string `json:"refund_policy_disclosure,omitempty"`
	RefundRefusalExplanation     *string `json:"refund_refusal_explanation,omitempty"`
	ServiceDate                  *string `json:"service_date,omitempty"`
	ServiceDocumentation         *File  `json:"service_documentation,omitempty"`
	ShippingAddress              *string `json:"shipping_address,omitempty"`
	ShippingCarrier              *string `json:"shipping_carrier,omitempty"`
	ShippingDate                 *string `json:"shipping_date,omitempty"`
	ShippingDocumentation        *File  `json:"shipping_documentation,omitempty"`
	ShippingTrackingNumber       *string `json:"shipping_tracking_number,omitempty"`
	UncategorizedFile            *File  `json:"uncategorized_file,omitempty"`
	UncategorizedText            *string `json:"uncategorized_text,omitempty"`
}

// UnmarshalJSON handles deserialization of a Dispute.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (d *Dispute) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		d.ID = &id
		return nil
	}

	type dispute Dispute
	var v dispute
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*d = Dispute(v)
	return nil
}
