gitmock
=======

[![Test Status](https://github.com/Songmu/gitmock/workflows/test/badge.svg?branch=main)][actions]
[![codecov.io](https://codecov.io/github/Songmu/gitmock/coverage.svg?branch=main)][codecov]
[![MIT License](https://img.shields.io/github/license/Songmu/gitmock)][license]
[![PkgGoDev](https://pkg.go.dev/badge/github.com/Songmu/gitmock)][PkgGoDev]

[actions]: https://github.com/Songmu/gitmock/actions?workflow=test
[codecov]: https://codecov.io/github/Songmu/gitmock?branch=main
[license]: https://github.com/Songmu/gitmock/blob/main/LICENSE
[PkgGoDev]: https://pkg.go.dev/github.com/Songmu/gitmock

## Description

Create mock git repository for testing.

## Synopsis

```go
gm, err := gitmock.New("")
if err != nil {
    log.Fatal(err)
}
defer os.RemoveAll(gm.RepoPath())
gm.Init() // shortcut of `gm.Do("init")
file := "hoge/fuga.txt"
gm.PutFile(file, "aaa\n")
gm.Add(file)
gm.Commit("-m", "initial commit")
out, stderr, err := gm.Status()
```

## Author

[Songmu](https://github.com/Songmu)
