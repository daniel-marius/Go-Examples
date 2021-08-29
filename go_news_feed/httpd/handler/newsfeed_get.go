package handler

import (
  "net/http"
  "go_news_feed/platform/newsfeed"
  "github.com/gin-gonic/gin"
)

func NewsFeedGet(feed newsfeed.Getter) gin.HandlerFunc {
  return func(c *gin.Context) {
    results := feed.GetAll()
    c.JSON(http.StatusOK, results)
  }
}
