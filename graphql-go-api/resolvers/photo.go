package resolvers

import (
	"graphql-go-api/models"
)

type Photo struct {
	photoData models.Photo
}

func (p *Photo) AlbumId() int32 {
	return p.photoData.AlbumId
}

func (p *Photo) Id() int32 {
	return p.photoData.Id
}

func (p *Photo) Title() string {
	return p.photoData.Title
}

func (p *Photo) Url() string {
	return p.photoData.Url
}

func (p *Photo) ThumbnailUrl() string {
	return p.photoData.ThumbnailUrl
}
