//go:generate wire
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/gin-gonic/gin"
	"github.com/ultra-supara/containerdev23/pkg/config"
	"github.com/ultra-supara/containerdev23/pkg/controllers/book"
	"github.com/ultra-supara/containerdev23/pkg/controllers/healthcheck"
	"github.com/ultra-supara/containerdev23/pkg/db"
	"github.com/ultra-supara/containerdev23/pkg/usecases"
)

var APISet = wire.NewSet(
		healthcheck.NewHealthcheckController,
		book.NewBookController,
)

var UsecaseSet = wire.NewSet(
		usecases.NewListBooks,
)

var DBSet = wire.NewSet(
		db.NewDatabaseCon,
		db.NewBookRepository,
)

// ProvideGinEngine is a provider function that returns a new instance of *gin.Engine.
func ProvideGinEngine() *gin.Engine {
    return gin.Default()
}

func Initialize(conf config.Config) (*gin.Engine, error) {
    wire.Build(
        APISet,
        DBSet,
        UsecaseSet,
        ProvideGinEngine,
    )
    return nil, nil
}
