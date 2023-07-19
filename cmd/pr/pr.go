package pr

import (
	"fmt"

	"github.com/WaterLemons2k/gh-watch-commits/api"
	"github.com/WaterLemons2k/gh-watch-commits/cmd/gh"
)

// Open open a pull request in a public repository on GitHub
//
// https://cli.github.com/manual/gh_pr_create
func Open(repo *api.RepositoryV3, title, body string, draft bool) error {
	for {
		commits, err := api.CompareCommits(repo.FullName(), repo.ParentOwner())
		if err != nil {
			return fmt.Errorf("failed to compare commits: %w", err)
		}

		// AheadBy greater than 0 means there are some commits behind the parent.
		// This means that PR can be created.
		if commits.AheadBy > 0 {
			// Generate title and body if they are not provided
			if title == "" {
				title = generateTitle(repo)
			}
			if body == "" {
				body = generateBody(repo)
			}

			args := []string{"pr", "create", "--repo", repo.FullName(), "--title", title, "--body", body, "--head", repo.FullParentBranch(), "--base", repo.ParentBranch(), "--no-maintainer-edit"}

			if draft {
				args = append(args, "--draft")
			}

			gh.Exec(args...)
			break

		} else {
			// PR can not be created. Make it creatable.
			fmt.Printf("There are %d commits ahead instead of behind %s. Make it behind...\n", commits.BehindBy, repo.FullBranch())

			// Update to the parent commit. Like `git reset HEAD~1`.
			err = api.UpdateReference(repo.FullName(), repo.ParentBranch(), commits.ParentSHA())
			if err != nil {
				return fmt.Errorf("failed to update to the parent commit: %w", err)
			}
		}
	}

	return nil
}

// generateTitle generate and return a title for opening PR
func generateTitle(repo *api.RepositoryV3) string {
	return "[watch-commits] from " + repo.FullParentName()
}

// generateBody generate and return a body for opening PR
func generateBody(repo *api.RepositoryV3) string {
	fpn := repo.FullParentName()

	return fmt.Sprintf(`Watching [%s](https://github.com/%s) Commits.

---

Created by [watch-commits](https://github.com/WaterLemons2k/gh-watch-commits).`,
		fpn, fpn)
}
