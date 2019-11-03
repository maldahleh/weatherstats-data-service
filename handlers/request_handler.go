package handlers

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"weatherstatsData/request"
	"weatherstatsData/utils"

	log "github.com/sirupsen/logrus"
)

const rootUrl = "https://dd.weather.gc.ca/climate/observations/daily/csv/"

type yearlyData map[string]monthlyData
type monthlyData map[string]dailyData
type dailyData map[string]pointData
type pointData map[string]string

type StationResponse struct {
	Points yearlyData
}

func RetrieveData(stationId string, station request.StationRequest, dataPoints *[]string) StationResponse {
	data := make(yearlyData)
	for yearRaw, months := range *station.Months {
		year := strconv.Itoa(yearRaw)

		for _, monthRaw := range months {
			month := stringifyMonth(monthRaw)
			filePath := stationId + "-" + year + "-" + month + ".csv"

			err := utils.DownloadFile(filePath, createUrl(stationId, *station.Province, year, month))
			if err != nil {
				log.Error("Failed to download file ", filePath, " error ", err)
				continue
			}

			csvFile, err := os.Open(filePath)
			if err != nil {
				log.Error("Failed to open file ", filePath, " error ", err)
				continue
			}

			r := csv.NewReader(csvFile)
			for {
				record, err := r.Read()
				if err == io.EOF {
					break
				}

				if err != nil {
					log.Error("Error encountered reading ", filePath, " error ", err)
					break
				}

				fmt.Println(record)
			}

			err = csvFile.Close()
			if err != nil {
				log.Error("Failed to close file ", filePath, " error ", err)
			}

			err = utils.DeleteFile(filePath)
			if err != nil {
				log.Error("Failed to delete file ", filePath, " error ", err)
			}
		}

	}

	return StationResponse{
		Points: data,
	}
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
