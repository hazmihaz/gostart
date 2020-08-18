package user

import (
	"github.com/hazmihaz/gostart/internal/domain"
	"github.com/hazmihaz/gostart/pkg/log"
	tou "github.com/hazmihaz/gostart/pkg/strtouint"

	"github.com/gofiber/fiber"
)

type paging struct {
	offset int
	limit  int
}

// RegisterHandlers registers handlers for different HTTP requests.
func RegisterHandlers(g fiber.Router, logger log.Logger, userService domain.UserService) {
	g.Get("/user", func(c *fiber.Ctx) {
		p := new(paging)

		if err := c.QueryParser(p); err != nil {
			logger.Error(err)
		}

		user, err := userService.Query(c.Context(), p.offset, p.limit)
		if err != nil {
			logger.Error(err)
		} else {
			c.JSON(&user)
		}
	})

	g.Get("/user/:id", func(c *fiber.Ctx) {
		id, err := tou.Parse(c.Params("id"))
		user, err := userService.Get(c.Context(), id)
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
