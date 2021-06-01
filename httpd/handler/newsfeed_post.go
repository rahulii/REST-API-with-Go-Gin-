package handler

import (
	"news/httpd/platform/newsfeed"

	"github.com/gin-gonic/gin"

	"net/http"
)

type newsfeedpostRequest struct {
	Title string `json:"title"`
	Post string	`json:"post"`
}

func NewsFeedPost(feed *newsfeed.Repo) gin.HandlerFunc{
	return func(c *gin.Context) {
		
		requestBody := newsfeedpostRequest{}
		c.Bind(&requestBody)

		//create new Item
		item := newsfeed.Item{
			Title: requestBody.Title,
			Post: requestBody.Title,
		}
		//add item to feed
		feed.Add(item)

		c.Status(http.StatusNoContent)
	
	}
}