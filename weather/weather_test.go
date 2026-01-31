package weather_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/IlyaRomanyuk/go-weather/geo"
	"github.com/IlyaRomanyuk/go-weather/weather"
)

func TestWeather(t *testing.T) {
	// Arrange - подготовка данных для теста
	expected := "Volgograd"
	format := 3

	geoData := geo.GeoDataStruct{
		City: expected,
	}

	// Act - запуск теста
	result, err := weather.GetWeather(&geoData, format)

	fmt.Println(result, "result")

	// Assert - сравнение результатов теста
	if err != nil {
		t.Error(err.Error())
	}

	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, получили %v", expected, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{"First negative test", 147},
	{"Second negative test", -1},
}

func TestWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange - подготовка данных для теста
			expected := "Volgograd"

			geoData := geo.GeoDataStruct{
				City: expected,
			}

			// Act - запуск теста
			result, err := weather.GetWeather(&geoData, tc.format)

			fmt.Println(result, "result!")

			// Assert - проверка результата с expected
			if err != weather.ErrorWrongFormat {
				t.Errorf("Ожидалось %v, получили %v", weather.ErrorWrongFormat, err)
			}
		})
	}
}
