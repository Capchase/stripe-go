package stripe

import "encoding/json"

// VerificationDocumentDetailsCode is a machine-readable code specifying the verification state of
// a document associated with a person.
type VerificationDocumentDetailsCode string

// List of values that IdentityVerificationDetailsCode can take.
const (
	VerificationDocumentDetailsCodeDocumentCorrupt               VerificationDocumentDetailsCode = "document_corrupt"
	VerificationDocumentDetailsCodeDocumentFailedCopy            VerificationDocumentDetailsCode = "document_failed_copy"
	VerificationDocumentDetailsCodeDocumentFailedGreyscale       VerificationDocumentDetailsCode = "document_failed_greyscale"
	VerificationDocumentDetailsCodeDocumentFailedOther           VerificationDocumentDetailsCode = "document_failed_other"
	VerificationDocumentDetailsCodeDocumentFailedTestMode        VerificationDocumentDetailsCode = "document_failed_test_mode"
	VerificationDocumentDetailsCodeDocumentFraudulent            VerificationDocumentDetailsCode = "document_fraudulent"
	VerificationDocumentDetailsCodeDocumentIDTypeNotSupported    VerificationDocumentDetailsCode = "document_id_type_not_supported"
	VerificationDocumentDetailsCodeDocumentIDCountryNotSupported VerificationDocumentDetailsCode = "document_id_country_not_supported"
	VerificationDocumentDetailsCodeDocumentManipulated           VerificationDocumentDetailsCode = "document_manipulated"
	VerificationDocumentDetailsCodeDocumentMissingBack           VerificationDocumentDetailsCode = "document_missing_back"
	VerificationDocumentDetailsCodeDocumentMissingFront          VerificationDocumentDetailsCode = "document_missing_front"
	VerificationDocumentDetailsCodeDocumentNotReadable           VerificationDocumentDetailsCode = "document_not_readable"
	VerificationDocumentDetailsCodeDocumentNotUploaded           VerificationDocumentDetailsCode = "document_not_uploaded"
	VerificationDocumentDetailsCodeDocumentTooLarge              VerificationDocumentDetailsCode = "document_too_large"
)

// PersonPoliticalExposure describes the political exposure of a given person.
type PersonPoliticalExposure string

// List of values that IdentityVerificationStatus can take.
const (
	PersonPoliticalExposureExisting PersonPoliticalExposure = "existing"
	PersonPoliticalExposureNone     PersonPoliticalExposure = "none"
)

// PersonVerificationDetailsCode is a machine-readable code specifying the verification state of a
// person.
type PersonVerificationDetailsCode string

// List of values that IdentityVerificationDetailsCode can take.
const (
	PersonVerificationDetailsCodeFailedKeyedIdentity PersonVerificationDetailsCode = "failed_keyed_identity"
	PersonVerificationDetailsCodeFailedOther         PersonVerificationDetailsCode = "failed_other"
	PersonVerificationDetailsCodeScanNameMismatch    PersonVerificationDetailsCode = "scan_name_mismatch"
)

// IdentityVerificationStatus describes the different statuses for identity verification.
type IdentityVerificationStatus string

// List of values that IdentityVerificationStatus can take.
const (
	IdentityVerificationStatusPending    IdentityVerificationStatus = "pending"
	IdentityVerificationStatusUnverified IdentityVerificationStatus = "unverified"
	IdentityVerificationStatusVerified   IdentityVerificationStatus = "verified"
)

// DOBParams represents a DOB during account creation/updates.
type DOBParams struct {
	Day   *int64 `form:"day"`
	Month *int64 `form:"month"`
	Year  *int64 `form:"year"`
}

// RelationshipParams is used to set the relationship between an account and a person.
type RelationshipParams struct {
	Director         *bool    `form:"director"`
	Executive        *bool    `form:"executive"`
	Owner            *bool    `form:"owner"`
	PercentOwnership *float64 `form:"percent_ownership"`
	Representative   *bool    `form:"representative"`
	Title            *string  `form:"title"`
}

// PersonVerificationDocumentParams represents the parameters available for the document verifying
// a person's identity.
type PersonVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// PersonVerificationParams is used to represent parameters associated with a person's verification
// details.
type PersonVerificationParams struct {
	AdditionalDocument *PersonVerificationDocumentParams `form:"additional_document"`
	Document           *PersonVerificationDocumentParams `form:"document"`
}

// PersonParams is the set of parameters that can be used when creating or updating a person.
// For more details see https://stripe.com/docs/api#create_person.
type PersonParams struct {
	Params            `form:"*"`
	Account           *string                   `form:"-"` // Included in URL
	Address           *AccountAddressParams     `form:"address"`
	AddressKana       *AccountAddressParams     `form:"address_kana"`
	AddressKanji      *AccountAddressParams     `form:"address_kanji"`
	DOB               *DOBParams                `form:"dob"`
	Email             *string                   `form:"email"`
	FirstName         *string                   `form:"first_name"`
	FirstNameKana     *string                   `form:"first_name_kana"`
	FirstNameKanji    *string                   `form:"first_name_kanji"`
	Gender            *string                   `form:"gender"`
	IDNumber          *string                   `form:"id_number"`
	LastName          *string                   `form:"last_name"`
	LastNameKana      *string                   `form:"last_name_kana"`
	LastNameKanji     *string                   `form:"last_name_kanji"`
	MaidenName        *string                   `form:"maiden_name"`
	Nationality       *string                   `form:"nationality"`
	PersonToken       *string                   `form:"person_token"`
	Phone             *string                   `form:"phone"`
	PoliticalExposure *string                   `form:"political_exposure"`
	Relationship      *RelationshipParams       `form:"relationship"`
	SSNLast4          *string                   `form:"ssn_last_4"`
	Verification      *PersonVerificationParams `form:"verification"`
}

// RelationshipListParams is used to filter persons by the relationship
type RelationshipListParams struct {
	Director       *bool `form:"director"`
	Executive      *bool `form:"executive"`
	Owner          *bool `form:"owner"`
	Representative *bool `form:"representative"`
}

// PersonListParams is the set of parameters that can be used when listing persons.
// For more detail see https://stripe.com/docs/api#list_persons.
type PersonListParams struct {
	ListParams   `form:"*"`
	Account      *string                 `form:"-"` // Included in URL
	Relationship *RelationshipListParams `form:"relationship"`
}

// DOB represents a Person's date of birth.
type DOB struct {
	Day   *int64 `json:"day,omitempty"`
	Month *int64 `json:"month,omitempty"`
	Year  *int64 `json:"year,omitempty"`
}

// Relationship represents how the Person relates to the business.
type Relationship struct {
	Director         *bool `json:"director,omitempty"`
	Executive        *bool `json:"executive,omitempty"`
	Owner            *bool `json:"owner,omitempty"`
	PercentOwnership *float64 `json:"percent_ownership,omitempty"`
	Representative   *bool `json:"representative,omitempty"`
	Title            *string `json:"title,omitempty"`
}

// Requirements represents what's missing to verify a Person.
type Requirements struct {
	CurrentlyDue        []string                    `json:"currently_due,omitempty"`
	Errors              []*AccountRequirementsError `json:"errors,omitempty"`
	EventuallyDue       []string                    `json:"eventually_due,omitempty"`
	PastDue             []string                    `json:"past_due,omitempty"`
	PendingVerification []string                    `json:"pending_verification,omitempty"`
}

// PersonVerificationDocument represents the documents associated with a Person.
type PersonVerificationDocument struct {
	Back        *File                           `json:"back,omitempty"`
	Details     *string `json:"details,omitempty"`
	DetailsCode VerificationDocumentDetailsCode `json:"details_code,omitempty"`
	Front       *File                           `json:"front,omitempty"`
}

// PersonVerification is the structure for a person's verification details.
type PersonVerification struct {
	AdditionalDocument *PersonVerificationDocument   `json:"additional_document,omitempty"`
	Details            *string `json:"details,omitempty"`
	DetailsCode        PersonVerificationDetailsCode `json:"details_code,omitempty"`
	Document           *PersonVerificationDocument   `json:"document,omitempty"`
	Status             IdentityVerificationStatus    `json:"status,omitempty"`
}

// Person is the resource representing a Stripe person.
// For more details see https://stripe.com/docs/api#persons.
type Person struct {
	APIResource
	Account           *string `json:"account,omitempty"`
	Address           *AccountAddress         `json:"address,omitempty"`
	AddressKana       *AccountAddress         `json:"address_kana,omitempty"`
	AddressKanji      *AccountAddress         `json:"address_kanji,omitempty"`
	Deleted           *bool `json:"deleted,omitempty"`
	DOB               *DOB                    `json:"dob,omitempty"`
	Email             *string `json:"email,omitempty"`
	FirstName         *string `json:"first_name,omitempty"`
	FirstNameKana     *string `json:"first_name_kana,omitempty"`
	FirstNameKanji    *string `json:"first_name_kanji,omitempty"`
	Gender            *string `json:"gender,omitempty"`
	ID                *string `json:"id,omitempty"`
	IDNumberProvided  *bool `json:"id_number_provided,omitempty"`
	LastName          *string `json:"last_name,omitempty"`
	LastNameKana      *string `json:"last_name_kana,omitempty"`
	LastNameKanji     *string `json:"last_name_kanji,omitempty"`
	MaidenName        *string `json:"maiden_name,omitempty"`
	Nationality       *string `json:"nationality,omitempty"`
	Metadata          map[string]string       `json:"metadata,omitempty"`
	Object            *string `json:"object,omitempty"`
	Phone             *string `json:"phone,omitempty"`
	PoliticalExposure PersonPoliticalExposure `json:"political_exposure,omitempty"`
	Relationship      *Relationship           `json:"relationship,omitempty"`
	Requirements      *Requirements           `json:"requirements,omitempty"`
	SSNLast4Provided  *bool `json:"ssn_last_4_provided,omitempty"`
	Verification      *PersonVerification     `json:"verification,omitempty"`
}

// PersonList is a list of persons as retrieved from a list endpoint.
type PersonList struct {
	APIResource
	ListMeta
	Data []*Person `json:"data,omitempty"`
}

// UnmarshalJSON handles deserialization of a Person.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Person) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = &id
		return nil
	}

	type person Person
	var v person
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Person(v)
	return nil
}
