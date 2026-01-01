package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/IlyaRomanyuk/go-weather/geo"
)

func GetWeather(geo *geo.GeoDataStruct, format int) string {

	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	return string(data)
}
