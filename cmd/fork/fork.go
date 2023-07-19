package fork

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/WaterLemons2k/gh-watch-commits/api"
	"github.com/WaterLemons2k/gh-watch-commits/cmd/repo"
)

func Repo(repoName, forkName, org string, defaultBranchOnly bool) (*api.RepositoryV3, error) {
	forkedRepo, err := api.ForkRepo(repoName, forkName, org, defaultBranchOnly)

	if err != nil {
		return nil, err
	}

	// From https://github.com/cli/cli/blob/7f3196fcd4976a29ca6b4eba6229e5d3a02fe145/pkg/cmd/repo/fork/fork.go#L199-L203
	//
	// This is weird. There is not an efficient way to determine via the GitHub API whether or not a
	// given user has forked a given repo. We noticed, also, that the create fork API endpoint just
	// returns the fork repo data even if it already exists -- with no change in status code or
	// anything. We thus check the created time to see if the repo is brand new or not; if it's not,
	// we assume the fork already existed and report an error.
	createdAgo := time.Since(forkedRepo.CreatedAt)
	if createdAgo > time.Minute {
		fmt.Fprintf(os.Stderr, "%s already exists\n", forkedRepo.FullName())
	} else {
		fmt.Fprintf(os.Stderr, "created fork %s\n", forkedRepo.FullName())

		// Wait 1s to make sure the GitHub API is ready to go
		time.Sleep(time.Second)
	}

	if forkName != "" && !strings.EqualFold(forkedRepo.Name, repo.NormalizeRepoName(forkName)) {
		forkedRepo, err = api.RenameRepo(forkedRepo, forkName)
		if err != nil {
			return nil, fmt.Errorf("could not rename fork: %w", err)
		}
		fmt.Fprintf(os.Stderr, "Renamed fork to %s\n", forkedRepo.FullName())
	}

	return forkedRepo, nil
}
