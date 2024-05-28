package gorouter

import (
	"net/http"
	"regexp"
	"strings"
)

var Methods = [5]string{"POST", "GET", "DELETE", "PUT", "PATCH"}

type Param struct {
	Key   string
	Value string
}

type Params []Param

type Handler func(http.ResponseWriter, *http.Request, Params) error

type Router struct {
	tree map[string]*Node
}

func (r *Router) Register(method string, path string, handler Handler) {
	isValidMethod := checkValidMethods(method)

	if !isValidMethod {
		panic("Method is not valid")
	}

	if len(path) < 1 || path[0] != '/' {
		panic("Path Should Not Be Empty Or Should Start With /")
	}

	if handler == nil {
		panic("Router Cannot Deal With Null Handle Function")
	}

	if r.tree == nil {
		r.tree = make(map[string]*Node)
	}

	root := r.tree[method]
	if root == nil {
		root = &Node{
			nType: "Root",
		}
		r.tree[method] = root
	}

	regex := regexp.MustCompile("^" + path + "$")
	r.AttachRoute(regex, method, handler)
}

func (r *Router) AttachRoute(regex *regexp.Regexp, method string, handler Handler) {
	root := r.tree[method]

	if root == nil {
		panic("There is no root created for " + method + " on the tree")
	}

	if root.next == nil {
		childNode := &Node{
			regex:    regex,
			handler:  handler,
			nType:    "Child",
			next:     nil,
			previous: root,
		}
		root.next = childNode
		return
	}
	attached := false
	current := root.next
	for !attached {
		if current.regex == regex {
			panic("Route Already Attached to " + method + " Method")
		}
		if current.next != nil {
			current = current.next
			continue
		}

		childNode := &Node{
			regex:    regex,
			handler:  handler,
			nType:    "Child",
			next:     nil,
			previous: current,
		}
		current.next = childNode
		attached = true
	}
}

func checkValidMethods(method string) bool {
	for _, v := range Methods {
		if v == strings.ToUpper(method) {
			return true
		}
	}
	return false
}

func New() *Router {
	return &Router{}

}
