package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/IlyaRomanyuk/go-weather/geo"
)

var ErrorWrongFormat = errors.New("WRONG_FORMAT")

func GetWeather(geo *geo.GeoDataStruct, format int) (string, error) {

	if format < 1 || format > 4 {
		return "", ErrorWrongFormat
	}

	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)

	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_URL")
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())

	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_HTTP")
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", errors.New("ERROR_READBODY")
	}

	return string(data), nil
}
