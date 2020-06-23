package user

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func RegisterHandlers(g *fiber.Group) {
	fmt.Println("routes")

	g.Get("/user", get)
}

func get(c *fiber.Ctx) {
	c.JSON(&User{
		1,
		"mi",
		"mi@mail.co",
		"$asdzxc",
	})
	// c.Send("Users")
}
