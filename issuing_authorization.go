package stripe

import "encoding/json"

// IssuingAuthorizationAuthorizationMethod is the list of possible values for the authorization method
// on an issuing authorization.
type IssuingAuthorizationAuthorizationMethod string

// List of values that IssuingAuthorizationAuthorizationMethod can take.
const (
	IssuingAuthorizationAuthorizationMethodChip        IssuingAuthorizationAuthorizationMethod = "chip"
	IssuingAuthorizationAuthorizationMethodContactless IssuingAuthorizationAuthorizationMethod = "contactless"
	IssuingAuthorizationAuthorizationMethodKeyedIn     IssuingAuthorizationAuthorizationMethod = "keyed_in"
	IssuingAuthorizationAuthorizationMethodOnline      IssuingAuthorizationAuthorizationMethod = "online"
	IssuingAuthorizationAuthorizationMethodSwipe       IssuingAuthorizationAuthorizationMethod = "swipe"
)

// IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity is the list of possible values
// for the entity that owns the authorization control.
type IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity string

// List of values that IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity can take.
const (
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntityAccount    IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity = "account"
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntityCard       IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity = "card"
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntityCardholder IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity = "cardholder"
)

// IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName is the list of possible values
// for the name associated with the authorization control.
type IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName string

// List of values that IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName can take.
const (
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlNameAllowedCategories IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName = "allowed_categories"
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlNameBlockedCategories IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName = "blocked_categories"
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlNameMaxAmount         IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName = "max_amount"
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlNameMaxApprovals      IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName = "max_approvals"
	IssuingAuthorizationRequestHistoryViolatedAuthorizationControlNameSpendingLimits    IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName = "spending_limits"
)

// IssuingAuthorizationRequestHistoryReason is the list of possible values for the request history
// reason on an issuing authorization.
type IssuingAuthorizationRequestHistoryReason string

// List of values that IssuingAuthorizationRequestHistoryReason can take.
const (
	IssuingAuthorizationRequestHistoryReasonAccountDisabled                IssuingAuthorizationRequestHistoryReason = "account_disabled"
	IssuingAuthorizationRequestHistoryReasonCardActive                     IssuingAuthorizationRequestHistoryReason = "card_active"
	IssuingAuthorizationRequestHistoryReasonCardInactive                   IssuingAuthorizationRequestHistoryReason = "card_inactive"
	IssuingAuthorizationRequestHistoryReasonCardholderInactive             IssuingAuthorizationRequestHistoryReason = "cardholder_inactive"
	IssuingAuthorizationRequestHistoryReasonCardholderVerificationRequired IssuingAuthorizationRequestHistoryReason = "cardholder_verification_required"
	IssuingAuthorizationRequestHistoryReasonInsufficientFunds              IssuingAuthorizationRequestHistoryReason = "insufficient_funds"
	IssuingAuthorizationRequestHistoryReasonNotAllowed                     IssuingAuthorizationRequestHistoryReason = "not_allowed"
	IssuingAuthorizationRequestHistoryReasonSpendingControls               IssuingAuthorizationRequestHistoryReason = "spending_controls"
	IssuingAuthorizationRequestHistoryReasonSuspectedFraud                 IssuingAuthorizationRequestHistoryReason = "suspected_fraud"
	IssuingAuthorizationRequestHistoryReasonVerificationFailed             IssuingAuthorizationRequestHistoryReason = "verification_failed"
	IssuingAuthorizationRequestHistoryReasonWebhookApproved                IssuingAuthorizationRequestHistoryReason = "webhook_approved"
	IssuingAuthorizationRequestHistoryReasonWebhookDeclined                IssuingAuthorizationRequestHistoryReason = "webhook_declined"
	IssuingAuthorizationRequestHistoryReasonWebhookTimeout                 IssuingAuthorizationRequestHistoryReason = "webhook_timeout"
)

// IssuingAuthorizationStatus is the possible values for status for an issuing authorization.
type IssuingAuthorizationStatus string

// List of values that IssuingAuthorizationStatus can take.
const (
	IssuingAuthorizationStatusClosed   IssuingAuthorizationStatus = "closed"
	IssuingAuthorizationStatusPending  IssuingAuthorizationStatus = "pending"
	IssuingAuthorizationStatusReversed IssuingAuthorizationStatus = "reversed"
)

// IssuingAuthorizationVerificationDataCheck is the list of possible values for result of a check
// for verification data on an issuing authorization.
type IssuingAuthorizationVerificationDataCheck string

// List of values that IssuingAuthorizationVerificationDataCheck can take.
const (
	IssuingAuthorizationVerificationDataCheckMatch       IssuingAuthorizationVerificationDataCheck = "match"
	IssuingAuthorizationVerificationDataCheckMismatch    IssuingAuthorizationVerificationDataCheck = "mismatch"
	IssuingAuthorizationVerificationDataCheckNotProvided IssuingAuthorizationVerificationDataCheck = "not_provided"
)

// IssuingAuthorizationWalletType is the list of possible values for the authorization's wallet provider.
type IssuingAuthorizationWalletType string

// List of values that IssuingAuthorizationWalletType can take.
const (
	IssuingAuthorizationWalletTypeApplePay   IssuingAuthorizationWalletType = "apple_pay"
	IssuingAuthorizationWalletTypeGooglePay  IssuingAuthorizationWalletType = "google_pay"
	IssuingAuthorizationWalletTypeSamsungPay IssuingAuthorizationWalletType = "samsung_pay"
)

// IssuingAuthorizationParams is the set of parameters that can be used when updating an issuing authorization.
type IssuingAuthorizationParams struct {
	Params `form:"*"`
}

// IssuingAuthorizationApproveParams is the set of parameters that can be used when approving an issuing authorization.
type IssuingAuthorizationApproveParams struct {
	Params `form:"*"`
	Amount *int64 `form:"amount"`
}

// IssuingAuthorizationDeclineParams is the set of parameters that can be used when declining an issuing authorization.
type IssuingAuthorizationDeclineParams struct {
	Params `form:"*"`
}

// IssuingAuthorizationListParams is the set of parameters that can be used when listing issuing authorizations.
type IssuingAuthorizationListParams struct {
	ListParams   `form:"*"`
	Card         *string           `form:"card"`
	Cardholder   *string           `form:"cardholder"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Status       *string           `form:"status"`
}

// IssuingAuthorizationAmountDetails is the resource representing the breakdown of the amount.
type IssuingAuthorizationAmountDetails struct {
	ATMFee *int64 `json:"atm_fee,omitempty"`
}

// IssuingAuthorizationMerchantData is the resource representing merchant data on Issuing APIs.
type IssuingAuthorizationMerchantData struct {
	Category   *string `json:"category,omitempty"`
	City       *string `json:"city,omitempty"`
	Country    *string `json:"country,omitempty"`
	Name       *string `json:"name,omitempty"`
	NetworkID  *string `json:"network_id,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	State      *string `json:"state,omitempty"`
}

// IssuingAuthorizationPendingRequest is the resource representing details about the pending authorization request.
type IssuingAuthorizationPendingRequest struct {
	Amount               *int64 `json:"amount,omitempty"`
	AmountDetails        *IssuingAuthorizationAmountDetails `json:"amount_details,omitempty"`
	Currency             Currency                           `json:"currency,omitempty"`
	IsAmountControllable *bool `json:"is_amount_controllable,omitempty"`
	MerchantAmount       *int64 `json:"merchant_amount,omitempty"`
	MerchantCurrency     Currency                           `json:"merchant_currency,omitempty"`
}

// IssuingAuthorizationRequestHistory is the resource representing a request history on an issuing authorization.
type IssuingAuthorizationRequestHistory struct {
	Amount           *int64 `json:"amount,omitempty"`
	AmountDetails    *IssuingAuthorizationAmountDetails       `json:"amount_details,omitempty"`
	Approved         *bool `json:"approved,omitempty"`
	Created          *int64 `json:"created,omitempty"`
	Currency         Currency                                 `json:"currency,omitempty"`
	MerchantAmount   *int64 `json:"merchant_amount,omitempty"`
	MerchantCurrency Currency                                 `json:"merchant_currency,omitempty"`
	Reason           IssuingAuthorizationRequestHistoryReason `json:"reason,omitempty"`
}

// IssuingAuthorizationVerificationData is the resource representing verification data on an issuing authorization.
type IssuingAuthorizationVerificationData struct {
	AddressLine1Check      IssuingAuthorizationVerificationDataCheck `json:"address_line1_check,omitempty"`
	AddressPostalCodeCheck IssuingAuthorizationVerificationDataCheck `json:"address_postal_code_check,omitempty"`
	CVCCheck               IssuingAuthorizationVerificationDataCheck `json:"cvc_check,omitempty"`
	ExpiryCheck            IssuingAuthorizationVerificationDataCheck `json:"expiry_check,omitempty"`
}

// IssuingAuthorization is the resource representing a Stripe issuing authorization.
type IssuingAuthorization struct {
	APIResource
	Amount              *int64 `json:"amount,omitempty"`
	AmountDetails       *IssuingAuthorizationAmountDetails      `json:"amount_details,omitempty"`
	Approved            *bool `json:"approved,omitempty"`
	AuthorizationMethod IssuingAuthorizationAuthorizationMethod `json:"authorization_method,omitempty"`
	BalanceTransactions []*BalanceTransaction                   `json:"balance_transactions,omitempty"`
	Card                *IssuingCard                            `json:"card,omitempty"`
	Cardholder          *IssuingCardholder                      `json:"cardholder,omitempty"`
	Created             *int64 `json:"created,omitempty"`
	Currency            Currency                                `json:"currency,omitempty"`
	ID                  *string `json:"id,omitempty"`
	Livemode            *bool `json:"livemode,omitempty"`
	MerchantAmount      *int64 `json:"merchant_amount,omitempty"`
	MerchantCurrency    Currency                                `json:"merchant_currency,omitempty"`
	MerchantData        *IssuingAuthorizationMerchantData       `json:"merchant_data,omitempty"`
	Metadata            map[string]string                       `json:"metadata,omitempty"`
	Object              *string `json:"object,omitempty"`
	PendingRequest      *IssuingAuthorizationPendingRequest     `json:"pending_request,omitempty"`
	RequestHistory      []*IssuingAuthorizationRequestHistory   `json:"request_history,omitempty"`
	Status              IssuingAuthorizationStatus              `json:"status,omitempty"`
	Transactions        []*IssuingTransaction                   `json:"transactions,omitempty"`
	VerificationData    *IssuingAuthorizationVerificationData   `json:"verification_data,omitempty"`
	Wallet              IssuingAuthorizationWalletType          `json:"wallet,omitempty"`
}

// IssuingAuthorizationList is a list of issuing authorizations as retrieved from a list endpoint.
type IssuingAuthorizationList struct {
	APIResource
	ListMeta
	Data []*IssuingAuthorization `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of an IssuingAuthorization.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingAuthorization) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = &id
		return nil
	}

	type issuingAuthorization IssuingAuthorization
	var v issuingAuthorization
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingAuthorization(v)
	return nil
}
