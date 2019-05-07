gitmock
=======

[![Build Status](https://travis-ci.org/Songmu/gitmock.svg?branch=master)][travis]
[![Coverage Status](https://coveralls.io/repos/Songmu/gitmock/badge.svg?branch=master)][coveralls]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDoc](https://godoc.org/github.com/Songmu/gitmock?status.svg)][godoc]

[travis]: https://travis-ci.org/Songmu/gitmock
[coveralls]: https://coveralls.io/r/Songmu/gitmock?branch=master
[license]: https://github.com/Songmu/gitmock/blob/master/LICENSE
[godoc]: https://godoc.org/github.com/Songmu/gitmock

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
