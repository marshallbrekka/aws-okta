package client

import "errors"

var (
	InvalidCredentialsError = errors.New("Okta credentials are not valid")
	InvalidSessionError     = errors.New("Okta session is not valid")
	UnexpectedResponseError = errors.New("Okta returned an unexpected response")
	NotImplementedError     = errors.New("Not implemented")
)