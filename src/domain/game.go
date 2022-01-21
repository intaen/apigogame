package domain

type GameService interface {
	PredictionByName(string) (*PredictName, error)
	DoYouDare() (*CheckDare, error)
	CheckFact() (*CheckFact, error)
	CheckImg() (*CheckImg, error)
	CheckJoke(string) (*CheckJoke, error)
}

type GameRepo interface {
	GetPredictAge(string) (*PredictAge, error)
	GetPredictNationality(string) (*PredictNationality, error)
	GetPredictGender(string) (*PredictGender, error)
	GetRandomActivity() (*RandomActivity, error)
	GetRandomFact() (*RandomFact, error)
	GetFactMath() (*FactMath, error)
	GetFactDog() ([]FactDog, error)
	GetDogImg() (*RandomDogImg, error)
	GetCatImg() (*RandomCatImg, error)
	GetDuckImg() (*RandomDuckImg, error)
	GetRandomJoke(mode string) (*RandomJoke, error)
}

type (
	InputPredictName struct {
		Name string `form:"name" json:"name" binding:"required"`
	}

	PredictName struct {
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

	InputCheckJoke struct {
		SafeMode bool `json:"safe_mode" example:"false"`
	}

	CheckJoke struct {
		Category string     `json:"category"`
		Joke     string     `json:"random_joke"`
		Flags    DetailFlag `json:"flags"`
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

	RandomFact struct {
		Text    string `json:"text"`
		Light   string `json:"light"`
		Dark    string `json:"dark"`
		Primary string `json:"primary"`
	}

	FactMath struct {
		Text   string `json:"text"`
		Number int    `json:"number"`
		Found  bool   `json:"bool"`
		Type   string `json:"type"`
	}

	FactDog struct {
		Fact string `json:"fact"`
	}

	RandomDogImg struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	RandomCatImg struct {
		ID     int     `json:"id"`
		URL    string  `json:"url"`
		WEBURL string  `json:"webpurl"`
		X      float64 `json:"x"`
		Y      float64 `json:"y"`
	}

	RandomDuckImg struct {
		Message string `json:"message"`
		URL     string `json:"url"`
	}

	RandomJoke struct {
		Error    bool       `json:"error"`
		Category string     `json:"category"`
		Type     string     `json:"type"`
		Setup    string     `json:"setup"`
		Delivery string     `json:"delivery"`
		Joke     string     `json:"joke"`
		Flags    DetailFlag `json:"flags"`
		ID       int        `json:"id"`
		Safe     bool       `json:"safe"`
		Language string     `json:"lang"`
	}

	DetailFlag struct {
		NSFW      bool `json:"nsfw,omitempty"`
		Religious bool `json:"religious,omitempty"`
		Political bool `json:"political,omitempty"`
		Racist    bool `json:"racist,omitempty"`
		Sexist    bool `json:"sexist,omitempty"`
		Explicit  bool `json:"explicit,omitempty"`
		Safe      bool `json:"safe,omitempty"`
	}
)
