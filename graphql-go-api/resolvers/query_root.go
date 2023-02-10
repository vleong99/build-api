package resolvers

import (
	"context"
	"encoding/json"
	"fmt"
	"graphql-go-api/models"
	"io"
	"net/http"
)

type QueryRoot struct{}

type PostArgs struct {
	Id int32
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

// func (_ *QueryRoot) photo() Photo {
// 	//insert
// }

// func (_ *QueryRoot) user() User {
// 	//insert
// }
