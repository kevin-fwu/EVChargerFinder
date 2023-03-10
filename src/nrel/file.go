package nrel

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"io"
	"os"
)

type InputStation struct {
	Name                   string   `json:"station_name"`
	PhoneNumber            string   `json:"station_phone"`
	StreetAddress          string   `json:"street_address"`
	City                   string   `json:"city"`
	State                  string   `json:"state"`
	Country                string   `json:"country"`
	Zip                    string   `json:"zip"`
	IntersectionDirections string   `json:"intersection_directions"`
	AccessTime             string   `json:"access_days_time"`
	Connectors             []string `json:"ev_connector_types"`
	Network                string   `json:"ev_network"`
	Pricing                string   `json:"ev_pricing"`
	FacilityType           string   `json:"facility_type"`
	GeocodeStatus          string   `json:"geocode_status"`
	Latitude               float64  `json:"latitude"`
	Longitude              float64  `json:"longitude"`
	RestrictedAccess       bool     `json:"restricted_access"`
	CntLevel2Chargers      int      `json:"ev_level2_evse_num"`
	CntLevel3Chargers      int      `json:"ev_dc_fast_num"`
}

type InputStationList struct {
	Stations []InputStation `json:"fuel_stations"`
}

func ParseFile(filename string) (*InputStationList, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	var stationList InputStationList
	err = decoder.Decode(&stationList)
	if err != nil {
		return nil, err
	}

	if len(stationList.Stations) == 0 {
		return nil, errors.New("station list is empty")
	}
	return &stationList, nil
}

func GetMd5(filename string) ([]byte, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
