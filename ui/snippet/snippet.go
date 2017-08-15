package snippet

import (
	"github.com/gu-io/gu"

	"github.com/gu-io/gu/trees"

	"github.com/gu-io/gu/trees/elems"

	"github.com/gu-io/gu/trees/property"
)

//go:generate go run generate.go

// Snippet defines a component which implements the gu.Renderable interface.
type Snippet struct {
	gu.Reactive
	services gu.Services
}

// New returns a new instance of Snippet component.
func New(services gu.Services) *Snippet {
	return &Snippet{
		services: services,
		Reactive: gu.NewReactive(),
	}
}

// Render returns the markup for this Snippet component.
func (sn Snippet) Render() *trees.Markup {
	return elems.Div(property.ClassAttr("component", "snippet"))
}

// Apply adds the giving components Render() result to the
// provided root.
func (sn Snippet) Apply(root *trees.Markup) {
	root.AddChild(sn.Render())
}
