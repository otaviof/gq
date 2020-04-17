package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const expected = "expected"

type writer struct {
	t *testing.T
}

func (w *writer) Write(p []byte) (int, error) {
	assert.Equal(w.t, expected, string(p))
	return len(p), nil
}

func TestGQ(t *testing.T) {
	payload := map[string]interface{}{
		"status": map[string]interface{}{
			"host": expected,
		},
	}

	gq, err := NewGQ("{{ .status.host }}", payload)
	require.NoError(t, err)
	require.NotNil(t, gq)

	wr := &writer{t: t}
	err = gq.Execute(wr)
	require.NoError(t, err)
}
