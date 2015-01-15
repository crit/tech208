package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/polds/MyIP"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		Layout:     "layout",
	}))

	storer.InitDb()
	cacher.InitCache()
	// cacher.Engine = "none"

	m.Use(storageCheck())

	m.Get("", func(out render.Render) {
		ip, _ := myip.GetMyIP()
		out.HTML(200, "index", ip)
	})

	m.Group("/api", func(m martini.Router) {

		m.Get("/people", func(out render.Render) {
			out.JSON(200, map[string][]Person{"people": PersonList()})
		})

		m.Put("/people", func(out render.Render, req *http.Request) {
			name, email := req.FormValue("name"), req.FormValue("email")

			if err := PersonCreate(name, email); err != nil {
				hadError(out, err)
			} else {
				out.Status(201)
			}
		})
	})

	m.NotFound(func(out render.Render) {
		out.HTML(400, "404", nil)
	})

	m.Run()
}

func storageCheck() martini.Handler {
	return func(out render.Render) {
		if storer.MissingDb() {
			out.HTML(500, "error", "Database is unavailable. Check server log for reason.")
			return
		}
	}
}

func alwaysError() martini.Handler {
	return func(out render.Render) {
		out.HTML(500, "error", "Database was unable to start!")
	}
}

func hadError(out render.Render, err error) {
	out.JSON(500, err.Error())
}
