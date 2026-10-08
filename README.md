# go-tmdb-api

A small Go client for the [TMDB API](https://www.themoviedb.org/documentation/api), covering the movie endpoints: lists, search, details, credits, reviews, similar movies, genres and discover.

Module `github.com/ralvarezdev/go-tmdb-api` (package `gotmdbapi`, Go 1.25.4).

## Installation

```bash
go get github.com/ralvarezdev/go-tmdb-api
```

## Usage

`NewClient` takes a TMDB API key or read access token and sends it on every request as `Authorization: Bearer <key>`. It returns `ErrEmptyAPIKey` for an empty key. Every method returns `(response, httpStatusCode, error)`.

```go
client, err := gotmdbapi.NewClient(os.Getenv("TMDB_API_KEY"))
if err != nil {
	panic(err)
}

// ctx, language, page, region
resp, status, err := client.GetMoviesNowPlaying(context.Background(), "en-US", 1, "")
```

Discover with filters:

```go
resp, status, err := client.DiscoverMovies(ctx, &gotmdbapi.DiscoverMoviesQueryParameters{
	Language:   "en-US",
	Page:       1,
	SortBy:     gotmdbapi.SortByPopularityDesc,
	WithGenres: []string{"28"},
})
```

## API

- **`GetMoviesNowPlaying`, `GetMoviesPopular`, `GetMoviesTopRated`, `GetMoviesUpcoming`** — movie lists.
- **`SearchMovies`** — `/search/movie`.
- **`SimilarMovies`, `GetMovieDetails`, `GetMovieCredits`, `GetMovieReviews`** — per-movie endpoints.
- **`GetGenresMovieList`** — `/genre/movie/list`.
- **`DiscoverMovies`** — `/discover/movie` with `DiscoverMoviesQueryParameters`.

Also exported: response models (`model.go`), `SortByEnum` and `WatchMonetizationTypeEnums` (`enums.go`), endpoint and image URL constants (`constants.go`), `Add*QueryParameter` helpers (`query.go`), and the errors `ErrNilClient`, `ErrEmptyAPIKey`, `ErrResponseParsing`.

## Development

The tests in `types_test.go` call the live TMDB API and read the key from `TMDB_API_KEY`:

```bash
export TMDB_API_KEY=<your key>
go test ./...
```

## License

GNU General Public License v3.0.
