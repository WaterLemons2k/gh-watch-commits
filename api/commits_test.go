package api

import (
	"reflect"
	"testing"
)

func TestCompareCommits(t *testing.T) {
	type args struct {
		repo        string
		parentOwner string
	}
	tests := []struct {
		name    string
		args    args
		want    *Commits
		wantErr bool
	}{
		{
			name: "non-existent repo and parent owner",
			args: args{
				repo:        "foo/bar",
				parentOwner: "bar",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareCommits(tt.args.repo, tt.args.parentOwner)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareCommits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareCommits() = %v, want %v", got, tt.want)
			}
		})
	}
}
