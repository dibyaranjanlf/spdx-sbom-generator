// SPDX-License-Identifier: Apache-2.0

package javamaven

import (
	"strings"
)

type command string

var (
	VersionCmd command = "java -version"
)

// Parse ...
func (c command) Parse() []string {
	cmd := strings.TrimSpace(string(c))
	return strings.Fields(cmd)
}
