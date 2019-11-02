package request

import "strings"

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

	return true
}