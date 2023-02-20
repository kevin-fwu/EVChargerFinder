package main

import (
	"bytes"
	"log"
	"os"

	"github.com/kevin-fwu/EVChargerFinder/nrel"
	"github.com/kevin-fwu/EVChargerFinder/util"
)

type StationLoader struct {
	token    string
	filename string
	md5sum   []byte
}

var stationLoader *StationLoader

func InitStationLoader(token, filename string) error {
	stationLoader = &StationLoader{token: token, filename: filename}

	stationList, err := nrel.ParseFile(stationLoader.filename)

	if err != nil {
		err = nrel.FetchData(stationLoader.token, stationLoader.filename)
		if err == nil {
			stationList, err = nrel.ParseFile(stationLoader.filename)
		}
		if err != nil {
			return err
		}
	}
	stationLoader.md5sum, err = nrel.GetMd5(stationLoader.filename)

	if err != nil {
		return err
	}

	sortedTree = buildTree(stationList)

	util.RegisterTimer(stationLoader)
	return nil
}

func (s *StationLoader) OnTimer86400Sec() {
	tmp := s.filename + ".tmp"
	if err := nrel.FetchData(s.token, tmp); err != nil {
		log.Printf("Failed to fetch updated NREL data: %+v\n", err)
	}
	md5sum, err := nrel.GetMd5(tmp)

	if err != nil {
		log.Printf("Failed to get md5 from updated NREL data: %+v\n", err)
	}

	if !bytes.Equal(md5sum, stationLoader.md5sum) {

		stationList, err := nrel.ParseFile(tmp)
		if err != nil {
			log.Printf("Failed to parse updated NREL data: %+v\n", err)
		}
		err = os.Rename(tmp, s.filename)
		if err != nil {
			log.Printf("Failed to overwrite NREL data file: %+v\n", err)
		}
		sortedTree = buildTree(stationList)
		s.md5sum = md5sum
	}
}
