package address_formater

import "io.parcely.address_validation/pkg/address_validator"

type AddressFormatter interface {
	Format(address address_validator.ValidatedAddress) string
}
