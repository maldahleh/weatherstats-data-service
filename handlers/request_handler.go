package handlers

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"weatherstatsData/request"

	log "github.com/sirupsen/logrus"
)

const rootUrl = "https://dd.weather.gc.ca/climate/observations/daily/csv/"

type YearlyData map[string]monthlyData
type monthlyData map[string]dailyData
type dailyData map[string]pointData
type pointData map[string]string

func RetrieveData(stationId string, station request.StationRequest, dataPoints *[]string) YearlyData {
	data := make(YearlyData)
	for yearRaw, months := range *station.Months {
		year := strconv.Itoa(yearRaw)

		monthData := make(monthlyData)
		for _, monthRaw := range months {
			month := stringifyMonth(monthRaw)
			filePath := stationId + "-" + year + "-" + month + ".csv"

			//err := utils.DownloadFile(filePath, createUrl(stationId, *station.Province, year, month))
			//if err != nil {
			//	log.Error("Failed to download file ", filePath, " error ", err)
			//	continue
			//}

			csvFile, err := os.Open(filePath)
			if err != nil {
				log.Error("Failed to open file ", filePath, " error ", err)
				continue
			}

			r := csv.NewReader(csvFile)
			r.TrimLeadingSpace = true

			dayData := make(dailyData)
			for {
				record, err := r.Read()
				if err == io.EOF {
					break
				}

				if err != nil {
					log.Error("Error encountered reading ", filePath, " error ", err)
					break
				}

				dayPointData := make(pointData)
				for _, requestedPoint := range *dataPoints {
					switch requestedPoint {
					case "maxtemp":
						dayPointData[requestedPoint] = record[5]
					case "mintemp":
						dayPointData[requestedPoint] = record[7]
					case "meantemp":
						dayPointData[requestedPoint] = record[9]
					case "rain":
						dayPointData[requestedPoint] = record[15]
					case "snow":
						dayPointData[requestedPoint] = record[17]
					case "precip":
						dayPointData[requestedPoint] = record[19]
					case "snowgrnd":
						dayPointData[requestedPoint] = record[21]
					case "maxgust":
						dayPointData[requestedPoint] = record[25]
					}
				}

				dayData[record[3]] = dayPointData
			}

			err = csvFile.Close()
			if err != nil {
				log.Error("Failed to close file ", filePath, " error ", err)
			}

			//err = utils.DeleteFile(filePath)
			//if err != nil {
			//	log.Error("Failed to delete file ", filePath, " error ", err)
			//}

			monthData[month] = dayData
		}

		data[year] = monthData
	}

	return data
}

func createUrl(stationId string, province string, year string, month string) string {
	return rootUrl + province + "/" + "climate_daily_" + province + "_" + stationId + "_" + year + "-" + month + "_PID.csv"
}

func stringifyMonth(month int) string {
	if month >= 10 {
		return strconv.Itoa(month)
	}

	return "0" + strconv.Itoa(month)
}
