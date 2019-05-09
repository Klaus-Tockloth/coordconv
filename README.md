# coordconv

Utility for **conv**erting **coord**inates between WGS84 Lon Lat, UTM and MGRS/UTMREF.

## Usage

``` TXT
$ ./coordconv

Program:
  Name    : ./coordconv
  Release : v0.1.0 - 2019/05/09
  Purpose : coordinates conversion
  Info    : Utility for converting coordinates between WGS84 Lon Lat, UTM and MGRS/UTMREF.

Supported conversions:
  UTMtoLL   : converts from UTM to LL
  UTMtoMGRS : converts from UTM to MGRS
  LLtoUTM   : converts from LL to UTM
  LLtoMGRS  : converts from LL to MGRS
  MGRStoUTM : converts from MGRS to UTM
  MGRStoLL  : converts from MGRS to LL

Data objects:
  UTM  : ZoneNumber ZoneLetter Easting Northing
  LL   : Longitude Latitude
  MGRS : String

Examples:
  ./coordconv UTMtoLL 32 U 399000 5757000
  ./coordconv UTMtoMGRS 32 U 399000 5757000
  ./coordconv LLtoUTM 7.53 51.95
  ./coordconv LLtoMGRS 7.53 51.95
  ./coordconv MGRStoUTM 32ULC989564
  ./coordconv MGRStoLL 32ULC9897356497

Abbreviations:
  Lon    : Longitude
  Lat    : Latitude
  MGRS   : Military Grid Reference System (same as UTMREF)
  UTM    : Universal Transverse Mercator
  UTMREF : UTM Reference System (same as MGRS)
  WGS84  : World Geodetic System 1984 (same as EPSG:4326)
```

## Output

```TXT
$ ./coordconv UTMtoLL 32 U 399000 5757000
32 U 399000 5757000 -> 7.53023117 51.95451906

$ ./coordconv UTMtoMGRS 32 U 399000 5757000
32 U 399000 5757000 -> 32ULC9900057000

$ ./coordconv LLtoUTM 7.53 51.95
7.53000000 51.95000000 -> 32 U 398973 5756497

$ ./coordconv LLtoMGRS 7.53 51.95
7.53000000 51.95000000 -> 32ULC9897356497

$ ./coordconv MGRStoUTM 32ULC989564
32ULC989564 -> 32 U 398900 5756400

$ ./coordconv MGRStoLL 32ULC9897356497
32ULC9897356497 -> 7.52998627 51.94999316
```

## Remarks

* See [releases](https://github.com/Klaus-Tockloth/coordconv/releases) for executable program.
* See package [coco](https://github.com/Klaus-Tockloth/coco) for Go library.
