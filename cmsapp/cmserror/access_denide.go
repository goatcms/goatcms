package cmserror

// AccessDenideError is a error object for access denide
type AccessDenideError struct {
	base error
}

// NewAccessDenideError create new AccessDenideError instance
func NewAccessDenideError(base error) AccessDenideError {
	return AccessDenideError{
		base: base,
	}
}

// Error return error message
func (err AccessDenideError) Error() string {
	return "Access denide"
}

// TranslateKey return key for message translation
func (err AccessDenideError) TranslateKey() string {
	return "access_denide"
}

// HTTPCode return http error code
func (err AccessDenideError) HTTPCode() int {
	return 403
}
