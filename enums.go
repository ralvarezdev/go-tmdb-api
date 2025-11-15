package gotmdbapi

type (
	// sortBy represents the sorting options for TMDB API requests
	sortBy string
	
	// watchMonetizationType represents the watch monetization types for TMDB API requests
	watchMonetizationType string
)

const (
	SortByPopularityAsc               sortBy = "popularity.asc"
	SortByValuePopularityDesc         sortBy = "popularity.desc"
	SortByValuePrimaryReleaseDateAsc  sortBy = "primary_release_date.asc"
	SortByValuePrimaryReleaseDateDesc sortBy = "primary_release_date.desc"
	SortByValueOriginalTitleAsc       sortBy = "original_title.asc"
	SortByValueOriginalTitleDesc      sortBy = "original_title.desc"
	SortByValueRevenueAsc             sortBy = "revenue.asc"
	SortByValueRevenueDesc            sortBy = "revenue.desc"
	SortByValueTitleAsc               sortBy = "title.asc"
	SortByValueTitleDesc              sortBy = "title.desc"
	SortByValueVoteAverageAsc         sortBy = "vote_average.asc"
	SortByValueVoteAverageDesc        sortBy = "vote_average.desc"
	SortByValueVoteCountAsc           sortBy = "vote_count.asc"
	SortByValueVoteCountDesc          sortBy      = "vote_count.desc"
)

const (
	WatchMonetizationTypeFlatrate watchMonetizationType = "flatrate"
	WatchMonetizationTypeFree     watchMonetizationType = "free"
	WatchMonetizationTypeAds      watchMonetizationType = "ads"
	WatchMonetizationTypeRent     watchMonetizationType = "rent"
	WatchMonetizationTypeBuy      watchMonetizationType = "buy"
)