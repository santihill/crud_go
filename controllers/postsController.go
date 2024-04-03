package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/santihill/crud_go/initializers"
	"github.com/santihill/crud_go/models"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Body   string
		Tittle string
	}
	c.Bind(&body)

	//Create a post
	post := models.Post{Tittle: body.Tittle, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id from URL
	id := c.Param("id")

	//Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	//Get the data off req body
	var body struct {
		Body   string
		Tittle string
	}
	c.Bind(&body)

	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Tittle: body.Tittle,
		Body:   body.Body,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	//Get the id off url
	id := c.Param("id")

	//Delete the posts
	initializers.DB.Delete(&models.Post{}, id)
	//Respond

	c.Status(200)
}
