package service

import (
	"apigogame/src/domain"
	"apigogame/src/util"
	"reflect"
	"strings"
)

type GameServiceImpl struct {
	gRepo     domain.GameRepo
	mdService domain.MasterDataService
}

func CreateGameServiceImpl(gRepo domain.GameRepo, mdService domain.MasterDataService) domain.GameService {
	return &GameServiceImpl{
		gRepo:     gRepo,
		mdService: mdService,
	}
}

func (g *GameServiceImpl) FindPredictionByName(name string) (*domain.PredictName, error) {
	// Get age
	age, err := g.gRepo.GetPredictAge(name)
	if err != nil {
		return nil, err
	}

	// Get nationality
	nat, err := g.gRepo.GetPredictNationality(name)
	if err != nil {
		return nil, err
	}

	// Get country name by country id from above
	for i, v := range nat.Countries {
		cname, _, err := g.mdService.FindCountryByID(v.CountryID)
		if err != nil {
			return nil, err
		}
		nat.Countries[i].CountryName = cname
	}

	// Get gender
	gender, err := g.gRepo.GetPredictGender(name)
	if err != nil {
		return nil, err
	}

	// Create response as needed
	result := domain.PredictName{
		PredictAge:         age.Age,
		PredictNationality: nat,
		PredictGender:      gender.Gender,
	}

	return &result, nil
}

func (g *GameServiceImpl) FindRandomActivity() (*domain.RandomActivity, error) {
	// Get activity
	bored, err := g.gRepo.GetBoredActivity()
	if err != nil {
		return nil, err
	}

	// Create response as needed
	result := domain.RandomActivity{
		RandomActivity: bored.Activity,
	}

	return &result, nil
}

func (g *GameServiceImpl) FindRandomFact() (*domain.RandomFact, error) {
	var facts []string

	// // Get fun fact
	// fun, err := g.gRepo.GetFunFact()
	// if err != nil {
	// 	return nil, err
	// }

	// Get fact abt math
	math, err := g.gRepo.GetFactMath()
	if err != nil {
		return nil, err
	}

	// API Error
	// Get fact abt dog
	dog, err := g.gRepo.GetFactDog()
	if err != nil {
		return nil, err
	}

	// Add facts to array to randomize
	facts = append(facts, math.Text, dog[0].Fact)
	fact := util.RandomData(facts)

	// Create response as needed
	result := domain.RandomFact{
		Fact: fact,
	}

	return &result, nil
}

func (g *GameServiceImpl) FindRandomImage() (*domain.RandomImg, error) {
	var imgs []string

	// Get dog image
	dog, err := g.gRepo.GetDogImg()
	if err != nil {
		return nil, err
	}

	// API Error
	// Get cat image
	cat, err := g.gRepo.GetCatImg()
	if err != nil {
		return nil, err
	}

	// Get duck image
	duck, err := g.gRepo.GetDuckImg()
	if err != nil {
		return nil, err
	}

	// Add image to array to randomize
	imgs = append(imgs, dog.Message, cat.URL, duck.URL)
	url := util.RandomData(imgs)

	// Create response as needed
	result := domain.RandomImg{
		URL: url,
	}

	return &result, nil
}

func (g *GameServiceImpl) FindRandomJoke(isSafe bool) (*domain.RandomJoke, error) {
	// Check if joke is safe or not
	mode := ""
	if isSafe {
		mode = "safe-mode"
	}

	// Get random joke
	data, err := g.gRepo.GetRandomJoke(mode)
	if err != nil {
		return nil, err
	}

	// Create response as needed
	var result domain.RandomJoke
	result.Category = data.Category
	result.Joke = data.Joke
	result.Flags.Safe = data.Safe

	// If joke type twopart, concate so it will be show as one var
	if data.Type == "twopart" {
		result.Joke = data.Setup + "\n" + data.Delivery
	}

	// Delete if joke containes quote to generalize response
	if strings.Contains(result.Joke, "\"") {
		result.Joke = strings.Replace(result.Joke, "\"", "", -1)
	}

	// Check flag, if it's false, we dont have to show it to the front
	e := reflect.ValueOf(&data.Flags).Elem()
	for i := 0; i < e.NumField(); i++ {
		if e.Field(i).Interface() == false {
			continue
		}
		result.Flags = data.Flags
	}

	return &result, nil
}
