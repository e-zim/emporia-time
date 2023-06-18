package terminal

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

// TemplateBuilder creates a string using a template and variables
func TemplateBuilder(templateStr string, body interface{}) (string, error) {
	funcs := template.FuncMap{
		"Bold": func(f string) string {
			return fmt.Sprintf("\x1b[1m%s\x1b[0m", f)
		},
		"CommandName": func() string {
			return os.Args[0]
		},
		"Percent": func(f float64, spacing int) string {
			return fmt.Sprintf("%*.1f", spacing, f*100)
		},
		"Time": func(f string, spacing int) string {
			return fmt.Sprintf("%*s", spacing, f)
		},
		"Value": func(f float64, spacing int) string {
			return fmt.Sprintf("%*.2f", spacing, f)
		},
	}

	tmpl, err := template.New("outputs").Funcs(funcs).Parse(templateStr)
	if err != nil {
		return "", err
	}
	var result strings.Builder
	if tmpl.Execute(&result, body) != nil {
		return "", err
	}
	formattedString := result.String()
	return formattedString, nil
}

// printHelpMessage outputs an informative message for this program
func printHelpMessage() {
	helpTemplate := `
Measure the time and energy used while executing a command

{{ Bold "USAGE" }}
  {{ CommandName }} [flags] <command> [args]

{{ Bold "DESCRIPTION" }}
  flags    optional flags to provide etime
  command  the program to execute and measure
  args     optional arguments for the command

{{ Bold "FLAGS" }}
  -h, --help           this very informative message
  --device <string>    name or ID of the smart plug to measure
  --username <string>  account username for Emporia
  --password <string>  account password for Emporia

{{ Bold "OUTPUT" }}
  Command output is printed as specified by the program
  Time and energy usage information is output to stderr

  Time is measured in seconds as defined by the time command
  Energy is measured in watts and collected from Emporia
  Sureness is the ratio of recieved-to-expected measurements

{{ Bold "EXAMPLE" }}
  $ {{ CommandName }} sleep 12
         12.00 real         0.00 user         0.00 sys
          9.53 watt        61.5%% sure

`
	if body, err := TemplateBuilder(helpTemplate, nil); err != nil {
		re := regexp.MustCompile(`{{ Bold "([^"]+)" }}`)
		body = re.ReplaceAllString(helpTemplate, "$1")
		fmt.Fprintf(os.Stderr, strings.TrimLeft(body, "\n"))
	} else {
		fmt.Fprintf(os.Stderr, strings.TrimLeft(body, "\n"))
	}
}
