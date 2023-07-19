package flags

import (
	"flag"
	"fmt"
	"os"
)

type Flag struct {
	Repo              string
	ForkName          string
	Org               string
	DefaultBranchOnly bool
	Title             string
	Body              string
	Draft             bool
}

// New returns a new Flag
func New() *Flag {
	repo := flag.String("R", "", "repository using the OWNER/REPO format")
	forkName := flag.String("fork-name", "", "Rename the forked repository")
	org := flag.String("org", "", "Create the fork in an organization")
	defaultBranchOnly := flag.Bool("default-branch-only", false, "Only include the default branch in the fork")
	title := flag.String("t", "", "Title for the pull request")
	body := flag.String("b", "", "Body for the pull request")
	draft := flag.Bool("d", false, "Mark pull request as a Draft")

	// https://stackoverflow.com/a/23726033
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), `
Usage:
  gh watch-commits [-R <repository>] [flags]
		
Flags:`)
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	f := Flag{
		Repo:              *repo,
		ForkName:          *forkName,
		Org:               *org,
		DefaultBranchOnly: *defaultBranchOnly,
		Title:             *title,
		Body:              *body,
		Draft:             *draft,
	}

	// Print usage when no repository is provided
	if f.Repo == "" {
		fmt.Fprintln(os.Stderr, "flag not provided: -R")
		flag.Usage()
		return nil
	}

	return &f
}
