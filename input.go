package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"

	"gopkg.in/yaml.v2"
)

const (
	// JSON format
	JSON = "json"
	// YAML format
	YAML = "yaml"
	// TOML format
	TOML = "toml"
)

// SupportedContentTypes list of support content-types.
var SupportedContentTypes = []string{YAML, JSON, TOML}

// Input for gq processing, comprises the bytes and content-type together.
type Input struct {
	b           []byte // input bufferr
	contentType string // content-type
}

// Unstructured format to extract from input
type Unstructured map[string]interface{}

// SlurpFile reads file bytes.
func (i *Input) SlurpFile(f *os.File) error {
	var err error
	i.b, err = ioutil.ReadAll(f)
	return err
}

// SlurpPath reads file from path.
func (i *Input) SlurpPath(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return i.SlurpFile(f)
}

// Unmarshal contents as unstructured format.
func (i *Input) Unmarshal() (map[string]interface{}, error) {
	var payload = map[string]interface{}{}
	var err error
	switch i.contentType {
	case YAML:
		err = yaml.Unmarshal(i.b, payload)
	case JSON:
		err = yaml.Unmarshal(i.b, payload)
	case TOML:
		err = toml.Unmarshal(i.b, &payload)
	default:
		return nil, fmt.Errorf("content-type '%s' is not supported", i.contentType)
	}
	return payload, err
}

// NewInput instantiate Input.
func NewInput(contentType string) *Input {
	return &Input{contentType: contentType}
}
