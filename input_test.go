package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func unmarshal(t *testing.T, input *Input) {
	payload, err := input.Unmarshal()
	assert.NoError(t, err)
	assert.NotNil(t, payload)
}

func TestInputYAML(t *testing.T) {
	input := NewInput(YAML)
	require.NotNil(t, input)

	err := input.SlurpPath("test/data/file.yaml")
	require.NoError(t, err)

	unmarshal(t, input)
}

func TestInputJSON(t *testing.T) {
	input := NewInput(JSON)
	require.NotNil(t, input)

	err := input.SlurpPath("test/data/file.json")
	require.NoError(t, err)

	unmarshal(t, input)
}

func TestInputStdin(t *testing.T) {
	_ = os.RemoveAll("/var/tmp/gq-e2e.yaml")

	stdin, err := os.Create("/var/tmp/gq-e2e.yaml")
	assert.NoError(t, err)
	stdin.WriteString("key: value")
	defer stdin.Close()

	input := NewInput(YAML)
	err = input.SlurpFile(stdin)
	assert.NoError(t, err)

	unmarshal(t, input)
}
