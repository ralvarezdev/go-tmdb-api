package gotmdbapi

type (
	// sortBy represents the sorting options for TMDB API requests
	sortBy string

	// watchMonetizationType represents the watch monetization types for TMDB API requests
	watchMonetizationType string
)

const (
	SortByPopularityAsc               sortBy = "popularity.asc"
	SortByPopularityDesc         sortBy = "popularity.desc"
	SortByPrimaryReleaseDateAsc  sortBy = "primary_release_date.asc"
	SortByPrimaryReleaseDateDesc sortBy = "primary_release_date.desc"
	SortByOriginalTitleAsc       sortBy = "original_title.asc"
	SortByOriginalTitleDesc      sortBy = "original_title.desc"
	SortByRevenueAsc             sortBy = "revenue.asc"
	SortByRevenueDesc            sortBy = "revenue.desc"
	SortByTitleAsc               sortBy = "title.asc"
	SortByTitleDesc              sortBy = "title.desc"
	SortByVoteAverageAsc         sortBy = "vote_average.asc"
	SortByVoteAverageDesc        sortBy = "vote_average.desc"
	SortByVoteCountAsc           sortBy = "vote_count.asc"
	SortByVoteCountDesc          sortBy = "vote_count.desc"
)

const (
	WatchMonetizationTypeFlatrate watchMonetizationType = "flatrate"
	WatchMonetizationTypeFree     watchMonetizationType = "free"
	WatchMonetizationTypeAds      watchMonetizationType = "ads"
	WatchMonetizationTypeRent     watchMonetizationType = "rent"
	WatchMonetizationTypeBuy      watchMonetizationType = "buy"
)
