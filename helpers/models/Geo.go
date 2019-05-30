package models

type Geo struct {
	ID        string  `json:"id"`
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}
