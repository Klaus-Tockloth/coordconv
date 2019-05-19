# coordconv

Utility for **conv**erting **coord**inates between WGS84 Lon Lat, UTM and MGRS/UTMREF.

## Usage

``` TXT
$ ./coordconv

Program:
  Name    : ./coordconv
  Release : v0.2.0 - 2019/05/10
  Purpose : coordinates conversion
  Info    : Utility for converting coordinates between WGS84 Lon Lat, UTM and MGRS/UTMREF.

Supported conversions:
  UTM2LatLon  : converts from UTM to LatLon
  UTM2MGRS    : converts from UTM to MGRS
  LatLon2UTM  : converts from LatLon to UTM
  LatLon2MGRS : converts from LatLon to MGRS
  MGRS2UTM    : converts from MGRS to UTM
  MGRS2LatLon : converts from MGRS to LatLon

Data objects:
  UTM    : ZoneNumber ZoneLetter Easting Northing
  LatLon : Longitude Latitude
  MGRS   : String

Examples:
  ./coordconv UTM2LatLon 32U 399000 5757000
  ./coordconv UTM2MGRS 32U 399000 5757000
  ./coordconv LatLon2UTM 51.95 7.53
  ./coordconv LatLon2MGRS 51.95 7.53
  ./coordconv MGRS2UTM 32ULC989564
  ./coordconv MGRS2LatLon 32ULC9897356497

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
$ ./coordconv UTM2LatLon 32U 399000 5757000
32U 399000 5757000 -> 51.954519 7.530231

$ ./coordconv UTM2MGRS 32U 399000 5757000
32U 399000 5757000 -> 32ULC9900057000

$ ./coordconv LatLon2UTM 51.95 7.53
51.950000 7.530000 -> 32U 398973 5756497

$ ./coordconv LatLon2MGRS 51.95 7.53
51.950000 7.530000 -> 32ULC9897356497

$ ./coordconv MGRS2UTM 32ULC989564
32ULC989564 -> 32U 398900 5756400

$ ./coordconv MGRS2LatLon 32ULC9897356497
32ULC9897356497 -> 51.949993 7.529986
```

## Remarks

* See [releases](https://github.com/Klaus-Tockloth/coordconv/releases) for executable program.
* See package [coco](https://github.com/Klaus-Tockloth/coco) for Go library.
* UTM format = zone number, zone letter (not hemisphere), easting, northing
