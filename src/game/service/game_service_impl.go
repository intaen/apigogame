package service

import (
	"apigogame/src/domain"
)

type GameServiceImpl struct {
	gRepo domain.GameRepo
}

func CreateGameServiceImpl(gRepo domain.GameRepo) domain.GameService {
	return &GameServiceImpl{
		gRepo: gRepo,
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

	///
	result := domain.GameName{
		PredictAge:         pAge.Age,
		PredictNationality: pNat.Country,
	}

	return &result, nil
}

func (g *GameServiceImpl) Punishment() (*domain.Punishment, error) {
	rAct, err := g.gRepo.GetRandomActivity()
	if err != nil {
		return nil, err
	}

	result := domain.Punishment{
		RandomActivity: rAct.Activity,
	}

	return &result, nil
}
