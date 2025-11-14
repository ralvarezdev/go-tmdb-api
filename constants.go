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
)
