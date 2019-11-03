# WeatherStats Data Service

This repo houses a Go microservice that responds to requests with historical climate data from Environment and Climate 
Change Canada's servers.

### Data Points
Acceptable data points are a combination of:

- maxtemp - Maximum temperature for the day
- mintemp - Minimum temperature for the day
- meantemp - Mean temperature for the day
- rain - Rain for the day in millimeters
- snow - Snow for the day in centimetres
- snowgrnd - Snow on the ground on the given day in cm
- precip - Total precipitation for the day in mm
- maxgust - Maximum wind gust for the day in km/h

### Endpoints
**/** - Returns the requested data as a JSON object

When making a request, you must include a proper payload in the body of your request for it to be parsed, here is a
sample request body:

```json
{
  "data": {
    "2300426": {
      "province": "NU",
      "months": {
        "2008": [
          7,
          8
        ]
      }
    }
  },
  "points": [
    "maxtemp",
    "rain"
  ]
}
```

The data dictionary identifies which climate stations you would like data from, the key of each item in the dictionary
must be a valid climate station ID, for each station you are expected to provide the following:
 
- the province key which identifies where the station resides using the province/territory's 2 digit code
- the months key which should contain a dictionary identifying which months and years you would like data from

The points key should be a list of datapoints you would like information for, for a list of valid data points click 
[here](###data-points)

### Additional Information

You are expected to provide valid data or you will receive an error as a response, for a complete list of weather 
stations including their province and months with available data, look 
[here](https://github.com/maldahleh/weatherstats-location-service)