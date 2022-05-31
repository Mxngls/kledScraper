package main

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func getRef(c *html.Node, in Result, l string) {

	// Create a refs variable for readability
	refs := in[len(in)-1].Senses[len(in[len(in)-1].Senses)-1].References

	if c.Data == "dl" {
		refs = append(refs, reference{})

	} else if CheckClass(c, fmt.Sprintf("manyLang%s", l)) && c.Data == "span" {
		// Get the type
		refs[len(refs)-1].Type = strings.TrimSpace(c.FirstChild.Data)

	} else if c.Data == "a" && CheckClass(c, "undL") {

		// Needed for reference to other entries
		refs[len(refs)-1].Values = append(refs[len(refs)-1].Values, value{})

		re := regexp.MustCompile("[0-9]+")
		id := c.Attr[0].Val
		id = re.FindAllString(id, -1)[0]

		// Set the id for future use
		// Get the reference itself
		refs[len(refs)-1].Values[len(refs[len(refs)-1].Values)-1].Id = id
		refs[len(refs)-1].Values[len(refs[len(refs)-1].Values)-1].Hangul = c.FirstChild.Data

	} else if c.Data == "dd" {

		// Simple usage reference
		refs[len(refs)-1].Values = append(refs[len(refs)-1].Values, value{})
		refs[len(refs)-1].Values[len(refs[len(refs)-1].Values)-1].Hangul = strings.TrimSpace(GetTextAll(c))

	} else if c.Data == " //.explain_list " {
		return
	}

	// Write to input
	in[len(in)-1].Senses[len(in[len(in)-1].Senses)-1].References = refs

	for e := c.FirstChild; e != nil; e = e.NextSibling {
		// Skip all commment nodes or nodes whose type is "script"
		if e.Type == html.CommentNode || e.Data == "script" {
			continue
		} else {
			getRef(e, in, l)
		}
	}
}
