package gotmdbapi

type (
	// DateRange represents a date range with maximum and minimum dates
	DateRange struct {
		Maximum string `json:"maximum"`
		Minimum string `json:"minimum"`
	}

	// Crew represents a crew member in a movie
	Crew struct {
		Adult              bool    `json:"adult"`
		Gender             int32   `json:"gender"`
		ID                 int32   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		CreditID           string  `json:"credit_id"`
		Department         string  `json:"department"`
		Job                string  `json:"job"`
	}

	// Cast represents a cast member in a movie
	Cast struct {
		Adult              bool    `json:"adult"`
		Gender             int32   `json:"gender"`
		ID                 int32   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		CreditID           string  `json:"credit_id"`
		CastID             int32   `json:"cast_id"`
		Character          string  `json:"character"`
		Order              int32   `json:"order"`
	}

	// Genre represents a movie genre
	Genre struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	}

	// ProductionCompany represents a movie production company
	ProductionCompany struct {
		ID            int32  `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	}

	// ProductionCountry represents a movie production country
	ProductionCountry struct {
		// nolint:revive
		ISO3166_1 string `json:"iso_3166_1"`
		Name      string `json:"name"`
	}

	// SpokenLanguage represents a spoken language in a movie
	SpokenLanguage struct {
		EnglishName string `json:"english_name"`
		// nolint:revive
		ISO639_1 string `json:"iso_639_1"`
		Name     string `json:"name"`
	}

	// SimpleMovie represents a simplified movie structure
	SimpleMovie struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDs         []int32 `json:"genre_ids"`
		ID               int32   `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		Popularity       float32 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Video            bool    `json:"video"`
		VoteAverage      float32 `json:"vote_average"`
		VoteCount        int32   `json:"vote_count"`
	}

	// DateMovieListResponse represents a movie list response with date range
	DateMovieListResponse struct {
		Dates        DateRange     `json:"dates"`
		Page         int32         `json:"page"`
		Results      []SimpleMovie `json:"results"`
		TotalPages   int32         `json:"total_pages"`
		TotalResults int32         `json:"total_results"`
	}

	// MovieListResponse represents a generic movie list response
	MovieListResponse struct {
		Page         int32         `json:"page"`
		Results      []SimpleMovie `json:"results"`
		TotalPages   int32         `json:"total_pages"`
		TotalResults int32         `json:"total_results"`
	}

	// MovieCreditsResponse represents a movie credits response
	MovieCreditsResponse struct {
		ID   int32  `json:"id"`
		Cast []Cast `json:"cast"`
		Crew []Crew `json:"crew"`
	}

	// MovieDetailsResponse represents a movie details response
	MovieDetailsResponse struct {
		Adult               bool                `json:"adult"`
		BackdropPath        string              `json:"backdrop_path"`
		BelongsToCollection interface{}         `json:"belongs_to_collection"`
		Budget              int64               `json:"budget"`
		Genres              []Genre             `json:"genres"`
		Homepage            string              `json:"homepage"`
		ID                  int32               `json:"id"`
		ImdbID              string              `json:"imdb_id"`
		OriginalLanguage    string              `json:"original_language"`
		OriginalTitle       string              `json:"original_title"`
		Overview            string              `json:"overview"`
		Popularity          float32             `json:"popularity"`
		PosterPath          string              `json:"poster_path"`
		ProductionCompanies []ProductionCompany `json:"production_companies"`
		ProductionCountries []ProductionCountry `json:"production_countries"`
		ReleaseDate         string              `json:"release_date"`
		Revenue             int64               `json:"revenue"`
		Runtime             int32               `json:"runtime"`
		SpokenLanguages     []SpokenLanguage    `json:"spoken_languages"`
		Status              string              `json:"status"`
		Tagline             string              `json:"tagline"`
		Title               string              `json:"title"`
		Video               bool                `json:"video"`
		VoteAverage         float32             `json:"vote_average"`
		VoteCount           int32               `json:"vote_count"`
	}

	// AuthorDetails represents the details of an author in a review
	AuthorDetails struct {
		Name       string  `json:"name"`
		Username   string  `json:"username"`
		AvatarPath string  `json:"avatar_path"`
		Rating     float32 `json:"rating"`
	}

	// Review represents a movie review
	Review struct {
		Author        string        `json:"author"`
		AuthorDetails AuthorDetails `json:"author_details"`
		Content       string        `json:"content"`
		CreatedAt     string        `json:"created_at"`
		ID            string        `json:"id"`
		UpdatedAt     string        `json:"updated_at"`
		URL           string        `json:"url"`
	}

	// MovieReviewsResponse represents a movie reviews response
	MovieReviewsResponse struct {
		ID           int32    `json:"id"`
		Page         int32    `json:"page"`
		Results      []Review `json:"results"`
		TotalPages   int32    `json:"total_pages"`
		TotalResults int32    `json:"total_results"`
	}
)
