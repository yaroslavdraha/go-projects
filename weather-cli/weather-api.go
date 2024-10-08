package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WeatherData struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type Current struct {
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TempC            float64   `json:"temp_c"`
	TempF            float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMph          float64   `json:"wind_mph"`
	WindKph          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDir          string    `json:"wind_dir"`
	PressureMb       float64   `json:"pressure_mb"`
	PressureIn       float64   `json:"pressure_in"`
	PrecipMm         float64   `json:"precip_mm"`
	PrecipIn         float64   `json:"precip_in"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelslikeC       float64   `json:"feelslike_c"`
	FeelslikeF       float64   `json:"feelslike_f"`
	WindchillC       float64   `json:"windchill_c"`
	WindchillF       float64   `json:"windchill_f"`
	HeatindexC       float64   `json:"heatindex_c"`
	HeatindexF       float64   `json:"heatindex_f"`
	DewpointC        float64   `json:"dewpoint_c"`
	DewpointF        float64   `json:"dewpoint_f"`
	VisKm            float64   `json:"vis_km"`
	VisMiles         float64   `json:"vis_miles"`
	Uv               float64   `json:"uv"`
	GustMph          float64   `json:"gust_mph"`
	GustKph          float64   `json:"gust_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CurrentWeather struct {
	Temp      float64
	Condition string
	Location  string
}

func getCurrentWeather(location string) (*CurrentWeather, error) {
	apiKey, err := getApiKey()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s?q=%s&key=%s", WEATHER_API_CURRENT_ENDPOINT, location, apiKey)
	response, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if response.StatusCode == http.StatusOK {
		var successResponse WeatherData
		if err := json.Unmarshal(body, &successResponse); err != nil {
			return nil, fmt.Errorf("error to parse success json: %w", err)
		}

		currentWeather := CurrentWeather{
			Location:  successResponse.Location.Name,
			Temp:      successResponse.Current.TempC,
			Condition: successResponse.Current.Condition.Text,
		}
		return &currentWeather, nil
	} else {
		var errorResp ErrorResponse
		if err := json.Unmarshal(body, &errorResp); err != nil {
			return nil, fmt.Errorf("error to parse error json: %w", err)
		}

		return nil, errors.New(errorResp.Error.Message)
	}
}

func getApiKey() (string, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")

	if apiKey == "" {
		return "", errors.New("empty API key")
	}

	return apiKey, nil
}
