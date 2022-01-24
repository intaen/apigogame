package service

import (
	"apigogame/src/domain"
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

type MasterDataServiceImpl struct {
	mdRepo domain.MasterDataRepo
}

func CreateMasterDataServiceImpl(mdRepo domain.MasterDataRepo) domain.MasterDataService {
	return &MasterDataServiceImpl{
		mdRepo: mdRepo,
	}
}

func (md *MasterDataServiceImpl) FindCountry(id, reg string) (*domain.Country, error) {
	var result domain.Country

	// Get all country
	if id == "" && reg == "" {
		listData, err := md.mdRepo.GetAllCountryName()
		if err != nil {
			return nil, err
		}

		// Check country id to result from above
		var data domain.DetailCountryIDs
		var country domain.DetailCountry
		e := reflect.ValueOf(&listData.Data).Elem()
		for i := 0; i < e.NumField(); i++ {
			dt, _ := json.Marshal(e.Field(i).Interface())
			json.Unmarshal(dt, &data)

			// Reassign data
			country.CountryID = e.Type().Field(i).Name
			country.CountryName = data.Country
			country.Region = data.Region
			result.DetailCountry = append(result.DetailCountry, country)
			result.Count = len(result.DetailCountry)
		}
		return &result, nil
	}

	if id != "" && reg != "" {
		cname, region, err := md.FindCountryByIDRegion(id, reg)
		if err != nil {
			return nil, err
		}

		// Reassign data
		var country domain.DetailCountry
		country.CountryID = id
		country.CountryName = cname
		country.Region = region
		result.DetailCountry = append(result.DetailCountry, country)
		result.Count = len(result.DetailCountry)
		return &result, nil
	}

	// Get country by id
	if id != "" {
		cname, region, err := md.FindCountryByID(id)
		if err != nil {
			return nil, err
		}

		// Reassign data
		var country domain.DetailCountry
		country.CountryName = cname
		country.Region = region
		result.DetailCountry = append(result.DetailCountry, country)
		result.Count = len(result.DetailCountry)
		return &result, nil
	}

	// Get country by region
	if reg != "" {
		cids, cnames, err := md.FindCountryByRegion(reg)
		if err != nil {
			return nil, err
		}

		// Reassign data
		var country domain.DetailCountry
		for i, v := range cids {
			country.CountryID = v
			country.CountryName = cnames[i]
			result.DetailCountry = append(result.DetailCountry, country)
			result.Count = len(cids)
		}

		country.Region = reg
		return &result, nil
	}

	return nil, errors.New("data not found")
}

func (md *MasterDataServiceImpl) FindCountryByID(id string) (string, string, error) {
	// Get data country
	listData, err := md.mdRepo.GetAllCountryName()
	if err != nil {
		return "", "", err
	}

	// Check country id to result from above
	var data domain.DetailCountryIDs
	e := reflect.ValueOf(&listData.Data).Elem()
	for i := 0; i < e.NumField(); i++ {
		if e.Type().Field(i).Name != id {
			continue
		}
		// Convert it to new struct
		dt, _ := json.Marshal(e.Field(i).Interface())
		json.Unmarshal(dt, &data)
	}

	return data.Country, data.Region, nil
}

func (md *MasterDataServiceImpl) FindCountryByRegion(reg string) ([]string, []string, error) {
	// Get data country
	listData, err := md.mdRepo.GetAllCountryName()
	if err != nil {
		return nil, nil, err
	}

	// Check region to result from above
	var listCountryID, listCountry []string
	var data domain.DetailCountryIDs
	e := reflect.ValueOf(&listData.Data).Elem()
	for i := 0; i < e.NumField(); i++ {
		// Convert it to new struct
		dt, _ := json.Marshal(e.Field(i).Interface())
		json.Unmarshal(dt, &data)

		if data.Region == strings.Title(strings.ToLower(reg)) {
			listCountryID = append(listCountryID, e.Type().Field(i).Name)
			listCountry = append(listCountry, data.Country)
		}
	}

	return listCountryID, listCountry, nil
}

func (md *MasterDataServiceImpl) FindCountryByIDRegion(id, reg string) (string, string, error) {
	// Get data country
	listData, err := md.mdRepo.GetAllCountryName()
	if err != nil {
		return "", "", err
	}

	// Check region to result from above
	var data domain.DetailCountryIDs
	e := reflect.ValueOf(&listData.Data).Elem()
	for i := 0; i < e.NumField(); i++ {
		if e.Type().Field(i).Name != id {
			continue
		}
		
		// Convert it to new struct
		dt, _ := json.Marshal(e.Field(i).Interface())
		json.Unmarshal(dt, &data)

		if data.Region != strings.Title(strings.ToLower(reg)) {
			continue
		}

		
	}

	return data.Country, data.Region, nil
}

func (md *MasterDataServiceImpl) FindAPI(auths []string, cat string) (*domain.PublicAPIs, error) {
	if len(auths) == 0 && cat == "" {
		listData, err := md.mdRepo.GetAllAPI("auth", "", "category", "")
		if err != nil {
			return nil, err
		}

		return listData, nil
	}

	if len(auths) != 0 && cat != "" {
		listData, err := md.FindAPIByAuthCategory(auths, cat)
		if err != nil {
			return nil, err
		}

		return listData, nil
	}

	// Get api by auth
	if len(auths) != 0 {
		listData, err := md.FindAPIByAuth(auths)
		if err != nil {
			return nil, err
		}
		return listData, nil
	}

	// Get api by category
	if cat != "" {
		listData, err := md.FindAPIByCategory(cat)
		if err != nil {
			return nil, err
		}
		return listData, nil
	}

	return nil, errors.New("data not found")
}

func (md *MasterDataServiceImpl) FindAPIByAuth(auths []string) (*domain.PublicAPIs, error) {
	// Get public api by auth, its needed to be loop because if isAuth true, there's 2 kind of auth and we have to look up for it twice
	var listData domain.PublicAPIs
	for _, v := range auths {
		listAPI, err := md.mdRepo.GetAllAPI("auth", v, "category", "")
		if err != nil {
			return nil, err
		}
		listData.TotalAuth = len(auths)
		listData.Auths = auths
		listData.Count += listAPI.Count
		listData.Entries = append(listData.Entries, listAPI.Entries...)
	}

	// Reassign data that match
	for _, v := range listData.Entries {
		if v.Auth == auths[listData.TotalAuth-1] {
			listData.Entries = append(listData.Entries, listData.Entries...)
		}
	}

	return &listData, nil
}

func (md *MasterDataServiceImpl) FindAPIByCategory(cat string) (*domain.PublicAPIs, error) {
	// Get public api by category
	listData, err := md.mdRepo.GetAllAPI("category", cat, "auth", "")
	if err != nil {
		return nil, err
	}

	return listData, nil
}

func (md *MasterDataServiceImpl) FindAPIByAuthCategory(auths []string, cat string) (*domain.PublicAPIs, error) {
	var listData domain.PublicAPIs
	for _, v := range auths {
		listAPI, err := md.mdRepo.GetAllAPI("auth", v, "category", cat)
		if err != nil {
			return nil, err
		}
		listData.TotalAuth = len(auths)
		listData.Auths = auths
		listData.Count += listAPI.Count
		listData.Entries = append(listData.Entries, listAPI.Entries...)
	}

	// Reassign data that match
	for _, v := range listData.Entries {
		if v.Auth == auths[listData.TotalAuth-1] {
			listData.Entries = append(listData.Entries, listData.Entries...)
		}
	}

	return &listData, nil
}
