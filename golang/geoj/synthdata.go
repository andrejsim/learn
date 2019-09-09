package main

import (
	"github.com/apex/log"
	geojson "github.com/paulmach/go.geojson"
)
func funcSyntheticGeojson() *geojson.FeatureCollection {
	raw_geojson := funcRawGeojson()

	fc_geojson, err := geojson.UnmarshalFeatureCollection(raw_geojson)
	if err != nil {
		log.Info(err.Error())
		return nil
	}

	return fc_geojson
}

func funcRawGeojson() []byte {
	raw_geojson := []byte(`
			{   
				"type": "FeatureCollection",
				"features": [
					{
						"type": "Feature",
						"properties": {
							"name": "europe",
							"description": "datasetbbox",
							"dataset": "radar"
						},
						"geometry": {
							"type": "Polygon",
							"coordinates": [
								[
									[
										-21.09375,
										32.84267363195431
									],
									[
										30.585937499999996,
										32.84267363195431
									],
									[
										30.585937499999996,
										60.930432202923335
									],
									[
										-21.09375,
										60.930432202923335
									],
									[
										-21.09375,
										32.84267363195431
									]
								]
							]
						}
					}
				]
			}`)
	return raw_geojson
}