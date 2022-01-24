package domain

type MasterDataService interface {
	FindCountry(string, string) (*Country, error)
	FindCountryByID(string) (string, string, error)
	FindCountryByRegion(string) ([]string, []string, error)
	FindAPI([]string, string) (*PublicAPIs, error)
	FindAPIByAuth([]string) (*PublicAPIs, error)
	FindAPIByCategory(string) (*PublicAPIs, error)
}

type MasterDataRepo interface {
	GetAllCountryName() (*Countries, error)
	//GetAllAPI(string, string) (*PublicAPIs, error)
	GetAllAPI(string, string, string, string) (*PublicAPIs, error)
}

type (
	Country struct {
		Count         int             `json:"count"`
		DetailCountry []DetailCountry `json:"countries,omitempty"`
	}
	DetailCountry struct {
		CountryID   string `json:"country_id,omitempty"`
		CountryName string `json:"country_name,omitempty"`
		Region      string `json:"region,omitempty"`
	}

	// API
	Countries struct {
		Status     string     `json:"status"`
		StatusCode int        `json:"status-code"`
		Version    string     `json:"version"`
		Access     string     `json:"public"`
		Data       CountryIDs `json:"data"`
	}
	CountryIDs struct {
		DZ DetailCountryIDs
		AO DetailCountryIDs
		BJ DetailCountryIDs
		BW DetailCountryIDs
		BF DetailCountryIDs
		BI DetailCountryIDs
		CV DetailCountryIDs
		CM DetailCountryIDs
		CF DetailCountryIDs
		TD DetailCountryIDs
		KM DetailCountryIDs
		CD DetailCountryIDs
		CG DetailCountryIDs
		CI DetailCountryIDs
		DJ DetailCountryIDs
		EG DetailCountryIDs
		GQ DetailCountryIDs
		ER DetailCountryIDs
		SZ DetailCountryIDs
		ET DetailCountryIDs
		GA DetailCountryIDs
		GM DetailCountryIDs
		GH DetailCountryIDs
		GN DetailCountryIDs
		GW DetailCountryIDs
		KE DetailCountryIDs
		LS DetailCountryIDs
		LR DetailCountryIDs
		LY DetailCountryIDs
		MG DetailCountryIDs
		MW DetailCountryIDs
		ML DetailCountryIDs
		MR DetailCountryIDs
		MU DetailCountryIDs
		YT DetailCountryIDs
		MA DetailCountryIDs
		MZ DetailCountryIDs
		NA DetailCountryIDs
		NE DetailCountryIDs
		NG DetailCountryIDs
		RE DetailCountryIDs
		RW DetailCountryIDs
		SH DetailCountryIDs
		ST DetailCountryIDs
		SN DetailCountryIDs
		SC DetailCountryIDs
		SL DetailCountryIDs
		SO DetailCountryIDs
		ZA DetailCountryIDs
		SS DetailCountryIDs
		SD DetailCountryIDs
		TZ DetailCountryIDs
		TG DetailCountryIDs
		TN DetailCountryIDs
		UG DetailCountryIDs
		EH DetailCountryIDs
		ZM DetailCountryIDs
		ZW DetailCountryIDs
		AQ DetailCountryIDs
		BV DetailCountryIDs
		TF DetailCountryIDs
		HM DetailCountryIDs
		GS DetailCountryIDs
		AF DetailCountryIDs
		AM DetailCountryIDs
		AZ DetailCountryIDs
		BD DetailCountryIDs
		BT DetailCountryIDs
		IO DetailCountryIDs
		KH DetailCountryIDs
		CN DetailCountryIDs
		GE DetailCountryIDs
		HK DetailCountryIDs
		IN DetailCountryIDs
		ID DetailCountryIDs
		JP DetailCountryIDs
		KZ DetailCountryIDs
		KP DetailCountryIDs
		KR DetailCountryIDs
		KG DetailCountryIDs
		LA DetailCountryIDs
		MO DetailCountryIDs
		MY DetailCountryIDs
		MV DetailCountryIDs
		MN DetailCountryIDs
		MM DetailCountryIDs
		NP DetailCountryIDs
		PK DetailCountryIDs
		PH DetailCountryIDs
		SG DetailCountryIDs
		LK DetailCountryIDs
		TW DetailCountryIDs
		TJ DetailCountryIDs
		TH DetailCountryIDs
		TL DetailCountryIDs
		TM DetailCountryIDs
		UZ DetailCountryIDs
		VN DetailCountryIDs
		BZ DetailCountryIDs
	}
	DetailCountryIDs struct {
		Country string `json:"country"`
		Region  string `json:"region"`
	}

	PublicAPIs struct {
		Auths     []string      `json:"-"`
		TotalAuth int           `json:"-"`
		Count     int           `json:"count"`
		Entries   []DetailEntry `json:"entries"`
	}
	DetailEntry struct {
		API         string `json:"API"`
		Description string `json:"Description"`
		Auth        string `json:"Auth"`
		HTTPS       bool   `json:"HTTPS"`
		Cors        string `json:"Cors"`
		Link        string `json:"Link"`
		Category    string `json:"Category"`
	}
)
