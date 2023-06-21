package router

import (
	"MovieDB/constants"
	"MovieDB/controller"

	"MovieDB/models"
	"MovieDB/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	addRoute(r)
	return r
}

func addRoute(r *gin.Engine) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(new(models.Movie))

	repository := repository.Repository{Db: db}
	moviec := controller.Moviecontroller{
		Repository: &repository,
	}

	routeM := r.Group(constants.Movies)
	{
		routeM.POST("", moviec.AddMovie)
		routeM.GET("", moviec.GetMovies)
		routeM.GET(constants.ID, moviec.GetByIdMovies)
		routeM.PATCH(constants.ID, moviec.UpdateMovie)
		routeM.DELETE(constants.ID, moviec.DeleteMovie)
	}

}
