package webapp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type WebApp struct {
	Router *Router
}

func NewWebApp() *WebApp {
	wp := &WebApp{
		Router: NewRouter(),
	}
	return wp
}

func (wp *WebApp) GET(path string, handler Handler) {
	wp.Router.Register("GET", path, handler)
}

func (wp *WebApp) POST(path string, handler Handler) {
	wp.Router.Register("POST", path, handler)
}

func (wp *WebApp) DELETE(path string, handler Handler) {
	wp.Router.Register("DELETE", path, handler)
}

func (wp *WebApp) PUT(path string, handler Handler) {
	wp.Router.Register("PUT", path, handler)
}

func (wp *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := string(r.URL.Path)
	method := strings.ToUpper(r.Method)

	found := wp.Router.FindRoute(path, method, w, r)

	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"path":   path,
		"method": method,
		"found":  fmt.Sprintf("%v", found),
	})

	if err != nil {
		log.Println("Could not encode the response ", err)
	}
}
