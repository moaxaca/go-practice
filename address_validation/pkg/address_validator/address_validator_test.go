package address_validator_test

import (
	"github.com/stretchr/testify/assert"
	"io.parcely.address_validation/pkg/address_formater"
	"io.parcely.address_validation/pkg/address_validator"
	"testing"
)

func TestAddressValidation(t *testing.T) {
	fixture := address_validator.ValidatedAddress{}
	fixture.AddressLines = []string {"110N Marina Dr." }
	fixture.Locality = "Long Beach"
	fixture.Region = "california"
	fixture.PostalCode = "90803"

	builder := address_validator.CreateBuilder()
	builder.WithInMemoryCache()
	builder.WithStub(fixture)

	formatter := address_formater.UsAddress{}
	validator, _ := builder.Build()

	sut := address_validator.UnvalidatedAddress{}
	sut.AddressLines = []string {"110N Marina Dr." }
	sut.Locality = "Long Beach"
	sut.Region = "california"
	sut.PostalCode = "90803"

	address, err := validator.Validate(sut)
	assert.Equal(t, formatter.Format(fixture), formatter.Format(address))
	assert.Nil(t, err)
}

func TestAddressValidatorBuilder(t *testing.T) {
	builder := address_validator.CreateBuilder()
	builder.WithInMemoryCache()
	_, stubErr := builder.Build()
	assert.Nil(t, stubErr)

	builder.WithInMemoryCache()
	_, inMemoryErr := builder.Build()
	assert.Nil(t, inMemoryErr)

	smartyCredentials := address_validator.SmartyStreetsCredentials{}
	smartyCredentials.AuthId = "FAKE_SMART_AUTH_ID"
	smartyCredentials.AuthToken = "FAKE_SMART_AUTH_TOKEN"
	builder.WithSmartyValidator(smartyCredentials)
	_, smartyErr := builder.Build()
	assert.Nil(t, smartyErr)
}
