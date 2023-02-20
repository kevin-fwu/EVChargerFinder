package nrel

import (
	"io"
	"net/http"
	"os"
)

func FetchData(token, file string) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://developer.nrel.gov/api/alt-fuel-stations/v1.json?access=public&fuel_type=ELEC", nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(token, "")
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	io.Copy(f, resp.Body)
	return nil
}
