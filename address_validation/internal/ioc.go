package internal

import (
	"io.parcely.address_validation/pkg/address_validator"
	"os"
)

type IocContainer struct {
	AddressValidator *address_validator.AddressValidator
}

func CreateIoc() IocContainer {
	ioc := IocContainer{}

	smartyCredentials := address_validator.SmartyStreetsCredentials{}
	smartyCredentials.AuthId = os.Getenv("SMART_AUTH_ID")
	smartyCredentials.AuthToken = os.Getenv("SMART_AUTH_TOKEN")
	addressValidatorBuilder := address_validator.CreateBuilder()
	addressValidatorBuilder.WithInMemoryCache()
	addressValidatorBuilder.WithSmartyValidator(smartyCredentials)
	addressValidator, _ := addressValidatorBuilder.Build()
	// TODO Error Handle
	ioc.AddressValidator = &addressValidator

	return ioc
}
