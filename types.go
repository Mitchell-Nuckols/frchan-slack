package main

// TBATeam : GET /team/{team_key}
type TBATeam struct {
	Key              string   `json:"key,omitempty"`
	TeamNumber       int      `json:"team_number,omitempty"`
	Nickname         string   `json:"nickname,omitempty"`
	Name             string   `json:"name,omitempty"`
	City             string   `json:"city,omitempty"`
	SateProv         string   `json:"state_prov,omitempty"`
	Country          string   `json:"country,omitempty"`
	Address          string   `json:"address,omitempty"`
	PostalCode       string   `json:"postal_code,omitempty"`
	GMapsPlaceID     string   `json:"gmaps_place_id,omitempty"`
	GMapsPlaceURL    string   `json:"gmaps_url,omitempty"`
	Lat              int      `json:"lat,omitempty"`
	Lng              int      `json:"lng,omitempty"`
	LocationName     string   `json:"location_name,omitempty"`
	Website          string   `json:"website,omitempty"`
	RookieYear       int      `json:"rookie_year,omitempty"`
	Motto            string   `json:"motto,omitempty"`
	HomeChampionship struct{} `json:"home_championship,omitempty"`
}
