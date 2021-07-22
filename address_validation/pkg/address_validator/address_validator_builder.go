package address_validator

type AddressValidatorBuilder struct {
	rootValidator *AddressValidator
	latestValidator *AddressValidator
}

func CreateBuilder() AddressValidatorBuilder {
	builder := AddressValidatorBuilder{}
	return builder
}

type SmartyStreetsCredentials struct {
	AuthId    string
	AuthToken string
}

func (b AddressValidatorBuilder) setValidator(validator AddressValidator) {
	if b.rootValidator == nil {
		b.rootValidator = &validator
	} else {
		existing := *b.latestValidator
		existing.setNext(validator)
	}
	b.latestValidator = &validator
}

func (b AddressValidatorBuilder) WithInMemoryCache() {
	validator := addressValidatorInMemoryCache{}
	b.setValidator(validator)
}

func (b AddressValidatorBuilder) WithSmartyValidator(creds SmartyStreetsCredentials) {
	validator := smartyStreetsValidator{}
	validator.authId = creds.AuthId
	validator.authToken = creds.AuthToken
	b.setValidator(validator)
}

func (b AddressValidatorBuilder) Build() AddressValidator {
	return *b.rootValidator
}
