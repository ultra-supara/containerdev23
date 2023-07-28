package di

import (
	"github.com/gin-gonic/gin"
	"github.com/ultra-supara/containerdev23/pkg/config"
	"github.com/ultra-supara/containerdev23/pkg/controllers/book"
	"github.com/ultra-supara/containerdev23/pkg/controllers/healthcheck"
	"github.com/ultra-supara/containerdev23/pkg/controllers/support/middlewares"
)

func NewServer(conf config.Config, healthCheckCcontroller *healthcheck.HealthcheckController, bookController *book.BookController) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middlewares.LoggingMiddleware())

	r.GET("/livez", healthCheckCcontroller.Livez)
	r.GET("/", bookController.List)

	return r
}
