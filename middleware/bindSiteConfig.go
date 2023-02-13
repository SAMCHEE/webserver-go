package middleware

import (
	"net/http"

	"github.com/SAMCHEE/webserver-go/siteconfig"
	"github.com/gofiber/fiber/v2"
)

func BindSiteConfig(c *fiber.Ctx) error {
	m := fiber.Map{}
	m["Site"] = siteconfig.Config
	if err := c.Bind(m); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Next()
}
