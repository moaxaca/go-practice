package address_validator

import (
	"context"
	"encoding/base64"
	"encoding/json"
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
	validate(address UnvalidatedAddress) (ValidatedAddress, error)
}

type AddressFormatter interface {
	format(address ValidatedAddress) string
}

type ValidatorInit struct {
	AuthId    string
	AuthToken string
}

func NewValidator(config ValidatorInit) AddressValidator {
	smarty := smartyStreetsValidator{}
	smarty.authId = config.AuthId
	smarty.authToken = config.AuthToken

	inMemory := addressValidatorInMemoryCache{}
	inMemory.cache = make(map[string]ValidatedAddress)
	inMemory.next = smarty

	return inMemory
}

// Address In-Memory Cache Layer
type addressValidatorInMemoryCache struct {
	cache map[string]ValidatedAddress
	next AddressValidator
}

func (av addressValidatorInMemoryCache) validate(address UnvalidatedAddress) (ValidatedAddress, error) {
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
	validated, err := av.next.validate(address)
	if err == nil && hashedAddress != "" {
		av.cache[hashedAddress] = validated
		fmt.Println("Cached address in memory")
	}
	return validated, err
}

// SmartStreets Validator
type smartyStreetsValidator struct {
	authId    string
	authToken string
	next      AddressValidator
}

func (av smartyStreetsValidator) validate(address UnvalidatedAddress) (ValidatedAddress, error) {
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
		return av.next.validate(address)
	}
	for i, input := range batch.Records() {
		fmt.Println("Smarty Results for input:", i)
		fmt.Println()
		for j, candidate := range input.Results {
			fmt.Println("  Candidate:", j)
			fmt.Println(" ", candidate.DeliveryLine1)
			fmt.Println(" ", candidate.LastLine)
			fmt.Println()
		}
	}
	validated := ValidatedAddress{}
	return validated, nil
}
