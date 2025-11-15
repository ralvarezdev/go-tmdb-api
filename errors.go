package gotmdbapi

import (
	"errors"
)

const (
	ErrBuildingRequest           = "error building TMDB API request: %v"
	ErrAnErrOcurredDuringRequest = "an error occurred during the TMDB API request: %v"
	ErrRequestFailed             = "TMDB API request failed with status code %d: %s"
)

var (
	ErrNilClient       = errors.New("TMDB API client is nil")
	ErrEmptyAPIKey     = errors.New("TMDB API key is nil or empty")
	ErrResponseParsing = errors.New("failed to parse TMDB API response")
)
