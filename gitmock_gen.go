package gitmock

// Clone Do git clone
func (gm *GitMock) Clone(args ...string) (string, string, error) {
	arg := []string{"clone"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Init Do git init
func (gm *GitMock) Init(args ...string) (string, string, error) {
	arg := []string{"init"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Add Do git add
func (gm *GitMock) Add(args ...string) (string, string, error) {
	arg := []string{"add"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Mv Do git mv
func (gm *GitMock) Mv(args ...string) (string, string, error) {
	arg := []string{"mv"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Reset Do git reset
func (gm *GitMock) Reset(args ...string) (string, string, error) {
	arg := []string{"reset"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Rm Do git rm
func (gm *GitMock) Rm(args ...string) (string, string, error) {
	arg := []string{"rm"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Log Do git log
func (gm *GitMock) Log(args ...string) (string, string, error) {
	arg := []string{"log"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Show Do git show
func (gm *GitMock) Show(args ...string) (string, string, error) {
	arg := []string{"show"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Status Do git status
func (gm *GitMock) Status(args ...string) (string, string, error) {
	arg := []string{"status"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Branch Do git branch
func (gm *GitMock) Branch(args ...string) (string, string, error) {
	arg := []string{"branch"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Checkout Do git checkout
func (gm *GitMock) Checkout(args ...string) (string, string, error) {
	arg := []string{"checkout"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Commit Do git commit
func (gm *GitMock) Commit(args ...string) (string, string, error) {
	arg := []string{"commit"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Diff Do git diff
func (gm *GitMock) Diff(args ...string) (string, string, error) {
	arg := []string{"diff"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Merge Do git merge
func (gm *GitMock) Merge(args ...string) (string, string, error) {
	arg := []string{"merge"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Rebase Do git rebase
func (gm *GitMock) Rebase(args ...string) (string, string, error) {
	arg := []string{"rebase"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// Tag Do git tag
func (gm *GitMock) Tag(args ...string) (string, string, error) {
	arg := []string{"tag"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}

// LsFiles Do git ls-files
func (gm *GitMock) LsFiles(args ...string) (string, string, error) {
	arg := []string{"ls-files"}
	arg = append(arg, args...)
	return gm.Do(arg...)
}
