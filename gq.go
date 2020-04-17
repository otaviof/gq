package main

import (
	"io"
	"text/template"
)

type GQ struct {
	t       *template.Template
	payload map[string]interface{}
}

func (g *GQ) Execute(wr io.Writer) error {
	return g.t.Execute(wr, g.payload)
}

func NewGQ(tmpl string, payload map[string]interface{}) (*GQ, error) {
	t, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		return nil, err
	}
	return &GQ{t: t, payload: payload}, nil
}
