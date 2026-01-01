package main

import (
	"flag"
	"fmt"

	"github.com/IlyaRomanyuk/go-weather/geo"
	"github.com/IlyaRomanyuk/go-weather/weather"
)

func main() {
	city := flag.String("city", "", "Введенный город")
	format := flag.Int("format", 1, "Формат вывода в консоль")

	flag.Parse()

	fmt.Println(*city, *format)

	geoData, err := geo.GetMyLocation(*city)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	weatherData := weather.GetWeather(geoData, *format)

	fmt.Println(weatherData)
}
