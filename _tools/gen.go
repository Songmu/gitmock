package main

import (
	"fmt"
	"io/ioutil"

	"github.com/danverbraganza/varcaser/varcaser"
)

var tmpl = `package gitmock
%s`

var funcTmpl = `
// %s Do git %s
func (gm *GitMock) %s(args ...string) (string, string, error) {
	arg := append([]string{"%s"}, args...)
	return gm.Do(arg...)
}
`

func main() {
	commands := []string{
		"clone",
		"init",
		"add",
		"mv",
		"reset",
		"rm",
		"log",
		"show",
		"status",
		"branch",
		"checkout",
		"commit",
		"diff",
		"merge",
		"rebase",
		"tag",
		"ls-files",
	}
	content := ""
	for _, com := range commands {
		funcName := varcaser.Caser{From: varcaser.KebabCase, To: varcaser.UpperCamelCase}.String(com)
		content += fmt.Sprintf(funcTmpl, funcName, com, funcName, com)
	}
	content = fmt.Sprintf(tmpl, content)
	ioutil.WriteFile("gitmock_gen.go", []byte(content), 0644)
}
