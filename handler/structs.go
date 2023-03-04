package handler

type UNI struct {
	Name     string   `json:"name"`
	Country  string   `json:"country"`
	Webpages []string `json:"web_pages"`
	Isocode  string   `json:"alpha_two_code"`
}

type NABUNI struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Language   map[string]interface{} `json:"languages"`
	GoogleMaps map[string]interface{} `json:"maps"`
	Borders    []string               `json:"borders"`
}

type DIAGNOSTIC struct {
	Universitiesapi string `json:"universitiesapi"`
	Countriesapi    string `json:"countriesapi"`
	Version         string `json:"version"`
	Uptime          string `json:"uptime"`
}
