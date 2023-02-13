package server

import "github.com/gofiber/template/html"

func engine() *html.Engine {
	e := html.New("./views", ".html")
	e.Reload(true) // html 소스 변경됐을때 서버 재시작 없이 핫리로드

	/*
		html에서 사용 할 커스텀함수 작성 하는 부분
		ex)
		e.AddFunc("Hello", func() string {
			return "world"
		})
	*/

	return e
}
