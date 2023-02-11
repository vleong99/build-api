package resolvers

import (
	"context"
	"encoding/json"
	"fmt"
	"graphql-go-api/models"
	"io"
	"log"
	"net/http"
)

type QueryRoot struct{}

type PostArgs struct {
	Id int32
}

type PhotoArgs struct {
	AlbumId int32
	Title   string
}

func (_ *QueryRoot) Post(ctx context.Context, args PostArgs) (*Post, error) {
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", args.Id))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var model models.Post

	err = json.Unmarshal(b, &model)
	if err != nil {
		return nil, err
	}

	return &Post{postData: model}, nil
}

func (_ *QueryRoot) Photo(ctx context.Context, args PhotoArgs) ([]*Photo, error) {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/photos")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var allPhotos []*Photo

	decoder := json.NewDecoder(resp.Body)

	_, err = decoder.Token()
	if err != nil {
		log.Fatal(err)
	}

	for decoder.More() {
		var photo models.Photo

		err = decoder.Decode(&photo)
		if err != nil {
			log.Fatal(err)
		}

		if (args.AlbumId == 0 || photo.AlbumId == args.AlbumId) && (args.Title == "" || photo.Title == args.Title) {
			allPhotos = append(allPhotos, &Photo{photoData: photo})
		}
	}

	_, err = decoder.Token()
	if err != nil {
		log.Fatal(err)
	}

	return allPhotos, nil

}

// func (_ *QueryRoot) user() User {
// 	//insert
// }
