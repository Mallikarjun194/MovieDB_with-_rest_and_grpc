package controller

import (
	"MovieDB/models"
	"MovieDB/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Moviecontroller struct {
	Repository repository.RepositoryI
}

func (m *Moviecontroller) AddMovie(c *gin.Context) {
	var movie models.Movie

	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	movie.Id = uuid.New().String()

	err = m.Repository.Create(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

func (m *Moviecontroller) GetMovies(c *gin.Context) {
	var movies []models.Movie
	err := m.Repository.QueryAll(&movies)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (m *Moviecontroller) GetByIdMovies(c *gin.Context) {
	var movie models.Movie
	id := c.Param("id")
	movie.Id = id
	err := m.Repository.Query(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (m *Moviecontroller) UpdateMovie(c *gin.Context) {
	var movie models.MovieUpdate
	id := c.Param("id")
	movie.Id = id
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	var movieModel models.Movie = models.Movie(movie)

	err = m.Repository.Update(&movieModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (m *Moviecontroller) DeleteMovie(c *gin.Context) {
	var movie models.MovieUpdate
	id := c.Param("id")
	movie.Id = id
	err := m.Repository.Delete(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	Message := models.Message{
		Message: fmt.Sprintf("movie %s deleted successfully ", movie.Id),
	}
	c.JSON(http.StatusOK, Message)
}
