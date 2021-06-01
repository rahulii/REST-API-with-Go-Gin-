package handler

import (
	"news/httpd/platform/newsfeed"

	"github.com/gin-gonic/gin"

	"net/http"
)

func NewsFeedGet(feed *newsfeed.Repo) gin.HandlerFunc{
	return func(c *gin.Context) {
	results := feed.GetAll()	
	c.JSON(http.StatusOK,results)
	
	}
}