# solar
Generate sunrise and sunset times for a given longitude, latitude, altitude and date.

## Usage

Longitude and latitude should be supplied as degrees west and north. Altitude should be supplied in meters (defaults to sea level). The date can be supplied in RFC3339 format (defaults to the current date).

    solar -la 54.9966 -lo 7.3086
    05:09:27 19:59:36

    solar -la 54.9966 -lo 7.3086 -alt 88.0
    05:06:57 20:02:06

    solar -la 54.9966 -lo 7.3086 -d 2016-12-25T00:00:00Z
    08:55:08 16:05:58

The times can optionally be returned as julian day values:

    solar -la 54.9966 -lo 7.3086 -fmt julian
    2457617.71490423 2457618.33306497
