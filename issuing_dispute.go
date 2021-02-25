package stripe

import "encoding/json"

// IssuingDisputeEvidenceReason is the list of allowed reasons for the evidence on an issuing dispute.
type IssuingDisputeEvidenceReason string

// List of values that IssuingDisputeEvidenceReason can take.
const (
	IssuingDisputeEvidenceReasonCanceled                  IssuingDisputeEvidenceReason = "canceled"
	IssuingDisputeEvidenceReasonDuplicate                 IssuingDisputeEvidenceReason = "duplicate"
	IssuingDisputeEvidenceReasonFraudulent                IssuingDisputeEvidenceReason = "fraudulent"
	IssuingDisputeEvidenceReasonMerchandiseNotAsDescribed IssuingDisputeEvidenceReason = "merchandise_not_as_described"
	IssuingDisputeEvidenceReasonNotReceived               IssuingDisputeEvidenceReason = "not_received"
	IssuingDisputeEvidenceReasonOther                     IssuingDisputeEvidenceReason = "other"
	IssuingDisputeEvidenceReasonServiceNotAsDescribed     IssuingDisputeEvidenceReason = "service_not_as_described"
)

// IssuingDisputeEvidenceCanceledProductType is the list of allowed product types on an issuing dispute of type canceled.
type IssuingDisputeEvidenceCanceledProductType string

// List of values that IssuingDisputeEvidenceProductType can take.
const (
	IssuingDisputeEvidenceCanceledProductTypeMerchandise IssuingDisputeEvidenceCanceledProductType = "merchandise"
	IssuingDisputeEvidenceCanceledProductTypeService     IssuingDisputeEvidenceCanceledProductType = "service"
)

// IssuingDisputeEvidenceCanceledReturnStatus is the list of allowed return status on an issuing dispute of type canceled.
type IssuingDisputeEvidenceCanceledReturnStatus string

// List of values that IssuingDisputeEvidenceCanceledReturnStatus can take.
const (
	IssuingDisputeEvidenceCanceledReturnStatusMerchantRejected IssuingDisputeEvidenceCanceledReturnStatus = "merchant_rejected"
	IssuingDisputeEvidenceCanceledReturnStatusSuccessful       IssuingDisputeEvidenceCanceledReturnStatus = "successful"
)

// IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatus is the list of allowed return status on an issuing dispute of type merchandise not as described.
type IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatus string

// List of values that IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatus can take.
const (
	IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatusMerchantRejected IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatus = "merchant_rejected"
	IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatusSuccessful       IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatus = "successful"
)

// IssuingDisputeEvidenceNotReceivedProductType is the list of allowed product types on an issuing dispute of type not received.
type IssuingDisputeEvidenceNotReceivedProductType string

// List of values that IssuingDisputeEvidenceNotReceivedProductType can take.
const (
	IssuingDisputeEvidenceNotReceivedProductTypeMerchandise IssuingDisputeEvidenceNotReceivedProductType = "merchandise"
	IssuingDisputeEvidenceNotReceivedProductTypeService     IssuingDisputeEvidenceNotReceivedProductType = "service"
)

// IssuingDisputeEvidenceOtherProductType is the list of allowed product types on an issuing dispute of type other.
type IssuingDisputeEvidenceOtherProductType string

// List of values that IssuingDisputeEvidenceNotReceivedProductType can take.
const (
	IssuingDisputeEvidenceOtherProductTypeMerchandise IssuingDisputeEvidenceOtherProductType = "merchandise"
	IssuingDisputeEvidenceOtherProductTypeService     IssuingDisputeEvidenceOtherProductType = "service"
)

// IssuingDisputeEvidenceServiceNotAsDescribedProductType is the list of allowed product types on an issuing dispute of type service not as described.
type IssuingDisputeEvidenceServiceNotAsDescribedProductType string

// List of values that IssuingDisputeEvidenceServiceNotAsDescribedProductType can take.
const (
	IssuingDisputeEvidenceServiceNotAsDescribedProductTypeMerchandise IssuingDisputeEvidenceServiceNotAsDescribedProductType = "merchandise"
	IssuingDisputeEvidenceServiceNotAsDescribedProductTypeService     IssuingDisputeEvidenceServiceNotAsDescribedProductType = "service"
)

// IssuingDisputeStatus is the list of allowed values for status on an issuing dispute.
type IssuingDisputeStatus string

// List of values that IssuingDisputeStatus can take.
const (
	IssuingDisputeStatusExpired     IssuingDisputeStatus = "expired"
	IssuingDisputeStatusLost        IssuingDisputeStatus = "lost"
	IssuingDisputeStatusSubmitted   IssuingDisputeStatus = "submitted"
	IssuingDisputeStatusUnsubmitted IssuingDisputeStatus = "unsubmitted"
	IssuingDisputeStatusWon         IssuingDisputeStatus = "won"
)

// IssuingDisputeEvidenceCanceledParams is the resource representing the evidence of an issuing dispute with a reason set to canceled.
type IssuingDisputeEvidenceCanceledParams struct {
	AdditionalDocumentation    *string `form:"additional_documentation"`
	CanceledAt                 *int64  `form:"canceled_at"`
	CancellationPolicyProvided *bool   `form:"cancellation_policy_provided"`
	CancellationReason         *string `form:"cancellation_reason"`
	ExpectedAt                 *int64  `form:"expected_at"`
	Explanation                *string `form:"explanation"`
	ProductDescription         *string `form:"product_description"`
	ProductType                *string `form:"product_type"`
	ReturnStatus               *string `form:"return_status"`
	ReturnedAt                 *int64  `form:"returned_at"`
}

// IssuingDisputeEvidenceDuplicateParams is the resource representing the evidence of an issuing dispute with a reason set to duplicate.
type IssuingDisputeEvidenceDuplicateParams struct {
	AdditionalDocumentation *string `form:"additional_documentation"`
	CardStatement           *string `form:"card_statement"`
	CashReceipt             *string `form:"cash_receipt"`
	CheckImage              *string `form:"check_image"`
	Explanation             *string `form:"explanation"`
	OriginalTransaction     *string `form:"original_transaction"`
}

// IssuingDisputeEvidenceFraudulentParams is the resource representing the evidence of an issuing dispute with a reason set to fraudulent.
type IssuingDisputeEvidenceFraudulentParams struct {
	AdditionalDocumentation *string `form:"additional_documentation"`
	Explanation             *string `form:"explanation"`
}

// IssuingDisputeEvidenceMerchandiseNotAsDescribedParams is the resource representing the evidence of an issuing dispute with a reason set to merchandise not as described.
type IssuingDisputeEvidenceMerchandiseNotAsDescribedParams struct {
	AdditionalDocumentation *string `form:"additional_documentation"`
	Explanation             *string `form:"explanation"`
	ReceivedAt              *int64  `form:"received_at"`
	ReturnDescription       *string `form:"return_description"`
	ReturnStatus            *string `form:"return_status"`
	ReturnedAt              *int64  `form:"returned_at"`
}

// IssuingDisputeEvidenceNotReceivedParams is the resource representing the evidence of an issuing dispute with a reason set to not received.
type IssuingDisputeEvidenceNotReceivedParams struct {
	AdditionalDocumentation *string `form:"additional_documentation"`
	ExpectedAt              *int64  `form:"expected_at"`
	Explanation             *string `form:"explanation"`
	ProductDescription      *string `form:"product_description"`
	ProductType             *string `form:"product_type"`
}

// IssuingDisputeEvidenceOtherParams is the resource representing the evidence of an issuing dispute with a reason set to other.
type IssuingDisputeEvidenceOtherParams struct {
	AdditionalDocumentation *string `form:"additional_documentation"`
	Explanation             *string `form:"explanation"`
	ProductDescription      *string `form:"product_description"`
	ProductType             *string `form:"product_type"`
}

// IssuingDisputeEvidenceServiceNotAsDescribedParams is the resource representing the evidence of an issuing dispute with a reason set to service not as described.
type IssuingDisputeEvidenceServiceNotAsDescribedParams struct {
	AdditionalDocumentation *string `form:"additional_documentation"`
	CanceledAt              *int64  `form:"canceled_at"`
	Explanation             *string `form:"explanation"`
	ProductDescription      *string `form:"product_description"`
	ProductType             *string `form:"product_type"`
}

// IssuingDisputeEvidenceParams is the set of parameters for the evidence on an issuing dispute
type IssuingDisputeEvidenceParams struct {
	Canceled                  *IssuingDisputeEvidenceCanceledParams                  `form:"canceled"`
	Duplicate                 *IssuingDisputeEvidenceDuplicateParams                 `form:"duplicate"`
	Fraudulent                *IssuingDisputeEvidenceFraudulentParams                `form:"fraudulent"`
	MerchandiseNotAsDescribed *IssuingDisputeEvidenceMerchandiseNotAsDescribedParams `form:"merchandise_not_as_described"`
	NotReceived               *IssuingDisputeEvidenceNotReceivedParams               `form:"not_received"`
	Other                     *IssuingDisputeEvidenceOtherParams                     `form:"other"`
	Reason                    *string                                                `form:"reason"`
	ServiceNotAsDescribed     *IssuingDisputeEvidenceServiceNotAsDescribedParams     `form:"service_not_as_described"`
}

// IssuingDisputeParams is the set of parameters that can be used when creating or updating an issuing dispute.
type IssuingDisputeParams struct {
	Params      `form:"*"`
	Evidence    *IssuingDisputeEvidenceParams `form:"evidence"`
	Transaction *string                       `form:"transaction"`
}

// IssuingDisputeListParams is the set of parameters that can be used when listing issuing dispute.
type IssuingDisputeListParams struct {
	ListParams  `form:"*"`
	Status      *string `form:"status"`
	Transaction *string `form:"transaction"`
}

// IssuingDisputeSubmitParams is the set of parameters that can be used when submitting an issuing dispute.
type IssuingDisputeSubmitParams struct {
	Params `form:"*"`
}

// IssuingDisputeEvidenceCanceled is the resource representing the evidence of an issuing dispute with a reason set to canceled.
type IssuingDisputeEvidenceCanceled struct {
	AdditionalDocumentation    *File                                      `json:"additional_documentation,omitempty"`
	CanceledAt                 *int64 `json:"canceled_at,omitempty"`
	CancellationPolicyProvided *bool `json:"cancellation_policy_provided,omitempty"`
	CancellationReason         *string `json:"cancellation_reason,omitempty"`
	ExpectedAt                 *int64 `json:"expected_at,omitempty"`
	Explanation                *string `json:"explanation,omitempty"`
	ProductDescription         *string `json:"product_description,omitempty"`
	ProductType                IssuingDisputeEvidenceCanceledProductType  `json:"product_type,omitempty"`
	ReturnStatus               IssuingDisputeEvidenceCanceledReturnStatus `json:"return_status,omitempty"`
	ReturnedAt                 *int64 `json:"returned_at,omitempty"`
}

// IssuingDisputeEvidenceDuplicate is the resource representing the evidence of an issuing dispute with a reason set to duplicate.
type IssuingDisputeEvidenceDuplicate struct {
	AdditionalDocumentation *File  `json:"additional_documentation,omitempty"`
	CardStatement           *File  `json:"card_statement,omitempty"`
	CashReceipt             *File  `json:"cash_receipt,omitempty"`
	CheckImage              *File  `json:"check_image,omitempty"`
	Explanation             *string `json:"explanation,omitempty"`
	OriginalTransaction     *string `json:"original_transaction,omitempty"`
}

// IssuingDisputeEvidenceFraudulent is the resource representing the evidence of an issuing dispute with a reason set to fraudulent.
type IssuingDisputeEvidenceFraudulent struct {
	AdditionalDocumentation *File  `json:"additional_documentation,omitempty"`
	Explanation             *string `json:"explanation,omitempty"`
}

// IssuingDisputeEvidenceMerchandiseNotAsDescribed is the resource representing the evidence of an issuing dispute with a reason set to merchandise not as described.
type IssuingDisputeEvidenceMerchandiseNotAsDescribed struct {
	AdditionalDocumentation *File                                                       `json:"additional_documentation,omitempty"`
	Explanation             *string `json:"explanation,omitempty"`
	ReceivedAt              *int64 `json:"received_at,omitempty"`
	ReturnDescription       *string `json:"return_description,omitempty"`
	ReturnStatus            IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatus `json:"return_status,omitempty"`
	ReturnedAt              *int64 `json:"returned_at,omitempty"`
}

// IssuingDisputeEvidenceNotReceived is the resource representing the evidence of an issuing dispute with a reason set to not received.
type IssuingDisputeEvidenceNotReceived struct {
	AdditionalDocumentation *File                                        `json:"additional_documentation,omitempty"`
	ExpectedAt              *int64 `json:"expected_at,omitempty"`
	Explanation             *string `json:"explanation,omitempty"`
	ProductDescription      *string `json:"product_description,omitempty"`
	ProductType             IssuingDisputeEvidenceNotReceivedProductType `json:"product_type,omitempty"`
}

// IssuingDisputeEvidenceOther is the resource representing the evidence of an issuing dispute with a reason set to other.
type IssuingDisputeEvidenceOther struct {
	AdditionalDocumentation *File                                  `json:"additional_documentation,omitempty"`
	Explanation             *string `json:"explanation,omitempty"`
	ProductDescription      *string `json:"product_description,omitempty"`
	ProductType             IssuingDisputeEvidenceOtherProductType `json:"product_type,omitempty"`
}

// IssuingDisputeEvidenceServiceNotAsDescribed is the resource representing the evidence of an issuing dispute with a reason set to service not as described.
type IssuingDisputeEvidenceServiceNotAsDescribed struct {
	AdditionalDocumentation *File                                                  `json:"additional_documentation,omitempty"`
	CanceledAt              *int64 `json:"canceled_at,omitempty"`
	Explanation             *string `json:"explanation,omitempty"`
	ProductDescription      *string `json:"product_description,omitempty"`
	ProductType             IssuingDisputeEvidenceServiceNotAsDescribedProductType `json:"product_type,omitempty"`
}

// IssuingDisputeEvidence is the resource representing the evidence of an issuing dispute.
type IssuingDisputeEvidence struct {
	Canceled                  *IssuingDisputeEvidenceCanceled                  `json:"canceled,omitempty"`
	Duplicate                 *IssuingDisputeEvidenceDuplicate                 `json:"duplicate,omitempty"`
	Fraudulent                *IssuingDisputeEvidenceFraudulent                `json:"fraudulent,omitempty"`
	MerchandiseNotAsDescribed *IssuingDisputeEvidenceMerchandiseNotAsDescribed `json:"merchandise_not_as_described,omitempty"`
	NotReceived               *IssuingDisputeEvidenceNotReceived               `json:"not_received,omitempty"`
	Other                     *IssuingDisputeEvidenceOther                     `json:"other,omitempty"`
	Reason                    IssuingDisputeEvidenceReason                     `json:"reason,omitempty"`
	ServiceNotAsDescribed     *IssuingDisputeEvidenceServiceNotAsDescribed     `json:"service_not_as_described,omitempty"`
}

// IssuingDispute is the resource representing an issuing dispute.
type IssuingDispute struct {
	APIResource
	Amount              *int64 `json:"amount,omitempty"`
	BalanceTransactions []*BalanceTransaction   `json:"balance_transactions,omitempty"`
	Created             *int64 `json:"created,omitempty"`
	Currency            Currency                `json:"currency,omitempty"`
	Evidence            *IssuingDisputeEvidence `json:"evidence,omitempty"`
	ID                  *string `json:"id,omitempty"`
	Livemode            *bool `json:"livemode,omitempty"`
	Metadata            map[string]string       `json:"metadata,omitempty"`
	Object              *string `json:"object,omitempty"`
	Status              *IssuingDisputeStatus   `json:"status,omitempty"`
	Transaction         *IssuingTransaction     `json:"transaction,omitempty"`
}

// IssuingDisputeList is a list of issuing disputes as retrieved from a list endpoint.
type IssuingDisputeList struct {
	APIResource
	ListMeta
	Data []*IssuingDispute `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of an IssuingDispute.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingDispute) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = &id
		return nil
	}

	type issuingDispute IssuingDispute
	var v issuingDispute
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingDispute(v)
	return nil
}
