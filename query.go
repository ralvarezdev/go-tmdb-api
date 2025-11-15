package gotmdbapi

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type (
	// DiscoverMoviesQueryParameters represents the query parameters for discovering movies
	DiscoverMoviesQueryParameters struct {
		Certification              string
		CertificationCountry       string
		CertificationGTE           string
		CertificationLTE           string
		IncludeAdult               bool
		IncludeVideo               bool
		Language                   string
		PrimaryReleaseYear         int32
		PrimaryReleaseYearGTE      int32
		PrimaryReleaseYearLTE      int32
		Page                       int32
		Region                     string
		ReleaseDateGTE             string
		ReleaseDateLTE             string
		SortBy                     sortBy
		VoteAverageGTE             float32
		VoteAverageLTE             float32
		VoteCountGTE               float32
		VoteCountLTE               float32
		WithGenres                 []string
		WithCompanies              []string
		WithKeywords               []string
		WithCast                   []string
		WithCrew                   []string
		WithPeople                 []string
		WithOriginCountry          string
		WithOriginalLanguage       string
		WatchRegion                string
		WithRuntimeGTE             int32
		WithRuntimeLTE             int32
		WithWatchMonetizationTypes []watchMonetizationType
		WithWatchProviders         []string
		WithoutCompanies           string
		WithoutGenres              string
		WithoutKeywords            string
		Year                       int32
	}
)

// AddLanguageQueryParameter adds the language query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - language: the language code (optional, defaults to "en-US")
func AddLanguageQueryParameter(
	query url.Values,
	language string,
) {
	if language != "" {
		query.Add(Language, language)
	}
}

// AddIncludeAdultQueryParameter adds the include adult query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - includeAdult: whether to include adult content
func AddIncludeAdultQueryParameter(
	query url.Values,
	includeAdult bool,
) {
	query.Add(IncludeAdult, strconv.FormatBool(includeAdult))
}

// AddIncludeVideoQueryParameter adds the include video query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - includeVideo: whether to include video content
func AddIncludeVideoQueryParameter(
	query url.Values,
	includeVideo bool,
) {
	query.Add(IncludeVideo, strconv.FormatBool(includeVideo))
}

// AddPageQueryParameter adds the page query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - page: the page number (optional, defaults to 1)
func AddPageQueryParameter(
	query url.Values,
	page int32,
) {
	if page > 0 {
		query.Add(Page, fmt.Sprintf("%d", page))
	}
}

// AddRegionQueryParameter adds the region query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - region: the region code (optional)
func AddRegionQueryParameter(
	query url.Values,
	region string,
) {
	if region != "" {
		query.Add(Region, region)
	}
}

// AddPrimaryReleaseYearQueryParameter adds the primary release year query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - primaryReleaseYear: the primary release year (optional)
func AddPrimaryReleaseYearQueryParameter(
	query url.Values,
	primaryReleaseYear int32,
) {
	if primaryReleaseYear > 0 {
		query.Add(PrimaryReleaseYear, fmt.Sprintf("%d", primaryReleaseYear))
	}
}

// AddPrimaryReleaseYearGTEQueryParameter adds the primary release year greater than or equal to query parameter to the
// HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - primaryReleaseYearGTE: the primary release year greater than or equal to (optional)
func AddPrimaryReleaseYearGTEQueryParameter(
	query url.Values,
	primaryReleaseYearGTE int32,
) {
	if primaryReleaseYearGTE > 0 {
		query.Add(PrimaryReleaseYearGTE, fmt.Sprintf("%d", primaryReleaseYearGTE))
	}
}

// AddPrimaryReleaseYearLTEQueryParameter adds the primary release year less than or equal to query parameter to the
// HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - primaryReleaseYearLTE: the primary release year less than or equal to (optional)
func AddPrimaryReleaseYearLTEQueryParameter(
	query url.Values,
	primaryReleaseYearLTE int32,
) {
	if primaryReleaseYearLTE > 0 {
		query.Add(PrimaryReleaseYearLTE, fmt.Sprintf("%d", primaryReleaseYearLTE))
	}
}

// AddYearQueryParameter adds the year query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - year: the year (optional)
func AddYearQueryParameter(
	query url.Values,
	year int32,
) {
	if year > 0 {
		query.Add(Year, fmt.Sprintf("%d", year))
	}
}

// AddCertificationQueryParameter adds the certification query parameter to the HTTP request query parameters
//
// Parameters:
//
// - query: the HTTP request query parameters
// - certification: the certification value
func AddCertificationQueryParameter(
	query url.Values,
	certification string,
) {
	query.Add(Certification, certification)
}

// AddCertificationCountryQueryParameter adds the certification country query parameter to the HTTP request query
// parameters
//
// Parameters:
//
// - req: the HTTP request query parameters
// - certificationCountry: the certification country value
func AddCertificationCountryQueryParameter(
	query url.Values,
	certificationCountry string,
) {
	query.Add(CertificationCountry, certificationCountry)
}

// AddCertificationGTEQueryParameter adds the certification.gte query parameter to the HTTP request query parameters
//
// Parameters:
//
// - query: the HTTP request query parameters
// - certificationGTE: the certification.gte value
func AddCertificationGTEQueryParameter(
	query url.Values,
	certificationGTE string,
) {
	query.Add(CerificationGTE, certificationGTE)
}

// AddCertificationLTEQueryParameter adds the certification.lte query parameter to the HTTP request query parameters
//
// Parameters:
//
// - query: the HTTP request query parameters
// - certificationLTE: the certification.lte value
func AddCertificationLTEQueryParameter(
	query url.Values,
	certificationLTE string,
) {
	query.Add(CertificationLTE, certificationLTE)
}

// AddReleaseDateGTEQueryParameter adds the release_date.gte query parameter to the HTTP request query parameters
//
// Parameters:
//
// - query: the HTTP request query parameters
// - releaseDateGTE: the release_date.gte value
func AddReleaseDateGTEQueryParameter(
	query url.Values,
	releaseDateGTE string,
) {
	query.Add(ReleaseDateGTE, releaseDateGTE)
}

// AddReleaseDateLTEQueryParameter adds the release_date.lte query parameter to the HTTP request query parameters
//
// Parameters:
//
// - query: the HTTP request query parameters
// - releaseDateLTE: the release_date.lte value
func AddReleaseDateLTEQueryParameter(
	query url.Values,
	releaseDateLTE string,
) {
	query.Add(ReleaseDateLTE, releaseDateLTE)
}

// AddSortByQueryParameter adds the sort_by query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - sortBy: the sort by value
func AddSortByQueryParameter(
	query url.Values,
	sortBy sortBy,
) {
	if sortBy != "" {
		query.Add(SortBy, string(sortBy))
	}
}

// AddVoteAverageGTEQueryParameter adds the vote_average.gte query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - voteAverageGTE: the vote average greater than or equal to value
func AddVoteAverageGTEQueryParameter(
	query url.Values,
	voteAverageGTE float32,
) {
	if voteAverageGTE > 0 {
		query.Add(VoteAverageGTE, fmt.Sprintf("%.1f", voteAverageGTE))
	}
}

// AddVoteAverageLTEQueryParameter adds the vote_average.lte query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - voteAverageLTE: the vote average less than or equal to value
func AddVoteAverageLTEQueryParameter(
	query url.Values,
	voteAverageLTE float32,
) {
	if voteAverageLTE > 0 {
		query.Add(VoteAverageLTE, fmt.Sprintf("%.1f", voteAverageLTE))
	}
}

// AddVoteCountGTEQueryParameter adds the vote_count.gte query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - voteCountGTE: the vote count greater than or equal to value
func AddVoteCountGTEQueryParameter(
	query url.Values,
	voteCountGTE float32,
) {
	if voteCountGTE > 0 {
		query.Add(VoteCountGTE, fmt.Sprintf("%.1f", voteCountGTE))
	}
}

// AddVoteCountLTEQueryParameter adds the vote_count.lte query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - voteCountLTE: the vote count less than or equal to value
func AddVoteCountLTEQueryParameter(
	query url.Values,
	voteCountLTE float32,
) {
	if voteCountLTE > 0 {
		query.Add(VoteCountLTE, fmt.Sprintf("%.1f", voteCountLTE))
	}
}

// AddWithGenresQueryParameter adds the with_genres query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withGenres: the list of genre IDs
func AddWithGenresQueryParameter(
	query url.Values,
	withGenres []string,
) {
	if len(withGenres) > 0 {
		query.Add(WithGenres, url.QueryEscape(strings.Join(withGenres, ",")))
	}
}

// AddWithCompaniesQueryParameter adds the with_companies query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withCompanies: the list of company IDs
func AddWithCompaniesQueryParameter(
	query url.Values,
	withCompanies []string,
) {
	if len(withCompanies) > 0 {
		query.Add(WithCompanies, url.QueryEscape(strings.Join(withCompanies, ",")))
	}
}

// AddWithKeywordsQueryParameter adds the with_keywords query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withKeywords: the list of keyword IDs
func AddWithKeywordsQueryParameter(
	query url.Values,
	withKeywords []string,
) {
	if len(withKeywords) > 0 {
		query.Add(WithKeywords, url.QueryEscape(strings.Join(withKeywords, ",")))
	}
}

// AddWithCastQueryParameter adds the with_cast query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withCast: the list of cast IDs
func AddWithCastQueryParameter(
	query url.Values,
	withCast []string,
) {
	if len(withCast) > 0 {
		query.Add(WithCast, url.QueryEscape(strings.Join(withCast, ",")))
	}
}

// AddWithCrewQueryParameter adds the with_crew query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withCrew: the list of crew IDs
func AddWithCrewQueryParameter(
	query url.Values,
	withCrew []string,
) {
	if len(withCrew) > 0 {
		query.Add(WithCrew, url.QueryEscape(strings.Join(withCrew, ",")))
	}
}

// AddWithPeopleQueryParameter adds the with_people query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withPeople: the list of people IDs
func AddWithPeopleQueryParameter(
	query url.Values,
	withPeople []string,
) {
	if len(withPeople) > 0 {
		query.Add(WithPeople, url.QueryEscape(strings.Join(withPeople, ",")))
	}
}

// AddWithOriginCountryQueryParameter adds the with_origin_country query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withOriginCountry: the origin country code
func AddWithOriginCountryQueryParameter(
	query url.Values,
	withOriginCountry string,
) {
	if withOriginCountry != "" {
		query.Add(WithOriginCountry, withOriginCountry)
	}
}

// AddWithOriginalLanguageQueryParameter adds the with_original_language query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withOriginalLanguage: the original language code
func AddWithOriginalLanguageQueryParameter(
	query url.Values,
	withOriginalLanguage string,
) {
	if withOriginalLanguage != "" {
		query.Add(WithOriginalLanguage, withOriginalLanguage)
	}
}

// AddWatchRegionQueryParameter adds the watch_region query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - watchRegion: the watch region code
func AddWatchRegionQueryParameter(
	query url.Values,
	watchRegion string,
) {
	if watchRegion != "" {
		query.Add(WatchRegion, watchRegion)
	}
}

// AddWithRuntimeGTEQueryParameter adds the with_runtime.gte query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withRuntimeGTE: the runtime greater than or equal to value
func AddWithRuntimeGTEQueryParameter(
	query url.Values,
	withRuntimeGTE int32,
) {
	if withRuntimeGTE > 0 {
		query.Add(WithRuntimeGTE, fmt.Sprintf("%d", withRuntimeGTE))
	}
}

// AddWithRuntimeLTEQueryParameter adds the with_runtime.lte query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withRuntimeLTE: the runtime less than or equal to value
func AddWithRuntimeLTEQueryParameter(
	query url.Values,
	withRuntimeLTE int32,
) {
	if withRuntimeLTE > 0 {
		query.Add(WithRuntimeLTE, fmt.Sprintf("%d", withRuntimeLTE))
	}
}

// AddWithWatchMonetizationTypesQueryParameter adds the with_watch_monetization_types query parameter to the HTTP
// request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withWatchMonetizationTypes: the list of watch monetization types
func AddWithWatchMonetizationTypesQueryParameter(
	query url.Values,
	withWatchMonetizationTypes []watchMonetizationType,
) {
	if len(withWatchMonetizationTypes) > 0 {
		strTypes := make([]string, len(withWatchMonetizationTypes))
		for i, t := range withWatchMonetizationTypes {
			strTypes[i] = string(t)
		}
		query.Add(WithWatchMonetizationTypes, url.QueryEscape(strings.Join(strTypes, ",")))
	}
}

// AddWithWatchProvidersQueryParameter adds the with_watch_providers query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withWatchProviders: the list of watch provider IDs
func AddWithWatchProvidersQueryParameter(
	query url.Values,
	withWatchProviders []string,
) {
	if len(withWatchProviders) > 0 {
		query.Add(WithWatchProviders, url.QueryEscape(strings.Join(withWatchProviders, ",")))
	}
}

// AddWithoutCompaniesQueryParameter adds the without_companies query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withoutCompanies: the list of company IDs to exclude
func AddWithoutCompaniesQueryParameter(
	query url.Values,
	withoutCompanies string,
) {
	if withoutCompanies != "" {
		query.Add(WithoutCompanies, withoutCompanies)
	}
}

// AddWithoutGenresQueryParameter adds the without_genres query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withoutGenres: the list of genre IDs to exclude
func AddWithoutGenresQueryParameter(
	query url.Values,
	withoutGenres string,
) {
	if withoutGenres != "" {
		query.Add(WithoutGenres, withoutGenres)
	}
}

// AddWithoutKeywordsQueryParameter adds the without_keywords query parameter to the HTTP request
//
// Parameters:
//
// - query: the HTTP request query parameters
// - withoutKeywords: the list of keyword IDs to exclude
func AddWithoutKeywordsQueryParameter(
	query url.Values,
	withoutKeywords string,
) {
	if withoutKeywords != "" {
		query.Add(WithoutKeywords, withoutKeywords)
	}
}

// AddMovieListsQueryParameters adds the query parameters for movie lists to the HTTP request
//
// Parameters:
//
// - req: the HTTP request
// - language: the language code (optional, defaults to "en-US")
// - page: the page number (optional, defaults to 1)
// - region: the region code (optional)
func AddMovieListsQueryParameters(
	req *http.Request,
	language string,
	page int32,
	region string,
) {
	q := req.URL.Query()
	AddLanguageQueryParameter(q, language)
	AddPageQueryParameter(q, page)
	AddRegionQueryParameter(q, region)
	req.URL.RawQuery = q.Encode()
}

// AddSearchMoviesQueryParameters adds the query parameters for searching movies to the HTTP request
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
func AddSearchMoviesQueryParameters(
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
	AddIncludeAdultQueryParameter(q, includeAdult)
	AddLanguageQueryParameter(q, language)
	AddPrimaryReleaseYearQueryParameter(q, primaryReleaseYear)
	AddPageQueryParameter(q, page)
	AddRegionQueryParameter(q, region)
	AddYearQueryParameter(q, year)
	req.URL.RawQuery = q.Encode()
}

// AddSimilarMoviesQueryParameters adds the query parameters for similar movies to the HTTP request
//
// Parameters:
//
// - req: the HTTP request
// - language: the language code (optional, defaults to "en-US")
// - page: the page number (optional, defaults to 1)
func AddSimilarMoviesQueryParameters(
	req *http.Request,
	language string,
	page int32,
) {
	q := req.URL.Query()
	AddLanguageQueryParameter(q, language)
	AddPageQueryParameter(q, page)
	req.URL.RawQuery = q.Encode()
}

// AddGenreMovieListQueryParameters adds the query parameters for genre movie lists to the HTTP request
//
// Parameters:
//
// - req: the HTTP request
// - queryParameters: the query parameters for discovering movies
func AddGenreMovieListQueryParameters(
	req *http.Request,
	queryParameters *DiscoverMoviesQueryParameters,
) {
	// If query parameters are nil, return
	if queryParameters == nil {
		return
	}

	q := req.URL.Query()
	AddCertificationQueryParameter(q, queryParameters.Certification)
	AddCertificationCountryQueryParameter(q, queryParameters.CertificationCountry)
	AddCertificationGTEQueryParameter(q, queryParameters.CertificationGTE)
	AddCertificationLTEQueryParameter(q, queryParameters.CertificationLTE)
	AddIncludeAdultQueryParameter(q, queryParameters.IncludeAdult)
	AddIncludeVideoQueryParameter(q, queryParameters.IncludeVideo)
	AddLanguageQueryParameter(q, queryParameters.Language)
	AddPageQueryParameter(q, queryParameters.Page)
	AddPrimaryReleaseYearQueryParameter(q, queryParameters.PrimaryReleaseYear)
	AddPrimaryReleaseYearGTEQueryParameter(q, queryParameters.PrimaryReleaseYearGTE)
	AddPrimaryReleaseYearLTEQueryParameter(q, queryParameters.PrimaryReleaseYearLTE)
	AddRegionQueryParameter(q, queryParameters.Region)
	AddReleaseDateGTEQueryParameter(q, queryParameters.ReleaseDateGTE)
	AddReleaseDateLTEQueryParameter(q, queryParameters.ReleaseDateLTE)
	AddSortByQueryParameter(q, queryParameters.SortBy)
	AddVoteAverageGTEQueryParameter(q, queryParameters.VoteAverageGTE)
	AddVoteAverageLTEQueryParameter(q, queryParameters.VoteAverageLTE)
	AddVoteCountGTEQueryParameter(q, queryParameters.VoteCountGTE)
	AddVoteCountLTEQueryParameter(q, queryParameters.VoteCountLTE)
	AddWithGenresQueryParameter(q, queryParameters.WithGenres)
	AddWithCompaniesQueryParameter(q, queryParameters.WithCompanies)
	AddWithKeywordsQueryParameter(q, queryParameters.WithKeywords)
	AddWithCastQueryParameter(q, queryParameters.WithCast)
	AddWithCrewQueryParameter(q, queryParameters.WithCrew)
	AddWithPeopleQueryParameter(q, queryParameters.WithPeople)
	AddWithOriginCountryQueryParameter(q, queryParameters.WithOriginCountry)
	AddWithOriginalLanguageQueryParameter(q, queryParameters.WithOriginalLanguage)
	AddWatchRegionQueryParameter(q, queryParameters.WatchRegion)
	AddWithRuntimeGTEQueryParameter(q, queryParameters.WithRuntimeGTE)
	AddWithRuntimeLTEQueryParameter(q, queryParameters.WithRuntimeLTE)
	AddWithWatchMonetizationTypesQueryParameter(q, queryParameters.WithWatchMonetizationTypes)
	AddWithWatchProvidersQueryParameter(q, queryParameters.WithWatchProviders)
	AddWithoutCompaniesQueryParameter(q, queryParameters.WithoutCompanies)
	AddWithoutGenresQueryParameter(q, queryParameters.WithoutGenres)
	AddWithoutKeywordsQueryParameter(q, queryParameters.WithoutKeywords)
	AddYearQueryParameter(q, queryParameters.Year)
	req.URL.RawQuery = q.Encode()
}
