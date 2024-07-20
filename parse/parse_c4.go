package parse

import (
	"regexp"
	"strings"
)

type C4Item struct {
	Name          string
	Type          string
	Parameters    []Parameter
	Documentation string
}

type Parameter struct {
	Name         string
	Optional     bool
	DefaultValue string
}

// text should be full text of a puml file containing c4 model definitions
func ExtractC4Items(text string) []C4Item {
	var out []C4Item
	var currDocs string
	var docsBuf []string

	procRe := regexp.MustCompile(`^\s*!(unquoted\s+)?procedure\s+(\w+)\((.*)\)`)

	for _, line := range strings.Split(text, "\n") {
		// end of documentation block
		if strings.HasPrefix(line, "' ##") {
			currDocs = strings.Join(docsBuf, "\n")
			docsBuf = []string{}
			continue
		}

		// handle doc comments
		if strings.HasPrefix(line, "'") {
			docsBuf = append(docsBuf, strings.TrimSpace(strings.TrimPrefix(line, "'")))
		} else if strings.HasPrefix(line, "!") {
			// handle procedure definitions
			if procMatch := procRe.FindStringSubmatch(line); procMatch != nil {
				out = append(out, C4Item{
					Name:          procMatch[2],
					Type:          "procedure", // TODO: possibly change this to currDocs and put something else in for documentation
					Parameters:    parseParameters(procMatch[3]),
					Documentation: currDocs,
				})
			}
			// TODO: plantuml functions
		}

		// reset docs buffer when non-comment line is found
		// TODO: use newly found comments for next single item only
		// - this will mainly be for functions once they are included
		if !strings.HasPrefix(line, "'") && len(docsBuf) > 0 {
			docsBuf = []string{}
		}
	}

	return out
}

func parseParameters(params string) []Parameter {
	params = strings.TrimSpace(params)
	if params == "" {
		return []Parameter{}
	}

	var out []Parameter

	for _, param := range strings.Split(params, ",") {
		param = strings.TrimSpace(param)
		if strings.Contains(param, "=") {
			parts := strings.SplitN(param, "=", 2)
			out = append(out, Parameter{
				Name:         strings.TrimSpace(parts[0]),
				Optional:     true,
				DefaultValue: strings.TrimSpace(parts[1]),
			})
		} else {
			out = append(out, Parameter{
				Name:         param,
				Optional:     false,
				DefaultValue: "",
			})
		}
	}

	return out
}
