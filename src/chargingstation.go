package main

import (
	"fmt"

	"github.com/kevin-fwu/EVChargerFinder/nrel"
	"github.com/kevin-fwu/EVChargerFinder/util"
)

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

func buildTree(stationList *nrel.InputStationList) *util.KdNode {

	var locList []util.Coordinates
	locMap := make(map[string]*Location)

	for _, inStation := range stationList.Stations {
		coordString := fmt.Sprintf("%f,%f", inStation.Latitude, inStation.Longitude)

		loc := locMap[coordString]

		if loc == nil {
			loc = &Location{
				StreetAddress:    inStation.StreetAddress,
				City:             inStation.City,
				State:            inStation.State,
				Country:          inStation.Country,
				Zip:              inStation.Zip,
				GeocodeStatus:    inStation.GeocodeStatus,
				Coordinates:      []float64{inStation.Latitude, inStation.Longitude},
				CoordinateString: coordString,
			}
			locMap[coordString] = loc
			locList = append(locList, loc)
		}

		cs := &ChargingStation{

			Name:                   inStation.Name,
			PhoneNumber:            inStation.PhoneNumber,
			IntersectionDirections: inStation.IntersectionDirections,
			AccessTime:             inStation.AccessTime,
			Connectors:             inStation.Connectors,
			Network:                inStation.Network,
			Pricing:                inStation.Pricing,
			FacilityType:           inStation.FacilityType,
			RestrictedAccess:       inStation.RestrictedAccess,
			CntLevel2Chargers:      inStation.CntLevel2Chargers,
			CntLevel3Chargers:      inStation.CntLevel3Chargers,
		}
		loc.Stations = append(loc.Stations, cs)
	}

	return util.CreateKdTree(locList, 2)
}
