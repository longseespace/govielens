package movielens

import (
	"fmt"
	"time"
)

// Date type
type Date time.Time

// MarshalJSON implements Marshaler
func (d Date) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(d).Format("2006-01-02"))
	return []byte(stamp), nil
}

// UnmarshalJSON implements Unmarshaler
func (d Date) UnmarshalJSON(b []byte) error {
	s := string(b)
	t, err := time.Parse("2006-01-02", s)
	d = Date(t)
	fmt.Println("dd", t, err)
	return err
}

// ExploreResponseWrapper represents response to an Explore() call
type ExploreResponseWrapper struct {
	Data *ExploreResponse `json:"data,omitempty"`
}

// ExploreResponse ExploreResponse Details
type ExploreResponse struct {
	Title         string          `json:"title,omitempty"`
	Description   string          `json:"description,omitempty"`
	URL           string          `json:"url,omitempty"`
	SearchResults []*SearchResult `json:"searchResults,omitempty"`
	Pager         *Pager          `json:"pager,omitempty"`
}

// SearchResult represents a single movie search result
type SearchResult struct {
	MovieID       int           `json:"movieId,omitempty"`
	Movie         Movie         `json:"movie,omitempty"`
	MovieUserData MovieUserData `json:"movieUserData,omitempty"`
}

// Movie represents a movie
type Movie struct {
	ID                int      `json:"movieId,omitempty"`
	TmdbID            int      `json:"tmdbMovieId,omitempty"`
	ImdbID            string   `json:"imdbMovieId,omitempty"`
	Title             string   `json:"title,omitempty"`
	OriginalTitle     string   `json:"originalTitle,omitempty"`
	MPAA              string   `json:"mpaa,omitempty"`
	Runtime           int      `json:"runtime,omitempty"`
	ReleaseDate       string   `json:"releaseDate,omitempty"`
	DVDReleaseDate    string   `json:"dvdReleaseDate,omitempty"`
	Genres            []string `json:"genres,omitempty"`
	Languages         []string `json:"languages,omitempty"`
	Directors         []string `json:"directors,omitempty"`
	Actors            []string `json:"actors,omitempty"`
	PosterPath        string   `json:"posterPath,omitempty"`
	BackdropPaths     []string `json:"backdropPaths,omitempty"`
	YoutubeTrailerIDs []string `json:"youtubeTrailerIds,omitempty"`
	PlotSummary       string   `json:"plotSummary,omitempty"`
	NumRatings        int      `json:"numRatings,omitempty"`
	AvgRating         float64  `json:"avgRating,omitempty"`
	ReleaseYear       int      `json:"releaseYear,omitempty"`
}

// MovieUserData represents user data on a movie
type MovieUserData struct {
	UserID            int     `json:"userId,omitempty"`
	MovieID           int     `json:"movieId,omitempty"`
	Rating            float64 `json:"rating,omitempty"`
	Prediction        float64 `json:"prediction,omitempty"`
	Wishlish          bool    `json:"wishlish,omitempty"`
	Hidden            bool    `json:"hidden,omitempty"`
	PredictionDetails PredictionDetail
}

// PredictionComponent represents prediction model component
type PredictionComponent struct {
	Prediction float64 `json:"pred,omitempty"`
	Type       string  `json:"type,omitempty"`
}

// PredictionDetail represents details of a prediction
type PredictionDetail struct {
	PrimaryPrediction     float64 `json:"primaryPrediction,omitempty"`
	PrimaryPredictionType string  `json:"primaryPredictionType,omitempty"`
	Components            []PredictionComponent
}
