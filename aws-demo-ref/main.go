package main

import (
	"github.com/crit/critical-go/cacher"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/polds/MyIP"
	"net/http"
	"os"
)

func main() {
	m := martini.Classic()

	opts := render.Options{}
	opts.Directory = os.Getenv("TEMPLATES")
	opts.Extensions = []string{".html"}
	opts.Layout = "layout"

	m.Use(render.Renderer(opts))

	storer.InitDb()
	MigratePerson()

	cacher.InitCache(cacher.Options{
		Hosts:  os.Getenv("MEMCACHE"),
		Engine: cacher.MEMCACHE,
	})

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
				out.JSON(500, err.Error())
			} else {
				out.Status(201)
			}
		})

	})

	m.Get("/pwd", func() string {
		dir, _ := os.Getwd()
		return dir
	})

	m.Get("/templates", func() string {
		return os.Getenv("TEMPLATES")
	})

	m.NotFound(func(out render.Render) {
		out.HTML(404, "404", nil)
	})

	m.Run()
}

func storageCheck() martini.Handler {
	return func(out render.Render) {
		if storer.MissingDb() {
			out.HTML(500, "error", "Database is unavailable. Check server log for reason.")
		}
	}
}
