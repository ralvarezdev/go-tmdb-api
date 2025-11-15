package gotmdbapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type (
	// Client is the TMDB API client
	Client struct {
		apiKey string
	}
)

// NewClient creates a new TMDB API client
//
// Parameters:
//
// - apiKey: the TMDB API key
//
// Returns:
//
// - *Client: the TMDB API client
// - error: if there was an error creating the client
func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, ErrEmptyAPIKey
	}

	return &Client{
		apiKey: apiKey,
	}, nil
}

// addAuthorizationToRequest adds the Authorization header to the HTTP request
//
// Parameters:
//
// - req: the HTTP request
func (c Client) addAuthorizationToRequest(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
}

// addLanguageQueryParameter adds the language query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - language: the language code (optional, defaults to "en-US")
func (c Client) addLanguageQueryParameter(
	query url.Values,
	language string,
) {
	if language != "" {
		query.Add(Language, language)
	}
}

// addIncludeAdultQueryParameter adds the include adult query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - includeAdult: whether to include adult content
func (c Client) addIncludeAdultQueryParameter(
	query url.Values,
	includeAdult bool,
) {
	query.Add(IncludeAdult, fmt.Sprintf("%t", includeAdult))
}

// addPageQueryParameter adds the page query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - page: the page number (optional, defaults to 1)
func (c Client) addPageQueryParameter(
	query url.Values,
	page int32,
) {
	if page > 0 {
		query.Add(Page, fmt.Sprintf("%d", page))
	}
}

// addRegionQueryParameter adds the region query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - region: the region code (optional)
func (c Client) addRegionQueryParameter(
	query url.Values,
	region string,
) {
	if region != "" {
		query.Add(Region, region)
	}
}

// addPrimaryReleaseYearQueryParameter adds the primary release year query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - primaryReleaseYear: the primary release year (optional)
func (c Client) addPrimaryReleaseYearQueryParameter(
	query url.Values,
	primaryReleaseYear int32,
) {
	if primaryReleaseYear > 0 {
		query.Add(PrimaryReleaseYear, fmt.Sprintf("%d", primaryReleaseYear))
	}
}

// addYearQueryParameter adds the year query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - year: the year (optional)
func (c Client) addYearQueryParameter(
	query url.Values,
	year int32,
) {
	if year > 0 {
		query.Add(Year, fmt.Sprintf("%d", year))
	}
}

// addMovieListsQueryParameters adds the query parameters for movie lists to the HTTP request
//
// Parameters:
//
// - req: the HTTP request
// - language: the language code (optional, defaults to "en-US")
// - page: the page number (optional, defaults to 1)
// - region: the region code (optional)
func (c Client) addMovieListsQueryParameters(
	req *http.Request,
	language string,
	page int32,
	region string,
) {
	q := req.URL.Query()
	c.addLanguageQueryParameter(q, language)
	c.addPageQueryParameter(q, page)
	c.addRegionQueryParameter(q, region)
	req.URL.RawQuery = q.Encode()
}

// addSearchMoviesQueryParameters adds the query parameters for searching movies to the HTTP request
//
// Parameters:
//
// - req: the HTTP request
// - query: the search query
// - includeAdult: whether to include adult content
// - language: the language code (optional, defaults to "en-US")
// - primaryReleaseYear: the primary release year (optional)
// - page: the page number (optional, defaults to 1)
// - region: the region code (optional)
// - year: the year (optional)
func (c Client) addSearchMoviesQueryParameters(
	req *http.Request,
	query string,
	includeAdult bool,
	language string,
	primaryReleaseYear int32,
	page int32,
	region string,
	year int32,
) {
	q := req.URL.Query()
	q.Add(Query, query)
	c.addIncludeAdultQueryParameter(q, includeAdult)
	c.addLanguageQueryParameter(q, language)
	c.addPrimaryReleaseYearQueryParameter(q, primaryReleaseYear)
	c.addPageQueryParameter(q, page)
	c.addRegionQueryParameter(q, region)
	c.addYearQueryParameter(q, year)
	req.URL.RawQuery = q.Encode()
}

// addSimilarMoviesQueryParameters adds the query parameters for similar movies to the HTTP request
//
// Parameters:
//
// - req: the HTTP request
// - language: the language code (optional, defaults to "en-US")
// - page: the page number (optional, defaults to 1)
func (c Client) addSimilarMoviesQueryParameters(
	req *http.Request,
	language string,
	page int32,
) {
	q := req.URL.Query()
	c.addLanguageQueryParameter(q, language)
	c.addPageQueryParameter(q, page)
	req.URL.RawQuery = q.Encode()
}

// GetMoviesNowPlaying fetches the list of movies that are now playing in theaters
//
// Parameters:
//
// - ctx: the context of the request
// - id: the movie ID
// - language: the language code (optional, defaults to "en-US")
// - page: the page number (optional, defaults to 1)
// - region: the region code (optional)
//
// Returns:
//
// - (*DataMovieListResponse): the response containing the list of movies
// - int: the HTTP status code
// - error: if there was an error fetching the movies
func (c Client) GetMoviesNowPlaying(
	ctx context.Context,
	language string,
	page int32,
	region string,
) (parsedResp *DateMovieListResponse, statusCode int, err error) {
	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, GetNowPlayingMoviesURL, http.NoBody)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(ErrBuildingRequest, err)
	}

	// Add the Authorization header
	c.addAuthorizationToRequest(req)

	// Add query parameters
	c.addMovieListsQueryParameters(req, language, page, region)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf(ErrAnErrOcurredDuringRequest, err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, resp.StatusCode, fmt.Errorf(ErrRequestFailed, resp.StatusCode, string(body))
	}

	// Parse the response
	parsedResp = &DateMovieListResponse{}
	if parseErr := json.NewDecoder(resp.Body).Decode(parsedResp); parseErr != nil {
		return nil, resp.StatusCode, ErrResponseParsing
	}
	return parsedResp, resp.StatusCode,  nil
}

// GetMoviesPopular fetches the list of popular movies
//
// Parameters:
//
// - ctx: the context of the request
// - language: the language code (optional, defaults to "en-US")
// - page: the page number (optional, defaults to 1)
// - region: the region code (optional)
//
// Returns:
//
// - (*MovieListResponse): the response containing the list of popular movies
// - int: the HTTP status code
// - error: if there was an error fetching the movies
func (c Client) GetMoviesPopular(
	ctx context.Context,
	language string,
	page int32,
	region string,
) (parsedResp *MovieListResponse, statusCode int, err error) {
	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, GetPopularMoviesURL, http.NoBody)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(ErrBuildingRequest, err)
	}

	// Add the Authorization header
	c.addAuthorizationToRequest(req)

	// Add query parameters
	c.addMovieListsQueryParameters(req, language, page, region)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf(ErrAnErrOcurredDuringRequest, err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, resp.StatusCode, fmt.Errorf(ErrRequestFailed, resp.StatusCode, string(body))
	}

	// Parse the response
	parsedResp = &MovieListResponse{}
	if parseErr := json.NewDecoder(resp.Body).Decode(parsedResp); parseErr != nil {
		return nil, resp.StatusCode, ErrResponseParsing
	}
	return parsedResp, resp.StatusCode, nil
}

// GetMoviesTopRated fetches the list of top-rated movies
//
// Parameters:
//
// - ctx: the context of the request
// - language: the language code (optional, defaults to "en-US")
// - page: the page number (optional, defaults to 1)
// - region: the region code (optional)
//
// Returns:
//
// - (*MovieListResponse): the response containing the list of top-rated movies
// - int: the HTTP status code
// - error: if there was an error fetching the movies
func (c Client) GetMoviesTopRated(
	ctx context.Context,
	language string,
	page int32,
	region string,
) (parsedResp *MovieListResponse, httpStatus int, err error) {
	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, GetTopRatedMoviesURL, http.NoBody)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(ErrBuildingRequest, err)
	}

	// Add the Authorization header
	c.addAuthorizationToRequest(req)

	// Add query parameters
	c.addMovieListsQueryParameters(req, language, page, region)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf(ErrAnErrOcurredDuringRequest, err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, resp.StatusCode, fmt.Errorf(ErrRequestFailed, resp.StatusCode, string(body))
	}

	// Parse the response
	parsedResp = &MovieListResponse{}
	if parseErr := json.NewDecoder(resp.Body).Decode(parsedResp); parseErr != nil {
		return nil, resp.StatusCode, ErrResponseParsing
	}
	return parsedResp,resp.StatusCode, nil
}

// GetMoviesUpcoming fetches the list of upcoming movies
//
// Parameters:
//
// - ctx: the context of the request
// - language: the language code (optional, defaults to "en-US")
// - page: the page number (optional, defaults to 1)
// - region: the region code (optional)
//
// Returns:
//
// - (*DateMovieListResponse): the response containing the list of upcoming movies
// - int: the HTTP status code
// - error: if there was an error fetching the movies
func (c Client) GetMoviesUpcoming(
	ctx context.Context,
	language string,
	page int32,
	region string,
) (parsedResp *DateMovieListResponse, statusCode int, err error) {
	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, GetUpcomingMoviesURL, http.NoBody)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(ErrBuildingRequest, err)
	}

	// Add the Authorization header
	c.addAuthorizationToRequest(req)

	// Add query parameters
	c.addMovieListsQueryParameters(req, language, page, region)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf(ErrAnErrOcurredDuringRequest, err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, resp.StatusCode, fmt.Errorf(ErrRequestFailed, resp.StatusCode, string(body))
	}

	// Parse the response
	parsedResp = &DateMovieListResponse{}
	if parseErr := json.NewDecoder(resp.Body).Decode(parsedResp); parseErr != nil {
		return nil, resp.StatusCode, ErrResponseParsing
	}
	return parsedResp, resp.StatusCode, nil
}

// SearchMovies searches for movies by query
//
// Parameters:
//
// - ctx: the context of the request
// - query: the search query
// - includeAdult: whether to include adult content
// - language: the language code (optional, defaults to "en-US")
// - primaryReleaseYear: the primary release year (optional)
// - page: the page number (optional, defaults to 1)
// - region: the region code (optional)
// - year: the year (optional)
//
// Returns:
//
// - (*MovieListResponse): the response containing the list of movies matching the search query
// - int: the HTTP status code
// - error: if there was an error searching for movies
func (c Client) SearchMovies(
	ctx context.Context,
	query string,
	includeAdult bool,
	language string,
	primaryReleaseYear int32,
	page int32,
	region string,
	year int32,
) (parsedresp *MovieListResponse, statusCode int, err error) {
	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, SearchMoviesURL, http.NoBody)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(ErrBuildingRequest, err)
	}

	// Add the Authorization header
	c.addAuthorizationToRequest(req)

	// Add query parameters
	c.addSearchMoviesQueryParameters(
		req,
		query,
		includeAdult,
		language,
		primaryReleaseYear,
		page,
		region,
		year,
	)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf(ErrAnErrOcurredDuringRequest, err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, resp.StatusCode, fmt.Errorf(ErrRequestFailed, resp.StatusCode, string(body))
	}

	// Parse the response
	var parsedResp MovieListResponse
	if parseErr := json.NewDecoder(resp.Body).Decode(&parsedResp); parseErr != nil {
		return nil, resp.StatusCode, ErrResponseParsing
	}
	return &parsedResp, resp.StatusCode, nil
}

// SimilarMovies fetches the list of movies similar to a given movie
//
// Parameters:
//
// - ctx: the context of the request
// - movieID: the ID of the movie
// - language: the language code (optional, defaults to "en-US")
// - page: the page number (optional, defaults to 1)
//
// Returns:
//
// - (*MovieListResponse): the response containing the list of similar movies
// - int: the HTTP status code
// - error: if there was an error fetching similar movies
func (c Client) SimilarMovies(
	ctx context.Context,
	movieID int32,
	language string,
	page int32,
) (parsedResp *MovieListResponse, statusCode int, err error) {
	// Create the HTTP request
	apiURL := fmt.Sprintf(SimilarMoviesURL, fmt.Sprintf("%d", movieID))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, http.NoBody)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(ErrBuildingRequest, err)
	}

	// Add the Authorization header
	c.addAuthorizationToRequest(req)

	// Add query parameters
	c.addSimilarMoviesQueryParameters(req, language, page)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf(ErrAnErrOcurredDuringRequest, err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, resp.StatusCode, fmt.Errorf(ErrRequestFailed, resp.StatusCode, string(body))
	}

	// Parse the response
	parsedResp = &MovieListResponse{}
	if parseErr := json.NewDecoder(resp.Body).Decode(parsedResp); parseErr != nil {
		return nil, resp.StatusCode, ErrResponseParsing
	}
	return parsedResp, resp.StatusCode, nil
}

// GetMovieCredits fetches the credits for a given movie
//
// Parameters:
//
// - ctx: the context of the request
// - movieID: the ID of the movie
// - language: the language code (optional, defaults to "en-US")
//
// Returns:
//
// - (*MovieCreditsResponse): the response containing the movie credits
// - int: the HTTP status code
// - error: if there was an error fetching the movie credits
func (c Client) GetMovieCredits(
	ctx context.Context,
	movieID int32,
	language string,
) (parsedResp *MovieCreditsResponse, statusCode int, err error) {
	// Create the HTTP request
	apiURL := fmt.Sprintf(GetMovieCreditsURL, fmt.Sprintf("%d", movieID))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, http.NoBody)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(ErrBuildingRequest, err)
	}

	// Add the Authorization header
	c.addAuthorizationToRequest(req)

	// Add query parameters
	q := req.URL.Query()
	c.addLanguageQueryParameter(q, language)
	req.URL.RawQuery = q.Encode()

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf(ErrAnErrOcurredDuringRequest, err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, resp.StatusCode, fmt.Errorf(ErrRequestFailed, resp.StatusCode, string(body))
	}

	// Parse the response
	parsedResp = &MovieCreditsResponse{}
	if parseErr := json.NewDecoder(resp.Body).Decode(parsedResp); parseErr != nil {
		return nil, resp.StatusCode, ErrResponseParsing
	}
	return parsedResp, resp.StatusCode, nil
}

// GetMovieDetails fetches the details of a given movie
//
// Parameters:
//
// - ctx: the context of the request
// - movieID: the ID of the movie
// - language: the language code (optional, defaults to "en-US")
//
// Returns:
//
// - (*MovieDetailsResponse): the response containing the movie details
// - int: the HTTP status code
// - error: if there was an error fetching the movie details
func (c Client) GetMovieDetails(
	ctx context.Context,
	movieID int32,
	language string,
) (parsedResp *MovieDetailsResponse, statusCode int, err error) {
	// Create the HTTP request
	apiURL := fmt.Sprintf(GetMovieDetailsURL, fmt.Sprintf("%d", movieID))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, http.NoBody)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(ErrBuildingRequest, err)
	}

	// Add the Authorization header
	c.addAuthorizationToRequest(req)

	// Add query parameters
	q := req.URL.Query()
	c.addLanguageQueryParameter(q, language)
	req.URL.RawQuery = q.Encode()

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf(ErrAnErrOcurredDuringRequest, err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, resp.StatusCode, fmt.Errorf(ErrRequestFailed, resp.StatusCode, string(body))
	}

	// Parse the response
	parsedResp = &MovieDetailsResponse{}
	if parseErr := json.NewDecoder(resp.Body).Decode(parsedResp); parseErr != nil {
		return nil, resp.StatusCode, ErrResponseParsing
	}
	return parsedResp, resp.StatusCode, nil
}

// GetMovieReviews fetches the reviews for a given movie
//
// Parameters:
//
// - ctx: the context of the request
// - movieID: the ID of the movie
// - language: the language code (optional, defaults to "en-US")
// - page: the page number (optional, defaults to 1)
//
// Returns:
//
// - (*MovieReviewsResponse): the response containing the movie reviews
// - int: the HTTP status code
// - error: if there was an error fetching the movie reviews
func (c Client) GetMovieReviews(
	ctx context.Context,
	movieID int32,
	language string,
	page int32,
) (parsedResp *MovieReviewsResponse, statusCode int, err error) {
	// Create the HTTP request
	apiURL := fmt.Sprintf(GetMovieReviewsURL, fmt.Sprintf("%d", movieID))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, http.NoBody)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(ErrBuildingRequest, err)
	}

	// Add the Authorization header
	c.addAuthorizationToRequest(req)

	// Add query parameters
	q := req.URL.Query()
	c.addLanguageQueryParameter(q, language)
	c.addPageQueryParameter(q, page)
	req.URL.RawQuery = q.Encode()

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf(ErrAnErrOcurredDuringRequest, err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, resp.StatusCode, fmt.Errorf(ErrRequestFailed, resp.StatusCode, string(body))
	}

	// Parse the response
	parsedResp = &MovieReviewsResponse{}
	if parseErr := json.NewDecoder(resp.Body).Decode(parsedResp); parseErr != nil {
		return nil, resp.StatusCode, ErrResponseParsing
	}
	return parsedResp, resp.StatusCode, nil
}

// GetGenresMovieList fetches the list of movie genres
// 
// Parameters:
// 
// - ctx: the context of the request
// - language: the language code (optional, defaults to "en-US")
// 
// Returns:
// 
// - (*GenreListResponse): the response containing the list of movie genres
// - int: the HTTP status code
// - error: if there was an error fetching the movie genres
func (c Client) GetGenresMovieList(
	ctx context.Context,
	language string,
) (parsedResp *GenreListResponse, statusCode int, err error) {
	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, GetGenresMovieListURL, http.NoBody)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(ErrBuildingRequest, err)
	}

	// Add the Authorization header
	c.addAuthorizationToRequest(req)

	// Add query parameters
	q := req.URL.Query()
	c.addLanguageQueryParameter(q, language)
	req.URL.RawQuery = q.Encode()

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf(ErrAnErrOcurredDuringRequest, err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, resp.StatusCode, fmt.Errorf(ErrRequestFailed, resp.StatusCode, string(body))
	}

	// Parse the response
	parsedResp = &GenreListResponse{}
	if parseErr := json.NewDecoder(resp.Body).Decode(parsedResp); parseErr != nil {
		return nil, resp.StatusCode, ErrResponseParsing
	}
	return parsedResp, resp.StatusCode, nil
}