package service

import (
	"apigogame/src/domain"
	"strconv"
)

type TwitterServiceImpl struct {
	tRepo domain.TwitterRepo
	gRepo domain.GameRepo
}

func CreateTwitterServiceImpl(tRepo domain.TwitterRepo, gRepo domain.GameRepo) domain.TwitterService {
	return &TwitterServiceImpl{
		tRepo: tRepo,
	}
}

func (t *TwitterServiceImpl) FindAllTweet(count int) ([]string, error) {
	var twts []string
	tws, err := t.tRepo.GetAllTweet(count)
	if err != nil {
		return nil, err
	}

	for _, v := range tws {
		twts = append(twts, v.Text)
	}

	return twts, nil
}

func (t *TwitterServiceImpl) CreateTweet(stat string) (string, error) {
	tw, err := t.tRepo.CreateTweet(stat)
	if err != nil {
		return "", err
	}

	// Get username and status id
	name := tw.User.Name
	id := strconv.Itoa(int(tw.ID))

	return "https://twitter.com/" + name + "/status/" + id, nil
}
