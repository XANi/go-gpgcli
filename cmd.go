package gpgcli

import (
//	"io"
	"os/exec"
	"bytes"
	//	"fmt"
)

func (g *Gpg) cmd (args ...string) (string, string, error) {
	var err error
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var cmdArgs []string
	cmdArgs = append(cmdArgs, g.gpgOptions...)
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command(`gpg` ,cmdArgs...)
	// make commands consistent
	cmd.Env = g.filteredEnv
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		return "", "", err
	}
	if err := cmd.Wait(); err != nil {
			return stdout.String(), stderr.String(), err
	}
	return stdout.String(), stderr.String(), err
}
