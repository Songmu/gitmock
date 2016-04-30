package gitmock_test

import (
	"fmt"
	"log"
	"os"

	"github.com/Songmu/gitmock"
)

func ExampleNew() {
	gm, err := gitmock.New()
	if err != nil {
		log.Print(err)
		return
	}
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
