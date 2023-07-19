package gh

import (
	"fmt"
	"os"
	"os/exec"
)

// Exec invokes a gh command in a subprocess.
//
// https://pkg.go.dev/github.com/cli/go-gh#Exec
func Exec(args ...string) {
	cmd := exec.Command("gh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "gh execution failed: %v", err)
		return
	}
}
