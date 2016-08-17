package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"
	"github.com/nstruthers/riseandset"
)

const (
	FmtHMS     = "hms"
	FmtJulDate = "julian"
)

func main() {
	longitude := flag.Float64("lo", 0.0, "longitude (degrees west)")
	latitude := flag.Float64("la", 0.0, "latitude (degrees north)")
	altitude := flag.Float64("alt", 0.0, "altitude (meters)")
	dateStr :=
		flag.String(
			"d",
			fmt.Sprintf("%v", time.Now().Format(time.RFC3339)),
			"date")
	outFmt := flag.String("fmt", FmtHMS, "output format: hms or julian")

	flag.Parse()

	date, err := time.Parse(time.RFC3339, *dateStr)
	if err != nil {
		fmt.Printf("Problem parsing RFC3339 date: %v", err)
		os.Exit(1)
	}

	riseJulian, setJulian :=
		riseandset.Times(
			julDate(date),
			*longitude,
			*latitude,
			*altitude)

	output(*outFmt, riseJulian, setJulian)
}

func output(outFmt string, riseJulian, setJulian float64) {
	switch outFmt {
	case FmtJulDate:
		fmt.Printf("%.8f %.8f\n", riseJulian, setJulian)
	default:
		rh, rm, rs := julDayToHMS(riseJulian)
		sh, sm, ss := julDayToHMS(setJulian)
		fmt.Printf("%02d:%02d:%02d %02d:%02d:%02d\n", rh, rm, rs, sh, sm, ss)
	}
}

// Yanked from http://aa.usno.navy.mil/faq/docs/JD_Formula.php
// Ignores the time component of the date
func julDate(t time.Time) int {
	year := int(t.Year())
	month := int(t.Month())
	day := int(t.Day())

	return (day - 32075 +
		1461*(year+4800+(month-14)/12)/4 +
		367*(month-2-(month-14)/12*12)/12 -
		3*((year+4900+(month-14)/12)/100)/4)
}

func julDayToHMS(jDay float64) (uint, uint, uint) {
	julDayAccum := math.Mod((0.5+jDay), 1.0) * 24.0

	h := math.Mod(math.Floor(julDayAccum), 24.0)
	julDayAccum -= h

	m := math.Mod(math.Floor(julDayAccum*60.0), 60.0)
	julDayAccum -= m / 60.0

	s := math.Mod(math.Floor(julDayAccum*3600.0), 60.0)

	return uint(h), uint(m), uint(s)
}
