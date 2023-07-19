package repo

import (
	"regexp"
	"strings"
)

var invalidCharactersRE = regexp.MustCompile(`[^\w._-]+`)

// NormalizeRepoName takes in the repo name the user inputted and normalizes it using the same logic as GitHub (GitHub.com/new)
//
// https://github.com/cli/cli/blob/2a4160a3a38d3c05a1395b32cd422d5fe1a8e92d/pkg/cmd/repo/shared/repo.go#L11
func NormalizeRepoName(repoName string) string {
	newName := invalidCharactersRE.ReplaceAllString(repoName, "-")
	return strings.TrimSuffix(newName, ".git")
}

// IsValid return true when extracts the repository information from the
// following string formats: "OWNER/REPO", otherwise return false.
//
// https://github.com/cli/go-gh/blob/7adca2a0702063e9dad0683f24e9a64d0c5ee6d9/pkg/repository/repository.go#L26
func IsValid(repo string) bool {
	parts := strings.SplitN(repo, "/", 3)

	if len(parts) != 2 {
		return false
	}

	for _, p := range parts {
		if len(p) == 0 {
			return false
		}
	}

	return true
}
