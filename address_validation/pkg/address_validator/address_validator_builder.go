package address_validator

import "errors"

type AddressValidatorBuilder struct {
	validators []*AddressValidator
}

func CreateBuilder() AddressValidatorBuilder {
	builder := AddressValidatorBuilder{}
	builder.validators = make([]*AddressValidator, 0)
	return builder
}

type SmartyStreetsCredentials struct {
	AuthId    string
	AuthToken string
}

func (b *AddressValidatorBuilder) WithInMemoryCache() {
	validator := &addressValidatorInMemoryCache{}
	validator.cache = make(map[string]ValidatedAddress)
	casted := AddressValidator(validator)
	b.validators = append(b.validators, &casted)
}

func (b *AddressValidatorBuilder) WithSmartyValidator(creds SmartyStreetsCredentials) {
	validator := &smartyStreetsValidator{}
	validator.authId = creds.AuthId
	validator.authToken = creds.AuthToken
	casted := AddressValidator(validator)
	b.validators = append(b.validators, &casted)
}

func (b *AddressValidatorBuilder) Build() (AddressValidator, error) {
	if len(b.validators) == 0 {
		return nil, errors.New("missing address validators")
	}
	var previous *AddressValidator = nil
	for _, validator := range b.validators {
		if previous != nil {
			previousValidator := *previous
			previousValidator.setNext(validator)
		}
		previous = validator
	}
	saved := b.validators
	b.validators = make([]*AddressValidator, 0)
	return *saved[0], nil
}
