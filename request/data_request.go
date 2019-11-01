package request

type StationRequest struct {
	Province *string        `json:"province"`
	Months   *map[int][]int `json:"months"`
}

type DataRequest struct {
	Data       *map[string]StationRequest `json:"data"`
	DataPoints *[]string                  `json:"points"`
}
