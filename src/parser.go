package main

import (
	"io"
	"strconv"

	"golang.org/x/net/html"
)

func ParseView(result_html io.Reader, id string, l string) (res Result, err error) {
	var r Result
	doc, err := html.Parse(result_html)
	dfsv(doc, &r, l, id)
	return r, err
}

func dfsv(n *html.Node, in *Result, l string, id string) {

	// Create a senses variable for readability
	senses := (*in)[len(*in)-1].Senses

	// Get the Hangul
	if CheckClass(n, "word_head") {

		// Append to array
		*in = append(*in, Ideom{})

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c != nil {
				if c.Type == html.TextNode {
					(*in)[len(*in)-1].Word = c.Data
				} else if c.Type == html.ElementNode {
					(*in)[len(*in)-1].HomonymNumber, _ = strconv.Atoi(GetTextAll(c))
					break
				}
			}
		}

		// Get the ideom itself and its type
	} else if CheckClass(n, "idiom_title printArea") {

		// Set the id for future reference
		(*in)[len(*in)-1].Id, _ = strconv.Atoi(id)
		(*in)[len(*in)-1].Hangul = GetTextAll(n.LastChild.PrevSibling)
		getType(n, *in, l)

	} else if CheckClass(n, "explain_list") && CheckClass(n.Parent, "explain_list_wrap") {
		// Initialize a slice to store the sense information in

		// Append a slice to the senses array
		senses = append(
			senses, sense{})

	} else if CheckClass(n, "subMultiTrans manyLang6 mb10 printArea") {
		// Get the translation
		senses[len(senses)-1].Translation = cleanStringSpecial([]byte(GetTextAll(n)))

	} else if CheckClass(n, "subSenseDef ml20 printArea") {
		// Get the Korean definition
		senses[len(senses)-1].KrDefinition = GetTextAll(n)

	} else if CheckClass(n, "subMultiSenseDef manyLang6 ml20 printArea") {
		// Get the definition
		senses[len(senses)-1].Definition = GetTextAll(n)

	} else if CheckClass(n, "dot printArea") &&
		CheckClass(n.PrevSibling.PrevSibling, "subMultiSenseDef manyLang6 ml20 printArea") &&
		n.FirstChild.NextSibling != nil {
		// Get the example
		senses[len(senses)-1].Example = GetContent(n.FirstChild.NextSibling, "b")

	} else if CheckClass(n, "heading_wrap dotted printArea") {
		// Get references
		getRef(n, *in, l)
	}

	// Write to input
	(*in)[len(*in)-1].Senses = senses

	// Traverse the tree of nodes vi depth-first search
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// Skip all commment nodes or nodes whose type is "script"
		if c.Type == html.CommentNode || c.Data == "script" {
			continue
		} else {
			dfsv(c, in, l, id)
		}
	}
}
