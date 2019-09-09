package main

import (
	"fmt"
	geojson "github.com/paulmach/go.geojson"
	"github.com/paulmach/orb"
	orbgeojson "github.com/paulmach/orb/geojson"
	orbplanar "github.com/paulmach/orb/planar"
)

func example() {
	fmt.Println("go.geojson")

	raw_geojson := funcRawGeojson()

	// go.geojson lightweight.
	fc_geojson, err := geojson.UnmarshalFeatureCollection(raw_geojson)
	fmt.Println(fc_geojson)
	fmt.Println(err)

	fc_orbgeojson, err := orbgeojson.UnmarshalFeatureCollection(raw_geojson)
	fmt.Println(fc_orbgeojson)
	fmt.Println(err)

	fmt.Println(isPointInsidePolygon(fc_orbgeojson, orb.Point{90.5, 90}))

	// https://github.com/paulmach/orb/tree/master/geojson
	// orb.geojson geometry arithmetic
	fmt.Println("orb")
	ring := orb.Ring{{0.0, 100.0}, {100.0, 100.0}, {100.0, 10.0}, {0.0, 0.0}, {0.0, 100.0}}
	polygon := orb.Polygon{ring}
	fc_orb := orbgeojson.NewFeatureCollection()
	fc_orb.Append(orbgeojson.NewFeature(polygon))

	fmt.Println(fc_orb.MarshalJSON())
	fmt.Println(isPointInsidePolygon(fc_orb, orb.Point{90.5, 90}))

}



func isPointInsidePolygon(fc *orbgeojson.FeatureCollection, point orb.Point) bool {
	for _, feature := range fc.Features {
		// Try on a MultiPolygon to begin
		multiPoly, isMulti := feature.Geometry.(orb.MultiPolygon)
		if isMulti {
			if orbplanar.MultiPolygonContains(multiPoly, point) {
				return true
			}
		} else {
			// Fallback to Polygon
			polygon, isPoly := feature.Geometry.(orb.Polygon)
			if isPoly {
				if orbplanar.PolygonContains(polygon, point) {
					return true
				}
			}
		}
	}
	return false
}
