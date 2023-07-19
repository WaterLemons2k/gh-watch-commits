package api

import "testing"

func TestUpdateReference(t *testing.T) {
	type args struct {
		repo   string
		branch string
		sha    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "non-existent repo, branch and sha",
			args: args{
				repo:   "foo/bar",
				branch: "main",
				sha:    "23b058bbbef50f01891b3487f7e9e8bcf62979cb",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateReference(tt.args.repo, tt.args.branch, tt.args.sha); (err != nil) != tt.wantErr {
				t.Errorf("UpdateReference() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
