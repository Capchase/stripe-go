package countryspec

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/Capchase/stripe-go/v72"
	_ "github.com/Capchase/stripe-go/v72/testing"
)

func TestCountrySpecGet(t *testing.T) {
	spec, err := Get("US", nil)
	assert.Nil(t, err)
	assert.NotNil(t, spec)
}

func TestCountrySpecList(t *testing.T) {
	i := List(&stripe.CountrySpecListParams{})

	// Verify that we can get at least one spec
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.CountrySpec())
	assert.NotNil(t, i.CountrySpecList())
}
