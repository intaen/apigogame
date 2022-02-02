package repo

import (
	"apigogame/src/domain"
	"apigogame/src/util"
	"errors"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type GameRepoImpl struct {
}

func CreateGameRepoImpl() domain.GameRepo {
	return &GameRepoImpl{}
}

func (g *GameRepoImpl) GetPredictAge(name string) (*domain.PredictAge, error) {
	// Consume third API
	client := resty.New()      // Create client
	var resp domain.PredictAge // Initialize new variable to catch response from 3rd party\
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("name", name).
		SetResult(&resp).
		Get(viper.GetString("url.predict-age"))
	util.Log3rdParty(res.Request.Method, res.Request.URL, "", string(res.Body()))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New("data not found")
	} else if res.IsSuccess() {
		return &resp, nil
	}

	return nil, nil
}

func (g *GameRepoImpl) GetPredictNationality(name string) (*domain.PredictNationality, error) {
	// Consume third API
	client := resty.New()              // Create client
	var resp domain.PredictNationality // Initialize new variable to catch response from 3rd party\
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("name", name).
		SetResult(&resp).
		Get(viper.GetString("url.predict-nat"))
	util.Log3rdParty(res.Request.Method, res.Request.URL, "", string(res.Body()))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New("data not found")
	} else if res.IsSuccess() {
		return &resp, nil
	}

	return nil, nil
}

func (g *GameRepoImpl) GetPredictGender(name string) (*domain.PredictGender, error) {
	// Consume third API
	client := resty.New()         // Create client
	var resp domain.PredictGender // Initialize new variable to catch response from 3rd party\
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("name", name).
		SetResult(&resp).
		Get(viper.GetString("url.predict-gender"))
	util.Log3rdParty(res.Request.Method, res.Request.URL, "", string(res.Body()))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New("data not found")
	} else if res.IsSuccess() {
		return &resp, nil
	}

	return nil, nil
}

func (g *GameRepoImpl) GetBoredActivity() (*domain.BoredActivity, error) {
	// Consume third API
	client := resty.New()         // Create client
	var resp domain.BoredActivity // Initialize new variable to catch response from 3rd party\
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&resp).
		Get(viper.GetString("url.random-act"))
	util.Log3rdParty(res.Request.Method, res.Request.URL, "", string(res.Body()))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New("data not found")
	} else if res.IsSuccess() {
		return &resp, nil
	}

	return nil, nil
}

func (g *GameRepoImpl) GetFunFact() (*domain.FactFun, error) {
	// Consume third API
	client := resty.New()   // Create client
	var resp domain.FactFun // Initialize new variable to catch response from 3rd party\
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&resp).
		Get(viper.GetString("url.fact-random"))
	if strings.Contains(string(res.Body()), "base64") {
		resp.Dark = resp.Dark[:50]
		resp.Light = resp.Light[:50]
		resp.Primary = resp.Primary[:50]
	}
	util.Log3rdParty(res.Request.Method, res.Request.URL, "", resp)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		resp.Text = "High heels were originally worn by men"
		util.LogError("data not found", "GetFunFact IsError")
		return &resp, nil
	} else if res.IsSuccess() {
		return &resp, nil
	}

	return nil, nil
}

func (g *GameRepoImpl) GetFactMath() (*domain.FactMath, error) {
	// Consume third API
	client := resty.New()    // Create client
	var resp domain.FactMath // Initialize new variable to catch response from 3rd party\
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&resp).
		Get(viper.GetString("url.fact-math"))
	util.Log3rdParty(res.Request.Method, res.Request.URL, "", string(res.Body()))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		resp.Text = "The bowler hat was invented as safety measure"
		util.LogError("data not found", "GetFactMath IsError")
		return &resp, nil
	} else if res.IsSuccess() {
		return &resp, nil
	}

	return nil, nil
}

func (g *GameRepoImpl) GetFactDog() ([]domain.FactDog, error) {
	// Consume third API
	client := resty.New()     // Create client
	var resp []domain.FactDog // Initialize new variable to catch response from 3rd party
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&resp).
		Get(viper.GetString("url.fact-dog"))
	util.Log3rdParty(res.Request.Method, res.Request.URL, "", string(res.Body()))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		resp[0].Fact = "Some cats are actually allergic to humans"
		util.LogError("data not found", "GetFactDog IsError")
		return resp, nil
	} else if res.IsSuccess() {
		return resp, nil
	}

	return nil, nil
}

func (g *GameRepoImpl) GetDogImg() (*domain.ImgDog, error) {
	// Consume third API
	client := resty.New()  // Create client
	var resp domain.ImgDog // Initialize new variable to catch response from 3rd party\
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&resp).
		Get(viper.GetString("url.random-dog"))
	util.Log3rdParty(res.Request.Method, res.Request.URL, "", string(res.Body()))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		resp.Message = "https://upload.wikimedia.org/wikipedia/commons/9/94/My_dog.jpg"
		util.LogError("data not found", "GetDogImg IsError")
		return &resp, nil
	} else if res.IsSuccess() {
		return &resp, nil
	}

	return nil, nil
}

func (g *GameRepoImpl) GetCatImg() (*domain.ImgCat, error) {
	// Consume third API
	client := resty.New()  // Create client
	var resp domain.ImgCat // Initialize new variable to catch response from 3rd party\
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&resp).
		Get(viper.GetString("url.random-cat"))
	util.Log3rdParty(res.Request.Method, res.Request.URL, "", string(res.Body()))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		resp.URL = "https://upload.wikimedia.org/wikipedia/commons/3/38/Adorable-animal-cat-20787.jpg"
		util.LogError("data not found", "GetCatImg IsError")
		return &resp, nil
	} else if res.IsSuccess() {
		return &resp, nil
	}

	return nil, nil
}

func (g *GameRepoImpl) GetDuckImg() (*domain.ImgDuck, error) {
	// Consume third API
	client := resty.New()   // Create client
	var resp domain.ImgDuck // Initialize new variable to catch response from 3rd party\
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&resp).
		Get(viper.GetString("url.random-duck"))
	util.Log3rdParty(res.Request.Method, res.Request.URL, "", string(res.Body()))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		resp.URL = "https://upload.wikimedia.org/wikipedia/commons/7/74/White_domesticated_duck%2C_stretching.jpg"
		util.LogError("data not found", "GetDuckImg IsError")
		return &resp, nil
	} else if res.IsSuccess() {
		return &resp, nil
	}

	return nil, nil
}

func (g *GameRepoImpl) GetRandomJoke(mode string) (*domain.AnyJoke, error) {
	// Consume third API
	client := resty.New()   // Create client
	var resp domain.AnyJoke // Initialize new variable to catch response from 3rd party
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&resp).
		Get(viper.GetString("url.random-joke") + mode)
	util.Log3rdParty(res.Request.Method, res.Request.URL, "", string(res.Body()))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New("data not found")
	} else if res.IsSuccess() {
		return &resp, nil
	}

	return nil, nil
}
