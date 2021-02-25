package stripe

import "encoding/json"

// List of values that BalanceTransactionStatus can take.
const (
	BalanceTransactionStatusAvailable BalanceTransactionStatus = "available"
	BalanceTransactionStatusPending   BalanceTransactionStatus = "pending"
)

// BalanceTransactionType is the list of allowed values for the balance transaction's type.
type BalanceTransactionType string

// List of values that BalanceTransactionType can take.
const (
	BalanceTransactionTypeAdjustment                      BalanceTransactionType = "adjustment"
	BalanceTransactionTypeAnticipationRepayment           BalanceTransactionType = "anticipation_repayment"
	BalanceTransactionTypeApplicationFee                  BalanceTransactionType = "application_fee"
	BalanceTransactionTypeApplicationFeeRefund            BalanceTransactionType = "application_fee_refund"
	BalanceTransactionTypeCharge                          BalanceTransactionType = "charge"
	BalanceTransactionTypeContribution                    BalanceTransactionType = "contribution"
	BalanceTransactionTypeIssuingAuthorizationHold        BalanceTransactionType = "issuing_authorization_hold"
	BalanceTransactionTypeIssuingAuthorizationRelease     BalanceTransactionType = "issuing_authorization_release"
	BalanceTransactionTypeIssuingAuthorizationDispute     BalanceTransactionType = "issuing_dispute"
	BalanceTransactionTypeIssuingAuthorizationTransaction BalanceTransactionType = "issuing_transaction"
	BalanceTransactionTypePayment                         BalanceTransactionType = "payment"
	BalanceTransactionTypePaymentFailureRefund            BalanceTransactionType = "payment_failure_refund"
	BalanceTransactionTypePaymentRefund                   BalanceTransactionType = "payment_refund"
	BalanceTransactionTypePayout                          BalanceTransactionType = "payout"
	BalanceTransactionTypePayoutCancel                    BalanceTransactionType = "payout_cancel"
	BalanceTransactionTypePayoutFailure                   BalanceTransactionType = "payout_failure"
	BalanceTransactionTypeRefund                          BalanceTransactionType = "refund"
	BalanceTransactionTypeStripeFee                       BalanceTransactionType = "stripe_fee"
	BalanceTransactionTypeTransfer                        BalanceTransactionType = "transfer"
	BalanceTransactionTypeTransferRefund                  BalanceTransactionType = "transfer_refund"
)

// BalanceTransactionSourceType consts represent valid balance transaction sources.
type BalanceTransactionSourceType string

// List of values that BalanceTransactionSourceType can take.
const (
	BalanceTransactionSourceTypeApplicationFee       BalanceTransactionSourceType = "application_fee"
	BalanceTransactionSourceTypeCharge               BalanceTransactionSourceType = "charge"
	BalanceTransactionSourceTypeDispute              BalanceTransactionSourceType = "dispute"
	BalanceTransactionSourceTypeIssuingAuthorization BalanceTransactionSourceType = "issuing.authorization"
	BalanceTransactionSourceTypeIssuingDispute       BalanceTransactionSourceType = "issuing.dispute"
	BalanceTransactionSourceTypeIssuingTransaction   BalanceTransactionSourceType = "issuing.transaction"
	BalanceTransactionSourceTypePayout               BalanceTransactionSourceType = "payout"
	BalanceTransactionSourceTypeRefund               BalanceTransactionSourceType = "refund"
	BalanceTransactionSourceTypeReversal             BalanceTransactionSourceType = "reversal"
	BalanceTransactionSourceTypeTransfer             BalanceTransactionSourceType = "transfer"
)

// BalanceTransactionReportingCategory represents reporting categories for balance transactions.
type BalanceTransactionReportingCategory string

// List of values that BalanceTransactionReportingCategory can take.
const (
	BalanceTransactionReportingCategoryAdvance                     BalanceTransactionReportingCategory = "advance"
	BalanceTransactionReportingCategoryAdvanceFunding              BalanceTransactionReportingCategory = "advance_funding"
	BalanceTransactionReportingCategoryCharge                      BalanceTransactionReportingCategory = "charge"
	BalanceTransactionReportingCategoryChargeFailure               BalanceTransactionReportingCategory = "charge_failure"
	BalanceTransactionReportingCategoryConnectCollectionTransfer   BalanceTransactionReportingCategory = "connect_collection_transfer"
	BalanceTransactionReportingCategoryConnectReservedFunds        BalanceTransactionReportingCategory = "connect_reserved_funds"
	BalanceTransactionReportingCategoryDispute                     BalanceTransactionReportingCategory = "dispute"
	BalanceTransactionReportingCategoryDisputeReversal             BalanceTransactionReportingCategory = "dispute_reversal"
	BalanceTransactionReportingCategoryFee                         BalanceTransactionReportingCategory = "fee"
	BalanceTransactionReportingCategoryIssuingAuthorizationHold    BalanceTransactionReportingCategory = "issuing_authorization_hold"
	BalanceTransactionReportingCategoryIssuingAuthorizationRelease BalanceTransactionReportingCategory = "issuing_authorization_release"
	BalanceTransactionReportingCategoryIssuingTransaction          BalanceTransactionReportingCategory = "issuing_transaction"
	BalanceTransactionReportingCategoryOtherAdjustment             BalanceTransactionReportingCategory = "other_adjustment"
	BalanceTransactionReportingCategoryPartialCaptureReversal      BalanceTransactionReportingCategory = "partial_capture_reversal"
	BalanceTransactionReportingCategoryPayout                      BalanceTransactionReportingCategory = "payout"
	BalanceTransactionReportingCategoryPayoutReversal              BalanceTransactionReportingCategory = "payout_reversal"
	BalanceTransactionReportingCategoryPlatformEarning             BalanceTransactionReportingCategory = "platform_earning"
	BalanceTransactionReportingCategoryPlatformEarningRefund       BalanceTransactionReportingCategory = "platform_earning_refund"
	BalanceTransactionReportingCategoryRefund                      BalanceTransactionReportingCategory = "refund"
	BalanceTransactionReportingCategoryRefundFailure               BalanceTransactionReportingCategory = "refund_failure"
	BalanceTransactionReportingCategoryRiskReservedFunds           BalanceTransactionReportingCategory = "risk_reserved_funds"
	BalanceTransactionReportingCategoryTax                         BalanceTransactionReportingCategory = "tax"
	BalanceTransactionReportingCategoryTopup                       BalanceTransactionReportingCategory = "topup"
	BalanceTransactionReportingCategoryTopupReversal               BalanceTransactionReportingCategory = "topup_reversal"
	BalanceTransactionReportingCategoryTransfer                    BalanceTransactionReportingCategory = "transfer"
	BalanceTransactionReportingCategoryTransferReversal            BalanceTransactionReportingCategory = "transfer_reversal"
)

// BalanceTransactionSource describes the source of a balance Transaction.
// The Type should indicate which object is fleshed out.
// For more details see https://stripe.com/docs/api#retrieve_balance_transaction
type BalanceTransactionSource struct {
	ApplicationFee       *ApplicationFee              `json:"-,omitempty"`
	Charge               *Charge                      `json:"-,omitempty"`
	Dispute              *Dispute                     `json:"-,omitempty"`
	ID                   *string `json:"id,omitempty"`
	IssuingAuthorization *IssuingAuthorization        `json:"-,omitempty"`
	IssuingDispute       *IssuingDispute              `json:"-,omitempty"`
	IssuingTransaction   *IssuingAuthorization        `json:"-,omitempty"`
	Payout               *Payout                      `json:"-,omitempty"`
	Refund               *Refund                      `json:"-,omitempty"`
	Reversal             *Reversal                    `json:"-,omitempty"`
	Transfer             *Transfer                    `json:"-,omitempty"`
	Type                 BalanceTransactionSourceType `json:"object,omitempty"`
}

// BalanceTransactionParams is the set of parameters that can be used when retrieving a transaction.
// For more details see https://stripe.com/docs/api#retrieve_balance_transaction.
type BalanceTransactionParams struct {
	Params `form:"*"`
}

// BalanceTransactionListParams is the set of parameters that can be used when listing balance transactions.
// For more details see https://stripe.com/docs/api/#balance_history.
type BalanceTransactionListParams struct {
	ListParams       `form:"*"`
	AvailableOn      *int64            `form:"available_on"`
	AvailableOnRange *RangeQueryParams `form:"available_on"`
	Created          *int64            `form:"created"`
	CreatedRange     *RangeQueryParams `form:"created"`
	Currency         *string           `form:"currency"`
	Payout           *string           `form:"payout"`
	Source           *string           `form:"source"`
	Type             *string           `form:"type"`
}

// BalanceTransaction is the resource representing the balance transaction.
// For more details see https://stripe.com/docs/api/#balance.
type BalanceTransaction struct {
	APIResource
	Amount            *int64 `json:"amount,omitempty"`
	AvailableOn       *int64 `json:"available_on,omitempty"`
	Created           *int64 `json:"created,omitempty"`
	Currency          Currency                            `json:"currency,omitempty"`
	Description       *string `json:"description,omitempty"`
	ExchangeRate      *float64 `json:"exchange_rate,omitempty"`
	ID                *string `json:"id,omitempty"`
	Fee               *int64 `json:"fee,omitempty"`
	FeeDetails        []*BalanceTransactionFee            `json:"fee_details,omitempty"`
	Net               *int64 `json:"net,omitempty"`
	ReportingCategory BalanceTransactionReportingCategory `json:"reporting_category,omitempty"`
	Source            *BalanceTransactionSource           `json:"source,omitempty"`
	Status            BalanceTransactionStatus            `json:"status,omitempty"`
	Type              BalanceTransactionType              `json:"type,omitempty"`
}

// BalanceTransactionList is a list of transactions as returned from a list endpoint.
type BalanceTransactionList struct {
	APIResource
	ListMeta
	Data []*BalanceTransaction `json:"data,omitempty"`
}

// BalanceTransactionFee is a structure that breaks down the fees in a transaction.
type BalanceTransactionFee struct {
	Amount      *int64 `json:"amount,omitempty"`
	Application *string `json:"application,omitempty"`
	Currency    Currency `json:"currency,omitempty"`
	Description *string `json:"description,omitempty"`
	Type        *string `json:"type,omitempty"`
}

// UnmarshalJSON handles deserialization of a Transaction.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *BalanceTransaction) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = &id
		return nil
	}

	type balanceTransaction BalanceTransaction
	var v balanceTransaction
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = BalanceTransaction(v)
	return nil
}

// UnmarshalJSON handles deserialization of a BalanceTransactionSource.
// This custom unmarshaling is needed because the specific
// type of transaction source it refers to is specified in the JSON
func (s *BalanceTransactionSource) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = &id
		return nil
	}

	type balanceTransactionSource BalanceTransactionSource
	var v balanceTransactionSource
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var err error
	*s = BalanceTransactionSource(v)

	switch s.Type {
	case BalanceTransactionSourceTypeApplicationFee:
		err = json.Unmarshal(data, &s.ApplicationFee)
	case BalanceTransactionSourceTypeCharge:
		err = json.Unmarshal(data, &s.Charge)
	case BalanceTransactionSourceTypeDispute:
		err = json.Unmarshal(data, &s.Dispute)
	case BalanceTransactionSourceTypeIssuingAuthorization:
		err = json.Unmarshal(data, &s.IssuingAuthorization)
	case BalanceTransactionSourceTypeIssuingDispute:
		err = json.Unmarshal(data, &s.IssuingDispute)
	case BalanceTransactionSourceTypeIssuingTransaction:
		err = json.Unmarshal(data, &s.IssuingTransaction)
	case BalanceTransactionSourceTypePayout:
		err = json.Unmarshal(data, &s.Payout)
	case BalanceTransactionSourceTypeRefund:
		err = json.Unmarshal(data, &s.Refund)
	case BalanceTransactionSourceTypeReversal:
		err = json.Unmarshal(data, &s.Reversal)
	case BalanceTransactionSourceTypeTransfer:
		err = json.Unmarshal(data, &s.Transfer)
	}

	return err
}
