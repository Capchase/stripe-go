package stripe

import (
	"encoding/json"
)

// SetupIntentCancellationReason is the list of allowed values for the cancelation reason.
type SetupIntentCancellationReason string

// List of values that SetupIntentCancellationReason can take.
const (
	SetupIntentCancellationReasonAbandoned           SetupIntentCancellationReason = "abandoned"
	SetupIntentCancellationReasonFailedInvoice       SetupIntentCancellationReason = "failed_invoice"
	SetupIntentCancellationReasonFraudulent          SetupIntentCancellationReason = "fraudulent"
	SetupIntentCancellationReasonRequestedByCustomer SetupIntentCancellationReason = "requested_by_customer"
)

// SetupIntentNextActionType is the list of allowed values for the next action's type.
type SetupIntentNextActionType string

// List of values that SetupIntentNextActionType can take.
const (
	SetupIntentNextActionTypeRedirectToURL SetupIntentNextActionType = "redirect_to_url"
)

// SetupIntentPaymentMethodOptionsCardRequestThreeDSecure is the list of allowed values controlling
// when to request 3D Secure on a SetupIntent.
type SetupIntentPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that SetupIntentNextActionType can take.
const (
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureAny       SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "any"
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureAutomatic SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
)

// SetupIntentStatus is the list of allowed values for the setup intent's status.
type SetupIntentStatus string

// List of values that SetupIntentStatus can take.
const (
	SetupIntentStatusCanceled              SetupIntentStatus = "canceled"
	SetupIntentStatusProcessing            SetupIntentStatus = "processing"
	SetupIntentStatusRequiresAction        SetupIntentStatus = "requires_action"
	SetupIntentStatusRequiresConfirmation  SetupIntentStatus = "requires_confirmation"
	SetupIntentStatusRequiresPaymentMethod SetupIntentStatus = "requires_payment_method"
	SetupIntentStatusSucceeded             SetupIntentStatus = "succeeded"
)

// SetupIntentUsage is the list of allowed values for the setup intent's usage.
type SetupIntentUsage string

// List of values that SetupIntentUsage can take.
const (
	SetupIntentUsageOffSession SetupIntentUsage = "off_session"
	SetupIntentUsageOnSession  SetupIntentUsage = "on_session"
)

// SetupIntentCancelParams is the set of parameters that can be used when canceling a setup intent.
type SetupIntentCancelParams struct {
	Params             `form:"*"`
	CancellationReason *string `form:"cancellation_reason"`
}

// SetupIntentConfirmParams is the set of parameters that can be used when confirming a setup intent.
type SetupIntentConfirmParams struct {
	Params               `form:"*"`
	MandateData          *SetupIntentMandateDataParams          `form:"mandate_data"`
	PaymentMethod        *string                                `form:"payment_method"`
	PaymentMethodOptions *SetupIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	ReturnURL            *string                                `form:"return_url"`
}

// SetupIntentMandateDataCustomerAcceptanceOfflineParams is the set of parameters for the customer
// acceptance of an offline mandate.
type SetupIntentMandateDataCustomerAcceptanceOfflineParams struct {
}

// SetupIntentMandateDataCustomerAcceptanceOnlineParams is the set of parameters for the customer
// acceptance of an online mandate.
type SetupIntentMandateDataCustomerAcceptanceOnlineParams struct {
	IPAddress *string `form:"ip_address"`
	UserAgent *string `form:"user_agent"`
}

// SetupIntentMandateDataCustomerAcceptanceParams is the set of parameters for the customer
// acceptance of a mandate.
type SetupIntentMandateDataCustomerAcceptanceParams struct {
	AcceptedAt int64                                                  `form:"accepted_at"`
	Offline    *SetupIntentMandateDataCustomerAcceptanceOfflineParams `form:"offline"`
	Online     *SetupIntentMandateDataCustomerAcceptanceOnlineParams  `form:"online"`
	Type       MandateCustomerAcceptanceType                          `form:"type"`
}

// SetupIntentMandateDataParams is the set of parameters controlling the creation of the mandate
// associated with this SetupIntent.
type SetupIntentMandateDataParams struct {
	CustomerAcceptance *SetupIntentMandateDataCustomerAcceptanceParams `form:"customer_acceptance"`
}

// SetupIntentPaymentMethodOptionsCardParams represents the card-specific options applied to a
// SetupIntent.
type SetupIntentPaymentMethodOptionsCardParams struct {
	MOTO                *bool   `form:"moto"`
	RequestThreeDSecure *string `form:"request_three_d_secure"`
}

// SetupIntentPaymentMethodOptionsParams represents the type-specific payment method options
// applied to a SetupIntent.
type SetupIntentPaymentMethodOptionsParams struct {
	Card *SetupIntentPaymentMethodOptionsCardParams `form:"card"`
}

// SetupIntentSingleUseParams represents the single-use mandate-specific parameters.
type SetupIntentSingleUseParams struct {
	Amount   *int64  `form:"amount"`
	Currency *string `form:"currency"`
}

// SetupIntentParams is the set of parameters that can be used when handling a setup intent.
type SetupIntentParams struct {
	Params               `form:"*"`
	Confirm              *bool                                  `form:"confirm"`
	Customer             *string                                `form:"customer"`
	Description          *string                                `form:"description"`
	MandateData          *SetupIntentMandateDataParams          `form:"mandate_data"`
	OnBehalfOf           *string                                `form:"on_behalf_of"`
	PaymentMethod        *string                                `form:"payment_method"`
	PaymentMethodOptions *SetupIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	PaymentMethodTypes   []*string                              `form:"payment_method_types"`
	ReturnURL            *string                                `form:"return_url"`
	SingleUse            *SetupIntentSingleUseParams            `form:"single_use"`
	Usage                *string                                `form:"usage"`
}

// SetupIntentListParams is the set of parameters that can be used when listing setup intents.
// For more details see https://stripe.com/docs/api#list_payouts.
type SetupIntentListParams struct {
	ListParams    `form:"*"`
	Created       *int64            `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	Customer      *string           `form:"customer"`
	PaymentMethod *string           `form:"payment_method"`
}

// SetupIntentNextActionRedirectToURL represents the resource for the next action of type
// "redirect_to_url".
type SetupIntentNextActionRedirectToURL struct {
	ReturnURL *string `json:"return_url,omitempty"`
	URL       *string `json:"url,omitempty"`
}

// SetupIntentNextAction represents the type of action to take on a setup intent.
type SetupIntentNextAction struct {
	RedirectToURL *SetupIntentNextActionRedirectToURL `json:"redirect_to_url,omitempty"`
	Type          SetupIntentNextActionType           `json:"type,omitempty"`
}

// SetupIntentPaymentMethodOptionsCard represents the card-specific options applied to a
// SetupIntent.
type SetupIntentPaymentMethodOptionsCard struct {
	RequestThreeDSecure SetupIntentPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure,omitempty"`
}

// SetupIntentPaymentMethodOptions represents the type-specific payment method options applied to a
// SetupIntent.
type SetupIntentPaymentMethodOptions struct {
	Card *SetupIntentPaymentMethodOptionsCard `json:"card,omitempty"`
}

// SetupIntent is the resource representing a Stripe payout.
// For more details see https://stripe.com/docs/api#payment_intents.
type SetupIntent struct {
	APIResource
	Application          *Application                     `json:"application,omitempty"`
	CancellationReason   SetupIntentCancellationReason    `json:"cancellation_reason,omitempty"`
	ClientSecret         *string `json:"client_secret,omitempty"`
	Created              *int64 `json:"created,omitempty"`
	Customer             *Customer                        `json:"customer,omitempty"`
	Description          *string `json:"description,omitempty"`
	ID                   *string `json:"id,omitempty"`
	LastSetupError       *Error                           `json:"last_setup_error,omitempty"`
	LatestAttempt        *SetupAttempt                    `json:"latest_attempt,omitempty"`
	Livemode             *bool `json:"livemode,omitempty"`
	Mandate              *Mandate                         `json:"mandate,omitempty"`
	Metadata             map[string]string                `json:"metadata,omitempty"`
	NextAction           *SetupIntentNextAction           `json:"next_action,omitempty"`
	Object               *string `json:"object,omitempty"`
	OnBehalfOf           *Account                         `json:"on_behalf_of,omitempty"`
	PaymentMethod        *PaymentMethod                   `json:"payment_method,omitempty"`
	PaymentMethodOptions *SetupIntentPaymentMethodOptions `json:"payment_method_options,omitempty"`
	PaymentMethodTypes   []string                         `json:"payment_method_types,omitempty"`
	SingleUseMandate     *Mandate                         `json:"single_use_mandate,omitempty"`
	Status               SetupIntentStatus                `json:"status,omitempty"`
	Usage                SetupIntentUsage                 `json:"usage,omitempty"`
}

// SetupIntentList is a list of setup intents as retrieved from a list endpoint.
type SetupIntentList struct {
	APIResource
	ListMeta
	Data []*SetupIntent `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of a SetupIntent.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *SetupIntent) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = &id
		return nil
	}

	type setupIntent SetupIntent
	var v setupIntent
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = SetupIntent(v)
	return nil
}
