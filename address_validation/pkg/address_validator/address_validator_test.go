package address_validator_test

import (
	"github.com/stretchr/testify/assert"
	"io.parcely.address_validation/pkg/address_formater"
	"io.parcely.address_validation/pkg/address_validator"

	"os"
	"testing"
)

// TODO Write explicit builder and mock tests
func TestAddressValidation(t *testing.T) {
	smartyCredentials := address_validator.SmartyStreetsCredentials{}
	smartyCredentials.AuthId = os.Getenv("SMART_AUTH_ID")
	smartyCredentials.AuthToken = os.Getenv("SMART_AUTH_TOKEN")

	builder := address_validator.CreateBuilder()
	builder.WithInMemoryCache()
	builder.WithSmartyValidator(smartyCredentials)

	formatter := address_formater.UsAddress{}
	validator, _ := builder.Build()

	sut := address_validator.UnvalidatedAddress{}
	sut.AddressLines = []string {"110N Marina Dr." }
	sut.Locality = "Long Beach"
	sut.Region = "california"
	sut.PostalCode = "90803"

	address, err := validator.Validate(sut)
	expected := "110N N Marina Dr , Long Beach, CA 90803"

	assert.Equal(t, expected, formatter.Format(address))
	assert.Nil(t, err)
	address, err = validator.Validate(sut)
	assert.Nil(t, err)
	assert.Equal(t, expected, formatter.Format(address))
}
