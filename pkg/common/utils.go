// Copyright 2021-2024 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
