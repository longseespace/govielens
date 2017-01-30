package movielens

import (
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

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

// MakeUserFromJSON creates a user from json data
func MakeUserFromJSON(j *simplejson.Json) *User {
	firstLoginString := j.GetPath("data", "details", "account", "firstLogin").MustString()
	firstLogin, _ := time.Parse(time.RFC3339, firstLoginString)
	user := &User{
		Email:               j.GetPath("data", "details", "account", "email").MustString(),
		UserName:            j.GetPath("data", "details", "account", "userName").MustString(),
		TimeAsMemberSeconds: j.GetPath("data", "details", "account", "timeAsMemberSeconds").MustInt64(),
		FirstLogin:          firstLogin,
		Preferences: &UserPreference{
			CanSendEmail:     j.GetPath("data", "details", "preferences", "canSendEmail").MustBool(),
			NumMoviesPerPage: j.GetPath("data", "details", "preferences", "numMoviesPerPage").MustInt(),
			MovieGroupTags:   j.GetPath("data", "details", "preferences", "movieGroupTags").MustStringArray(),
		},
		Recommender: &UserRecommender{
			EngineID:     j.GetPath("data", "details", "recommender", "engineId").MustString(),
			EngineWeight: j.GetPath("data", "details", "recommender", "engineWeight").MustFloat64(),
			PopWeight:    j.GetPath("data", "details", "recommender", "popWeight").MustFloat64(),
		},
	}
	return user
}
