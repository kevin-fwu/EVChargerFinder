package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/kevin-fwu/EVChargerFinder/util"
)

// Lets get charging!

func main() {

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
	util.InitSignalHandler()

	err = InitStationLoader(conf.Nrel.Token, conf.Nrel.File)

	if err != nil {
		fmt.Printf("Failed to init station loader, error: %+v\n", err)
		return
	}

	if !math.IsNaN(*latitudeArg) && !math.IsNaN(*longitudeArg) && !math.IsNaN(*distArg) {
		parms := &ReqParams{Latitude: *latitudeArg, Longitude: *longitudeArg, Distance: *distArg, CountLimit: *limitArg}
		list := findClosest(parms)

		json.NewEncoder(os.Stdout).Encode(list)
	} else if conf.Server.Address != "" {
		listen(conf.Server.Address, conf.Server.Ssl.Cert, conf.Server.Ssl.Key)
	}
}
