package resolvers

import "graphql-go-api/models"

type Address struct {
	addressData models.Address
}

func (a *Address) Street() string {
	return a.addressData.Street
}

func (a *Address) City() string {
	return a.addressData.City
}

func (a *Address) Geo() *Geo {
	return &Geo{geoData: a.addressData.Geo}
}
