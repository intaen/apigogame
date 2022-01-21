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

func (g *GameServiceImpl) PredictionByName(name string) (*domain.PredictName, error) {
	pAge, err := g.gRepo.GetPredictAge(name)
	if err != nil {
		return nil, err
	}

	pNat, err := g.gRepo.GetPredictNationality(name)
	if err != nil {
		return nil, err
	}

	for i, v := range pNat.Countries {
		cname, err := g.mdService.FindCountryByID(v.CountryID)
		if err != nil {
			return nil, err
		}
		pNat.Countries[i].CountryName = cname
	}

	pGen, err := g.gRepo.GetPredictGender(name)
	if err != nil {
		return nil, err
	}

	///
	result := domain.PredictName{
		PredictAge:         pAge.Age,
		PredictNationality: pNat,
		PredictGender:      pGen.Gender,
	}

	return &result, nil
}

func (g *GameServiceImpl) DoYouDare() (*domain.CheckDare, error) {
	rAct, err := g.gRepo.GetRandomActivity()
	if err != nil {
		return nil, err
	}

	///
	result := domain.CheckDare{
		RandomActivity: rAct.Activity,
	}

	return &result, nil
}

func (g *GameServiceImpl) CheckFact() (*domain.CheckFact, error) {
	var facts []string
	fun, err := g.gRepo.GetRandomFact()
	if err != nil {
		return nil, err
	}

	math, err := g.gRepo.GetFactMath()
	if err != nil {
		return nil, err
	}

	dog, err := g.gRepo.GetFactDog()
	if err != nil {
		return nil, err
	}

	facts = append(facts, fun.Text, math.Text, dog[0].Fact)
	fact := util.RandomData(facts)

	///
	result := domain.CheckFact{
		Fact: fact,
	}

	return &result, nil
}

func (g *GameServiceImpl) CheckImg() (*domain.CheckImg, error) {
	var imgs []string
	dog, err := g.gRepo.GetDogImg()
	if err != nil {
		return nil, err
	}

	cat, err := g.gRepo.GetCatImg()
	if err != nil {
		return nil, err
	}

	duck, err := g.gRepo.GetDuckImg()
	if err != nil {
		return nil, err
	}

	// Create array to random data
	imgs = append(imgs, dog.Message, cat.URL, duck.URL)
	url := util.RandomData(imgs)

	///
	result := domain.CheckImg{
		URL: url,
	}

	return &result, nil
}

func (g *GameServiceImpl) CheckJoke(mode string) (*domain.CheckJoke, error) {
	data, err := g.gRepo.GetRandomJoke(mode)
	if err != nil {
		return nil, err
	}

	///
	var result domain.CheckJoke
	result.Category = data.Category
	result.Joke = data.Joke
	result.Flags.Safe = data.Safe

	if data.Type == "twopart" {
		result.Joke = data.Setup + "\n" + data.Delivery
	}

	if strings.Contains(result.Joke, "\"") {
		result.Joke = strings.Replace(result.Joke, "\"", "", -1)
	}

	e := reflect.ValueOf(&data.Flags).Elem()
	for i := 0; i < e.NumField(); i++ {
		if e.Field(i).Interface() == false {
			continue
		}
		result.Flags = data.Flags
	}

	return &result, nil
}
