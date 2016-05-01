package gitmock

// Clone Do git clone
func (gm *GitMock) Clone(args ...string) (string, string, error) {
	arg := append([]string{"clone"}, args...)
	return gm.Do(arg...)
}

// Init Do git init
func (gm *GitMock) Init(args ...string) (string, string, error) {
	arg := append([]string{"init"}, args...)
	return gm.Do(arg...)
}

// Add Do git add
func (gm *GitMock) Add(args ...string) (string, string, error) {
	arg := append([]string{"add"}, args...)
	return gm.Do(arg...)
}

// Mv Do git mv
func (gm *GitMock) Mv(args ...string) (string, string, error) {
	arg := append([]string{"mv"}, args...)
	return gm.Do(arg...)
}

// Reset Do git reset
func (gm *GitMock) Reset(args ...string) (string, string, error) {
	arg := append([]string{"reset"}, args...)
	return gm.Do(arg...)
}

// Rm Do git rm
func (gm *GitMock) Rm(args ...string) (string, string, error) {
	arg := append([]string{"rm"}, args...)
	return gm.Do(arg...)
}

// Log Do git log
func (gm *GitMock) Log(args ...string) (string, string, error) {
	arg := append([]string{"log"}, args...)
	return gm.Do(arg...)
}

// Show Do git show
func (gm *GitMock) Show(args ...string) (string, string, error) {
	arg := append([]string{"show"}, args...)
	return gm.Do(arg...)
}

// Status Do git status
func (gm *GitMock) Status(args ...string) (string, string, error) {
	arg := append([]string{"status"}, args...)
	return gm.Do(arg...)
}

// Branch Do git branch
func (gm *GitMock) Branch(args ...string) (string, string, error) {
	arg := append([]string{"branch"}, args...)
	return gm.Do(arg...)
}

// Checkout Do git checkout
func (gm *GitMock) Checkout(args ...string) (string, string, error) {
	arg := append([]string{"checkout"}, args...)
	return gm.Do(arg...)
}

// Commit Do git commit
func (gm *GitMock) Commit(args ...string) (string, string, error) {
	arg := append([]string{"commit"}, args...)
	return gm.Do(arg...)
}

// Diff Do git diff
func (gm *GitMock) Diff(args ...string) (string, string, error) {
	arg := append([]string{"diff"}, args...)
	return gm.Do(arg...)
}

// Merge Do git merge
func (gm *GitMock) Merge(args ...string) (string, string, error) {
	arg := append([]string{"merge"}, args...)
	return gm.Do(arg...)
}

// Rebase Do git rebase
func (gm *GitMock) Rebase(args ...string) (string, string, error) {
	arg := append([]string{"rebase"}, args...)
	return gm.Do(arg...)
}

// Tag Do git tag
func (gm *GitMock) Tag(args ...string) (string, string, error) {
	arg := append([]string{"tag"}, args...)
	return gm.Do(arg...)
}

// LsFiles Do git ls-files
func (gm *GitMock) LsFiles(args ...string) (string, string, error) {
	arg := append([]string{"ls-files"}, args...)
	return gm.Do(arg...)
}
