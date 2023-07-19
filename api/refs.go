package api

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
)

// referencesBody is the update a reference body from GitHub API
type referencesBody struct {
	SHA   string `json:"sha"`
	Force bool   `json:"force"`
}

// ResetToParentCommit uses GitHub API to update to a commit.
//
// A Git reference (git ref) is a file that contains a Git commit SHA-1 hash.
//
// You can use this as `git reset HEAD~1`.
//
// https://docs.github.com/rest/git/refs?apiVersion=2022-11-28#update-a-reference
func UpdateReference(repo, parentBranch, parentSha string) error {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return fmt.Errorf("couldn't create client: %w", err)
	}

	// repos/OWNER/REPO/git/refs/heads/main
	path := fmt.Sprintf("repos/%s/git/refs/heads/%s", repo, parentBranch)

	opts := referencesBody{
		SHA:   parentSha,
		Force: true,
	}

	body := &bytes.Buffer{}
	enc := json.NewEncoder(body)
	if err := enc.Encode(opts); err != nil {
		return err
	}

	err = client.Patch(path, body, nil)
	if err != nil {
		return err
	}

	return nil
}
