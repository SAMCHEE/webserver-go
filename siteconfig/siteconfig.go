package siteconfig

import (
	"fmt"
	"time"
)

var Config *config

type config struct {
	Author        string
	Title         string
	Description   string
	Keywords      string
	Domain        string
	Host          string
	Port          uint32
	DatePublished time.Time
	DateModified  time.Time
}

func Init(modeDev bool, port uint32) error {
	c := &config{}
	c.Author = "go"
	c.Title = "my webserver"
	c.Description = "my webserver is .."
	c.Keywords = "my,webserver,keywords"
	c.Domain = "foobar.com"
	if modeDev {
		c.Host = fmt.Sprintf("http://localhost:%d", port)
	} else {
		c.Host = fmt.Sprintf("https://%s", c.Domain)
	}
	c.Port = port
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return err
	}
	c.DatePublished = time.Date(2023, time.Month(2), 13, 0, 0, 0, 0, loc)
	c.DateModified = time.Date(2023, time.Month(2), 13, 0, 0, 0, 0, loc)
	Config = c
	return nil
}
