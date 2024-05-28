package gorouter

import "regexp"

type Node struct {
	regex    *regexp.Regexp
	handler  Handler
	nType    string
	next     *Node
	previous *Node
}
