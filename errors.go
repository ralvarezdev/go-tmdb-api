package gotmdbapi

import (
	"errors"
)

var (
	ErrNilClient       = errors.New("TMDB API client is nil")
	ErrNilBearerToken  = errors.New("TMDB API bearer token is nil")
	ErrRequestFailed   = errors.New("TMDB API request failed")
	ErrResponseParsing = errors.New("failed to parse TMDB API response")
)
