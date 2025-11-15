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
