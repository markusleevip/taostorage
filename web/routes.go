package taoweb

import (
	mux "github.com/julienschmidt/httprouter"
	"github.com/markusleevip/taostorage/web/action"
	"github.com/markusleevip/taostorage/web/album"
	"github.com/markusleevip/taostorage/web/upload"
)

type Route struct {
	Method string
	Path   string
	Handle mux.Handle // httprouter package as mux
}

var ()
type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		Index,
	},
	Route{
		"GET",
		"/posts",
		PostIndex,
	},
	Route{
		Method: "GET",
		Path:   "/test",
		Handle: action.TestTaodb,
	},
	Route{
		Method: "GET",
		Path:   "/albums/:prePath",
		Handle: album.List,
	},
	Route{
		Method: "GET",
		Path:   "/show/:filePath/:fileName",
		Handle: album.Show,
	},
	Route{
		Method: "POST",
		Path:   "/upload/",
		Handle: upload.Controller{}.Upload,
	},

}

func NewRouter() *mux.Router {

	router := mux.New()

	for _, route := range routes {

		router.Handle(route.Method, route.Path, route.Handle)

	}

	return router
}
