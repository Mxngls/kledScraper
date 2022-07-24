package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func getRefs(c *html.Node, r []ref, l string) {

	if CheckClass(c, fmt.Sprintf("manyLang%s", l)) && c.Data == "span" {
		r[len(r)-1].Type = strings.TrimSpace(c.FirstChild.Data)

	} else if c.Data == "a" && CheckClass(c, "undL") {
		re := regexp.MustCompile("[0-9]+")
		id := c.Attr[0].Val
		id = re.FindAllString(id, -1)[0]
		r[len(r)-1].Id, _ = strconv.Atoi(id)
		r[len(r)-1].Value = GetContent(c.Parent, "sup")

	} else if c.Data == "dd" {
		for a := c.FirstChild; a != nil; a = a.NextSibling {
			if a.Type == html.CommentNode || a.Data == "script" || len(strings.TrimSpace(a.Data)) == 0 {
				continue
			} else {
				r[len(r)-1].Value = strings.TrimSpace(c.FirstChild.Data)
				break
			}
		}

	} else if c.Data == " 자세히 보기 끝 " || c.Data == " //.idiom_adage " {
		return
	}

	for e := c.FirstChild; e != nil; e = e.NextSibling {
		// Skip all commment nodes or nodes whose type is "script"
		if e.Type == html.CommentNode || e.Data == "script" {
			continue
		} else {
			getRefs(e, r, l)
		}
	}
}
