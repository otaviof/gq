package main

import (
	"io"
	"text/template"
)

// GQ holds the template and payload.
type GQ struct {
	t       *template.Template     // go template
	payload map[string]interface{} // payload instance
}

// Execute render template results on informed writer.
func (g *GQ) Execute(wr io.Writer) error {
	return g.t.Execute(wr, g.payload)
}

// NewGQ instantiate GQ, preparing go template for execution.
func NewGQ(tmpl string, payload map[string]interface{}) (*GQ, error) {
	t, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		return nil, err
	}
	return &GQ{t: t, payload: payload}, nil
}
