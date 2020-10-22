package rest

import (
	"database/sql"
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/iwanjunaid/basesvc/adapter/controller"
	_ "github.com/iwanjunaid/basesvc/docs"
	"github.com/iwanjunaid/basesvc/infrastructure/rest/group"
	"github.com/iwanjunaid/basesvc/registry"
)

type RestImpl struct {
	port          string
	db            *sql.DB
	router        *fiber.App
	appController *controller.AppController
}

// @title BaseSVC API
// @version 1.0
// @description This is a sample basesvc server api.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func NewRest(port string, db *sql.DB) *RestImpl {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use("/swagger", swagger.Handler)

	registry := registry.NewRegistry(db)
	appController := registry.NewAppController()

	r := &RestImpl{
		db:            db,
		port:          port,
		router:        app,
		appController: &appController,
	}

	group.InitRoot(r)
	group.InitAuthorV1(r)

	return r
}

func (r *RestImpl) Serve() {
	if err := r.router.Listen(r.port); err != nil {
		log.Fatalln(err)
	}
}

func (r *RestImpl) GetRouter() *fiber.App {
	return r.router
}

func (r *RestImpl) GetAppController() *controller.AppController {
	return r.appController
}