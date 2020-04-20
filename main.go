package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// rootCmd defines the gq command elements.
var rootCmd = &cobra.Command{
	Use:   "gq <template> [file]...",
	Args:  validateArgs,
	Run:   run,
	Short: "gq: Command-line Go Templates processor for YAML/JSON",
	Long: `### gq

Go templates based swissarmy knife for structured text formats, like YAML/JSON. It can process
multiple files directly, or via standard input.

Examples:
	$ gq '{{ index . "current-context" }}' ~/.kube/config
	$ cat ~/.kube/config |gq '{{ index . "current-context" }}'
	$ gq --type=json '{{ range .HttpHeaders }}{{ printf "%s\n" . }}{{ end }}' ~/.docker/config.json
	$ gq --type=toml '{{ range $c := .constraint }}{{ printf "%s\n" $c.name }}{{ end }}' Gopkg.toml
	`,
}

// init declare command-line arguments.
func init() {
	flags := rootCmd.PersistentFlags()

	flags.String("type", YAML, fmt.Sprintf(
		"File type, accepted formats: \"%v\"", strings.Join(SupportedContentTypes, ", ")))

	if err := viper.BindPFlags(flags); err != nil {
		panic(err)
	}
}

// exit when on error display a final message and error.
func exit(msg string, err error) {
	fmt.Printf("[ERROR] %s: %s\n", msg, err)
	os.Exit(1)
}

// validateArgs makes sure the arguments are informed accordingly.
func validateArgs(cmd *cobra.Command, args []string) error {
	if args == nil || len(args) == 0 {
		exit("not enough arguments", fmt.Errorf("use --help for options"))
	}
	return nil
}

// splitTemplateAndFiles split args slice, first entry is dedicated to hold go template, and all
// subsequent entries are files.
func splitTemplateAndFiles(args []string) (string, []string) {
	return args[0], args[1:]
}

// process go template against input.
func process(tmpl string, input *Input) {
	payload, err := input.Unmarshal()
	if err != nil {
		exit("error unmarshaling file contents", err)
	}

	g, err := NewGQ(tmpl, payload)
	if err != nil {
		exit("template validation error", err)
	}
	if err = g.Execute(os.Stdout); err != nil {
		exit("error processing template", err)
	}
}

// run execute the primary objective of this application.
func run(cmd *cobra.Command, args []string) {
	tmpl, filePaths := splitTemplateAndFiles(args)
	contentType := viper.GetString("type")

	if len(filePaths) == 0 {
		input := NewInput(contentType)
		if err := input.SlurpFile(os.Stdin); err != nil {
			exit("reading standard input", err)
		}
		process(tmpl, input)
	} else {
		for _, filePath := range filePaths {
			input := NewInput(contentType)
			if err := input.SlurpPath(filePath); err != nil {
				exit(fmt.Sprintf("error reading file '%s' contents", filePath), err)
			}
			process(tmpl, input)
		}
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
