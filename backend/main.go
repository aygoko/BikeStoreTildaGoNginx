package main

import (
	"flag"
	"log"

	"github.com/gofiber/cors/v2"
	"github.com/gofiber/fiber/v2"

	pkgHttp "github.com/aygoko/BikeStoreTildaGoNginx/backend/api/user"
	repository "github.com/aygoko/BikeStoreTildaGoNginx/backend/repository/ram_storage"
	"github.com/aygoko/BikeStoreTildaGoNginx/backend/usecases/service"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP server address")
	flag.Parse()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := pkgHttp.NewUserHandler(userService)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: []string{"GET", "POST", "PUT"},
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(fiber.Logger())
	app.Use(fiber.Recovery())

	log.Printf("Starting HTTP server on %s", *addr)
	err := app.Listen(*addr)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
