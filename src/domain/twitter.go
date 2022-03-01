package domain

import (
	"github.com/dghubble/go-twitter/twitter"
)

type TwitterService interface {
	FindAllTweet(int) ([]string, error)
	CreateTweet(string) (string, error)
}

type TwitterRepo interface {
	GetAllTweet(int) ([]twitter.Tweet, error)
	CreateTweet(string) (*twitter.Tweet, error)
}

type (
	InputTimeline struct {
		TotalDataShow int `json:"total_data_show"`
	}
)
