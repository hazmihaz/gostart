package user

import (
	"fmt"

	"github.com/hazmihaz/gostart/internal/domain"
	"github.com/hazmihaz/gostart/pkg/log"
	tui "github.com/hazmihaz/gostart/pkg/strtouint"

	"github.com/gofiber/fiber"
)

type paging struct {
	Offset int
	Limit  int
}

// RegisterHandlers registers handlers for different HTTP requests.
func RegisterHandlers(g fiber.Router, logger log.Logger, userService domain.UserService) {
	g.Get("/user", func(c *fiber.Ctx) {
		p := new(paging)

		if err := c.QueryParser(p); err != nil {
			logger.Error(err)
		}

		user, err := userService.Query(c.Context(), p.Offset, p.Limit)
		if err != nil {
			logger.Error(err)
		} else {
			c.JSON(&user)
		}
	})

	g.Get("/user/count", func(c *fiber.Ctx) {
		count, err := userService.Count(c.Context())
		fmt.Print(count)

		if err != nil {
			logger.Error(err)
		} else {
			c.JSON(&count)
		}
	})

	g.Get("/user/:id", func(c *fiber.Ctx) {
		id, err := tui.Parse(c.Params("id"))
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

	g.Put("/user", func(c *fiber.Ctx) {
		user := domain.User{}

		if err := c.BodyParser(&user); err != nil {
			logger.Error(err)
			c.Next(err)
			return
		}

		err := userService.Update(c.Context(), user)
		if err != nil {
			logger.Error(err)
			c.Next(err)
			return
		}

		user, err = userService.Get(c.Context(), user.ID)
		if err != nil {
			logger.Error(err)
			c.Next(err)
			return
		}

		c.JSON(&user)
	})

	g.Delete("/user/:id", func(c *fiber.Ctx) {
		id, err := tui.Parse(c.Params("id"))
		err = userService.Delete(c.Context(), id)
		if err != nil {
			logger.Error(err)
			c.Next(err)
		} else {
			c.JSON("success")
		}
	})

}
