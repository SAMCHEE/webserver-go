package route

import (
	"errors"
	"fmt"
	"time"

	"github.com/SAMCHEE/webserver-go/siteconfig"
	"github.com/gofiber/fiber/v2"
)

type page struct {
	Author        string
	Title         string
	Description   string
	Keywords      string
	Path          string
	URL           string
	DatePublished time.Time
	DateModified  time.Time
	// TODO: Thumbnail
}

type pageConfigParams struct {
	Ctx           *fiber.Ctx
	Author        string
	Title         string
	Description   string
	Keywords      string
	DatePublished time.Time
	DateModified  time.Time
	// TODO: Thumbnail
}

func pageConfig(arg *pageConfigParams) (fiber.Map, error) {
	if arg.Ctx == nil {
		return fiber.Map{}, errors.New("Nil pointer: (*pageConfigParams).Ctx")
	}
	m := fiber.Map{}
	path := arg.Ctx.Path()
	m["Page"] = &page{
		Author:        arg.Author,
		Title:         arg.Title,
		Description:   arg.Description,
		Keywords:      arg.Keywords,
		Path:          path,
		URL:           fmt.Sprintf("%s%s", siteconfig.Config.Host, path),
		DatePublished: arg.DatePublished,
		DateModified:  arg.DateModified,
		// TODO: Thumbnail
	}
	return m, nil
}
