package gitmock_test

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Songmu/gitmock"
)

func ExampleNew() {
	gm, err := gitmock.New("")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(gm.RepoPath())
	gm.Init()
	file := "hoge/fuga.txt"
	gm.PutFile(file, "aaa\n")
	gm.Add(file)
	gm.Commit("-m", "initial commit")
	out, _, _ := gm.Status()
	if !strings.HasPrefix(out, "#") {
		out = "# " + out
	}
	fmt.Print(out)
	// Output:
	// # On branch master
	// nothing to commit, working tree clean
}
