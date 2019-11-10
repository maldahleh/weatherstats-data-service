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

const root = "https://dd.weather.gc.ca/climate/observations/daily/csv/"

type YearlyData map[string]monthlyData
type monthlyData map[string]dailyData
type dailyData map[string]pointData
type pointData map[string]string

type dataChannel struct {
	year  string
	month string
	data  dailyData
}

func RetrieveData(stationId string, station request.StationRequest, dataPoints *[]string) YearlyData {
	data := make(YearlyData)
	channel := make(chan dataChannel)

	channelCalls := 0
	for yearRaw, months := range *station.Months {
		year := strconv.Itoa(yearRaw)

		monthData := make(monthlyData)
		for _, monthRaw := range months {
			month := stringifyMonth(monthRaw)
			channelCalls++

			go downloadData(stationId, *station.Province, year, month, dataPoints, channel)
		}

		data[year] = monthData
	}

	for i := 0; i < channelCalls; i++ {
		channelData := <-channel
		if channelData.data == nil {
			continue
		}

		data[channelData.year][channelData.month] = channelData.data
	}

	return data
}

func downloadData(stationId string, province string, year string, month string, points *[] string, channel chan dataChannel) {
	path := stationId + "-" + year + "-" + month + ".csv"

	err := utils.DownloadFile(path, createUrl(stationId, province, year, month))
	if err != nil {
		log.Error("failed to download file", path, "error", err)
		channel <- dataChannel{
			year:  "",
			month: "",
			data:  nil,
		}

		return
	}

	file, err := os.Open(path)
	if err != nil {
		log.Error("failed to open file", path, "error", err)
		utils.DeleteFile(path)
		channel <- dataChannel{
			year:  "",
			month: "",
			data:  nil,
		}

		return
	}

	scanner := bufio.NewScanner(file)
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
		for _, requestedPoint := range *points {
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
		log.Fatal("error reading", file, "error", err)
	}

	err = file.Close()
	if err != nil {
		log.Error("failed to close file", path, "error", err)
	}

	utils.DeleteFile(path)
	channel <- dataChannel{
		year:  year,
		month: month,
		data:  dayData,
	}
}

func createUrl(stationId string, province string, year string, month string) string {
	return root + province + "/" + "climate_daily_" + province + "_" + stationId + "_" + year + "-" + month + "_PID.csv"
}

func stringifyMonth(month int) string {
	if month >= 10 {
		return strconv.Itoa(month)
	}

	return "0" + strconv.Itoa(month)
}
