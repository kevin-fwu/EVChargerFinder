package main

type Location struct {
	StreetAddress    string
	City             string
	State            string
	Country          string
	Zip              string
	GeocodeStatus    string
	Coordinates      []float64
	CoordinateString string
	Stations         []*ChargingStation
}

type ChargingStation struct {
	Name                   string
	PhoneNumber            string
	IntersectionDirections string
	AccessTime             string
	Connectors             []string
	Network                string
	Pricing                string
	FacilityType           string
	RestrictedAccess       bool
	CntLevel2Chargers      int
	CntLevel3Chargers      int
}

func (loc *Location) GetCoordinates() []float64 {
	return loc.Coordinates
}
