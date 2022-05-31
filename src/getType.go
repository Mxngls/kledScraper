package main

import (
	"golang.org/x/net/html"
)

func getType(c *html.Node, in Result, l string) {

	if CheckClass(c, "manyLang6") {
		in[len(in)-1].Type = GetTextAll(c)
	}

	for e := c.FirstChild; e != nil; e = e.NextSibling {
		// Skip all commment nodes or nodes whose type is "script"
		if e.Type == html.CommentNode || e.Data == "script" {
			continue
		} else {
			getType(e, in, l)
		}
	}
}
