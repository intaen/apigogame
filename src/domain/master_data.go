package domain

type MasterDataService interface {
	FindCountryByID(string) (string, error)
}

type MasterDataRepo interface {
	GetCountryName() (*Countries, error)
}

type (
	Countries struct {
		Status     string      `json:"status"`
		StatusCode int         `json:"status-code"`
		Version    string      `json:"version"`
		Access     string      `json:"public"`
		Data       CountriesID `json:"data"`
	}
	CountriesID struct {
		DZ DetailCountriesID
		AO DetailCountriesID
		BJ DetailCountriesID
		BW DetailCountriesID
		BF DetailCountriesID
		BI DetailCountriesID
		CV DetailCountriesID
		CM DetailCountriesID
		CF DetailCountriesID
		TD DetailCountriesID
		KM DetailCountriesID
		CD DetailCountriesID
		CG DetailCountriesID
		CI DetailCountriesID
		DJ DetailCountriesID
		EG DetailCountriesID
		GQ DetailCountriesID
		ER DetailCountriesID
		SZ DetailCountriesID
		ET DetailCountriesID
		GA DetailCountriesID
		GM DetailCountriesID
		GH DetailCountriesID
		GN DetailCountriesID
		GW DetailCountriesID
		KE DetailCountriesID
		LS DetailCountriesID
		LR DetailCountriesID
		LY DetailCountriesID
		MG DetailCountriesID
		MW DetailCountriesID
		ML DetailCountriesID
		MR DetailCountriesID
		MU DetailCountriesID
		YT DetailCountriesID
		MA DetailCountriesID
		MZ DetailCountriesID
		NA DetailCountriesID
		NE DetailCountriesID
		NG DetailCountriesID
		RE DetailCountriesID
		RW DetailCountriesID
		SH DetailCountriesID
		ST DetailCountriesID
		SN DetailCountriesID
		SC DetailCountriesID
		SL DetailCountriesID
		SO DetailCountriesID
		ZA DetailCountriesID
		SS DetailCountriesID
		SD DetailCountriesID
		TZ DetailCountriesID
		TG DetailCountriesID
		TN DetailCountriesID
		UG DetailCountriesID
		EH DetailCountriesID
		ZM DetailCountriesID
		ZW DetailCountriesID
		AQ DetailCountriesID
		BV DetailCountriesID
		TF DetailCountriesID
		HM DetailCountriesID
		GS DetailCountriesID
		AF DetailCountriesID
		AM DetailCountriesID
		AZ DetailCountriesID
		BD DetailCountriesID
		BT DetailCountriesID
		IO DetailCountriesID
		KH DetailCountriesID
		CN DetailCountriesID
		GE DetailCountriesID
		HK DetailCountriesID
		IN DetailCountriesID
		ID DetailCountriesID
		JP DetailCountriesID
		KZ DetailCountriesID
		KP DetailCountriesID
		KR DetailCountriesID
		KG DetailCountriesID
		LA DetailCountriesID
		MO DetailCountriesID
		MY DetailCountriesID
		MV DetailCountriesID
		MN DetailCountriesID
		MM DetailCountriesID
		NP DetailCountriesID
		PK DetailCountriesID
		PH DetailCountriesID
		SG DetailCountriesID
		LK DetailCountriesID
		TW DetailCountriesID
		TJ DetailCountriesID
		TH DetailCountriesID
		TL DetailCountriesID
		TM DetailCountriesID
		UZ DetailCountriesID
		VN DetailCountriesID
		BZ DetailCountriesID
	}
	DetailCountriesID struct {
		Country string `json:"country"`
		Region  string `json:"region"`
	}
)
