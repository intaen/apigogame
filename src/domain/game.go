package domain

type GameService interface {
	FindPredictionByName(string) (*PredictName, error)
	FindRandomActivity() (*RandomActivity, error)
	FindRandomFact() (*RandomFact, error)
	FindRandomImage() (*RandomImg, error)
	FindRandomJoke(bool) (*RandomJoke, error)
}

type GameRepo interface {
	GetPredictAge(string) (*PredictAge, error)
	GetPredictNationality(string) (*PredictNationality, error)
	GetPredictGender(string) (*PredictGender, error)
	GetBoredActivity() (*BoredActivity, error)
	GetFunFact() (*FactFun, error)
	GetFactMath() (*FactMath, error)
	GetFactDog() ([]FactDog, error)
	GetDogImg() (*ImgDog, error)
	GetCatImg() (*ImgCat, error)
	GetDuckImg() (*ImgDuck, error)
	GetRandomJoke(mode string) (*AnyJoke, error)
}

type (
	// Prediction
	InputPredictName struct {
		Name string `form:"name" json:"name" binding:"required"`
	}
	PredictName struct {
		PredictAge         int         `json:"predict_age"`
		PredictNationality interface{} `json:"predict_nationality"`
		PredictGender      string      `json:"predict_gender"`
	}

	// Dare or Dare
	RandomActivity struct {
		RandomActivity string `json:"random_activity"`
	}

	// Today Fact
	RandomFact struct {
		Fact string `json:"random_fact"`
	}

	// Today Fun
	RandomImg struct {
		URL string `json:"url"`
	}
	InputRandomJoke struct {
		IsSafe bool `json:"is_safe" example:"false"`
	}
	RandomJoke struct {
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

	BoredActivity struct {
		Activity      string  `json:"activity"`
		Type          string  `json:"type"`
		Participants  int     `json:"participants"`
		Price         float32 `json:"price"`
		Link          string  `json:"link"`
		Key           string  `json:"key"`
		Accessibility float32 `json:"accessibility"`
	}

	FactFun struct {
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

	ImgDog struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	ImgCat struct {
		ID     int     `json:"id"`
		URL    string  `json:"url"`
		WEBURL string  `json:"webpurl"`
		X      float64 `json:"x"`
		Y      float64 `json:"y"`
	}

	ImgDuck struct {
		Message string `json:"message"`
		URL     string `json:"url"`
	}

	AnyJoke struct {
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
