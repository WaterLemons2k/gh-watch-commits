package repo

import (
	"testing"
)

// https://github.com/cli/cli/blob/2a4160a3a38d3c05a1395b32cd422d5fe1a8e92d/pkg/cmd/repo/shared/repo_test.go#L7
func TestNormalizeRepoName(t *testing.T) {
	// confirmed using GitHub.com/new
	tests := []struct {
		LocalName      string
		NormalizedName string
	}{
		{
			LocalName:      "cli",
			NormalizedName: "cli",
		},
		{
			LocalName:      "cli.git",
			NormalizedName: "cli",
		},
		{
			LocalName:      "@-#$^",
			NormalizedName: "---",
		},
		{
			LocalName:      "[cli]",
			NormalizedName: "-cli-",
		},
		{
			LocalName:      "Hello World, I'm a new repo!",
			NormalizedName: "Hello-World-I-m-a-new-repo-",
		},
		{
			LocalName:      " @E3H*(#$#_$-ZVp,n.7lGq*_eMa-(-zAZSJYg!",
			NormalizedName: "-E3H-_--ZVp-n.7lGq-_eMa---zAZSJYg-",
		},
		{
			LocalName:      "I'm a crazy .git repo name .git.git .git",
			NormalizedName: "I-m-a-crazy-.git-repo-name-.git.git-",
		},
	}
	for _, tt := range tests {
		output := NormalizeRepoName(tt.LocalName)
		if output != tt.NormalizedName {
			t.Errorf("Expected %q, got %q", tt.NormalizedName, output)
		}
	}
}

// https://github.com/cli/go-gh/blob/7adca2a0702063e9dad0683f24e9a64d0c5ee6d9/pkg/repository/repository_test.go#L10
func TestIsValid(t *testing.T) {
	type args struct {
		repo string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "OWNER/REPO combo",
			args: args{
				repo: "cli/cli",
			},
			want: true,
		},
		{
			name: "too few elements",
			args: args{
				repo: "cli",
			},
			want: false,
		},
		{
			name: "too many elements",
			args: args{
				repo: "a/b/c",
			},
			want: false,
		},
		{
			name: "blank value",
			args: args{
				repo: "a/",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.repo); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
