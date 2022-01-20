package domain

type GameService interface {
	PredictionByName(string) (*GameName, error)
	DoYouDare() (*CheckDare, error)
	CheckFact() (*CheckFact, error)
	CheckImg() (*CheckImg, error)
}

type GameRepo interface {
	GetPredictAge(string) (*PredictAge, error)
	GetPredictNationality(string) (*PredictNationality, error)
	GetPredictGender(string) (*PredictGender, error)
	GetRandomActivity() (*RandomActivity, error)
	GetFactMath() (*FactMath, error)
	GetRandomImg() (*RandomDogImg, error)
}

type (
	InputGameName struct {
		Name string `form:"name" json:"name" binding:"required"`
	}

	GameName struct {
		PredictAge         int         `json:"predict_age"`
		PredictNationality interface{} `json:"predict_nationality"`
		PredictGender      string      `json:"predict_gender"`
	}

	CheckDare struct {
		RandomActivity string `json:"random_activity"`
	}

	CheckFact struct {
		Fact string `json:"random_fact"`
	}

	CheckImg struct {
		URL string `json:"url"`
	}

	// API
	PredictAge struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Count int    `json:"count"`
	}

	PredictNationality struct {
		Name      string            `json:"name"`
		Countries []DetailCountries `json:"country"`
	}
	DetailCountries struct {
		CountryID   string  `json:"country_id"`
		CountryName string  `json:"country_name"`
		Probability float64 `json:"probability"`
	}

	PredictGender struct {
		Name        string  `json:"name"`
		Gender      string  `json:"gender"`
		Probability float64 `json:"probability"`
		Count       int     `json:"count"`
	}

	RandomActivity struct {
		Activity      string  `json:"activity"`
		Type          string  `json:"type"`
		Participants  int     `json:"participants"`
		Price         float32 `json:"price"`
		Link          string  `json:"link"`
		Key           string  `json:"key"`
		Accessibility float32 `json:"accessibility"`
	}

	FactMath struct {
		Text   string `json:"text"`
		Number int    `json:"number"`
		Found  bool   `json:"bool"`
		Type   string `json:"type"`
	}

	RandomDogImg struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}
)
