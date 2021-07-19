package address_validator

import (
	"fmt"
	"os"
	"testing"
)

func TestInMemoryAddressCache(t *testing.T) {
	config := ValidatorInit{}
	config.AuthId = os.Getenv("SMART_AUTH_ID")
	config.AuthToken = os.Getenv("SMART_AUTH_TOKEN")

	uat := UnvalidatedAddress{}
	uat.AddressLines = []string {"110 N Marina Dr" }
	uat.Locality = "Long Beach"
	uat.Region = "california"
	uat.PostalCode = "90803"

	validator := NewValidator(config)
	_, err := validator.validate(uat)
	if err != nil {
		fmt.Println(err)
	}

	_, err = validator.validate(uat)
}