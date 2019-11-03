package handlers

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"weatherstatsData/request"
	"weatherstatsData/utils"

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

			err := utils.DownloadFile(filePath, createUrl(stationId, *station.Province, year, month))
			if err != nil {
				log.Error("failed to download file", filePath, "error", err)
				continue
			}

			csvFile, err := os.Open(filePath)
			if err != nil {
				log.Error("failed to open file", filePath, "error", err)
				utils.DeleteFile(filePath)
				continue
			}

			scanner := bufio.NewScanner(csvFile)
			dayData := make(dailyData)
			headerLine := 0
			for scanner.Scan() {
				if headerLine < 25 {
					headerLine++
					continue
				}

				text := scanner.Text()
				if text == "" {
					continue
				}

				dayPointData := make(pointData)
				record := strings.Split(text, ",")
				for _, requestedPoint := range *dataPoints {
					switch requestedPoint {
					case "maxtemp":
						dayPointData[requestedPoint] = strings.Trim(record[5], "\"")
					case "mintemp":
						dayPointData[requestedPoint] = strings.Trim(record[7], "\"")
					case "meantemp":
						dayPointData[requestedPoint] = strings.Trim(record[9], "\"")
					case "rain":
						dayPointData[requestedPoint] = strings.Trim(record[15], "\"")
					case "snow":
						dayPointData[requestedPoint] = strings.Trim(record[17], "\"")
					case "precip":
						dayPointData[requestedPoint] = strings.Trim(record[19], "\"")
					case "snowgrnd":
						dayPointData[requestedPoint] = strings.Trim(record[21], "\"")
					case "maxgust":
						dayPointData[requestedPoint] = strings.Trim(record[25], "\"")
					}
				}

				dayData[strings.Trim(record[3], "\"")] = dayPointData
			}

			if err := scanner.Err(); err != nil {
				log.Fatal("error reading", csvFile, "error", err)
			}

			err = csvFile.Close()
			if err != nil {
				log.Error("failed to close file", filePath, "error", err)
			}

			utils.DeleteFile(filePath)
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
