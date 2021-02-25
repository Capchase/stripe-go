package stripe

import "encoding/json"

// TransferSourceType is the list of allowed values for the transfer's source_type field.
type TransferSourceType string

// List of values that TransferSourceType can take.
const (
	TransferSourceTypeBankAccount TransferSourceType = "bank_account"
	TransferSourceTypeCard        TransferSourceType = "card"
	TransferSourceTypeFPX         TransferSourceType = "fpx"
)

// TransferDestination describes the destination of a Transfer.
// The Type should indicate which object is fleshed out
// For more details see https://stripe.com/docs/api/go#transfer_object
type TransferDestination struct {
	Account *Account `json:"-,omitempty"`
	ID      *string `json:"id,omitempty"`
}

// TransferParams is the set of parameters that can be used when creating or updating a transfer.
// For more details see https://stripe.com/docs/api#create_transfer and https://stripe.com/docs/api#update_transfer.
type TransferParams struct {
	Params            `form:"*"`
	Amount            *int64  `form:"amount"`
	Currency          *string `form:"currency"`
	Description       *string `form:"description"`
	Destination       *string `form:"destination"`
	SourceTransaction *string `form:"source_transaction"`
	SourceType        *string `form:"source_type"`
	TransferGroup     *string `form:"transfer_group"`
}

// TransferListParams is the set of parameters that can be used when listing transfers.
// For more details see https://stripe.com/docs/api#list_transfers.
type TransferListParams struct {
	ListParams    `form:"*"`
	Created       *int64            `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	Destination   *string           `form:"destination"`
	TransferGroup *string           `form:"transfer_group"`
}

// Transfer is the resource representing a Stripe transfer.
// For more details see https://stripe.com/docs/api#transfers.
type Transfer struct {
	APIResource
	Amount             *int64 `json:"amount,omitempty"`
	AmountReversed     *int64 `json:"amount_reversed,omitempty"`
	BalanceTransaction *BalanceTransaction       `json:"balance_transaction,omitempty"`
	Created            *int64 `json:"created,omitempty"`
	Currency           Currency                  `json:"currency,omitempty"`
	Description        *string `json:"description,omitempty"`
	Destination        *TransferDestination      `json:"destination,omitempty"`
	DestinationPayment *Charge                   `json:"destination_payment,omitempty"`
	ID                 *string `json:"id,omitempty"`
	Livemode           *bool `json:"livemode,omitempty"`
	Metadata           map[string]string         `json:"metadata,omitempty"`
	Reversals          *ReversalList             `json:"reversals,omitempty"`
	Reversed           *bool `json:"reversed,omitempty"`
	SourceTransaction  *BalanceTransactionSource `json:"source_transaction,omitempty"`
	SourceType         TransferSourceType        `json:"source_type,omitempty"`
	TransferGroup      *string `json:"transfer_group,omitempty"`
}

// TransferList is a list of transfers as retrieved from a list endpoint.
type TransferList struct {
	APIResource
	ListMeta
	Data []*Transfer `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of a Transfer.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *Transfer) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = &id
		return nil
	}

	type transfer Transfer
	var v transfer
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = Transfer(v)
	return nil
}

// UnmarshalJSON handles deserialization of a TransferDestination.
// This custom unmarshaling is needed because the specific
// type of destination it refers to is specified in the JSON
func (d *TransferDestination) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		d.ID = &id
		return nil
	}

	type transferDestination TransferDestination
	var v transferDestination
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*d = TransferDestination(v)
	return json.Unmarshal(data, &d.Account)
}
