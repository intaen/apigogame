package repo

import (
	"apigogame/src/domain"
	"apigogame/src/util"
	"errors"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type MasterDataRepoImpl struct {
}

func CreateMasterDataRepoImpl() domain.MasterDataRepo {
	return &MasterDataRepoImpl{}
}

func (md *MasterDataRepoImpl) GetCountryName() (*domain.Countries, error) {
	// Consume third API
	client := resty.New()     // Create client
	var resp domain.Countries // Initialize new variable to catch response from 3rd party\
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&resp).
		Get(viper.GetString("url.countries"))
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
