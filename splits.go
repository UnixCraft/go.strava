package strava

type Split struct {
	Distance            float64 `json:"distance"`
	ElapsedTime         float64 `json:"elapsed_time"`
	ElevationDifference float64 `json:"elevation_difference"`
	MovingTime          float64 `json:"moving_time"`
	Split               int     `json:"split"`
}
