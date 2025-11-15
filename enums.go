package gotmdbapi

type (
	// SortByEnum represents the sorting options for TMDB API requests
	SortByEnum string

	// WatchMonetizationTypeEnums represents the watch monetization types for TMDB API requests
	WatchMonetizationTypeEnums string
)

const (
	SortByPopularityAsc          SortByEnum = "popularity.asc"
	SortByPopularityDesc         SortByEnum = "popularity.desc"
	SortByPrimaryReleaseDateAsc  SortByEnum = "primary_release_date.asc"
	SortByPrimaryReleaseDateDesc SortByEnum = "primary_release_date.desc"
	SortByOriginalTitleAsc       SortByEnum = "original_title.asc"
	SortByOriginalTitleDesc      SortByEnum = "original_title.desc"
	SortByRevenueAsc             SortByEnum = "revenue.asc"
	SortByRevenueDesc            SortByEnum = "revenue.desc"
	SortByTitleAsc               SortByEnum = "title.asc"
	SortByTitleDesc              SortByEnum = "title.desc"
	SortByVoteAverageAsc         SortByEnum = "vote_average.asc"
	SortByVoteAverageDesc        SortByEnum = "vote_average.desc"
	SortByVoteCountAsc           SortByEnum = "vote_count.asc"
	SortByVoteCountDesc          SortByEnum = "vote_count.desc"
)

const (
	WatchMonetizationTypeFlatrate WatchMonetizationTypeEnums = "flatrate"
	WatchMonetizationTypeFree     WatchMonetizationTypeEnums = "free"
	WatchMonetizationTypeAds      WatchMonetizationTypeEnums = "ads"
	WatchMonetizationTypeRent     WatchMonetizationTypeEnums = "rent"
	WatchMonetizationTypeBuy      WatchMonetizationTypeEnums = "buy"
)
