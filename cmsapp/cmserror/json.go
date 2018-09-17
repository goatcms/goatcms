package cmserror

// JSONError is a error object for access denide
type JSONError struct {
	base     error
	httpCode int
	json     string
}

// NewJSONError create new AccessDenideError instance
func NewJSONError(base error, httpCode int, json string) JSONError {
	return JSONError{
		base:     base,
		httpCode: httpCode,
		json:     json,
	}
}

// Error return error message
func (err JSONError) Error() string {
	return err.json
}

// HTTPCode return http error code
func (err JSONError) HTTPCode() int {
	return err.httpCode
}
