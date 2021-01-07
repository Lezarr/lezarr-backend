package movies

import (
	"time"

	"gorm.io/gorm"
)

// MovieSearchResults represent the search results returned by a search request
type MovieSearchResults struct {
	Page         int64   `json:"page"`
	TotalResults int64   `json:"total_results"`
	TotalPages   int64   `json:"total_pages"`
	Results      []Movie `json:"results"`
}

// MovieVideos holds the data gathered from TMDB's /movie/{movie_id}/videos
// and an additional URL field
type MovieVideos struct {
	gorm.Model `json:"-"`

	MovieID uint `json:"-"`

	Iso639_1  string `json:"iso_639_1"`
	Iso3166_1 string `json:"iso_3166_1"`
	ID        string `json:"key" gorm:"primarykey"`
	Name      string `json:"name"`
	Site      string `json:"site"`
	Size      int    `json:"size"`
	Type      string `json:"type"`
	URL       string `json:"url"`
}

// MovieGenres represent a TMDB genre
type MovieGenres struct {
	gorm.Model `json:"-"`

	MovieID uint `json:"-"`

	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// Movie represent a movie in Lezarr's database
type Movie struct {
	// Lezarr fields
	LezarrStatus string    `json:"lezarr_status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Resolution   int       `json:"resolution"`
	Language     string    `json:"language"`
	Path         string    `json:"path"`
	Score        int       `json:"score"`

	// Custom metadata (data gathered from other endpoints)
	Videos []MovieVideos `json:"videos" gorm:"foreignKey:MovieID"`

	// TheMovieDB fields - basically what TMDB returns for a movie
	Adult            bool          `json:"adult"`
	BackdropPath     string        `json:"backdrop_path"`
	Budget           int64         `json:"budget"`
	Genres           []MovieGenres `json:"genres" gorm:"foreignKey:MovieID"`
	Homepage         string        `json:"homepage"`
	ID               int64         `json:"id"`
	IMDbID           string        `json:"imdb_id"`
	OriginalLanguage string        `json:"original_language"`
	OriginalTitle    string        `json:"original_title"`
	Overview         string        `json:"overview"`
	Popularity       float32       `json:"popularity"`
	PosterPath       string        `json:"poster_path"`
	ReleaseDate      string        `json:"release_date"`
	Revenue          int64         `json:"revenue"`
	Runtime          int           `json:"runtime"`
	Status           string        `json:"status"`
	Tagline          string        `json:"tagline"`
	Title            string        `json:"title"`
	Video            bool          `json:"video"`
	VoteAverage      float32       `json:"vote_average"`
	VoteCount        int64         `json:"vote_count"`
}
