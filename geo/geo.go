package geo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GeoDataStruct struct {
	City string `json:"city"`
}

type CityCheckResponse struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoDataStruct, error) {

	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			panic("Такого города нет!")
		}
		return &GeoDataStruct{
			City: city,
		}, nil
	}

	resp, err := http.Get("https://ipapi.co/json")

	fmt.Println(resp.StatusCode, "resp.StatusCode")

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var geo GeoDataStruct

	json.Unmarshal(data, &geo)

	return &geo, nil

}

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{"city": city})

	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var checkCity CityCheckResponse
	json.Unmarshal(body, &checkCity)

	return !checkCity.Error
}
