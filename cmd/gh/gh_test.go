package gh

import "testing"

func TestExec(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "exec gh --version",
			args: args{
				args: []string{"--version"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Exec(tt.args.args...)
		})
	}
}
