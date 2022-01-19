package domain

type GameService interface {
	PredictionByName(string) (*GameName, error)
	Punishment() (*Punishment, error)
}

type GameRepo interface {
	GetPredictAge(string) (*PredictAge, error)
	GetPredictNationality(string) (*PredictNationality, error)
	GetRandomActivity() (*RandomActivity, error)
}

type (
	InputGameName struct {
		Name string `form:"name" json:"name"`
	}

	ExampleGameName struct {
		Name string `json:"Name" example:"Admin"`
	}

	GameName struct {
		PredictAge         int         `json:"predict_age"`
		PredictNationality interface{} `json:"predict_nationality"`
	}

	Punishment struct {
		RandomActivity string
	}

	// API
	PredictAge struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Count int    `json:"count"`
	}

	PredictNationality struct {
		Name    string          `json:"name"`
		Country []DetailCountry `json:"country"`
	}
	DetailCountry struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
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
)
