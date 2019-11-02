package request

import (
	"strings"
)

var provinces = [...]string{
	"AB",
	"BC",
	"MB",
	"NB",
	"NL",
	"NS",
	"NT",
	"NU",
	"ON",
	"PE",
	"QC",
	"SK",
	"YT",
}

var dataPoints = [...]string{
	"maxtemp",
	"mintemp",
	"meantemp",
	"rain",
	"snow",
	"precip",
	"maxgust",
}

type StationRequest struct {
	Province *string        `json:"province"`
	Months   *map[int][]int `json:"months"`
}

func (r *StationRequest) validateProvince() bool {
	*r.Province = strings.ToUpper(*r.Province)
	for _, value := range provinces {
		if value == *r.Province {
			return true
		}
	}

	return false
}

func (r *StationRequest) validateMonths() bool {
	for key, value := range *r.Months {
		if key < 1900 || key > 2019 {
			return false
		}

		for _, month := range value {
			if month < 1 || month > 12 {
				return false
			}
		}
	}

	return true
}

func (r *StationRequest) validate() bool {
	return r.Province != nil && r.Months != nil && r.validateProvince() && r.validateMonths()
}

type DataRequest struct {
	Data       *map[string]StationRequest `json:"data"`
	DataPoints *[]string                  `json:"points"`
}

func (r *DataRequest) Validate() bool {
	if r.Data == nil {
		return false
	}

	for _, v := range *r.Data {
		if !v.validate() {
			return false
		}
	}

	if r.DataPoints == nil {
		return false
	}

	points := make([]string, len(*r.DataPoints))
	for k, v := range *r.DataPoints {
		loweredPoint := strings.ToLower(v)
		if !validPoint(loweredPoint) {
			return false
		}

		points[k] = loweredPoint
	}

	r.DataPoints = &points
	return true
}

func validPoint(point string) bool {
	for _, loopPoint := range dataPoints {
		if loopPoint == point {
			return true
		}
	}

	return false
}
