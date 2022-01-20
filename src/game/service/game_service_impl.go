package service

import (
	"apigogame/src/domain"
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

func (g *GameServiceImpl) PredictionByName(name string) (*domain.GameName, error) {
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
	result := domain.GameName{
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
	rAct, err := g.gRepo.GetFactMath()
	if err != nil {
		return nil, err
	}

	///
	result := domain.CheckFact{
		Fact: rAct.Text,
	}

	return &result, nil
}

func (g *GameServiceImpl) CheckImg() (*domain.CheckImg, error) {
	data, err := g.gRepo.GetRandomImg()
	if err != nil {
		return nil, err
	}

	///
	result := domain.CheckImg{
		URL: data.Message,
	}

	return &result, nil
}
