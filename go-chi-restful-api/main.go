package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	port := "8080" //allows server to listen for and receive incoming requests from clients

	if fromEnv := os.Getenv("PORT"); fromEnv != "" { //allow user to set PORT number in terminal
		port = fromEnv
	}

	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Mount("/posts", postsRouter())

	log.Fatal(http.ListenAndServe(":"+port, r))

}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"

// 	"github.com/go-chi/chi"
// )

// type Posts struct {
// 	UserID int    `json:"userId"`
// 	ID     int    `json:"id"`
// 	Title  string `json:"title"`
// 	Body   string `json:"body"`
// }

// func main() {

// 	router := chi.NewRouter()
// 	router.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
// 		tests := r.URL.Query()["test"]
// 		if len(tests) > 0 {
// 			fmt.Println("hello")
// 		}
// 		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		defer resp.Body.Close()

// 		data, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		var posts []Posts
// 		w.Write(data)
// 		json.Unmarshal(data, &posts)

// 		fmt.Println(posts)
// 	})

// 	http.ListenAndServe(":8080", router)

// }

// func main() {
// 	port := "8080"

// 	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
// 		port = fromEnv
// 	}

// 	log.Printf("Starting up on http://localhost:%s", port)

// 	r := chi.NewRouter()

// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/plain")
// 		w.Write([]byte("Hello World!"))
// 	})

// 	r.Mount("/posts", postsResource{}.Routes())

// 	log.Fatal(http.ListenAndServe(":"+port, r))
// }
