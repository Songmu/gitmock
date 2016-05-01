package gitmock

//go:generate go run _tools/gen.go
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
)

// New creates a new git mock repository.
// If gitPath is the empty string, `git` is used.
func New(gitPath string) (*GitMock, error) {
	return (&GitMock{gitPath: gitPath}).Initialize()
}

// GitMock is git mock repository
type GitMock struct {
	repoPath string
	gitPath  string
	user     string
	email    string
}

// Initialize the GitMock
func (gm *GitMock) Initialize() (*GitMock, error) {
	if gm.gitPath == "" {
		gm.gitPath = "git"
	}

	cmd := exec.Command(gm.gitPath, "version")
	cmd.Env = append(os.Environ(), "LANG=C")
	var b bytes.Buffer
	cmd.Stdout = &b
	err := cmd.Run()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new GitMock")
	}
	gitVer := b.String()
	arr := strings.Fields(gitVer)
	if len(arr) != 3 || arr[0] != "git" || arr[1] != "version" {
		return nil, fmt.Errorf("output of `git version` looks strange: %s", gitVer)
	}
	verArr := strings.Split(arr[2], ".")
	if len(verArr) < 3 {
		return nil, fmt.Errorf("git version [%s] looks strange", arr[2])
	}
	semv := strings.Join(verArr[:3], ".")
	ver, err := semver.NewVersion(semv)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("git version [%s] looks strange", arr[2]))
	}
	c, _ := semver.NewConstraint(">= 1.8.5")
	if !c.Check(ver) {
		return nil, fmt.Errorf("git 1.8.5 or later required")
	}

	err = exec.Command(gm.gitPath, "config", "user.name").Run()
	if err != nil {
		gm.user = "gitmock"
	}
	err = exec.Command(gm.gitPath, "config", "user.email").Run()
	if err != nil {
		gm.email = "gitmock@example.com"
	}

	return gm, nil
}

// RepoPath returns repository path
func (gm *GitMock) RepoPath() string {
	if gm.repoPath == "" {
		dir, err := ioutil.TempDir("", "gitmock-")
		if err != nil {
			log.Fatal(err.Error())
		}
		gm.repoPath = dir
	}
	return gm.repoPath
}

func (gm *GitMock) gitProg() string {
	if gm.gitPath != "" {
		return gm.gitPath
	}
	return "git"
}

func (gm *GitMock) env() (ret []string) {
	if gm.user != "" {
		envs := []string{"GIT_AUTHOR_NAME", "GIT_COMMITTER_NAME"}
		for _, v := range envs {
			if env := os.Getenv(v); env == "" {
				ret = append(ret, fmt.Sprintf("%s=%s", v, gm.user))
			}
		}
	}
	if gm.email != "" {
		envs := []string{"GIT_AUTHOR_EMAIL", "GIT_COMMITTER_EMAIL"}
		for _, v := range envs {
			if env := os.Getenv(v); env == "" {
				ret = append(ret, fmt.Sprintf("%s=%s", v, gm.email))
			}
		}
	}
	return ret
}

// Do the git command
func (gm *GitMock) Do(args ...string) (string, string, error) {
	arg := append([]string{"-C", gm.RepoPath()}, args...)
	cmd := exec.Command(gm.gitProg(), arg...)
	cmd.Env = append(os.Environ(), "LANG=C")
	cmd.Env = append(cmd.Env, gm.env()...)

	var bout, berr bytes.Buffer
	cmd.Stdout = &bout
	cmd.Stderr = &berr
	err := cmd.Run()
	return bout.String(), berr.String(), err
}

// PutFile put a file to repo
func (gm *GitMock) PutFile(file, content string) error {
	repo := gm.RepoPath()
	fpath := filepath.Join(repo, file)
	err := os.MkdirAll(filepath.Dir(fpath), 0755)
	if err != nil {
		return err
	}
	c := []byte(content)
	return ioutil.WriteFile(fpath, c, 0644)
}
