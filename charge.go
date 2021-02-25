package stripe

import (
	"encoding/json"
)

// ChargeFraudUserReport is the list of allowed values for reporting fraud.
type ChargeFraudUserReport string

// List of values that ChargeFraudUserReport can take.
const (
	ChargeFraudUserReportFraudulent ChargeFraudUserReport = "fraudulent"
	ChargeFraudUserReportSafe       ChargeFraudUserReport = "safe"
)

// ChargeFraudStripeReport is the list of allowed values for reporting fraud.
type ChargeFraudStripeReport string

// List of values that ChargeFraudStripeReport can take.
const (
	ChargeFraudStripeReportFraudulent ChargeFraudStripeReport = "fraudulent"
)

// ChargePaymentMethodDetailsCardPresentReceiptAccountType indicates the type of account backing a card present transaction.
type ChargePaymentMethodDetailsCardPresentReceiptAccountType string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResult can take.
const (
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeChecking ChargePaymentMethodDetailsCardPresentReceiptAccountType = "checking"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeCredit   ChargePaymentMethodDetailsCardPresentReceiptAccountType = "credit"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypePrepaid  ChargePaymentMethodDetailsCardPresentReceiptAccountType = "prepaid"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeUnknown  ChargePaymentMethodDetailsCardPresentReceiptAccountType = "unknown"
)

// ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow indicates the type of 3D Secure
// authentication performed.
type ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow can take.
const (
	ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlowChallenge    ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "challenge"
	ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlowFrictionless ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "frictionless"
)

// ChargePaymentMethodDetailsCardThreeDSecureResult indicates the outcome of 3D Secure authentication.
type ChargePaymentMethodDetailsCardThreeDSecureResult string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResult can take.
const (
	ChargePaymentMethodDetailsCardThreeDSecureResultAttemptAcknowledged ChargePaymentMethodDetailsCardThreeDSecureResult = "attempt_acknowledged"
	ChargePaymentMethodDetailsCardThreeDSecureResultAuthenticated       ChargePaymentMethodDetailsCardThreeDSecureResult = "authenticated"
	ChargePaymentMethodDetailsCardThreeDSecureResultFailed              ChargePaymentMethodDetailsCardThreeDSecureResult = "failed"
	ChargePaymentMethodDetailsCardThreeDSecureResultNotSupported        ChargePaymentMethodDetailsCardThreeDSecureResult = "not_supported"
	ChargePaymentMethodDetailsCardThreeDSecureResultProcessingError     ChargePaymentMethodDetailsCardThreeDSecureResult = "processing_error"
)

// ChargePaymentMethodDetailsCardThreeDSecureResultReason represents dditional information about why
// 3D Secure succeeded or failed
type ChargePaymentMethodDetailsCardThreeDSecureResultReason string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResultReason can take.
const (
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonAbandoned           ChargePaymentMethodDetailsCardThreeDSecureResultReason = "abandoned"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonBypassed            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "bypassed"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonCanceled            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "canceled"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonCardNotEnrolled     ChargePaymentMethodDetailsCardThreeDSecureResultReason = "card_not_enrolled"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonNetworkNotSupported ChargePaymentMethodDetailsCardThreeDSecureResultReason = "network_not_supported"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonProtocolError       ChargePaymentMethodDetailsCardThreeDSecureResultReason = "protocol_error"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonRejected            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "rejected"
)

// ChargePaymentMethodDetailsType is the type of the PaymentMethod associated with the Charge's
// payment method details.
type ChargePaymentMethodDetailsType string

// List of values that ChargePaymentMethodDetailsType can take.
const (
	ChargePaymentMethodDetailsTypeAchCreditTransfer ChargePaymentMethodDetailsType = "ach_credit_transfer"
	ChargePaymentMethodDetailsTypeAchDebit          ChargePaymentMethodDetailsType = "ach_debit"
	ChargePaymentMethodDetailsTypeAcssDebit         ChargePaymentMethodDetailsType = "acss_debit"
	ChargePaymentMethodDetailsTypeAlipay            ChargePaymentMethodDetailsType = "alipay"
	ChargePaymentMethodDetailsTypeAUBECSDebit       ChargePaymentMethodDetailsType = "au_becs_debit"
	ChargePaymentMethodDetailsTypeBACSDebit         ChargePaymentMethodDetailsType = "bacs_debit"
	ChargePaymentMethodDetailsTypeBancontact        ChargePaymentMethodDetailsType = "bancontact"
	ChargePaymentMethodDetailsTypeCard              ChargePaymentMethodDetailsType = "card"
	ChargePaymentMethodDetailsTypeCardPresent       ChargePaymentMethodDetailsType = "card_present"
	ChargePaymentMethodDetailsTypeEps               ChargePaymentMethodDetailsType = "eps"
	ChargePaymentMethodDetailsTypeFPX               ChargePaymentMethodDetailsType = "fpx"
	ChargePaymentMethodDetailsTypeGiropay           ChargePaymentMethodDetailsType = "giropay"
	ChargePaymentMethodDetailsTypeGrabpay           ChargePaymentMethodDetailsType = "grabpay"
	ChargePaymentMethodDetailsTypeIdeal             ChargePaymentMethodDetailsType = "ideal"
	ChargePaymentMethodDetailsTypeInteracPresent    ChargePaymentMethodDetailsType = "interac_present"
	ChargePaymentMethodDetailsTypeKlarna            ChargePaymentMethodDetailsType = "klarna"
	ChargePaymentMethodDetailsTypeMultibanco        ChargePaymentMethodDetailsType = "multibanco"
	ChargePaymentMethodDetailsTypeP24               ChargePaymentMethodDetailsType = "p24"
	ChargePaymentMethodDetailsTypeSepaDebit         ChargePaymentMethodDetailsType = "sepa_debit"
	ChargePaymentMethodDetailsTypeSofort            ChargePaymentMethodDetailsType = "sofort"
	ChargePaymentMethodDetailsTypeStripeAccount     ChargePaymentMethodDetailsType = "stripe_account"
	ChargePaymentMethodDetailsTypeWechat            ChargePaymentMethodDetailsType = "wechat"
)

// ChargeLevel3LineItemsParams is the set of parameters that represent a line item on level III data.
type ChargeLevel3LineItemsParams struct {
	DiscountAmount     *int64  `form:"discount_amount"`
	ProductCode        *string `form:"product_code"`
	ProductDescription *string `form:"product_description"`
	Quantity           *int64  `form:"quantity"`
	TaxAmount          *int64  `form:"tax_amount"`
	UnitCost           *int64  `form:"unit_cost"`
}

// ChargeLevel3Params is the set of parameters that can be used for the Level III data.
type ChargeLevel3Params struct {
	CustomerReference  *string                        `form:"customer_reference"`
	LineItems          []*ChargeLevel3LineItemsParams `form:"line_items"`
	MerchantReference  *string                        `form:"merchant_reference"`
	ShippingAddressZip *string                        `form:"shipping_address_zip"`
	ShippingFromZip    *string                        `form:"shipping_from_zip"`
	ShippingAmount     *int64                         `form:"shipping_amount"`
}

// ChargeTransferDataParams is the set of parameters allowed for the transfer_data hash.
type ChargeTransferDataParams struct {
	Amount *int64 `form:"amount"`
	// This parameter can only be used on Charge creation.
	Destination *string `form:"destination"`
}

// ChargeParams is the set of parameters that can be used when creating or updating a charge.
type ChargeParams struct {
	Params                    `form:"*"`
	Amount                    *int64                    `form:"amount"`
	ApplicationFeeAmount      *int64                    `form:"application_fee_amount"`
	Capture                   *bool                     `form:"capture"`
	Currency                  *string                   `form:"currency"`
	Customer                  *string                   `form:"customer"`
	Description               *string                   `form:"description"`
	Destination               *DestinationParams        `form:"destination"`
	ExchangeRate              *float64                  `form:"exchange_rate"`
	FraudDetails              *FraudDetailsParams       `form:"fraud_details"`
	Level3                    *ChargeLevel3Params       `form:"level3"`
	OnBehalfOf                *string                   `form:"on_behalf_of"`
	ReceiptEmail              *string                   `form:"receipt_email"`
	Shipping                  *ShippingDetailsParams    `form:"shipping"`
	Source                    *SourceParams             `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	StatementDescriptor       *string                   `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                   `form:"statement_descriptor_suffix"`
	TransferData              *ChargeTransferDataParams `form:"transfer_data"`
	TransferGroup             *string                   `form:"transfer_group"`
}

// ShippingDetailsParams is the structure containing shipping information as parameters
type ShippingDetailsParams struct {
	Address        *AddressParams `form:"address"`
	Carrier        *string        `form:"carrier"`
	Name           *string        `form:"name"`
	Phone          *string        `form:"phone"`
	TrackingNumber *string        `form:"tracking_number"`
}

// SetSource adds valid sources to a ChargeParams object,
// returning an error for unsupported sources.
func (p *ChargeParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
}

// DestinationParams describes the parameters available for the destination hash when creating a charge.
type DestinationParams struct {
	Account *string `form:"account"`
	Amount  *int64  `form:"amount"`
}

// FraudDetailsParams provides information on the fraud details for a charge.
type FraudDetailsParams struct {
	UserReport *string `form:"user_report"`
}

// ChargeListParams is the set of parameters that can be used when listing charges.
type ChargeListParams struct {
	ListParams    `form:"*"`
	Created       *int64            `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	Customer      *string           `form:"customer"`
	PaymentIntent *string           `form:"payment_intent"`
	TransferGroup *string           `form:"transfer_group"`
}

// CaptureParams is the set of parameters that can be used when capturing a charge.
type CaptureParams struct {
	Params                    `form:"*"`
	Amount                    *int64                    `form:"amount"`
	ApplicationFeeAmount      *int64                    `form:"application_fee_amount"`
	ExchangeRate              *float64                  `form:"exchange_rate"`
	ReceiptEmail              *string                   `form:"receipt_email"`
	StatementDescriptor       *string                   `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                   `form:"statement_descriptor_suffix"`
	TransferGroup             *string                   `form:"transfer_group"`
	TransferData              *ChargeTransferDataParams `form:"transfer_data"`
}

// ChargeLevel3LineItem represents a line item on level III data.
// This is in private beta and would be empty for most integrations
type ChargeLevel3LineItem struct {
	DiscountAmount     *int64 `json:"discount_amount,omitempty"`
	ProductCode        *string `json:"product_code,omitempty"`
	ProductDescription *string `json:"product_description,omitempty"`
	Quantity           *int64 `json:"quantity,omitempty"`
	TaxAmount          *int64 `json:"tax_amount,omitempty"`
	UnitCost           *int64 `json:"unit_cost,omitempty"`
}

// ChargeLevel3 represents the Level III data.
// This is in private beta and would be empty for most integrations
type ChargeLevel3 struct {
	CustomerReference  *string `json:"customer_reference,omitempty"`
	LineItems          []*ChargeLevel3LineItem `json:"line_items,omitempty"`
	MerchantReference  *string `json:"merchant_reference,omitempty"`
	ShippingAddressZip *string `json:"shipping_address_zip,omitempty"`
	ShippingFromZip    *string `json:"shipping_from_zip,omitempty"`
	ShippingAmount     *int64 `json:"shipping_amount,omitempty"`
}

// ChargePaymentMethodDetailsAchCreditTransfer represents details about the ACH Credit Transfer
// PaymentMethod.
type ChargePaymentMethodDetailsAchCreditTransfer struct {
	AccountNumber *string `json:"account_number,omitempty"`
	BankName      *string `json:"bank_name,omitempty"`
	RoutingNumber *string `json:"routing_number,omitempty"`
	SwiftCode     *string `json:"swift_code,omitempty"`
}

// ChargePaymentMethodDetailsAchDebit represents details about the ACH Debit PaymentMethod.
type ChargePaymentMethodDetailsAchDebit struct {
	AccountHolderType BankAccountAccountHolderType `json:"account_holder_type,omitempty"`
	BankName          *string `json:"bank_name,omitempty"`
	Country           *string `json:"country,omitempty"`
	Fingerprint       *string `json:"fingerprint,omitempty"`
	Last4             *string `json:"last4,omitempty"`
	RoutingNumber     *string `json:"routing_number,omitempty"`
}

// ChargePaymentMethodDetailsAcssDebit represents details about the ACSS Debit PaymentMethod.
type ChargePaymentMethodDetailsAcssDebit struct {
	BankName          *string `json:"bank_name,omitempty"`
	Fingerprint       *string `json:"fingerprint,omitempty"`
	InstitutionNumber *string `json:"institution_number,omitempty"`
	Last4             *string `json:"last4,omitempty"`
	Mandate           *string `json:"mandate,omitempty"`
	TransitNumber     *string `json:"transit_number,omitempty"`
}

// ChargePaymentMethodDetailsAfterpayClearpay represents details about the AfterpayClearpay PaymentMethod.
type ChargePaymentMethodDetailsAfterpayClearpay struct{}

// ChargePaymentMethodDetailsAlipay represents details about the Alipay PaymentMethod.
type ChargePaymentMethodDetailsAlipay struct {
	Fingerprint   *string `json:"fingerprint,omitempty"`
	TransactionID *string `json:"transaction_id,omitempty"`
}

// ChargePaymentMethodDetailsAUBECSDebit represents details about the AU BECS DD PaymentMethod.
type ChargePaymentMethodDetailsAUBECSDebit struct {
	BSBNumber   *string `json:"bsb_number,omitempty"`
	Fingerprint *string `json:"fingerprint,omitempty"`
	Last4       *string `json:"last4,omitempty"`
	Mandate     *string `json:"mandate,omitempty"`
}

// ChargePaymentMethodDetailsBACSDebit represents details about the BECS Debit PaymentMethod.
type ChargePaymentMethodDetailsBACSDebit struct {
	Fingerprint *string `json:"fingerprint,omitempty"`
	Last4       *string `json:"last4,omitempty"`
	Mandate     *string `json:"mandate,omitempty"`
	SortCode    *string `json:"sort_code,omitempty"`
}

// ChargePaymentMethodDetailsBancontact represents details about the Bancontact PaymentMethod.
type ChargePaymentMethodDetailsBancontact struct {
	BankCode                  *string `json:"bank_code,omitempty"`
	BankName                  *string `json:"bank_name,omitempty"`
	Bic                       *string `json:"bic,omitempty"`
	IbanLast4                 *string `json:"iban_last4,omitempty"`
	PreferredLanguage         *string `json:"preferred_language,omitempty"`
	VerifiedName              *string `json:"verified_name,omitempty"`
	GeneratedSepaDebit        *PaymentMethod `json:"generated_sepa_debit,omitempty"`
	GeneratedSepaDebitMandate *Mandate       `json:"generated_sepa_debit_mandate,omitempty"`
}

// ChargePaymentMethodDetailsCardChecks represents the checks associated with the charge's Card
// PaymentMethod.
type ChargePaymentMethodDetailsCardChecks struct {
	AddressLine1Check      CardVerification `json:"address_line1_check,omitempty"`
	AddressPostalCodeCheck CardVerification `json:"address_postal_code_check,omitempty"`
	CVCCheck               CardVerification `json:"cvc_check,omitempty"`
}

// ChargePaymentMethodDetailsCardInstallments represents details about the installment plan chosen
// for this charge.
type ChargePaymentMethodDetailsCardInstallments struct {
	Plan *PaymentIntentPaymentMethodOptionsCardInstallmentsPlan `json:"plan,omitempty"`
}

// ChargePaymentMethodDetailsCardThreeDSecure represents details about 3DS associated with the
// charge's PaymentMethod.
type ChargePaymentMethodDetailsCardThreeDSecure struct {
	AuthenticationFlow ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow `json:"authentication_flow,omitempty"`
	Result             ChargePaymentMethodDetailsCardThreeDSecureResult             `json:"result,omitempty"`
	ResultReason       ChargePaymentMethodDetailsCardThreeDSecureResultReason       `json:"result_reason,omitempty"`
	Version            *string `json:"version,omitempty"`
}

// ChargePaymentMethodDetailsCardWalletAmexExpressCheckout represents the details of the Amex
// Express Checkout wallet.
type ChargePaymentMethodDetailsCardWalletAmexExpressCheckout struct {
}

// ChargePaymentMethodDetailsCardWalletApplePay represents the details of the Apple Pay wallet.
type ChargePaymentMethodDetailsCardWalletApplePay struct {
}

// ChargePaymentMethodDetailsCardWalletGooglePay represents the details of the Google Pay wallet.
type ChargePaymentMethodDetailsCardWalletGooglePay struct {
}

// ChargePaymentMethodDetailsCardWalletMasterpass represents the details of the Masterpass wallet.
type ChargePaymentMethodDetailsCardWalletMasterpass struct {
	BillingAddress  *Address `json:"billing_address,omitempty"`
	Email           *string `json:"email,omitempty"`
	Name            *string `json:"name,omitempty"`
	ShippingAddress *Address `json:"shipping_address,omitempty"`
}

// ChargePaymentMethodDetailsCardWalletSamsungPay represents the details of the Samsung Pay wallet.
type ChargePaymentMethodDetailsCardWalletSamsungPay struct {
}

// ChargePaymentMethodDetailsCardWalletVisaCheckout represents the details of the Visa Checkout
// wallet.
type ChargePaymentMethodDetailsCardWalletVisaCheckout struct {
	BillingAddress  *Address `json:"billing_address,omitempty"`
	Email           *string `json:"email,omitempty"`
	Name            *string `json:"name,omitempty"`
	ShippingAddress *Address `json:"shipping_address,omitempty"`
}

// ChargePaymentMethodDetailsCardWallet represents the details of the card wallet if this Card
// PaymentMethod is part of a card wallet.
type ChargePaymentMethodDetailsCardWallet struct {
	AmexExpressCheckout *ChargePaymentMethodDetailsCardWalletAmexExpressCheckout `json:"amex_express_checkout,omitempty"`
	ApplePay            *ChargePaymentMethodDetailsCardWalletApplePay            `json:"apple_pay,omitempty"`
	DynamicLast4        *string `json:"dynamic_last4,omitempty"`
	GooglePay           *ChargePaymentMethodDetailsCardWalletGooglePay           `json:"google_pay,omitempty"`
	Masterpass          *ChargePaymentMethodDetailsCardWalletMasterpass          `json:"masterpass,omitempty"`
	SamsungPay          *ChargePaymentMethodDetailsCardWalletSamsungPay          `json:"samsung_pay,omitempty"`
	Type                PaymentMethodCardWalletType                              `json:"type,omitempty"`
	VisaCheckout        *ChargePaymentMethodDetailsCardWalletVisaCheckout        `json:"visa_checkout,omitempty"`
}

// ChargePaymentMethodDetailsCard represents details about the Card PaymentMethod.
type ChargePaymentMethodDetailsCard struct {
	Brand        PaymentMethodCardBrand                      `json:"brand,omitempty"`
	Checks       *ChargePaymentMethodDetailsCardChecks       `json:"checks,omitempty"`
	Country      *string `json:"country,omitempty"`
	ExpMonth     *uint64 `json:"exp_month,omitempty"`
	ExpYear      *uint64 `json:"exp_year,omitempty"`
	Fingerprint  *string `json:"fingerprint,omitempty"`
	Funding      CardFunding                                 `json:"funding,omitempty"`
	Installments *ChargePaymentMethodDetailsCardInstallments `json:"installments,omitempty"`
	Last4        *string `json:"last4,omitempty"`
	Network      PaymentMethodCardNetwork                    `json:"network,omitempty"`
	MOTO         *bool `json:"moto,omitempty"`
	ThreeDSecure *ChargePaymentMethodDetailsCardThreeDSecure `json:"three_d_secure,omitempty"`
	Wallet       *ChargePaymentMethodDetailsCardWallet       `json:"wallet,omitempty"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description *string `json:"description,omitempty"`
	IIN         *string `json:"iin,omitempty"`
	Issuer      *string `json:"issuer,omitempty"`
}

// ChargePaymentMethodDetailsCardPresentReceipt represents details about the receipt on a
// Card Present PaymentMethod.
type ChargePaymentMethodDetailsCardPresentReceipt struct {
	AccountType                  ChargePaymentMethodDetailsCardPresentReceiptAccountType `json:"account_type,omitempty"`
	ApplicationCryptogram        *string `json:"application_cryptogram,omitempty"`
	ApplicationPreferredName     *string `json:"application_preferred_name,omitempty"`
	AuthorizationCode            *string `json:"authorization_code,omitempty"`
	AuthorizationResponseCode    *string `json:"authorization_response_code,omitempty"`
	CardholderVerificationMethod *string `json:"cardholder_verification_method,omitempty"`
	DedicatedFileName            *string `json:"dedicated_file_name,omitempty"`
	TerminalVerificationResults  *string `json:"terminal_verification_results,omitempty"`
	TransactionStatusInformation *string `json:"transaction_status_information,omitempty"`
}

// ChargePaymentMethodDetailsCardPresent represents details about the Card Present PaymentMethod.
type ChargePaymentMethodDetailsCardPresent struct {
	Brand          PaymentMethodCardBrand                        `json:"brand,omitempty"`
	CardholderName *string `json:"cardholder_name,omitempty"`
	Country        *string `json:"country,omitempty"`
	EmvAuthData    *string `json:"emv_auth_data,omitempty"`
	ExpMonth       *uint64 `json:"exp_month,omitempty"`
	ExpYear        *uint64 `json:"exp_year,omitempty"`
	Fingerprint    *string `json:"fingerprint,omitempty"`
	Funding        CardFunding                                   `json:"funding,omitempty"`
	GeneratedCard  *string `json:"generated_card,omitempty"`
	Last4          *string `json:"last4,omitempty"`
	Network        PaymentMethodCardNetwork                      `json:"network,omitempty"`
	ReadMethod     *string `json:"read_method,omitempty"`
	Receipt        *ChargePaymentMethodDetailsCardPresentReceipt `json:"receipt,omitempty"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description *string `json:"description,omitempty"`
	IIN         *string `json:"iin,omitempty"`
	Issuer      *string `json:"issuer,omitempty"`
}

// ChargePaymentMethodDetailsEps represents details about the EPS PaymentMethod.
type ChargePaymentMethodDetailsEps struct {
	Bank         *string `json:"bank,omitempty"`
	VerifiedName *string `json:"verified_name,omitempty"`
}

// ChargePaymentMethodDetailsFPX represents details about the FPX PaymentMethod.
type ChargePaymentMethodDetailsFPX struct {
	AccountHolderType PaymentMethodFPXAccountHolderType `json:"account_holder_type,omitempty"`
	Bank              *string `json:"bank,omitempty"`
	TransactionID     *string `json:"transaction_id,omitempty"`
}

// ChargePaymentMethodDetailsGiropay represents details about the Giropay PaymentMethod.
type ChargePaymentMethodDetailsGiropay struct {
	BankCode     *string `json:"bank_code,omitempty"`
	BankName     *string `json:"bank_name,omitempty"`
	Bic          *string `json:"bic,omitempty"`
	VerifiedName *string `json:"verified_name,omitempty"`
}

// ChargePaymentMethodDetailsGrabpay represents details about the Grabpay PaymentMethod.
type ChargePaymentMethodDetailsGrabpay struct {
	TransactionID *string `json:"transaction_id,omitempty"`
}

// ChargePaymentMethodDetailsIdeal represents details about the Ideal PaymentMethod.
type ChargePaymentMethodDetailsIdeal struct {
	Bank                      *string `json:"bank,omitempty"`
	Bic                       *string `json:"bic,omitempty"`
	IbanLast4                 *string `json:"iban_last4,omitempty"`
	VerifiedName              *string `json:"verified_name,omitempty"`
	GeneratedSepaDebit        *PaymentMethod `json:"generated_sepa_debit,omitempty"`
	GeneratedSepaDebitMandate *Mandate       `json:"generated_sepa_debit_mandate,omitempty"`
}

// ChargePaymentMethodDetailsInteracPresent represents details about the InteracPresent PaymentMethod.
type ChargePaymentMethodDetailsInteracPresent struct {
	Brand            *string `json:"brand,omitempty"`
	CardholderName   *string `json:"cardholder_name,omitempty"`
	Country          *string `json:"country,omitempty"`
	EmvAuthData      *string `json:"emv_auth_data,omitempty"`
	ExpMonth         *int64 `json:"exp_month,omitempty"`
	ExpYear          *int64 `json:"exp_year,omitempty"`
	Fingerprint      *string `json:"fingerprint,omitempty"`
	Funding          *string `json:"funding,omitempty"`
	GeneratedCard    *string `json:"generated_card,omitempty"`
	Last4            *string `json:"last4,omitempty"`
	Network          *string `json:"network,omitempty"`
	PreferredLocales []string                                         `json:"preferred_locales,omitempty"`
	ReadMethod       *string `json:"read_method,omitempty"`
	Receipt          *ChargePaymentMethodDetailsInteracPresentReceipt `json:"receipt,omitempty"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description *string `json:"description,omitempty"`
	IIN         *string `json:"iin,omitempty"`
	Issuer      *string `json:"issuer,omitempty"`
}

// ChargePaymentMethodDetailsInteracPresentReceipt represents details about the InteracPresent Receipt.
type ChargePaymentMethodDetailsInteracPresentReceipt struct {
	AccountType                  *string `json:"account_type,omitempty"`
	ApplicationCryptogram        *string `json:"application_cryptogram,omitempty"`
	ApplicationPreferredName     *string `json:"application_preferred_name,omitempty"`
	AuthorizationCode            *string `json:"authorization_code,omitempty"`
	AuthorizationResponseCode    *string `json:"authorization_response_code,omitempty"`
	CardholderVerificationMethod *string `json:"cardholder_verification_method,omitempty"`
	DedicatedFileName            *string `json:"dedicated_file_name,omitempty"`
	TerminalVerificationResults  *string `json:"terminal_verification_results,omitempty"`
	TransactionStatusInformation *string `json:"transaction_status_information,omitempty"`
}

// ChargePaymentMethodDetailsKlarna represents details for the Klarna
// PaymentMethod.
type ChargePaymentMethodDetailsKlarna struct {
}

// ChargePaymentMethodDetailsMultibanco represents details about the Multibanco PaymentMethod.
type ChargePaymentMethodDetailsMultibanco struct {
	Entity    *string `json:"entity,omitempty"`
	Reference *string `json:"reference,omitempty"`
}

// ChargePaymentMethodDetailsOXXO represents details about the OXXO PaymentMethod.
type ChargePaymentMethodDetailsOXXO struct {
	Number *string `json:"number,omitempty"`
}

// ChargePaymentMethodDetailsP24 represents details about the P24 PaymentMethod.
type ChargePaymentMethodDetailsP24 struct {
	Bank         *string `json:"bank,omitempty"`
	Reference    *string `json:"reference,omitempty"`
	VerifiedName *string `json:"verified_name,omitempty"`
}

// ChargePaymentMethodDetailsSepaDebit represents details about the Sepa Debit PaymentMethod.
type ChargePaymentMethodDetailsSepaDebit struct {
	BankCode    *string `json:"bank_code,omitempty"`
	BranchCode  *string `json:"branch_code,omitempty"`
	Country     *string `json:"country,omitempty"`
	Fingerprint *string `json:"fingerprint,omitempty"`
	Last4       *string `json:"last4,omitempty"`
	Mandate     *Mandate `json:"mandate,omitempty"`
}

// ChargePaymentMethodDetailsSofort represents details about the Sofort PaymentMethod.
type ChargePaymentMethodDetailsSofort struct {
	BankCode                  *string `json:"bank_code,omitempty"`
	BankName                  *string `json:"bank_name,omitempty"`
	Bic                       *string `json:"bic,omitempty"`
	Country                   *string `json:"country,omitempty"`
	IbanLast4                 *string `json:"iban_last4,omitempty"`
	VerifiedName              *string `json:"verified_name,omitempty"`
	GeneratedSepaDebit        *PaymentMethod `json:"generated_sepa_debit,omitempty"`
	GeneratedSepaDebitMandate *Mandate       `json:"generated_sepa_debit_mandate,omitempty"`
}

// ChargePaymentMethodDetailsStripeAccount represents details about the StripeAccount PaymentMethod.
type ChargePaymentMethodDetailsStripeAccount struct {
}

// ChargePaymentMethodDetailsWechat represents details about the Wechat PaymentMethod.
type ChargePaymentMethodDetailsWechat struct {
}

// ChargePaymentMethodDetails represents the details about the PaymentMethod associated with the
// charge.
type ChargePaymentMethodDetails struct {
	AchCreditTransfer *ChargePaymentMethodDetailsAchCreditTransfer `json:"ach_credit_transfer,omitempty"`
	AchDebit          *ChargePaymentMethodDetailsAchDebit          `json:"ach_debit,omitempty"`
	AcssDebit         *ChargePaymentMethodDetailsAcssDebit         `json:"acss_debit,omitempty"`
	AfterpayClearpay  *ChargePaymentMethodDetailsAfterpayClearpay  `json:"afterpay_clearpay,omitempty"`
	Alipay            *ChargePaymentMethodDetailsAlipay            `json:"alipay,omitempty"`
	AUBECSDebit       *ChargePaymentMethodDetailsAUBECSDebit       `json:"au_becs_debit,omitempty"`
	BACSDebit         *ChargePaymentMethodDetailsBACSDebit         `json:"bacs_debit,omitempty"`
	Bancontact        *ChargePaymentMethodDetailsBancontact        `json:"bancontact,omitempty"`
	Card              *ChargePaymentMethodDetailsCard              `json:"card,omitempty"`
	CardPresent       *ChargePaymentMethodDetailsCardPresent       `json:"card_present,omitempty"`
	Eps               *ChargePaymentMethodDetailsEps               `json:"eps,omitempty"`
	FPX               *ChargePaymentMethodDetailsFPX               `json:"fpx,omitempty"`
	Giropay           *ChargePaymentMethodDetailsGiropay           `json:"giropay,omitempty"`
	Grabpay           *ChargePaymentMethodDetailsGrabpay           `json:"grabpay,omitempty"`
	Ideal             *ChargePaymentMethodDetailsIdeal             `json:"ideal,omitempty"`
	Klarna            *ChargePaymentMethodDetailsKlarna            `json:"klarna,omitempty"`
	Multibanco        *ChargePaymentMethodDetailsMultibanco        `json:"multibanco,omitempty"`
	OXXO              *ChargePaymentMethodDetailsOXXO              `json:"oxxo,omitempty"`
	P24               *ChargePaymentMethodDetailsP24               `json:"p24,omitempty"`
	SepaDebit         *ChargePaymentMethodDetailsSepaDebit         `json:"sepa_debit,omitempty"`
	Sofort            *ChargePaymentMethodDetailsSofort            `json:"sofort,omitempty"`
	StripeAccount     *ChargePaymentMethodDetailsStripeAccount     `json:"stripe_account,omitempty"`
	Type              ChargePaymentMethodDetailsType               `json:"type,omitempty"`
	Wechat            *ChargePaymentMethodDetailsWechat            `json:"wechat,omitempty"`
}

// ChargeTransferData represents the information for the transfer_data associated with a charge.
type ChargeTransferData struct {
	Amount      int64    `form:"amount"`
	Destination *Account `json:"destination,omitempty"`
}

// Charge is the resource representing a Stripe charge.
// For more details see https://stripe.com/docs/api#charges.
type Charge struct {
	APIResource
	Amount                        *int64 `json:"amount,omitempty"`
	AmountCaptured                *int64 `json:"amount_captured,omitempty"`
	AmountRefunded                *int64 `json:"amount_refunded,omitempty"`
	Application                   *Application                `json:"application,omitempty"`
	ApplicationFee                *ApplicationFee             `json:"application_fee,omitempty"`
	ApplicationFeeAmount          *int64 `json:"application_fee_amount,omitempty"`
	AuthorizationCode             *string `json:"authorization_code,omitempty"`
	BalanceTransaction            *BalanceTransaction         `json:"balance_transaction,omitempty"`
	BillingDetails                *BillingDetails             `json:"billing_details,omitempty"`
	CalculatedStatementDescriptor *string `json:"calculated_statement_descriptor,omitempty"`
	Captured                      *bool `json:"captured,omitempty"`
	Created                       *int64 `json:"created,omitempty"`
	Currency                      Currency                    `json:"currency,omitempty"`
	Customer                      *Customer                   `json:"customer,omitempty"`
	Description                   *string `json:"description,omitempty"`
	Destination                   *Account                    `json:"destination,omitempty"`
	Dispute                       *Dispute                    `json:"dispute,omitempty"`
	Disputed                      *bool `json:"disputed,omitempty"`
	FailureCode                   *string `json:"failure_code,omitempty"`
	FailureMessage                *string `json:"failure_message,omitempty"`
	FraudDetails                  *FraudDetails               `json:"fraud_details,omitempty"`
	ID                            *string `json:"id,omitempty"`
	Invoice                       *Invoice                    `json:"invoice,omitempty"`
	Level3                        ChargeLevel3                `json:"level3,omitempty"`
	Livemode                      *bool `json:"livemode,omitempty"`
	Metadata                      map[string]string           `json:"metadata,omitempty"`
	OnBehalfOf                    *Account                    `json:"on_behalf_of,omitempty"`
	Outcome                       *ChargeOutcome              `json:"outcome,omitempty"`
	Paid                          *bool `json:"paid,omitempty"`
	PaymentIntent                 *PaymentIntent              `json:"payment_intent,omitempty"`
	PaymentMethod                 *string `json:"payment_method,omitempty"`
	PaymentMethodDetails          *ChargePaymentMethodDetails `json:"payment_method_details,omitempty"`
	ReceiptEmail                  *string `json:"receipt_email,omitempty"`
	ReceiptNumber                 *string `json:"receipt_number,omitempty"`
	ReceiptURL                    *string `json:"receipt_url,omitempty"`
	Refunded                      *bool `json:"refunded,omitempty"`
	Refunds                       *RefundList                 `json:"refunds,omitempty"`
	Review                        *Review                     `json:"review,omitempty"`
	Shipping                      *ShippingDetails            `json:"shipping,omitempty"`
	Source                        *PaymentSource              `json:"source,omitempty"`
	SourceTransfer                *Transfer                   `json:"source_transfer,omitempty"`
	StatementDescriptor           *string `json:"statement_descriptor,omitempty"`
	StatementDescriptorSuffix     *string `json:"statement_descriptor_suffix,omitempty"`
	Status                        *string `json:"status,omitempty"`
	Transfer                      *Transfer                   `json:"transfer,omitempty"`
	TransferData                  *ChargeTransferData         `json:"transfer_data,omitempty"`
	TransferGroup                 *string `json:"transfer_group,omitempty"`
}

// UnmarshalJSON handles deserialization of a charge.
// This custom unmarshaling is needed because the resulting
// property may be an ID or the full struct if it was expanded.
func (c *Charge) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = &id
		return nil
	}

	type charge Charge
	var v charge
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Charge(v)
	return nil
}

// ChargeList is a list of charges as retrieved from a list endpoint.
type ChargeList struct {
	APIResource
	ListMeta
	Data []*Charge `json:"data,omitempty"`
}

// FraudDetails is the structure detailing fraud status.
type FraudDetails struct {
	UserReport   ChargeFraudUserReport   `json:"user_report,omitempty"`
	StripeReport ChargeFraudStripeReport `json:"stripe_report,omitempty"`
}

// ChargeOutcomeRule tells you the Radar rule that blocked the charge, if any.
type ChargeOutcomeRule struct {
	Action    *string `json:"action,omitempty"`
	ID        *string `json:"id,omitempty"`
	Predicate *string `json:"predicate,omitempty"`
}

// ChargeOutcome is the charge's outcome that details whether a payment
// was accepted and why.
type ChargeOutcome struct {
	NetworkStatus *string `json:"network_status,omitempty"`
	Reason        *string `json:"reason,omitempty"`
	RiskLevel     *string `json:"risk_level,omitempty"`
	RiskScore     *int64 `json:"risk_score,omitempty"`
	Rule          *ChargeOutcomeRule `json:"rule,omitempty"`
	SellerMessage *string `json:"seller_message,omitempty"`
	Type          *string `json:"type,omitempty"`
}

// ShippingDetails is the structure containing shipping information.
type ShippingDetails struct {
	Address        *Address `json:"address,omitempty"`
	Carrier        *string `json:"carrier,omitempty"`
	Name           *string `json:"name,omitempty"`
	Phone          *string `json:"phone,omitempty"`
	TrackingNumber *string `json:"tracking_number,omitempty"`
}

var depth = -1

// UnmarshalJSON handles deserialization of a ChargeOutcomeRule.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *ChargeOutcomeRule) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = &id
		return nil
	}

	type chargeOutcomeRule ChargeOutcomeRule
	var v chargeOutcomeRule
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = ChargeOutcomeRule(v)
	return nil
}
