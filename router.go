package webapp

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

var Methods = [5]string{"POST", "GET", "DELETE", "PUT", "PATCH"}

type Param struct {
	Key   string
	Value string
}

type Node struct {
	regex    *regexp.Regexp
	handler  Handler
	nType    string
	next     *Node
	previous *Node
}

type Params []Param

type Handler func(http.ResponseWriter, *http.Request)

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
	fmt.Println("Router Registered Successfully", regex)
}

func (r *Router) FindRoute(path string, method string, res http.ResponseWriter, req *http.Request) bool {
	root := r.tree[method]

	if root == nil {
		fmt.Println("Not Found")
	}

	regex := regexp.MustCompile("^" + path + "$")

	last := false
	current := root
	for !last {
		if reflect.DeepEqual(current.regex, regex) {
			current.handler(res, req)
			return true
		}
		if current.next == nil {
			last = true
		}
		current = current.next
	}
	return false
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
		if reflect.DeepEqual(current.regex, regex) {
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

func NewRouter() *Router {
	return &Router{}

}
