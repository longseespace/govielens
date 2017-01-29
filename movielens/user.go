package movielens

import "time"

// User stores user details
type User struct {
	Email               string           `json:"email,omitempty"`
	UserName            string           `json:"username,omitempty"`
	TimeAsMemberSeconds int64            `json:"timeAsMemberSeconds,omitempty"`
	FirstLogin          time.Time        `json:"firstLogin,omitempty"`
	NumRatings          int              `json:"numRatings,omitempty"`
	Preferences         *UserPreference  `json:"preferences,omitempty"`
	Recommender         *UserRecommender `json:"recommender,omitempty"`
}

// UserPreference stores various preferences
type UserPreference struct {
	CanSendEmail     bool     `json:"canSendEmail,omitempty"`
	NumMoviesPerPage int      `json:"numMoviesPerPage,omitempty"`
	MovieGroupTags   []string `json:"movieGroupTags"`
}

// UserRecommender stores recommender information
type UserRecommender struct {
	EngineID     string  `json:"engineId,omitempty"`
	EngineWeight float64 `json:"engineWeight,omitempty"`
	PopWeight    float64 `json:"popWeight,omitempty"`
}
