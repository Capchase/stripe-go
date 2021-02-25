package stripe

import "encoding/json"

// List of values that MandateStatus can take.
const (
	MandateCustomerAcceptanceTypeOffline MandateCustomerAcceptanceType = "offline"
	MandateCustomerAcceptanceTypeOnline  MandateCustomerAcceptanceType = "online"
)

// MandateCustomerAcceptanceType is the list of allowed values for the type of customer acceptance
// for a given mandate.
type MandateCustomerAcceptanceType string

// List of values that MandateStatus can take.
const (
	MandateStatusActive   MandateStatus = "active"
	MandateStatusInactive MandateStatus = "inactive"
	MandateStatusPending  MandateStatus = "pending"
)

// MandatePaymentMethodDetailsBACSDebitNetworkStatus is the list of allowed values for the status
// with the network for a given mandate.
type MandatePaymentMethodDetailsBACSDebitNetworkStatus string

// List of values that MandateStatus can take.
const (
	MandatePaymentMethodDetailsBACSDebitNetworkStatusAccepted MandatePaymentMethodDetailsBACSDebitNetworkStatus = "accepted"
	MandatePaymentMethodDetailsBACSDebitNetworkStatusPending  MandatePaymentMethodDetailsBACSDebitNetworkStatus = "pending"
	MandatePaymentMethodDetailsBACSDebitNetworkStatusRefused  MandatePaymentMethodDetailsBACSDebitNetworkStatus = "refused"
	MandatePaymentMethodDetailsBACSDebitNetworkStatusRevoked  MandatePaymentMethodDetailsBACSDebitNetworkStatus = "revoked"
)

// MandateStatus is the list of allowed values for the mandate status.
type MandateStatus string

// List of values that MandateType can take.
const (
	MandateTypeMultiUse  MandateType = "multi_use"
	MandateTypeSingleUse MandateType = "single_use"
)

// MandateType is the list of allowed values for the mandate type.
type MandateType string

// MandateParams is the set of parameters that can be used when retrieving a mandate.
type MandateParams struct {
	Params `form:"*"`
}

// MandateCustomerAcceptanceOffline represents details about the customer acceptance of an offline
// mandate.
type MandateCustomerAcceptanceOffline struct {
}

// MandateCustomerAcceptanceOnline represents details about the customer acceptance of an online
// mandate.
type MandateCustomerAcceptanceOnline struct {
	IPAddress *string `json:"ip_address,omitempty"`
	UserAgent *string `json:"user_agent,omitempty"`
}

// MandateCustomerAcceptance represents details about the customer acceptance for a mandate.
type MandateCustomerAcceptance struct {
	AcceptedAt *int64 `json:"accepted_at,omitempty"`
	Offline    *MandateCustomerAcceptanceOffline `json:"offline,omitempty"`
	Online     *MandateCustomerAcceptanceOnline  `json:"online,omitempty"`
	Type       MandateCustomerAcceptanceType     `json:"type,omitempty"`
}

// MandateMultiUse represents details about a multi-use mandate.
type MandateMultiUse struct {
}

// MandatePaymentMethodDetailsAUBECSDebit represents details about the Australia BECS debit account
// associated with this mandate.
type MandatePaymentMethodDetailsAUBECSDebit struct {
	URL *string `json:"url,omitempty"`
}

// MandatePaymentMethodDetailsBACSDebit represents details about the BACS debit account
// associated with this mandate.
type MandatePaymentMethodDetailsBACSDebit struct {
	NetworkStatus MandatePaymentMethodDetailsBACSDebitNetworkStatus `json:"network_status,omitempty"`
	Reference     *string `json:"reference,omitempty"`
	URL           *string `json:"url,omitempty"`
}

// MandatePaymentMethodDetailsCard represents details about the card associated with this mandate.
type MandatePaymentMethodDetailsCard struct {
}

// MandatePaymentMethodDetailsSepaDebit represents details about the SEPA debit bank account
// associated with this mandate.
type MandatePaymentMethodDetailsSepaDebit struct {
	Reference *string `json:"reference,omitempty"`
	URL       *string `json:"url,omitempty"`
}

// MandatePaymentMethodDetails represents details about the payment method associated with this
// mandate.
type MandatePaymentMethodDetails struct {
	AUBECSDebit *MandatePaymentMethodDetailsAUBECSDebit `json:"au_becs_debit,omitempty"`
	BACSDebit   *MandatePaymentMethodDetailsBACSDebit   `json:"bacs_debit,omitempty"`
	Card        *MandatePaymentMethodDetailsCard        `json:"card,omitempty"`
	SepaDebit   *MandatePaymentMethodDetailsSepaDebit   `json:"sepa_debit,omitempty"`
	Type        PaymentMethodType                       `json:"type,omitempty"`
}

// MandateSingleUse represents details about a single-use mandate.
type MandateSingleUse struct {
	Amount   *int64 `json:"amount,omitempty"`
	Currency Currency `json:"currency,omitempty"`
}

// Mandate is the resource representing a Mandate.
type Mandate struct {
	APIResource
	CustomerAcceptance   *MandateCustomerAcceptance   `json:"customer_acceptance,omitempty"`
	ID                   *string `json:"id,omitempty"`
	Livemode             *bool `json:"livemode,omitempty"`
	MultiUse             *MandateMultiUse             `json:"multi_use,omitempty"`
	Object               *string `json:"object,omitempty"`
	PaymentMethod        *PaymentMethod               `json:"payment_method,omitempty"`
	PaymentMethodDetails *MandatePaymentMethodDetails `json:"payment_method_details,omitempty"`
	SingleUse            *MandateSingleUse            `json:"single_use,omitempty"`
	Status               MandateStatus                `json:"status,omitempty"`
	Type                 MandateType                  `json:"type,omitempty"`
}

// UnmarshalJSON handles deserialization of a Mandate.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *Mandate) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = &id
		return nil
	}

	type ma Mandate
	var v ma
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = Mandate(v)
	return nil
}
