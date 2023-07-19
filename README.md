# watch-commits

✨ A GitHub (`gh`) CLI extension to watching commits in a repository.

English | [简体中文](README.zh-CN.md)

## Motivation

GitHub has been missing a way to watching commits in a repository, but https://stackoverflow.com/a/42600376 offers another solution: via pull request.

This extension is also the CLI version of this solution.

## Installation

```sh
gh extension install WaterLemons2k/gh-watch-commits
```

## Usage

Run:

```sh
gh watch-commits -R <repository>
```

Replace `<repository>` with the repository to which you want to watching commits.

If you have successfully created a pull request, please note the following:

- do **NOT** merge this pull request
  - if you merged it, don't worry, just run this  command again to open another pull request.
- under the Email section of your [Notifications](https://github.com/settings/notifications) settings enable:
  - Comments on Issues and Pull Requests
  - Pull Request reviews
  - Pull Request pushes

That's it. You will receive email notifications about every commit on default branch.

Run `gh watch-commits -h` for more help:

```
Usage:
  gh watch-commits [-R <repository>] [flags]
		
Flags:
  -R string
    	repository using the OWNER/REPO format
  -b string
    	Body for the pull request
  -d	Mark pull request as a Draft
  -default-branch-only
    	Only include the default branch in the fork
  -fork-name string
    	Rename the forked repository
  -org string
    	Create the fork in an organization
  -t string
    	Title for the pull request
```
