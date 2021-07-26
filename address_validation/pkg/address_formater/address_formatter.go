package address_formater

import (
	"fmt"
	"io.parcely.address_validation/pkg/address_validator"
	"strings"
)

type Address = address_validator.ValidatedAddress

type AddressFormatter interface {
	Format(address Address) string
}

type UsAddress struct {}

func mapStr(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func capitalizeAllWordsFirstChar (phrase string, sep string) string {
	return strings.Join(mapStr(strings.Split(phrase, sep), strings.Title), sep)
}

func (f *UsAddress) Format(address Address) string {
	addressLineString := strings.Join(address.AddressLines[:], " ")

	return fmt.Sprintf(
		"%s, %s, %s %s",
		addressLineString,
		capitalizeAllWordsFirstChar(address.Locality, " "),
		capitalizeAllWordsFirstChar(address.Region, " "),
		address.PostalCode,
	)
}
