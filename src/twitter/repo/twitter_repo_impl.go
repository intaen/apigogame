package repo

import (
	"errors"
	"fmt"

	"apigogame/src/domain"

	"github.com/dghubble/go-twitter/twitter"
)

type TwitterRepoImpl struct {
	twt *twitter.Client
}

func CreateTwitterRepoImpl(twt *twitter.Client) domain.TwitterRepo {
	return &TwitterRepoImpl{
		twt: twt,
	}
}

func (t *TwitterRepoImpl) GetAllTweet(count int) ([]twitter.Tweet, error) {
	if t.twt == nil {
		return nil, errors.New("can't access this")
	}

	tweets, r, err := t.twt.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count:           count,
		IncludeEntities: twitter.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	if r.StatusCode >= 299 {
		fmt.Println("StatusCode", r.StatusCode)
		return nil, errors.New("error with status " + r.Status)
	}

	return tweets, nil
}

func (t *TwitterRepoImpl) CreateTweet(stat string) (*twitter.Tweet, error) {
	if t.twt == nil {
		return nil, errors.New("can't access this")
	}

	tweet, r, err := t.twt.Statuses.Update(stat, &twitter.StatusUpdateParams{
		Status: stat,
	})
	if err != nil {
		return nil, err
	}

	if r.StatusCode >= 299 {
		fmt.Println("StatusCode", r.StatusCode)
		return nil, errors.New("error with status " + r.Status)
	}

	return tweet, nil
}
