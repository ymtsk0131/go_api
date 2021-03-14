package main

import (
	"fmt"
	"go_api/config"
	"go_api/open_weather"
	"go_api/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := open_weather.New(config.Config.Apikey)
	dailyForecast, err := apiClient.GetDailyForecast("Tokyo")
	if err != nil {
		log.Printf("action:main err=%s", err.Error())
	}
	fmt.Println(dailyForecast.City.Name)
	for _, val := range dailyForecast.List {
		fmt.Printf("%s %s\n", val.DtTxt, val.Weather[0].Main)
	}
}
