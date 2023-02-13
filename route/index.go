package route

import (
	"net/http"

	"github.com/SAMCHEE/webserver-go/util"
	"github.com/gofiber/fiber/v2"
)

type indexHandler struct{}

// GET /
func (*indexHandler) index(c *fiber.Ctx) error {
	// response html template
	// ./views/~
	m := fiber.Map{} // Bind
	m["Title"] = "my webserver"
	m["Description"] = "my webserver is ..."
	m["Keywords"] = "my,webserver,golang"
	m["Author"] = "go"
	m["Canonical"] = c.Path()
	return c.Status(http.StatusOK).Render("index", m, "layout/global")
}

// GET /response-string
func (*indexHandler) resString(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("Hello, index")
}

// GET /response-json
func (*indexHandler) resJSON(c *fiber.Ctx) error {
	// response json
	return c.Status(http.StatusOK).JSON(util.ResponseOK("Hello, json"))
}

// BaseURL = /
func HandleIndex(r fiber.Router) {
	h := &indexHandler{}
	r.Get("/", h.index)
	r.Get("/response-string", h.resString)
	r.Get("/response-json", h.resJSON)
}
