package resolvers

import "graphql-go-api/models"

type Geo struct {
	geoData models.Geo
}

func (g *Geo) Lat() string {
	return g.geoData.Lat
}

func (g *Geo) Lng() string {
	return g.geoData.Lng
}
