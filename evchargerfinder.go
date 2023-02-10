package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/kevin-fwu/EVChargerFinder/nrel"
)

// Lets get charging!

func main() {
	fmt.Println("Hello World!")

	confArg := flag.String("conf", "", "The JSON Configuration File.")
	latitudeArg := flag.Float64("latitude", math.NaN(), "Latitude to check.")
	longitudeArg := flag.Float64("longitude", math.NaN(), "Longitude to check.")
	distArg := flag.Float64("distance", math.NaN(), "Distance to check.")
	limitArg := flag.Int("limit", 0, "Max number of locations to return.")

	flag.Parse()

	if *confArg == "" {
		fmt.Println("Missing config file.")
		return
	}

	conf, err := LoadConf(*confArg)

	if err != nil {
		fmt.Printf("Failed to load config file, error: %+v\n", err)
		return
	}

	stationList, err := nrel.ParseFile(conf.Nrel.File)

	if err != nil {
		err = nrel.FetchData(conf.Nrel.Token, conf.Nrel.File)
		if err == nil {
			stationList, err = nrel.ParseFile(conf.Nrel.File)
		}
		if err != nil {
			fmt.Printf("Failed to retrieve NREL data file, error: %+v\n", err)
			return
		}
	}

	sortedTree = buildTree(stationList)

	if conf.Server.Address != "" {
		listen(conf.Server.Address, conf.Server.Ssl.Cert, conf.Server.Ssl.Key)
	} else if !math.IsNaN(*latitudeArg) && !math.IsNaN(*longitudeArg) && !math.IsNaN(*distArg) {
		list := findClosest(&Point{coords: []float64{*latitudeArg, *longitudeArg}}, *distArg, *limitArg)

		json.NewEncoder(os.Stdout).Encode(list)
	}
}
