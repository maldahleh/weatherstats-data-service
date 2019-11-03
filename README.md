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