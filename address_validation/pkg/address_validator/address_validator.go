package address_validator

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"
	"github.com/smartystreets/smartystreets-go-sdk/wireup"
)

type UnvalidatedAddress struct {
	AddressLines []string
	Region       string // State | Territory
	Locality     string // City
	PostalCode   string // PostalCode / ZipCode
	Country      string
}

type ValidatedAddress struct {
	Uuid         string
	AddressLines []string
	Region       string            // State | Territory
	Locality     string            // City
	PostalCode   string            // PostalCode / ZipCode
	CountryId    string            // Country ID
	Meta         map[string]string // Integration & Source Information
}

type AddressValidator interface {
	Validate(address UnvalidatedAddress) (ValidatedAddress, error)
	setNext(validator *AddressValidator)
	hasNext() bool
}

// Address In-Memory Cache Layer
type addressValidatorInMemoryCache struct {
	cache map[string]ValidatedAddress
	next *AddressValidator
}

func (av addressValidatorInMemoryCache) Validate(address UnvalidatedAddress) (ValidatedAddress, error) {
	addressJson, err := json.Marshal(address)
	var hashedAddress = ""
	if err == nil {
		hashedAddress = base64.URLEncoding.EncodeToString(addressJson)
	}
	if hashedAddress != "" {
		if validatedAddress, ok := av.cache[hashedAddress]; ok {
			fmt.Println("Address found in cache")
			return validatedAddress, nil
		}
	}
	next := *av.next
	validated, err := next.Validate(address)
	if err == nil && hashedAddress != "" {
		av.cache[hashedAddress] = validated
		fmt.Println("Cached address in memory")
	}
	return validated, err
}

func (av *addressValidatorInMemoryCache) setNext(validator *AddressValidator) {
	av.next = validator
}

func (av *addressValidatorInMemoryCache) hasNext() bool {
	return av.next != nil
}

// SmartStreets Validator
type smartyStreetsValidator struct {
	authId    string
	authToken string
	next      *AddressValidator
}

func (av smartyStreetsValidator) Validate(address UnvalidatedAddress) (ValidatedAddress, error) {
	client := wireup.BuildUSStreetAPIClient(
		wireup.SecretKeyCredential(av.authId, av.authToken),
	)
	lookup := &street.Lookup{
		Street:        strings.Join(address.AddressLines, " "),
		City:          address.Locality,
		State:         address.Region,
		ZIPCode:       address.PostalCode,
		MaxCandidates: 1,
		MatchStrategy: street.MatchStrict,
	}
	batch := street.NewBatch()
	batch.Append(lookup)
	err := client.SendBatchWithContext(context.Background(), batch)
	if err != nil {
		if av.next == nil {
			return ValidatedAddress{}, err
		}
		next := *av.next
		return next.Validate(address)
	}
	validated := ValidatedAddress{}
	for _, input := range batch.Records() {
		if len(input.Results) == 0 {
			return validated, errors.New("unable to validate address")
		}
		for _, candidate := range input.Results {
			validated.AddressLines = []string {
				candidate.DeliveryLine1,
				candidate.DeliveryLine2,
			}
			validated.PostalCode = candidate.Components.ZIPCode
			validated.Locality = candidate.Components.CityName
			validated.Region = candidate.Components.StateAbbreviation
			validated.CountryId = "1"
			validated.Meta = make(map[string]string)
			validated.Meta["integration"] = "smarty"
			validated.Meta["smarty_street_number"] = candidate.Components.PrimaryNumber
			validated.Meta["smarty_street_pre_direction"] = candidate.Components.StreetPredirection
			validated.Meta["smarty_street_name"] = candidate.Components.StreetName
			validated.Meta["smarty_street_post_direction"] = candidate.Components.StreetPostdirection
			validated.Meta["smarty_street_suffix"] = candidate.Components.StreetSuffix
			validated.Meta["smarty_street_secondary_number"] = candidate.Components.SecondaryNumber
			validated.Meta["smarty_street_secondary_designator"] = candidate.Components.SecondaryDesignator
			validated.Meta["smarty_zip"] = candidate.Components.ZIPCode
			validated.Meta["smarty_zip_plus_four"] = candidate.Components.Plus4Code
			validated.Meta["smarty_delivery_point"] = candidate.Components.DeliveryPoint
			validated.Meta["smarty_delivery_point_check_digit"] = candidate.Components.DeliveryPointCheckDigit
			validated.Meta["smarty_county"] = candidate.Metadata.CountyName
			validated.Meta["smarty_carrier_route"] = candidate.Metadata.CarrierRoute
			validated.Meta["smarty_congressional_district"] = candidate.Metadata.CongressionalDistrict
		}
	}
	return validated, nil
}

func (av *smartyStreetsValidator) setNext(validator *AddressValidator) {
	av.next = validator
}

func (av *smartyStreetsValidator) hasNext() bool {
	return av.next != nil
}
