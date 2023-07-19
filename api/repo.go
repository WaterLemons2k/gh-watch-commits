package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cli/go-gh/v2/pkg/api"
)

// RepositoryV3 is the repository result from GitHub API v3
//
// https://github.com/cli/cli/blob/8c7935e6cec01cc910e40bad6ec9999508eb4d83/api/queries_repo.go#L498
type RepositoryV3 struct {
	Name      string
	CreatedAt time.Time `json:"created_at"`
	Owner     struct {
		Login string
	}
	Parent        *RepositoryV3
	DefaultBranch string `json:"default_branch"`
}

// FullName serializes a GitHub repository into an "OWNER/REPO" string
//
// https://github.com/cli/cli/blob/8c7935e6cec01cc910e40bad6ec9999508eb4d83/internal/ghrepo/repo.go#L35
func (r *RepositoryV3) FullName() string {
	return fmt.Sprintf("%s/%s", r.owner(), r.Name)
}

// ParentFullName serializes a parent GitHub repository into
// an "OWNER/REPO" string
func (r *RepositoryV3) FullParentName() string {
	return fmt.Sprintf("%s/%s", r.ParentOwner(), r.ParentName())
}

// ParentName return the parent name of a GitHub repository
func (r *RepositoryV3) ParentName() string {
	return r.Parent.Name
}

// FullBranch seriallizes a GitHub branch into an "OWNER:BRANCH" string
func (r *RepositoryV3) FullBranch() string {
	return fmt.Sprintf("%s:%s", r.owner(), r.ParentBranch())
}

// FullParentBranch seriallizes a parent GitHub branch into
// an "OWNER:BRANCH" string
func (r *RepositoryV3) FullParentBranch() string {
	return fmt.Sprintf("%s:%s", r.ParentOwner(), r.ParentBranch())
}

// ParentBranch return the default parent branch of a GitHub repository
//
// use default parent branch to avoid default branch may not return
// real default branch when creating fork, but return "main"
func (r *RepositoryV3) ParentBranch() string {
	return r.Parent.DefaultBranch
}

// Owner return the owner of a GitHub repository
func (r *RepositoryV3) owner() string {
	return r.Owner.Login
}

// ParentOwner return the parent owner of a GitHub repository
func (r *RepositoryV3) ParentOwner() string {
	return r.Parent.Owner.Login
}

// ForkRepo forks the Repository on GitHub and returns the new Repository
//
// https://docs.github.com/rest/repos/forks?apiVersion=2022-11-28#create-a-fork
// https://github.com/cli/cli/blob/8c7935e6cec01cc910e40bad6ec9999508eb4d83/api/queries_repo.go#L511
func ForkRepo(repo, forkName, org string, defaultBranchOnly bool) (*RepositoryV3, error) {
	// repos/OWNER/REPO/forks
	path := fmt.Sprintf("repos/%s/forks", repo)

	params := map[string]interface{}{}
	if org != "" {
		params["organization"] = org
	}
	if forkName != "" {
		params["name"] = forkName
	}
	if defaultBranchOnly {
		params["default_branch_only"] = true
	}

	client, err := api.DefaultRESTClient()
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	enc := json.NewEncoder(body)
	if err := enc.Encode(params); err != nil {
		return nil, err
	}

	result := RepositoryV3{}
	err = client.Post(path, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// RenameRepo renames the repository on GitHub and returns
// the renamed repository
//
// https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#update-a-repository
// https://github.com/cli/cli/blob/trunk/api/queries_repo.go#L558
func RenameRepo(repo *RepositoryV3, newRepoName string) (*RepositoryV3, error) {
	path := fmt.Sprintf("repos/%s", repo.FullName())

	input := map[string]string{"name": newRepoName}
	body := &bytes.Buffer{}
	enc := json.NewEncoder(body)
	if err := enc.Encode(input); err != nil {
		return nil, err
	}

	client, err := api.DefaultRESTClient()
	if err != nil {
		return nil, err
	}

	result := RepositoryV3{}
	err = client.Patch(path, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
