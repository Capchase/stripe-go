package stripe

import "encoding/json"

// IssuingTransactionType is the type of an issuing transaction.
type IssuingTransactionType string

// List of values that IssuingTransactionType can take.
const (
	IssuingTransactionTypeCapture IssuingTransactionType = "capture"
	IssuingTransactionTypeRefund  IssuingTransactionType = "refund"
)

// IssuingTransactionPurchaseDetailsFuelType is the type of fuel purchased in a transaction.
type IssuingTransactionPurchaseDetailsFuelType string

// List of values that IssuingTransactionType can take.
const (
	IssuingTransactionPurchaseDetailsFuelTypeDiesel          IssuingTransactionPurchaseDetailsFuelType = "diesel"
	IssuingTransactionPurchaseDetailsFuelTypeOther           IssuingTransactionPurchaseDetailsFuelType = "other"
	IssuingTransactionPurchaseDetailsFuelTypeUnleadedPlus    IssuingTransactionPurchaseDetailsFuelType = "unleaded_plus"
	IssuingTransactionPurchaseDetailsFuelTypeUnleadedRegular IssuingTransactionPurchaseDetailsFuelType = "unleaded_regular"
	IssuingTransactionPurchaseDetailsFuelTypeUnleadedSuper   IssuingTransactionPurchaseDetailsFuelType = "unleaded_super"
)

// IssuingTransactionPurchaseDetailsFuelUnit is the unit of fuel purchased in a transaction.
type IssuingTransactionPurchaseDetailsFuelUnit string

// List of values that IssuingTransactionPurchaseDetailsFuelUnit can take.
const (
	IssuingTransactionPurchaseDetailsFuelUnitLiter    IssuingTransactionPurchaseDetailsFuelUnit = "liter"
	IssuingTransactionPurchaseDetailsFuelUnitUSGallon IssuingTransactionPurchaseDetailsFuelUnit = "us_gallon"
)

// IssuingTransactionParams is the set of parameters that can be used when creating or updating an issuing transaction.
type IssuingTransactionParams struct {
	Params `form:"*"`
}

// IssuingTransactionListParams is the set of parameters that can be used when listing issuing transactions.
type IssuingTransactionListParams struct {
	ListParams   `form:"*"`
	Card         *string           `form:"card"`
	Cardholder   *string           `form:"cardholder"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Type         *string           `form:"type"`
}

// IssuingTransactionAmountDetails is the resource representing the breakdown of the amount.
type IssuingTransactionAmountDetails struct {
	ATMFee *int64 `json:"atm_fee,omitempty"`
}

// IssuingTransactionPurchaseDetailsFlightSegment contains extra information about the flight in this transaction.
type IssuingTransactionPurchaseDetailsFlightSegment struct {
	ArrivalAirportCode   *string `json:"arrival_airport_code,omitempty"`
	Carrier              *string `json:"carrier,omitempty"`
	DepartureAirportCode *string `json:"departure_airport_code,omitempty"`
	FlightNumber         *string `json:"flight_number,omitempty"`
	ServiceClass         *string `json:"service_class,omitempty"`
	StopoverAllowed      *bool `json:"stopover_allowed,omitempty"`
}

// IssuingTransactionPurchaseDetailsFlight contains extra information about the flight in this transaction.
type IssuingTransactionPurchaseDetailsFlight struct {
	DepartureAt   *int64 `json:"departure_at,omitempty"`
	PassengerName *string `json:"passenger_name,omitempty"`
	Refundable    *bool `json:"refundable,omitempty"`
	Segments      []*IssuingTransactionPurchaseDetailsFlightSegment `json:"segments,omitempty"`
	TravelAgency  *string `json:"travel_agency,omitempty"`
}

// IssuingTransactionPurchaseDetailsFuel contains extra information about the fuel purchase in this transaction.
type IssuingTransactionPurchaseDetailsFuel struct {
	Type            IssuingTransactionPurchaseDetailsFuelType `json:"type,omitempty"`
	Unit            IssuingTransactionPurchaseDetailsFuelUnit `json:"unit,omitempty"`
	UnitCostDecimal *float64 `json:"unit_cost_decimal,string"`
	VolumeDecimal   *float64 `json:"volume_decimal,string"`
}

// IssuingTransactionPurchaseDetailsLodging contains extra information about the lodging purchase in this transaction.
type IssuingTransactionPurchaseDetailsLodging struct {
	CheckInAt *int64 `json:"check_in_at,omitempty"`
	Nights    *int64 `json:"nights,omitempty"`
}

// IssuingTransactionPurchaseDetailsReceipt contains extra information about the line items this transaction.
type IssuingTransactionPurchaseDetailsReceipt struct {
	Description *string `json:"description,omitempty"`
	Quantity    *float64 `json:"quantity,omitempty"`
	Total       *int64 `json:"total,omitempty"`
	UnitCost    *int64 `json:"unit_cost,omitempty"`
}

// IssuingTransactionPurchaseDetails contains extra information provided by the merchant.
type IssuingTransactionPurchaseDetails struct {
	Flight    *IssuingTransactionPurchaseDetailsFlight    `json:"flight,omitempty"`
	Fuel      *IssuingTransactionPurchaseDetailsFuel      `json:"fuel,omitempty"`
	Lodging   *IssuingTransactionPurchaseDetailsLodging   `json:"lodging,omitempty"`
	Receipt   []*IssuingTransactionPurchaseDetailsReceipt `json:"receipt,omitempty"`
	Reference *string `json:"reference,omitempty"`
}

// IssuingTransaction is the resource representing a Stripe issuing transaction.
type IssuingTransaction struct {
	APIResource
	Amount             *int64 `json:"amount,omitempty"`
	AmountDetails      *IssuingTransactionAmountDetails   `json:"amount_details,omitempty"`
	Authorization      *IssuingAuthorization              `json:"authorization,omitempty"`
	BalanceTransaction *BalanceTransaction                `json:"balance_transaction,omitempty"`
	Card               *IssuingCard                       `json:"card,omitempty"`
	Cardholder         *IssuingCardholder                 `json:"cardholder,omitempty"`
	Created            *int64 `json:"created,omitempty"`
	Currency           Currency                           `json:"currency,omitempty"`
	Dispute            *IssuingDispute                    `json:"dispute,omitempty"`
	ID                 *string `json:"id,omitempty"`
	Livemode           *bool `json:"livemode,omitempty"`
	MerchantData       *IssuingAuthorizationMerchantData  `json:"merchant_data,omitempty"`
	MerchantAmount     *int64 `json:"merchant_amount,omitempty"`
	MerchantCurrency   Currency                           `json:"merchant_currency,omitempty"`
	Metadata           map[string]string                  `json:"metadata,omitempty"`
	Object             *string `json:"object,omitempty"`
	PurchaseDetails    *IssuingTransactionPurchaseDetails `json:"purchase_details,omitempty"`
	Type               IssuingTransactionType             `json:"type,omitempty"`
}

// IssuingTransactionList is a list of issuing transactions as retrieved from a list endpoint.
type IssuingTransactionList struct {
	APIResource
	ListMeta
	Data []*IssuingTransaction `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of an IssuingTransaction.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingTransaction) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = &id
		return nil
	}

	type issuingTransaction IssuingTransaction
	var v issuingTransaction
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingTransaction(v)
	return nil
}
