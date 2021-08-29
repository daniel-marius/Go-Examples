package main

import (
  "go_news_feed/httpd/handler"
  "github.com/gin-gonic/gin"
  // "fmt"
  "go_news_feed/platform/newsfeed"
)

func main() {
  feed := newsfeed.New()
  r := gin.Default()

  r.GET("/newsfeed", handler.NewsFeedGet(feed))
  r.POST("/newsfeed", handler.NewsFeedPost(feed))

  r.Run()

  // feed := newsfeed.New()
  // fmt.Println(feed)
  // feed.Add(newsfeed.Item{"Hello", "Golang"})
  // fmt.Println(feed)
}
