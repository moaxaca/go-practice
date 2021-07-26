package controllers

import (
	"github.com/gin-gonic/gin"
	"io.parcely.address_validation/pkg/address_validator"
)

// AddressValidationRequest represents the user for this application
//
// A user is the security principal for this application.
// It's also used as one of main axes for reporting.
//
// A user can have friends with whom they can share what they like.
//
// swagger:parameters validateAddress
type AddressValidationRequest struct {
	// Address Lines
	//
	// min items: 1
	// max items: 5
	// unique: true
	// in: body
	// example: ["Address"]
	AddressLines []string `json:"addressLines,omitempty"`
	Region       string            `json:"region,omitempty"`
	Locality     string            `json:"locality,omitempty"`
	PostalCode   string            `json:"postalCode,omitempty"`
	CountryCode    string          `json:"countryCode,omitempty"`
}

// AddressValidationResponseSuccess represents the user for this application
//
// A user is the security principal for this application.
// It's also used as one of main axes for reporting.
//
// A user can have friends with whom they can share what they like.
//
// swagger:response address_validation_response_success
type AddressValidationResponseSuccess struct {
	Uuid         string `json:"uuid,omitempty"`
	AddressLines []string `json:"addressLines,omitempty"`
	Region       string            `json:"region,omitempty"`
	Locality     string            `json:"locality,omitempty"`
	PostalCode   string            `json:"postalCode,omitempty"`
	CountryCode    string          `json:"countryCode,omitempty"`
	Meta         map[string]string `json:"meta,omitempty"`
}

type addressHandlers struct {
	addressValidator address_validator.AddressValidator
}

// swagger:route GET /validate AddressManagement validateAddress
//
// Validate Address
//
// This will take a structured address and returns a
// structured address which has been cross-validated
// against external systems.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Security:
//       api_key:
//
//     Responses:
//       default: server_error
//       200: address_validation_response_success
//       401: access_error
//       404: miss_resource_error
//       422: validation_error
func (a *addressHandlers) index(c *gin.Context) {
	request := new(AddressValidationRequest)
	err := c.Bind(request)
	if err != nil {
		// TODO
		return
	}
	unvalidated := address_validator.UnvalidatedAddress{}
	unvalidated.AddressLines = request.AddressLines
	unvalidated.Locality = request.Locality
	unvalidated.Region = request.Region
	unvalidated.PostalCode = request.PostalCode
	unvalidated.Country = request.CountryCode
	validate, err := a.addressValidator.Validate(unvalidated)
	if err != nil {
		// TODO
		return
	}
	response := AddressValidationResponseSuccess{}
	response.AddressLines = validate.AddressLines
	response.Locality = validate.Locality
	response.Region = validate.Region
	response.PostalCode = validate.PostalCode
	response.CountryCode = validate.CountryCode
	c.JSON(200, response)
}

func RegisterAddressRoutes(router *gin.RouterGroup, validator address_validator.AddressValidator) {
	handlers := addressHandlers{}
	handlers.addressValidator = validator
	router.POST("/", handlers.index)
}
