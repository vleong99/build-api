package main

import (
	"io"
	"net/http"
)

//GET /posts - Read list of posts

func List(w http.ResponseWriter, r http.Request) {
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.URL.RawQuery = r.URL.Query().Encode()

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//POST /posts - Create new post

func Create(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.WriteHeader(201)

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GET /posts/{id} - Read a single post by id

func Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)

	// 	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/", nil)
	// 	if err != nil {
	// 		http.Error(w,err.Error(), 404)
	// 		return
	// 	}
	// req.URL.RawQuery

	// resp, err := http.DefaultClient.Do(req)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	// w.Header().Add("Content-Type", "application/json")
	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
