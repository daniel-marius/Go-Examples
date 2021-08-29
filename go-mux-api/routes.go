package main

import (
  "math/rand"
  "encoding/json"
  "net/http"
  "go-mux-api/repository"
  "go-mux-api/entity"
)

// var (
//   posts []Post
// )

var (
  repo repository.PostRepository = repository.NewPostRepository()
)

// Run when the application starts
// func init() {
//   posts = []Post{Post{Id: 1, Title: "Title 1", Text: "Text 1"}}
// }

func getPosts(resp http.ResponseWriter, req *http.Request) {
  resp.Header().Set("Content-Type", "application/json")
  posts, err := repo.FindAll()
  // result, err := json.Marshal(posts)
  if err != nil {
    resp.WriteHeader(http.StatusInternalServerError)
    resp.Write([]byte(`{"error": "Error marshalling the posts array"}`))
    return
  }
  resp.WriteHeader(http.StatusOK)
  json.NewEncoder(response).Encode(posts)
  //cresp.Write(result)
}

func addPost(resp http.ResponseWriter, req *http.Request) {
  resp.Header().Set("Content-Type", "application/json")
  // var post Post
  var post entity.Post

  // Decode the body
  err := json.NewDecoder(req.Body).Decode(&post)
  if err != nil {
    resp.WriteHeader(http.StatusInternalServerError)
    resp.Write([]byte(`{"error": "Error unmarshalling the request"}`))
    return
  }
  // post.Id = len(posts) + 1
  post.Id = rand.Int63()
  //cposts = append(posts, post)
  repo.Save(&post)
  resp.WriteHeader(http.StatusOK)
  json.NewEncoder(response).Encode(post)
  // result, err := json.Marshal(posts)
  // resp.Write(result)
}
