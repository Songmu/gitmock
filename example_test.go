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
	gm.Do("init")
	file := "hoge/fuga.txt"
	gm.PutFile(file, "aaa\n")
	gm.Do("add", file)
	gm.Do("commit", "-m", "initial commit")

	fmt.Println("done")
	// Output:
	// done
}
