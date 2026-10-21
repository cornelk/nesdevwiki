package main

import (
	"iter"
	"slices"
	"strings"

	"golang.org/x/net/html"
)

func isElement(n *html.Node, tag string) bool {
	return n.Type == html.ElementNode && n.Data == tag
}

func attr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

func setAttr(n *html.Node, key, val string) {
	for i := range n.Attr {
		if n.Attr[i].Key == key {
			n.Attr[i].Val = val
			return
		}
	}
	n.Attr = append(n.Attr, html.Attribute{Key: key, Val: val})
}

func removeAttr(n *html.Node, key string) {
	n.Attr = slices.DeleteFunc(n.Attr, func(a html.Attribute) bool { return a.Key == key })
}

func hasClass(n *html.Node, class string) bool {
	return slices.Contains(strings.Fields(attr(n, "class")), class)
}

// descendants yields every node below n in document order.
func descendants(n *html.Node) iter.Seq[*html.Node] {
	return func(yield func(*html.Node) bool) {
		var walk func(*html.Node) bool
		walk = func(parent *html.Node) bool {
			for c := parent.FirstChild; c != nil; {
				next := c.NextSibling // allows callers to remove c while iterating
				if !yield(c) || !walk(c) {
					return false
				}
				c = next
			}
			return true
		}
		walk(n)
	}
}

// find returns the first node below n that matches.
func find(n *html.Node, match func(*html.Node) bool) *html.Node {
	for c := range descendants(n) {
		if match(c) {
			return c
		}
	}
	return nil
}

// findAll collects all matching nodes below n.
func findAll(n *html.Node, match func(*html.Node) bool) []*html.Node {
	var out []*html.Node
	for c := range descendants(n) {
		if match(c) {
			out = append(out, c)
		}
	}
	return out
}

// text returns the concatenated text content of n.
func text(n *html.Node) string {
	var b strings.Builder
	for c := range descendants(n) {
		if c.Type == html.TextNode {
			b.WriteString(c.Data)
		}
	}
	return b.String()
}

func remove(n *html.Node) {
	if n.Parent != nil {
		n.Parent.RemoveChild(n)
	}
}

// unwrap replaces n with its children.
func unwrap(n *html.Node) {
	parent := n.Parent
	if parent == nil {
		return
	}
	for c := n.FirstChild; c != nil; {
		next := c.NextSibling
		n.RemoveChild(c)
		parent.InsertBefore(c, n)
		c = next
	}
	parent.RemoveChild(n)
}

// prependChild inserts child as the first child of n.
func prependChild(n, child *html.Node) {
	if n.FirstChild == nil {
		n.AppendChild(child)
		return
	}
	n.InsertBefore(child, n.FirstChild)
}
