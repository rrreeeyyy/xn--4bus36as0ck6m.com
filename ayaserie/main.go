package ayaserie

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func init() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.NotFound(func(request *http.Request, render render.Render) {
		render.Redirect("http://riemelo.itigo.jp" + request.URL.Path)
	})

	http.Handle("/", m)
}
