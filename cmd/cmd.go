package cmd

import (
	"fmt"
	"os"

	"github.com/WaterLemons2k/gh-watch-commits/cmd/flags"
	"github.com/WaterLemons2k/gh-watch-commits/cmd/fork"
	"github.com/WaterLemons2k/gh-watch-commits/cmd/pr"
	"github.com/WaterLemons2k/gh-watch-commits/cmd/repo"
)

func Run() {
	flag := flags.New()

	if !repo.IsValid(flag.Repo) {
		fmt.Fprintf(os.Stderr, `expected the "OWNER/REPO" format, got "%s"`, flag.Repo)
		return
	}

	forkedRepo, err := fork.Repo(flag.Repo, flag.ForkName, flag.Org, flag.DefaultBranchOnly)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to fork: %v", err)
		return
	}

	err = pr.Open(forkedRepo, flag.Title, flag.Body, flag.Draft)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create pull request: %v", err)
	}
}
