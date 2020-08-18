package main

import (
	"fmt"

	"github.com/hazmihaz/gostart/pkg/log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/hazmihaz/gostart/internal/auth"
	"github.com/hazmihaz/gostart/internal/domain"
	"github.com/hazmihaz/gostart/internal/user"
)

var version = "0.0.1"

func main() {
	// create root logger tagged with server version
	logger := log.New().With(nil, "version", version)

	app := fiber.New()
	app.Settings.ErrorHandler = errorHandler
	app.Use(middleware.Recover())

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/", handler)

	// init database
	db, err := gorm.Open("mysql", "root:pass@(localhost:3306)/gostart?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logger.Errorf("Error connecting to database. ", err)
	}
	defer db.Close()

	// auto migrate
	db.AutoMigrate(&domain.User{})

	userrep := user.NewRepository(db, logger)
	usersvc := user.NewService(userrep, logger)

	user.RegisterHandlers(v1, logger, usersvc)
	auth.RegisterHandlers(v1)

	logger.Error(app.Listen(3300))
}

func handler(c *fiber.Ctx) {
	c.Send("Go Rest Starter API V1")
}

type errorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func errorHandler(c *fiber.Ctx, err error) {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	// Check if it's an fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	fmt.Printf(err.Error())

	// Return HTTP response
	er := errorResponse{
		code,
		err.Error(),
	}
	// c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	c.Status(code).JSON(&er)
	//SendString(err.Error())
}
