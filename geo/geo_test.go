package geo_test

import (
	"fmt"
	"testing"

	"github.com/IlyaRomanyuk/go-weather/geo"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка данных для теста
	city := "Paris"
	expected := geo.GeoDataStruct{
		City: "Paris",
	}

	// Act - запуск теста
	got, err := geo.GetMyLocation(city)

	fmt.Println(got)

	// Assert - сравнение результатов теста
	if err != nil {
		t.Error(err.Error())
	}

	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получили %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Londonff"

	_, err := geo.GetMyLocation(city)

	if err != geo.ErrorNoCity {
		t.Errorf("Ожидалось %v, получили %v", geo.ErrorNoCity, err)
	}

}
