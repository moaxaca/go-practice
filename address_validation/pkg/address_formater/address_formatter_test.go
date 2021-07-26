package address_formater_test

import (
	"github.com/stretchr/testify/assert"
	"io.parcely.address_validation/pkg/address_formater"
	"testing"
)

func TestUsAddressFormatter_Format(t *testing.T) {
	sut := address_formater.Address{}
	sut.AddressLines = []string {"110 N Marina Dr" }
	sut.Locality = "long beach"
	sut.Region = "california"
	sut.PostalCode = "90803"
	formatter := address_formater.UsAddress{}

	expected := "110 N Marina Dr, Long Beach, California 90803"
	assert.Equal(t, expected, formatter.Format(sut))
}
