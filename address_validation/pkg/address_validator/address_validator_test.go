package address_validator_test

import (
	"fmt"
	"io.parcely.address_validation/pkg/address_validator"

	"os"
	"testing"
)

func TestInMemoryAddressCache(t *testing.T) {
	smartyCreds := address_validator.SmartyStreetsCredentials{}
	smartyCreds.AuthId = os.Getenv("SMART_AUTH_ID")
	smartyCreds.AuthToken = os.Getenv("SMART_AUTH_TOKEN")

	builder := address_validator.CreateBuilder()
	builder.WithInMemoryCache()
	builder.WithSmartyValidator(smartyCreds)
	validator, _ := builder.Build()

	uat := address_validator.UnvalidatedAddress{}
	uat.AddressLines = []string {"110 N Marina Dr" }
	uat.Locality = "Long Beach"
	uat.Region = "california"
	uat.PostalCode = "90803"

	address, err := validator.Validate(uat)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(address)
	address, err = validator.Validate(uat)
	fmt.Println(address)
	// TODO VALIDATE
}