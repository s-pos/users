package main

import (
	"spos/users/controllers"
	"spos/users/repository"
	"spos/users/routes"
	"spos/users/usecases"

	"github.com/s-pos/go-utils/adapter"
	"github.com/s-pos/go-utils/config"
	"github.com/s-pos/go-utils/middleware"
	"github.com/s-pos/go-utils/utils/server"
)

func init() {
	serviceName := "users"

	config.Load(serviceName)
}

func main() {
	log := config.Logrus()
	timezone := config.Timezone()

	db := adapter.DBConnection()
	redis := adapter.GetClientRedis()

	mdl := middleware.NewMiddleware(redis, log, timezone)

	repo := repository.NewRepository(db)
	uc := usecases.NewUsecases(repo)
	ctrl := controllers.NewController(uc)

	router := routes.NewRouter(mdl, ctrl)

	log.Fatal(server.Wrapper(router.Router()))
}
