package api

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
)

// commits is the Commits result from GitHub API
//
// https://docs.github.com/rest/commits/commits?apiVersion=2022-11-28
type Commits struct {
	BaseCommit struct {
		Parents []struct {
			SHA string
		}
	} `json:"base_commit"`
	AheadBy  int `json:"ahead_by"`
	BehindBy int `json:"behind_by"`
}

// ParentSHA return the first parent's SHA of a GitHub commit
func (c *Commits) ParentSHA() string {
	return c.BaseCommit.Parents[0].SHA
}

// CompareCommits compares two commits against one another.
//
// https://docs.github.com/en/rest/commits/commits?apiVersion=2022-11-28#compare-two-commits
func CompareCommits(repo, parentOwner string) (*Commits, error) {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return nil, fmt.Errorf("couldn't create client: %w", err)
	}

	// repos/OWNER/REPO/compare/HEAD...OWNER:HEAD?per_page=1
	path := fmt.Sprintf("repos/%s/compare/HEAD...%s:HEAD?per_page=1", repo, parentOwner)

	result := Commits{}
	err = client.Get(path, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
