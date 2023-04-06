package link

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

// Link represents a link in a html
type Link struct {
	Href string
	Text string
}

// Parse will rake an html and return slice of Link
func Parse(r io.Reader) ([]Link, error) {

	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)

	links := make([]Link, len(nodes))

	for i, _ := range nodes {
		links[i] = buildLink(nodes[i])
	}

	return links, nil
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var ret string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}

	return strings.Join(strings.Fields(ret), " ")

}

func buildLink(n *html.Node) Link {
	var ret Link

	for _, attribute := range n.Attr {

		if attribute.Key == "href" {
			ret.Href = attribute.Val

			break
		}
	}

	ret.Text = text(n)

	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {

		return []*html.Node{n}
	}

	var ret []*html.Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		ret = append(ret, linkNodes(c)...)
	}

	return ret
}
