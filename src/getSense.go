package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func getSenses(n *html.Node, s []sense, l string) {

	if CheckClass(n, fmt.Sprintf("multiTrans manyLang%s mb10 printArea", l)) ||
		CheckClass(n, fmt.Sprintf("subMultiTrans manyLang%s mb10 printArea", l)) {
		s[len(s)-1].Translation = cleanStringSpecial([]byte(GetTextAll(n)))

	} else if CheckClass(n, "senseDef ml20 printArea") ||
		CheckClass(n, "subSenseDef ml20 printArea") {
		// Get the korean definition
		s[len(s)-1].KrDefinition = GetTextSingle(n.LastChild)

	} else if CheckClass(n, fmt.Sprintf("multiSenseDef manyLang%s ml20 printArea", l)) ||
		CheckClass(n, fmt.Sprintf("subMultiSenseDef manyLang%s ml20 printArea", l)) {
		// Get the english definition
		s[len(s)-1].Definition = GetTextAll(n)

	} else if CheckClass(n, "dot printArea") {
		// Get the examples
		for e := n.FirstChild; e != nil; e = e.NextSibling {
			if e.Type == html.CommentNode || e.Data == "script" || len(strings.TrimSpace(e.Data)) == 0 {
				continue

			} else {
				var ex = example{}
				if e.Data == "li" {
					s[len(s)-1].Examples = append(s[len(s)-1].Examples, ex)
					v := GetContent(e, "b")
					s[len(s)-1].Examples[len(s[len(s)-1].Examples)-1].Value = v
					s[len(s)-1].Examples[len(s[len(s)-1].Examples)-1].Type = GetTextAll(e)
				} else if e.Parent.Data != "ul" {
					break
				}
			}
		}

	} else if CheckClass(n, "heading_wrap printArea") ||
		CheckClass(n, "heading_wrap dotted printArea") {
		s[len(s)-1].Reference = append(s[len(s)-1].Reference, ref{})
		getRefs(n, s[len(s)-1].Reference, l)
	}

	for a := n.FirstChild; a != nil; a = a.NextSibling {
		// Skip all commment nodes or nodes whose type is "script"
		if a.Type == html.CommentNode || a.Data == "script" || len(strings.TrimSpace(a.Data)) == 0 {
			continue
		} else {
			getSenses(a, s, l)
		}
	}
}
