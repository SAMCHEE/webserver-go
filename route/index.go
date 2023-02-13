package route

import (
	"fmt"
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
	// Site의 config, host(localhost or domain)은 siteconfig 패키지 따로 만들어서 불러오는데
	// 우선 여기까지 기본 틀 먼저 익히신다음 더 작성 해볼게용 ㅎㅎ
	m["Title"] = "my webserver"
	m["Description"] = "my webserver is ..."
	m["Keywords"] = "my,webserver,golang"
	m["Author"] = "go"
	m["Canonical"] = fmt.Sprintf("http://localhost:4000%s", c.Path())
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
