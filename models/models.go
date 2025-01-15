package models

type WeatherDuration struct {
	Hours   int
	Minutes int
	Seconds int
}

type WeatherAverages struct {
	Temperature float64
	Humidity    float64
}
