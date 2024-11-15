package controllers

import (
	"errors"
	"log"

	"gorm.io/gorm"

	"github.com/VarthanV/hyper-todo/models"
	hyper "github.com/VarthanV/hyper/core"
)

type Controller struct {
	DB *gorm.DB
}

type createPostRequest struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

type errorResponse struct {
	Message string `json:"message,omitempty"`
}

func (c *Controller) CreatePost(w hyper.ResponseWriter, r *hyper.Request) {
	var (
		request = createPostRequest{}
	)

	err := r.Bind(&request)
	if err != nil {
		log.Println("error in binding request ", err)
		w.WriteJSON(400, errorResponse{Message: "invalid request"})
		return
	}

	p := models.Post{
		Title: request.Title,
		Body:  request.Body,
	}
	err = c.DB.Model(&models.Post{}).Create(&p).Error
	if err != nil {
		log.Println("error in creating post ", err)
		w.WriteJSON(500, errorResponse{Message: errors.Join(errors.New("unable to insert in db"), err).Error()})
		return
	}

	w.WriteJSON(200, p)
}

func (c *Controller) GetAllPosts(w hyper.ResponseWriter, r *hyper.Request) {
	var (
		posts = []models.Post{}
	)

	err := c.DB.Model(&models.Post{}).Find(&posts).Error
	if err != nil {
		log.Println("error in getting all posts ", err)
		w.WriteJSON(500, errorResponse{Message: "error in getting all posts"})
		return
	}
	w.WriteJSON(200, posts)
}
