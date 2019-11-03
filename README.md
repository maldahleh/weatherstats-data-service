# WeatherStats Data Service

This repo houses a Go microservice that responds to requests with historical climate data from Environment and Climate 
Change Canada's servers.

### Table of Contents
- [WeatherStats Data Service](#weatherstats-data-service)
    + [Data Points](#data-points)
    + [Endpoints](#endpoints)
      - [Sample Response](#sample-response)
    + [Additional Information](#additional-information)

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

When making a request, you must include a proper payload in the body of your request for it to be responded to, here is
 a sample request body:

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
    "mintemp"
  ]
}
```

The data dictionary identifies which climate stations you would like data from, the key of each item in the dictionary
must be a valid climate station ID, for each station you are expected to provide the following:
 
- the province key which identifies where the station resides using the province/territory's 2 digit code
- the months key which should contain a dictionary identifying which months and years you would like data from

The points key should be a list of data points you would like information for, for a list of valid data points click 
[here](###data-points)

#### Sample Response
Here is a sample JSON response for the above request
```json
{
  "2300426": {
    "2008": {
      "07": {
        "25": {
          "maxtemp": "9.4",
          "mintemp": "7.6"
        },
        "26": {
          "maxtemp": "20.0",
          "mintemp": "8.0"
        },
        "27": {
          "maxtemp": "16.0",
          "mintemp": "9.7"
        },
        "28": {
          "maxtemp": "16.3",
          "mintemp": "7.4"
        },
        "29": {
          "maxtemp": "13.2",
          "mintemp": "6.9"
        },
        "30": {
          "maxtemp": "13.4",
          "mintemp": "9.0"
        },
        "31": {
          "maxtemp": "13.9",
          "mintemp": "8.7"
        }
      },
      "08": {
        "01": {
          "maxtemp": "13.5",
          "mintemp": "8.2"
        },
        "02": {
          "maxtemp": "12.1",
          "mintemp": "7.3"
        },
        "03": {
          "maxtemp": "9.8",
          "mintemp": "8.2"
        },
        "04": {
          "maxtemp": "11.3",
          "mintemp": "8.4"
        },
        "05": {
          "maxtemp": "14.7",
          "mintemp": "9.1"
        },
        "06": {
          "maxtemp": "19.6",
          "mintemp": "10.6"
        },
        "07": {
          "maxtemp": "18.4",
          "mintemp": "9.4"
        },
        "08": {
          "maxtemp": "17.4",
          "mintemp": "9.3"
        },
        "09": {
          "maxtemp": "16.9",
          "mintemp": "8.9"
        },
        "10": {
          "maxtemp": "18.3",
          "mintemp": "10.6"
        },
        "11": {
          "maxtemp": "15.7",
          "mintemp": "10.2"
        },
        "12": {
          "maxtemp": "19.6",
          "mintemp": "11.8"
        },
        "13": {
          "maxtemp": "14.6",
          "mintemp": "10.4"
        },
        "14": {
          "maxtemp": "18.8",
          "mintemp": "10.1"
        },
        "15": {
          "maxtemp": "16.6",
          "mintemp": "9.0"
        },
        "16": {
          "maxtemp": "14.5",
          "mintemp": "7.2"
        },
        "17": {
          "maxtemp": "18.9",
          "mintemp": "8.2"
        },
        "18": {
          "maxtemp": "16.7",
          "mintemp": "10.9"
        },
        "19": {
          "maxtemp": "24.7",
          "mintemp": "10.6"
        },
        "20": {
          "maxtemp": "11.1",
          "mintemp": "4.7"
        },
        "21": {
          "maxtemp": "8.7",
          "mintemp": "4.6"
        },
        "22": {
          "maxtemp": "11.3",
          "mintemp": "5.7"
        },
        "23": {
          "maxtemp": "12.2",
          "mintemp": "4.6"
        },
        "24": {
          "maxtemp": "19.5",
          "mintemp": "5.0"
        },
        "25": {
          "maxtemp": "9.9",
          "mintemp": "7.5"
        },
        "26": {
          "maxtemp": "10.3",
          "mintemp": "7.7"
        },
        "27": {
          "maxtemp": "12.4",
          "mintemp": "7.4"
        },
        "28": {
          "maxtemp": "10.3",
          "mintemp": "7.5"
        },
        "29": {
          "maxtemp": "14.9",
          "mintemp": "5.2"
        },
        "30": {
          "maxtemp": "13.2",
          "mintemp": "6.1"
        },
        "31": {
          "maxtemp": "13.6",
          "mintemp": "7.8"
        }
      }
    }
  }
}
```

### Additional Information

You are expected to provide valid data or you will receive an error as a response, for a complete list of weather 
stations including their province and months with available data, look 
[here](https://github.com/maldahleh/weatherstats-location-service)