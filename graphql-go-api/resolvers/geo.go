package resolvers

import "graphql-go-api/models"

type Geo struct {
	geoData models.Geo
}

func (g *Geo) Lat() float64 {
	return g.geoData.Lat
}

func (g *Geo) Lng() float64 {
	return g.geoData.Lng
}
