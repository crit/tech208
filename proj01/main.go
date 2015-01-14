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

	if err := dbInit(); err != nil {
		m.Use(AlwaysError())
	} else {
		m.Use(StorageCheck())
	}

	m.Get("", Page)
	m.Post("/register", Register)
	m.Get("/current", Current)
	m.NotFound(NotFound)

	m.Run()
}

func StorageCheck() martini.Handler {
	return func(out render.Render) {
		if dbMissing() {
			out.HTML(500, "error", "Database is unavailable. Check server log for reason.")
			return
		}
	}
}

func AlwaysError() martini.Handler {
	return func(out render.Render) {
		out.HTML(500, "error", "Database is unavailable. Check server log for reason.")
	}
}

func Page(out render.Render) {
	ip, _ := myip.GetMyIP()
	out.HTML(200, "index", ip)
}

func Register(out render.Render, req *http.Request) {
	name := req.FormValue("name")
	email := req.FormValue("email")

	if name == "" {
		out.Redirect("/?status=error")
		return
	}

	db.Exec("insert into people (name, email) values (?, ?)", name, email)

	out.Redirect("/?status=success")
	// ToDo: Flush MEMCACHED!
}

func Current(out render.Render) {
	// ToDo: Run this against MEMCACHED
	out.JSON(200, map[string][]string{"current": ListPeople()})
}

func NotFound(out render.Render) {
	out.HTML(400, "404", nil)
}
