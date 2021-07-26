package helpers

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response validation_error
type ValidationError struct {
	// The error message
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		// Example: Expected type int
		Message string
		// An optional field name to which this validation applies
		FieldName string
	}
}

// A AccessError is an error that is used when the required input fails access validation.
// swagger:response access_error
type AccessError struct {
	// The error message
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		// Example: Expected type int
		Message string
		// An optional field name to which this validation applies
		FieldName string
	}
}

// A MissingResourceError is an error that is used when the target resource isn't found.
// swagger:response miss_resource_error
type MissingResourceError struct {
	// The error message
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		// Example: Expected type int
		Message string
		// An optional field name to which this validation applies
		FieldName string
	}
}

// A ServerError is an error that is used when the target resource isn't found.
// swagger:response server_error
type ServerError struct {
	// The error message
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		// Example: Expected type int
		Message string
		// An optional field name to which this validation applies
		FieldName string
	}
}
