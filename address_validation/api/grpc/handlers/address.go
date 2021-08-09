package handlers

import (
	"context"
	pb "io.parcely.address_validation/api/grpc/proto"
	"io.parcely.address_validation/pkg/address_validator"
)

type addressHandle struct {
	addressValidator address_validator.AddressValidator
	pb.UnimplementedAddressServer
}

func (s *addressHandle) Validate(_ context.Context, req *pb.AddressValidationRequest) (*pb.AddressValidationResponse, error) {
	unvalidated := address_validator.UnvalidatedAddress{}
	unvalidated.AddressLines = req.GetAddressLines()
	unvalidated.Locality = req.GetLocality()
	unvalidated.Region = req.GetRegion()
	unvalidated.PostalCode = req.GetPostalCode()
	unvalidated.Country = req.GetCountryCode()
	validate, err := s.addressValidator.Validate(unvalidated)
	if err != nil {
		failure := pb.AddressValidationResponseFailure{}
		failure.Input = req
		failure.ErrorCode = pb.ErrorCode_UNABLE_TO_VALIDATE_ADDRESS
		res := pb.AddressValidationResponse_ErrorResponse{ErrorResponse: &failure}
		return &pb.AddressValidationResponse{Response: &res}, nil
	}
	success := pb.AddressValidationResponseSuccess{}
	success.Uuid = validate.Uuid
	success.AddressLines = validate.AddressLines
	success.Locality = validate.Locality
	success.Region = validate.Region
	success.PostalCode = validate.PostalCode
	success.CountryCode = validate.CountryCode
	res := pb.AddressValidationResponse_SuccessResponse{SuccessResponse: &success}
	return &pb.AddressValidationResponse{Response: &res}, nil
}

func CreateAddressHandler(av address_validator.AddressValidator) *addressHandle {
	handle := addressHandle{}
	handle.addressValidator = av
	return &handle
}
