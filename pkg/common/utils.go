package common

import (
	"bytes"
	"text/template"
)

// templateToString takes a Go template string and a map[string]string of values to parse,
// and returns a string in the template format
func templateToString(t string, m map[string]string) (string, error) {
	var err error
	var buf *bytes.Buffer

	tmpl, err := template.New("").Parse(t)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(buf, m)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
