package user

import (
	"strconv"

	"github.com/hazmihaz/gostart/internal/domain"
	"github.com/hazmihaz/gostart/pkg/log"

	"github.com/gofiber/fiber"
)

// RegisterHandlers registers handlers for different HTTP requests.
func RegisterHandlers(g *fiber.Group, logger log.Logger, userService domain.UserService) {
	g.Get("/user", func(c *fiber.Ctx) {
		user, err := userService.Get(c.Context(), 1)
		if err != nil {
			logger.Error(err)
		} else {
			c.JSON(&user)
		}
	})

	g.Get("/user/:id", func(c *fiber.Ctx) {
		id, err := strconv.ParseUint(c.Params("id"), 10, 64)
		idu := uint(id)
		user, err := userService.Get(c.Context(), idu)
		if err != nil {
			logger.Error(err)
			c.Next(err)
		} else {
			c.JSON(&user)
		}
	})

	g.Post("/user", func(c *fiber.Ctx) {
		user := domain.User{}

		if err := c.BodyParser(&user); err != nil {
			logger.Error(err)
			c.Next(err)
			return
		}

		user, err := userService.Create(c.Context(), user)
		if err != nil {
			logger.Error(err)
			c.Next(err)
			return
		}

		c.JSON(&user)
	})
}
