package service

import (
	"apigogame/src/domain"
	"encoding/json"
	"fmt"
	"reflect"
)

type MasterDataServiceImpl struct {
	mdRepo domain.MasterDataRepo
}

func CreateMasterDataServiceImpl(mdRepo domain.MasterDataRepo) domain.MasterDataService {
	return &MasterDataServiceImpl{
		mdRepo: mdRepo,
	}
}

func (md *MasterDataServiceImpl) FindCountryByID(id string) (string, error) {
	listData, err := md.mdRepo.GetCountryName()
	if err != nil {
		return "", err
	}

	e := reflect.ValueOf(&listData.Data).Elem()

	var detailCID domain.DetailCountriesID
	for i := 0; i < e.NumField(); i++ {
		if e.Type().Field(i).Name != id {
			fmt.Println(e.Field(i).Interface())
			continue
		}
		dt, _ := json.Marshal(e.Field(i).Interface())
		json.Unmarshal(dt, &detailCID)
	}

	return detailCID.Country, nil
}
