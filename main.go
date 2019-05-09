/*
Purpose:
- coordinates conversion

Description:
- Utility for converting coordinates between WGS84 Lon Lat, UTM and MGRS/UTMREF.

Releases:
- v0.1.0 - 2019/05/09 : initial release

Author:
- Klaus Tockloth

Copyright and license:
- Copyright (c) 2019 Klaus Tockloth
- MIT license

Permission is hereby granted, free of charge, to any person obtaining a copy of this software
and associated documentation files (the Software), to deal in the Software without restriction,
including without limitation the rights to use, copy, modify, merge, publish, distribute,
sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or
substantial portions of the Software.

The software is provided 'as is', without warranty of any kind, express or implied, including
but not limited to the warranties of merchantability, fitness for a particular purpose and
noninfringement. In no event shall the authors or copyright holders be liable for any claim,
damages or other liability, whether in an action of contract, tort or otherwise, arising from,
out of or in connection with the software or the use or other dealings in the software.

Contact (eMail):
- freizeitkarte@googlemail.com

Remarks:
- NN

Links:
- NN
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Klaus-Tockloth/coco"
)

// general program info
var (
	progName    = os.Args[0]
	progVersion = "v0.1.0"
	progDate    = "2019/05/09"
	progPurpose = "coordinates conversion"
	progInfo    = "Utility for converting coordinates between WGS84 Lon Lat, UTM and MGRS/UTMREF."
)

/*
main starts this program
*/
func main() {

	var ll coco.LL
	var utm coco.UTM
	var accuracy int
	var mgrs coco.MGRS
	var err error

	if len(os.Args) < 3 {
		printUsage()
	}

	conversion := os.Args[1]

	switch strings.ToUpper(conversion) {
	case "UTMTOLL":
		if len(os.Args) < 6 {
			printUsage()
		}
		utm, err = parseUTM(os.Args[2], os.Args[3], os.Args[4], os.Args[5])
		if err != nil {
			fmt.Printf("%s -> %v\n", utm, err)
			os.Exit(1)
		}
		ll, err = utm.ToLL()
		if err != nil {
			fmt.Printf("%s -> %v\n", utm, err)
			os.Exit(1)
		} else {
			fmt.Printf("%s -> %s\n", utm, ll)
		}

	case "UTMTOMGRS":
		if len(os.Args) < 6 {
			printUsage()
		}
		utm, err = parseUTM(os.Args[2], os.Args[3], os.Args[4], os.Args[5])
		if err != nil {
			fmt.Printf("%s -> %v\n", utm, err)
			os.Exit(1)
		}
		mgrs = utm.ToMGRS(1)
		fmt.Printf("%s -> %s\n", utm, mgrs)

	case "LLTOUTM":
		if len(os.Args) < 4 {
			printUsage()
		}
		ll, err = parseLL(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Printf("%s -> %v\n", ll, err)
			os.Exit(1)
		}
		utm = ll.ToUTM()
		fmt.Printf("%s -> %s\n", ll, utm)

	case "LLTOMGRS":
		if len(os.Args) < 4 {
			printUsage()
		}
		ll, err = parseLL(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Printf("%s -> %v\n", ll, err)
			os.Exit(1)
		}
		mgrs, err = ll.ToMGRS(1)
		if err != nil {
			fmt.Printf("%s -> %v\n", ll, err)
			os.Exit(1)
		} else {
			fmt.Printf("%s -> %s\n", ll, mgrs)
		}

	case "MGRSTOUTM":
		if len(os.Args) < 3 {
			printUsage()
		}
		mgrs = coco.MGRS(os.Args[2])
		utm, accuracy, err = mgrs.ToUTM()
		_ = accuracy
		if err != nil {
			fmt.Printf("%s -> %v\n", mgrs, err)
			os.Exit(1)
		} else {
			fmt.Printf("%s -> %s\n", mgrs, utm)
		}

	case "MGRSTOLL":
		if len(os.Args) < 3 {
			printUsage()
		}
		mgrs = coco.MGRS(os.Args[2])
		ll, accuracy, err = mgrs.ToLL()
		_ = accuracy
		if err != nil {
			fmt.Printf("%s -> %v\n", mgrs, err)
			os.Exit(1)
		} else {
			fmt.Printf("%s -> %s\n", mgrs, ll)
		}

	default:
		fmt.Printf("conversion <%s> not supported\n", conversion)
		os.Exit(1)
	}

	os.Exit(0)
}

/*
parseUTM parses string set into UTM object.
*/
func parseUTM(zoneNumber, zoneLetter, easting, northing string) (coco.UTM, error) {

	var utm coco.UTM
	var err error

	value, err := strconv.ParseInt(zoneNumber, 10, 0)
	if err != nil {
		return utm, fmt.Errorf("error <%v> parsing UTM zone number, input=%v", err, zoneNumber)
	}
	utm.ZoneNumber = int(value)

	utm.ZoneLetter = byte(zoneLetter[0])

	utm.Easting, err = strconv.ParseFloat(easting, 64)
	if err != nil {
		return utm, fmt.Errorf("error <%v> parsing UTM easting, input=%v", err, easting)
	}

	utm.Northing, err = strconv.ParseFloat(northing, 64)
	if err != nil {
		return utm, fmt.Errorf("error <%v> parsing UTM northing, input=%v", err, northing)
	}

	return utm, nil
}

/*
printUsage prints the usage of this program.
*/
func printUsage() {

	fmt.Printf("\nProgram:\n")
	fmt.Printf("  Name    : %s\n", progName)
	fmt.Printf("  Release : %s - %s\n", progVersion, progDate)
	fmt.Printf("  Purpose : %s\n", progPurpose)
	fmt.Printf("  Info    : %s\n", progInfo)

	fmt.Printf("\nSupported conversions:\n")
	fmt.Printf("  UTMtoLL   : converts from UTM to LL\n")
	fmt.Printf("  UTMtoMGRS : converts from UTM to MGRS\n")
	fmt.Printf("  LLtoUTM   : converts from LL to UTM\n")
	fmt.Printf("  LLtoMGRS  : converts from LL to MGRS\n")
	fmt.Printf("  MGRStoUTM : converts from MGRS to UTM\n")
	fmt.Printf("  MGRStoLL  : converts from MGRS to LL\n")

	fmt.Printf("\nData objects:\n")
	fmt.Printf("  UTM  : ZoneNumber ZoneLetter Easting Northing\n")
	fmt.Printf("  LL   : Longitude Latitude\n")
	fmt.Printf("  MGRS : String\n")

	fmt.Printf("\nExamples:\n")
	fmt.Printf("  %s UTMtoLL 32 U 399000 5757000\n", progName)
	fmt.Printf("  %s UTMtoMGRS 32 U 399000 5757000\n", progName)
	fmt.Printf("  %s LLtoUTM 7.53 51.95\n", progName)
	fmt.Printf("  %s LLtoMGRS 7.53 51.95\n", progName)
	fmt.Printf("  %s MGRStoUTM 32ULC989564\n", progName)
	fmt.Printf("  %s MGRStoLL 32ULC9897356497\n", progName)

	fmt.Printf("\nAbbreviations:\n")
	fmt.Printf("  Lon    : Longitude\n")
	fmt.Printf("  Lat    : Latitude\n")
	fmt.Printf("  MGRS   : Military Grid Reference System (same as UTMREF)\n")
	fmt.Printf("  UTM    : Universal Transverse Mercator\n")
	fmt.Printf("  UTMREF : UTM Reference System (same as MGRS)\n")
	fmt.Printf("  WGS84  : World Geodetic System 1984 (same as EPSG:4326)\n")

	fmt.Printf("\n")
	os.Exit(1)
}

/*
parseLL parses string set into LL object.
*/
func parseLL(lon, lat string) (coco.LL, error) {

	var ll coco.LL
	var err error

	ll.Lon, err = strconv.ParseFloat(lon, 64)
	if err != nil {
		return ll, fmt.Errorf("error <%v> parsing LL lon, input=%v", err, lon)
	}

	ll.Lat, err = strconv.ParseFloat(lat, 64)
	if err != nil {
		return ll, fmt.Errorf("error <%v> parsing LL lat, input=%v", err, lat)
	}

	return ll, nil
}

/*
formatLL formats Lon/Lat object.
*/
func formatLL(ll coco.LL) string {

	return fmt.Sprintf("%.8f %.8f", ll.Lon, ll.Lat)
}

/*
formatUTM formats UTM object.
*/
func formatUTM(utm coco.UTM) string {

	return fmt.Sprintf("%d %c %.0f %.0f", utm.ZoneNumber, utm.ZoneLetter, utm.Easting, utm.Northing)
}
