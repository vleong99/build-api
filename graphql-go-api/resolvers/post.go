package resolvers

import (
	"fmt"
	"graphql-go-api/models"
)

type Post struct {
	postData models.Post
}

func (p *Post) UserId() int32 {
	fmt.Println("called user ID")
	return p.postData.UserID
}

func (p *Post) Id() int32 {
	fmt.Println("called ID")
	return p.postData.ID
}

func (p *Post) Title() string {
	return p.postData.Title
}

func (p *Post) Body() string {
	return p.postData.Body
}

// func (p *Post) Comments() ([]*Comment, error) {
// 	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d/comments", p.postData.ID))
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer resp.Body.Close()

// 	b, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var model []models.Comment

// 	err = json.Unmarshal(b, &model)
// 	if err != nil {
// 		return nil, err
// 	}

// 	allComments := make([]*Comment, len(model))

// 	for i := 0; i < len(allComments); i++ {
// 		allComments[i] = &Comment{commentData: model[i]}
// 	}

// 	return allComments, nil
// }
