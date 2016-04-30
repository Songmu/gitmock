package gitmock_test

import (
	"fmt"
	"os"

	"github.com/Songmu/gitmock"
)

func ExampleNew() {
	gm, _ := gitmock.New()
	repo := gm.RepoPath()
	defer os.RemoveAll(repo)
	gm.Init()
	file := "hoge/fuga.txt"
	gm.PutFile(file, "aaa\n")
	gm.Add(file)
	gm.Commit("-m", "initial commit")

	fmt.Println("done")
	// Output:
	// done
}
