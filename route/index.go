package route

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SAMCHEE/webserver-go/siteconfig"
	"github.com/gofiber/fiber/v2"
)

type indexHandler struct{}

// GET /
func (*indexHandler) index(c *fiber.Ctx) error {
	// response html template
	// ./views/~
	m, err := pageConfig(&pageConfigParams{
		Ctx:           c,
		Author:        siteconfig.Config.Author,
		Title:         siteconfig.Config.Title,
		Description:   siteconfig.Config.Description,
		Keywords:      siteconfig.Config.Keywords,
		DatePublished: siteconfig.Config.DatePublished,
		DateModified:  siteconfig.Config.DateModified,
	})
	if err != nil {
		c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(http.StatusOK).Render("index", m, "layout/global")
}

// GET /robots.txt
func (*indexHandler) robots(c *fiber.Ctx) error {
	var ss []string
	ss = append(ss, "User-agent: *")
	ss = append(ss, "Allow: /")
	ss = append(ss, fmt.Sprintf("Sitemap: %s/sitemap.txt", siteconfig.Config.Host))
	return c.Status(http.StatusOK).SendString(strings.Join(ss, "\n"))
}

// GET /sitemap.txt
func (*indexHandler) sitemap(c *fiber.Ctx) error {
	var ss []string
	ss = append(ss, siteconfig.Config.Host)
	return c.Status(http.StatusOK).SendString(strings.Join(ss, "\n"))
}

// BaseURL = /
func HandleIndex(r fiber.Router) {
	h := &indexHandler{}
	r.Get("/", h.index)
	r.Get("/robots.txt", h.robots)
	r.Get("/sitemap.txt", h.sitemap)
}
