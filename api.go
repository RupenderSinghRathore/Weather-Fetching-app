package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var url = "https://api.openweathermap.org/data/2.5/forecast?q=Delhi,IN&units=metric&appid=d3b51840b75dbd2c3a5f3c74c900ab82"

type WeatherData struct {
	Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	Cnt     float64 `json:"cnt"`
	List    []struct {
		Dt   float64 `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  float64 `json:"pressure"`
			SeaLevel  float64 `json:"sea_level"`
			GrndLevel float64 `json:"grnd_level"`
			Humidity  float64 `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			ID          float64 `json:"id"`
			Main        string  `json:"main"`
			Description string  `json:"description"`
			Icon        string  `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All float64 `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   float64 `json:"deg"`
			Gust  float64 `json:"gust"`
		} `json:"wind"`
		Visibility float64 `json:"visibility"`
		Pop        float64 `json:"pop"`
		Sys        struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
}

func FetchData() (*WeatherData, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var data WeatherData
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
