package gotmdbapi

const (
	// Query is the query parameter for search queries
	Query = "query"

	// IncludeAdult is the query parameter for including adult content
	IncludeAdult = "include_adult"

	// Page is the query parameter for pagination
	Page = "page"

	// Language is the query parameter for language
	Language = "language"

	// Region is the query parameter for region
	Region = "region"

	// Year is the query parameter for year
	Year = "year"

	// PrimaryReleaseYear is the query parameter for primary release year
	PrimaryReleaseYear = "primary_release_year"
	
	// PrimaryReleaseYearGTE is the query parameter for primary release year greater than or equal to
	PrimaryReleaseYearGTE = "primary_release_year.gte"
	
	// PrimaryReleaseYearLTE is the query parameter for primary release year less than or equal to
	PrimaryReleaseYearLTE = "primary_release_year.lte"
	
	// Certification is the query parameter for certification
	Certification = "certification"
	
	// CertificationCountry is the query parameter for certification country
	CertificationCountry = "certification_country"
	
	// CerificationGTE is the query parameter for certification greater than or equal to
	CerificationGTE = "certification.gte"
	
	// CertificationLTE is the query parameter for certification less than or equal to
	CertificationLTE = "certification.lte"
	
	// IncludeVideo is the query parameter for including videos
	IncludeVideo = "include_video"
	
	// ReleaseDateGTE is the query parameter for release date greater than or equal to
	ReleaseDateGTE = "release_date.gte"
	
	// ReleaseDateLTE is the query parameter for release date less than or equal to
	ReleaseDateLTE = "release_date.lte"
	
	// SortBy is the query parameter for sorting results
	SortBy = "sort_by"
	
	// VoteAverageGTE is the query parameter for vote average greater than or equal to
	VoteAverageGTE = "vote_average.gte"
	
	// VoteAverageLTE is the query parameter for vote average less than or equal to
	VoteAverageLTE = "vote_average.lte"

	// VoteCountGTE is the query parameter for vote count greater than or equal to
	VoteCountGTE = "vote_count.gte"
	
	// VoteCountLTE is the query parameter for vote count less than or equal to
	VoteCountLTE = "vote_count.lte"
	
	// WithGenres is the query parameter for filtering by genres
	WithGenres = "with_genres"
	
	// WithCompanies is the query parameter for filtering by companies
	WithCompanies = "with_companies"
	
	// WithKeywords is the query parameter for filtering by keywords
	WithKeywords = "with_keywords"
	
	// WithCast is the query parameter for filtering by cast
	WithCast = "with_cast"
	
	// WithCrew is the query parameter for filtering by crew
	WithCrew = "with_crew"
	
	// WithPeople is the query parameter for filtering by people
	WithPeople = "with_people"
	
	// WithOriginCountry is the query parameter for filtering by origin country
	WithOriginCountry = "with_origin_country"
	
	// WithOriginalLanguage is the query parameter for filtering by original language
	WithOriginalLanguage = "with_original_language"
	
	// WatchRegion is the query parameter for watch region
	WatchRegion = "watch_region"
	
	// WithReleaseType is the query parameter for release type
	// WithReleaseType = "with_release_type" // I don't know what means each value
	
	// WithRuntimeGTE is the query parameter for runtime greater than or equal to
	WithRuntimeGTE = "with_runtime.gte"
	
	// WithRuntimeLTE is the query parameter for runtime less than or equal to
	WithRuntimeLTE = "with_runtime.lte"
	
	// WithWatchMonetizationTypes is the query parameter for watch monetization types
	WithWatchMonetizationTypes = "with_watch_monetization_types"
	
	// WithWatchProviders is the query parameter for watch providers
	WithWatchProviders = "with_watch_providers"
	
	// WithoutCompanies is the query parameter for excluding companies
	WithoutCompanies = "without_companies"
	
	// WithoutGenres is the query parameter for excluding genres
	WithoutGenres = "without_genres"
	
	// WithoutKeywords is the query parameter for excluding keywords
	WithoutKeywords = "without_keywords"
	
	// WithoutWatchProviders is the query parameter for excluding watch providers
	WithoutWatchProviders = "without_watch_providers"
)


const (
	// ImageOriginalQualityURL is the image original quality URL
	ImageOriginalQualityURL = "https://image.tmdb.org/t/p/original/%s"
	
	// ImageVariableQualityURL is the image variable quality URL
	ImageVariableQualityURL = "https://image.tmdb.org/t/p/w%d/%s"
)

const (
	// GetMovieCreditsURL is the TMDB API URL for getting movie credits
	GetMovieCreditsURL = "https://api.themoviedb.org/3/movie/%s/credits"

	// GetMovieDetailsURL is the TMDB API URL for getting movie details
	GetMovieDetailsURL = "https://api.themoviedb.org/3/movie/%s"

	// GetMovieReviewsURL is the TMDB API URL for getting movie reviews
	GetMovieReviewsURL = "https://api.themoviedb.org/3/movie/%s/reviews"

	// GetNowPlayingMoviesURL is the TMDB API URL for getting now playing movies
	GetNowPlayingMoviesURL = "https://api.themoviedb.org/3/movie/now_playing"

	// GetPopularMoviesURL is the TMDB API URL for getting popular movies
	GetPopularMoviesURL = "https://api.themoviedb.org/3/movie/popular"

	// GetTopRatedMoviesURL is the TMDB API URL for getting top rated movies
	GetTopRatedMoviesURL = "https://api.themoviedb.org/3/movie/top_rated"

	// GetUpcomingMoviesURL is the TMDB API URL for getting upcoming movies
	GetUpcomingMoviesURL = "https://api.themoviedb.org/3/movie/upcoming"

	// SearchMoviesURL is the TMDB API URL for searching movies
	SearchMoviesURL = "https://api.themoviedb.org/3/search/movie"

	// SimilarMoviesURL is the TMDB API URL for getting similar movies
	SimilarMoviesURL = "https://api.themoviedb.org/3/movie/%s/similar"
	
	// DiscoverMoviesURL is the TMDB API URL for discovering movies
	DiscoverMoviesURL = "https://api.themoviedb.org/3/discover/movie"
	
	// GetGenresMovieListURL is the TMDB API URL for getting the list of movie genres
	GetGenresMovieListURL = "https://api.themoviedb.org/3/genre/movie/list"
)
