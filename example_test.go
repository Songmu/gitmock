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

	fmt.Println("done")
	// Output:
	// done
}
