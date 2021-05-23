package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type CoWin struct {
	Sessions []Sessions `json:"sessions"`
}
type Sessions struct {
	CenterID               int      `json:"center_id"`
	Name                   string   `json:"name"`
	Address                string   `json:"address"`
	StateName              string   `json:"state_name"`
	DistrictName           string   `json:"district_name"`
	BlockName              string   `json:"block_name"`
	Pincode                int      `json:"pincode"`
	From                   string   `json:"from"`
	To                     string   `json:"to"`
	Lat                    int      `json:"lat"`
	Long                   int      `json:"long"`
	FeeType                string   `json:"fee_type"`
	SessionID              string   `json:"session_id"`
	Date                   string   `json:"date"`
	AvailableCapacityDose1 int      `json:"available_capacity_dose1"`
	AvailableCapacityDose2 int      `json:"available_capacity_dose2"`
	AvailableCapacity      int      `json:"available_capacity"`
	Fee                    string   `json:"fee"`
	MinAgeLimit            int      `json:"min_age_limit"`
	Vaccine                string   `json:"vaccine"`
	Slots                  []string `json:"slots"`
}

func GetJson(url string, target interface{}) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X)"+
		"AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1")
	req.Header.Add("Accept-Language", "hi_IN")
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, target)

	if err != nil {
		return err
	}

	return nil
}
