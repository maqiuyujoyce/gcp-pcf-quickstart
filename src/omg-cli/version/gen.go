// +build ignore

/*
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

func main() {
	f, err := os.Create("version_gen.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	semver := os.Getenv("SEMVER")
	if semver == "" {
		fmt.Print("Set SEMVER to specify a version number. Using dev.\n")
		semver = "dev"
	}

	var revision []byte
	if revision, err = exec.Command("git", "log", "--pretty=format:%h", "-n", "1").Output(); err != nil {
		log.Fatal("Error fetching git revision", err)
	}

	versionTemplate.Execute(f, struct {
		Semver   string
		Revision string
	}{
		Semver:   semver,
		Revision: fmt.Sprintf("%s", revision),
	})

}

var versionTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
/*
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package version

const semver = "{{ .Semver }}"
const revision = "{{ .Revision }}"
`))