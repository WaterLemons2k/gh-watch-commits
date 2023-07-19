package api

import (
	"reflect"
	"testing"
)

func TestForkRepo(t *testing.T) {
	type args struct {
		repo              string
		forkName          string
		org               string
		defaultBranchOnly bool
	}
	tests := []struct {
		name    string
		args    args
		want    *RepositoryV3
		wantErr bool
	}{
		{
			name: "non-existent repo",
			args: args{
				repo: "foo/bar",
			},
			wantErr: true,
		},
		{
			name: "non-existent org",
			args: args{
				repo: "cli/cli",
				org:  "foo",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ForkRepo(tt.args.repo, tt.args.forkName, tt.args.org, tt.args.defaultBranchOnly)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForkRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForkRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRenameRepo(t *testing.T) {
	type args struct {
		repo        *RepositoryV3
		newRepoName string
	}
	tests := []struct {
		name    string
		args    args
		want    *RepositoryV3
		wantErr bool
	}{
		{
			name: "non-existent repo name",
			args: args{
				repo: &RepositoryV3{
					Name: "bar",
					Owner: struct{ Login string }{
						Login: "foo",
					},
				},
				newRepoName: "foo/bar1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RenameRepo(tt.args.repo, tt.args.newRepoName)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenameRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RenameRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}
